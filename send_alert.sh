alert='[
  {
    "labels": {
       "alertname": "DiskRunningFull",
       "dev": "sda1",
       "instance": "example1"
     },
     "annotations": {
        "info": "The disk sda1 is running full",
        "summary": "please check the instance example1"
      }
  }
]'
curl -XPOST -d"$alert" http://localhost:9001/api/v1/alerts
curl -XPOST -d"$alert" http://localhost:9002/api/v1/alerts
curl -XPOST -d"$alert" http://localhost:9003/api/v1/alerts
