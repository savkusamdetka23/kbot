# Open-telemetry observability

Sample configuration for Kbot that send logs to [OpenTelemetry Collector] and metrics to [OpenTelemetry Collector] or [Prometheus].

## Prerequisites

- [Docker]
- [Docker Compose]

## How to run

```bash
export TELE_TOKEN=<TOKEN>
docker-compose -f otel/docker-compose.yaml up 
```
## Examples of Loki and Prometheus charts.
### Total requests on Prometheus
![Alt text](../img/image10.png)

### Memory usage
![Alt text](../img/image-11.png)

### Logs
![Alt text](../img/image-2.png)
![Alt text](../img/image-4.png)
![Alt text](../img/image-3.png)