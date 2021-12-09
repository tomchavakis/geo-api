# How-to

## Start Prometheus-Grafana

update the targets from the `prometheus.yml` file according to your network setup

```
  - job_name: app
    scrape_interval: 5s
    static_configs:
      - targets: ['192.168.0.108:3001']
```

```
docker-compose up
```
navigate to http:localhost:3000 and make the appropriate queries

## Start the Application

```
make run
```