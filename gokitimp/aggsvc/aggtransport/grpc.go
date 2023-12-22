package aggtransport

import (
	"context"
	"github.com/cmkqwerty/freight-fare-engine/gokitimp/aggsvc/aggendpoint"
	"github.com/cmkqwerty/freight-fare-engine/gokitimp/aggsvc/aggservice"
	"github.com/cmkqwerty/freight-fare-engine/gokitimp/aggsvc/pb"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/transport"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"time"
)

type grpcServer struct {
	aggregate grpctransport.Handler
	calculate grpctransport.Handler
	pb.UnimplementedAggregatorServer
}

func NewGRPCServer(endpoints aggendpoint.Set, logger log.Logger) pb.AggregatorServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	return &grpcServer{
		aggregate: grpctransport.NewServer(
			endpoints.AggregateEndpoint,
			decodeGRPCAggregateRequest,
			encodeGRPCAggregateResponse,
			options...,
		),
		calculate: grpctransport.NewServer(
			endpoints.CalculateEndpoint,
			decodeGRPCGetInvoiceRequest,
			encodeGRPCGetInvoiceResponse,
			options...,
		),
	}
}

func (s *grpcServer) Aggregate(ctx context.Context, req *pb.AggregateRequest) (*pb.AggregateReply, error) {
	_, rep, err := s.aggregate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AggregateReply), nil
}

func (s *grpcServer) GetInvoice(ctx context.Context, req *pb.GetInvoiceRequest) (*pb.GetInvoiceReply, error) {
	_, rep, err := s.calculate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetInvoiceReply), nil
}

func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) aggservice.Service {
	limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))

	var options []grpctransport.ClientOption

	var aggEndpoint endpoint.Endpoint
	{
		aggEndpoint = grpctransport.NewClient(
			conn,
			"pb.Aggregator",
			"Aggregate",
			encodeGRPCAggregateRequest,
			decodeGRPCAggregateResponse,
			pb.AggregateReply{},
			options...,
		).Endpoint()
		aggEndpoint = limiter(aggEndpoint)
		aggEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Aggregate",
			Timeout: 30 * time.Second,
		}))(aggEndpoint)
	}

	var calcEndpoint endpoint.Endpoint
	{
		calcEndpoint = grpctransport.NewClient(
			conn,
			"pb.Aggregator",
			"GetInvoice",
			encodeGRPCGetInvoiceRequest,
			decodeGRPCGetInvoiceResponse,
			pb.GetInvoiceReply{},
			options...,
		).Endpoint()
		calcEndpoint = limiter(calcEndpoint)
		calcEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Calculate",
			Timeout: 10 * time.Second,
		}))(calcEndpoint)
	}

	return aggendpoint.Set{
		AggregateEndpoint: aggEndpoint,
		CalculateEndpoint: calcEndpoint,
	}
}

func decodeGRPCAggregateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.AggregateRequest)
	return aggendpoint.AggregateRequest{
		Value: req.Value,
		OBUID: int(req.ObuID),
		Unix:  req.Unix,
	}, nil
}

func decodeGRPCGetInvoiceRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetInvoiceRequest)
	return aggendpoint.CalculateRequest{OBUID: int(req.ObuID)}, nil
}

func decodeGRPCAggregateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	_ = grpcReply.(*pb.AggregateReply)
	return aggendpoint.AggregateResponse{}, nil
}

func decodeGRPCGetInvoiceResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetInvoiceReply)
	return aggendpoint.CalculateResponse{
		OBUID:         int(reply.ObuID),
		TotalDistance: reply.TotalDistance,
		TotalAmount:   reply.TotalAmount,
	}, nil
}

func encodeGRPCAggregateResponse(_ context.Context, response interface{}) (interface{}, error) {
	_ = response.(aggendpoint.AggregateResponse)
	return &pb.AggregateReply{}, nil
}

func encodeGRPCGetInvoiceResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(aggendpoint.CalculateResponse)
	return &pb.GetInvoiceReply{
		ObuID:         int32(resp.OBUID),
		TotalDistance: resp.TotalDistance,
		TotalAmount:   resp.TotalAmount,
	}, nil
}

func encodeGRPCAggregateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(aggendpoint.AggregateRequest)
	return &pb.AggregateRequest{
		ObuID: int32(req.OBUID),
		Value: req.Value,
		Unix:  req.Unix,
	}, nil
}

func encodeGRPCGetInvoiceRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(aggendpoint.CalculateRequest)
	return &pb.GetInvoiceRequest{ObuID: int32(req.OBUID)}, nil
}
