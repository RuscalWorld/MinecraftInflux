package main

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

var InfluxClient influxdb2.Client

func SetupInfluxClient() {
	config := AppConfig.InfluxConfig
	InfluxClient = influxdb2.NewClient(config.URL, config.Token)
}

func GetWriteAPI() api.WriteAPI {
	config := AppConfig.InfluxConfig
	return InfluxClient.WriteAPI(config.Organization, config.Bucket)
}
