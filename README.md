# Cloudflare exporter
Export [Cloudflare](https://www.cloudflare.com/) zone statistics to [Prometheus](https://prometheus.io/).

## Usage
Run the binary:
```shell
./cloudflare-exporter --web.listen-address=:9178
```

## Exported Metrics
| Metric | Description | Labels |
| ------ | ------- | ------ |

## Flags
```shell
./cloudflare-exporter --help
```

| Flag | Description | Default |
| ---- | ----------- | ------- |
| web.listen-address | The address to listen on for HTTP requests.| `:9178` |
| web.metrics-path | Path under which to expose metrics. | `/metrics` |
| web.health-path | Path under which to expose exporter health. | `/health` |