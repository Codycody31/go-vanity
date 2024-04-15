# Go Vanity URL Server

The Go Vanity URL Server simplifies creating custom URLs for Go packages, particularly those hosted on GitHub. It acts as a redirection service, facilitating `go get` operations for your customized package paths.

## Features

- Configurable via a simple YAML file.
- Supports custom paths for Go packages.
- Designed for straightforward enhancements and easy maintenance.

## Getting Started

### Prerequisites

- **Go Installation**: This project requires Go 1.22 or higher. You can verify your Go version with:

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

## Docker Usage

### Prerequisites

- **Docker Installed**: Ensure you have Docker installed on your system. You can check by running `docker --version`. If not installed, you can download it from the official Docker website.

### Getting the Docker Image

You can pull the pre-built Docker image from the Docker Hub using the following command:

```bash
docker pull insidiousfiddler/vanity
```

### Running the Server with Docker

1. **Using a Local Configuration File**:
   To run the server using a local configuration file, you can mount the configuration directory to the Docker container. Make sure your `config.yaml` file is in a suitable directory, then use the following command to start the server:

   ```bash
   docker run -p 8080:8080 -v /path/to/your/config/directory:/etc/vanity insidiousfiddler/vanity
   ```

   This command mounts your local directory containing the `config.yaml` at `/etc/vanity` inside the container, which is the default path the Docker image expects for the configuration file.

2. **Using a Remote Configuration File**:
   If you prefer to fetch the configuration from a remote URL, set the `VANITY_CONFIG_URL` environment variable when running the Docker container:

   ```bash
   docker run -p 8080:8080 -e VANITY_CONFIG_URL="https://example.com/config.yaml" insidiousfiddler/vanity
   ```

   This setup is useful for centralized configuration management, allowing you to update the configuration without rebuilding or restarting containers manually.

### Custom Docker Configuration

For more advanced Docker configurations, such as network settings, logging configurations, or using Docker Compose, you can create a `docker-compose.yml` file or extend the Docker run commands with additional parameters as per your requirements.

## Contributing

We welcome contributions from the community. To contribute:

1. Fork the repository.
2. Create a new branch for your feature (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -am 'Add some AmazingFeature'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).
5. Submit a pull request.

## License

This project is licensed under the MIT License - see the `LICENSE` file for details.
