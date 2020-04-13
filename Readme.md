# iot-home

## Description

This is a system for metering room conditions with Kibana using NatureRemo API.
Metering items are humidity, temperature, and illuminance.

## System

Go / Rabbit MQ / ElasticSearch / Kibana

![iot-home](https://user-images.githubusercontent.com/19683276/79072208-bbb45200-7d1a-11ea-8a69-49d61a79d61e.png)

## Starting Local

```bash
$ cp .env.sample .env // edit your information
$ docker-compose up
```

## RabbitMQ Local

*RabbitMQ endpoint*

`http://localhost:5672/`


*RabbitMQ Management Admin Panel*

`http://localhost:15672/`

## ElasticSearch Local Settings

```bash
$ curl -H "Content-Type: application/json" -XPUT 'http://localhost:9200/natureremo' -d @datastore/mapping.json
$ curl -XGET "http://localhost:9200/natureremo/_mapping?pretty"
{
  "natureremo" : {
    "mappings" : {
      "properties" : {
        "Timestamp" : {
          "type" : "date"
        },
        "id" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "humidity" : {
          "type" : "float"
        },
        "lux" : {
          "type" : "float"
        },
        "temp" : {
          "type" : "float"
        }
      }
    }
  }
}
```

## Kibana Local

*UI Dashboard*

`http://localhost:5601/`

## Cron Job

*Mac*

```bash
$ cp job/iot-home.plist /Users/kosa3/Library/LaunchAgents
# Load plist
$ launchctl load ~/Library/LaunchAgents/iot-home.plist
# UnLoad plist
$ launchctl unload ~/Library/LaunchAgents/iot-home.plist
```