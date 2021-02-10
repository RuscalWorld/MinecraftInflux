package main

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	InfluxConfig InfluxConfig `yaml:"influx"`
	PingConfig   PingConfig   `yaml:"ping"`
}

type InfluxConfig struct {
	URL          string `yaml:"url"`
	Organization string `yaml:"organization"`
	Token        string `yaml:"token"`
	Bucket       string `yaml:"bucket"`
}

type PingConfig struct {
	Interval int      `yaml:"interval"`
	Servers  []Server `yaml:"servers"`
}

type Server struct {
	Address string `yaml:"address"`
	Name    string `yaml:"name"`
}

var AppConfig = &Config{}

func LoadConfig() {
	log.Println("Loading config from", ConfigPath)

	file, err := os.Open(ConfigPath)
	if err != nil {
		if os.IsNotExist(err) {
			SaveDefaultConfig()
			log.Println("Saved default config at", ConfigPath)
		} else {
			log.Fatalln("Unable to open config file:", err)
			return
		}
	}

	defer file.Close()
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln("Unable to read config file:", err)
		return
	}

	err = yaml.Unmarshal(contents, AppConfig)
	if err != nil {
		log.Fatalln("Unable to parse config file:", err)
		return
	}

	log.Println("Successfully loaded config")
}

func SaveDefaultConfig() *os.File {
	contents, err := yaml.Marshal(Config{
		InfluxConfig: InfluxConfig{
			URL:          "http://localhost:8086",
			Organization: "MinecraftInflux",
			Bucket:       "Minecraft",
		},
		PingConfig: PingConfig{
			Interval: 60,
			Servers: []Server{
				{
					Address: "localhost:25565",
					Name:    "Example server",
				},
			},
		},
	})

	if err != nil {
		log.Fatalln("Unable to encode default config:", err)
		return nil
	}

	file, err := os.Create(ConfigPath)
	if err != nil {
		log.Fatalln("Unable to create config file:", err)
		return nil
	}

	defer file.Close()
	_, err = file.WriteString(string(contents))
	if err != nil {
		log.Fatalln("Unable to save default config:", err)
		return nil
	}

	return file
}
