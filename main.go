package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/urfave/cli"
	"github.com/whatupdave/mcping"
)

var ConfigPath string

func main() {
	app := &cli.App{
		Name: "MinecraftInflux",
		Commands: []cli.Command{
			{
				Name:   "start",
				Action: Start,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "config",
						Value:       "/etc/minecraft-influx/config.yml",
						Destination: &ConfigPath,
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Unable to run app:", err)
	}
}

func Start(_ *cli.Context) {
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
