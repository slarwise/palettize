# Palettize

Fit an image to a colorscheme.

![Example](https://github.com/slarwise/palettize/assets/25964718/ed3e7225-9d5d-432e-b25c-d79425865a90)

## Usage

### Web version

Start the backend server:

```sh
cd ./backend
go run .
```

Open `frontend/index.html` in your browser:

```sh
open ./frontend/index.html
```

### CLI

[See the CLI README](/cli/README.md)

## Observability

Use grafana's LGTM stack:

- Traces: Tempo
- Metrics (TODO): Prometheus or Mimir
- Logs (TODO): Loki
- Visualization (TODO): Grafana

## Development

Start a local cluster, with a local image registry. Examples:

- Docker Desktop: Enable kubernetes
- k3d: `k3d create my-cluster --registry-create my-registry`

Start all services by running `tilt up`. Make some requests to the backend to
see what happens:

```sh
curl http://localhost:3001/colorschemes -d '{"name": "cool", "color": "red"}'
curl http://localhost:3001/colorschemes/cool
curl http://localhost:3001/colorschemes/does-not-exist
```
