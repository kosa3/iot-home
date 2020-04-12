# iot-home

## Description

This is a system for metering room conditions with Kibana using NatureRemo API.
Metering items are humidity, temperature, and illuminance.

## System

Go / Rabbit MQ / ElasticSearch / Kibana

## Starting Local

```bash
$ cp .env.sample .env // edit your information
$ docker-compose up
```

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
        "created_at" : {
          "type" : "date"
        },
        "humidity" : {
          "type" : "float"
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