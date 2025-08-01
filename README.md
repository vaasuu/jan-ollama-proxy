# jan-ollama-proxy

A tiny HTTP proxy for [Jan (jan.ai)](https://jan.ai) that removes headers unsupported by [Ollama](https://ollama.com).

Jan includes an `Origin: null` header in its requests, which causes Ollama to return `403 Forbidden`.  
This proxy strips the `Origin` header and forwards the request to your local Ollama server.

See related issue: [menloresearch/jan#5474](https://github.com/menloresearch/jan/issues/5474)

## Features

- Removes the `Origin` header (which Ollama rejects)
- Forwards all requests to a specified backend (default: `http://localhost:11434`)
- Listens on a configurable port (default: `8080`)
- Minimal: written in Go using only the standard library

## Usage

1. Install Go if you haven't already.
2. Clone or download this repository.
3. Run the proxy:

```sh
go run main.go
```

Or specify custom port and backend:

```sh
go run main.go --port=8090 --backend=http://127.0.0.1:11434
```

You can build the Go binary with `go build ./main.go`.

### `--help` argument

```
./main --help
Usage of ./main:
  -backend string
    	Backend URL to proxy to (default "http://localhost:11434")
  -port string
    	Port for the proxy to listen on (default "8080")
```

4. In Jan, set your LLM API URL to match the proxy (e.g. `http://localhost:8080`)

## Diagram

```
Jan → proxy → Ollama
     [Proxy removes Origin header]
```

## Flags

| Flag        | Description                | Default                  |
| ----------- | -------------------------- | ------------------------ |
| `--port`    | Port the proxy listens on  | `8080`                   |
| `--backend` | URL to forward requests to | `http://localhost:11434` |

## License

MIT
