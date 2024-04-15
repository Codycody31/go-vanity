# Go Vanity URL Server

The Go Vanity URL Server simplifies creating custom URLs for Go packages, particularly those hosted on GitHub. It acts as a redirection service, facilitating `go get` operations for your customized package paths.

## Features

- Configurable via a simple YAML file.
- Supports custom paths for Go packages.
- Designed for straightforward enhancements and easy maintenance.

## Getting Started

### Prerequisites

- **Go Installation**: This project requires Go 1.11 or higher due to its dependency on Go modules. You can verify your Go version with:

  ```bash
  go version
  ```

### Installation

1. **Clone the Repository**:
   Clone the project to your local environment:

   ```bash
   git clone https://github.com/Codycody31/go-vanity.git
   cd go-vanity
   ```

2. **Build the Server**:
   Compile the server from the source code:

   ```bash
   go build -o vanity
   ```

### Running the Server

To run the server:

```bash
./vanity
```

By default, the server utilizes `config.yaml` from the current directory and listens on port 8080. You can access `http://localhost:8080/your-package-path` to test the redirection.

### Configuration

The server's behavior is controlled through the `config.yaml` file. Here’s an example configuration that specifies the domain and package repositories:

```yaml
domain: "go.example.com"
packages:
  - path: "mylib"
    repo: "https://github.com/username/mylib"
    vcs: "git"
  - path: "myotherlib"
    repo: "https://github.com/username/myotherlib"
    vcs: "git"
```

- `domain`: Your custom domain for hosting Go packages.
- `packages`: A list of packages with their paths and repository URLs.

**Environment Variables**:

- `VANITY_CONFIG`: Set this to specify a local path to an alternative configuration file.
- `VANITY_CONFIG_URL`: Set this to specify a URL where the configuration file can be fetched, useful for centralized configuration management.

### Environment Setup

You can set environment variables directly or include them in a startup script:

```bash
export VANITY_CONFIG="/path/to/your/config.yaml"
export VANITY_CONFIG_URL="https://example.com/config.yaml"
./vanity
```

## Contributing

We welcome contributions from the community. To contribute:

1. Fork the repository.
2. Create a new branch for your feature (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -am 'Add some AmazingFeature'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).
5. Submit a pull request.

## License

This project is licensed under the MIT License - see the `LICENSE` file for details.
