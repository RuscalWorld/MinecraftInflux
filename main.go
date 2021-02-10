package main

import (
	"fmt"
	"log"
	"time"

	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/whatupdave/mcping"
)

func main() {
	LoadConfig()

	SetupInfluxClient()
	defer InfluxClient.Close()

	writeAPI := GetWriteAPI()

	for {
		playersPoint := write.NewPointWithMeasurement("players")
		latencyPoint := write.NewPointWithMeasurement("latency")

		for _, server := range AppConfig.PingConfig.Servers {
			response, err := mcping.Ping(server.Address)
			if err != nil {
				log.Println(fmt.Sprintf("Unable to ping %s (%s): %s", server.Name, server.Address, err))
				continue
			}

			playersPoint.AddTag("server", server.Name)
			playersPoint.AddField("players", response.Online)
			writeAPI.WritePoint(playersPoint)

			latencyPoint.AddTag("server", server.Name)
			latencyPoint.AddField("latency", response.Latency)
			writeAPI.WritePoint(latencyPoint)
		}

		time.Sleep(time.Duration(AppConfig.PingConfig.Interval) * time.Second)
	}
}
