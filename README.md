# MinecraftInflux
![](https://img.shields.io/github/license/RuscalWorld/MinecraftInflux)
![](https://img.shields.io/github/go-mod/go-version/RuscalWorld/MinecraftInflux)
![](https://img.shields.io/github/workflow/status/RuscalWorld/MinecraftInflux/build)
---
Simple program that allows you to save Minecraft server metrics to InfluxDB

## Installing
At the moment you can only use Docker to run MinecraftInflux

### Using Docker

```shell
docker run -d -v /etc/minecraft-influx:/etc/minecraft-influx --restart unless-stopped ruscalworld/minecraft-influx
```

## Configuring
Before configuring you should start MinecraftInflux at least once to generate default config file.
Now config for MinecraftInflux is located in `/etc/minecraft-influx` directory.
You can use any text editor to edit it.

```shell
nano /etc/minecraft-influx/config.yml
```

### InfluxDB

> **Only InfluxDB 2.0+ is supported!**

| Parameter | Description | Default value |
| --- | --- | --- |
| `url` | Address of your InfluxDB instance | `http://localhost:8086` |
| `organization` | Name of your organization in InfluxDB | `MinecraftInflux` |
| `token` | Token with read and write access that should be used | ` ` |
| `bucket` | Name of the bucket to use | `Minecraft` |

### Ping

| Parameter | Description | Default value |
| --- | --- | --- |
| `interval` | Interval between pings in seconds | `60` |
| `servers` | Array with servers to ping |  |

### Server

| Parameter | Description | Default value |
| --- | --- | --- |
| `address` | Server IP with port that should be pinged | `localhost:25565` |
| `name` | Name to identify this server in metrics | `Example server` |

## Default config

```yaml
influx:
  url: http://localhost:8086
  organization: MinecraftInflux
  token: ""
  bucket: Minecraft
ping:
  interval: 60
  servers:
  - address: localhost:25565
    name: Example server
```