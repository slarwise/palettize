# Palettize CLI

Make an image use only the colors of a specific colorscheme.

## Usage

### CLI

```sh
go run ./main.go  --output out.png --palette examples/palette.png examples/input.png
```

### Docker

```sh
docker build -t palettize-backend:dev .
docker run --rm -w /workspace -v ./examples:/workspace/data palettize-backend:dev --palette data/palette.png --output data/output.png data/input.png
```
