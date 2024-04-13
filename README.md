# Go Vanity URL Server

This Go Vanity URL Server provides a simple way to use custom URLs for your Go packages hosted on GitHub. It serves as a redirection point that allows `go get` to fetch packages from custom paths.

## Features

- Easy configuration with a YAML file.
- Supports custom package paths.
- Simple and clear structure for easy enhancements and maintenance.

## Getting Started

### Prerequisites

Before you start, ensure you have Go installed on your machine. This project uses Go modules, so Go 1.11+ is required. You can check your Go version using:

```bash
go version
```

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/yourusername/go-vanity.git
cd go-vanity
```

Build the project:

```bash
go build
```

### Running the Server

To start the server, run:

```bash
./go-vanity
```

By default, the server reads the configuration from `config.yaml` and starts on port 8080. You can visit `http://localhost:8080/your-package-path` to see the redirection metadata.

### Configuration

The server is configured via a `config.yaml` file located in the root directory. Here's an example configuration:

```yaml
packages:
  - path: "go.example.com/mylib"
    repo: "https://github.com/username/mylib"
  - path: "go.example.com/myotherlib"
    repo: "https://github.com/username/myotherlib"
```

Each item under `packages` should include:
- `path`: The custom URL path for your Go package.
- `repo`: The actual GitHub repository URL where the Go package is hosted.

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Project Link: [https://github.com/yourusername/go-vanity](https://github.com/yourusername/go-vanity)