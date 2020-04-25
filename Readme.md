# iot-home

## Description

This is a system for metering room conditions with Kibana using NatureRemo API.
Metering items are humidity, temperature, and illuminance.

## System

Go / Rabbit MQ / ElasticSearch / Kibana

![iot-home](https://user-images.githubusercontent.com/19683276/79072208-bbb45200-7d1a-11ea-8a69-49d61a79d61e.png)

## System on GKE

![k8s](https://user-images.githubusercontent.com/19683276/80268457-f71f2b00-86e1-11ea-95db-656262afce93.jpeg)


```bash
$ gcloud container clusters create iot-home --num-nodes=2
$ kubectl create secret generic nature-key --from-literal=nature-token='YOUR_NATURE_ACCESS_TOKEN'
$ kubectl apply -f iot-home.yaml
```

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
