package main

import (
	"github.com/cmkqwerty/freight-fare-engine/types"
	"github.com/gorilla/websocket"
	"log"
	"math"
	"math/rand"
	"time"
)

const wsEndpoint = "ws://localhost:30000/ws"

var sendInterval = time.Second

func generateCord() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()

	return n + f
}

func generateLatLong() (float64, float64) {
	return generateCord(), generateCord()
}

func generateOBUIDS(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}

	return ids
}

func main() {
	obuIDs := generateOBUIDS(20)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		for i := 0; i < len(obuIDs); i++ {
			lat, long := generateLatLong()
			data := types.OBUData{
				OBUID:     obuIDs[i],
				Latitude:  lat,
				Longitude: long,
			}

			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
		}
		time.Sleep(sendInterval)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
