# Palettize

Fit an image to a colorscheme.

## Usage

- Web version
- [CLI](/cli/README.md)

## Observability

Use grafana's LGTM stack:

- Traces: Tempo
- Metrics (TODO): Prometheus or Mimir
- Logs (TODO): Loki
- Visualization (TODO): Grafana

## Development

Start all services by running `tilt up`. Make some requests to the backend to
see what happens:

```sh
curl http://localhost:3001/colorschemes -d '{"name": "cool", "color": "red"}'
curl http://localhost:3001/colorschemes/cool
curl http://localhost:3001/colorschemes/does-not-exist
```
