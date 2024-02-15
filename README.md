# Go Gin

Gin gonic starter with zerolog, viper, gorm, jwt basic setup.

[![Build Status][build-status-image]][build-status]
[![license][license-image]][repository-url]
[![GitHub repo size][repo-size-image]][repository-url]
[![docker][docker-image]][docker-url]
[![tag][tag-image]][rle-url]

## Development

If you want to develop with this project, you can follow the steps below.

1. Clone the repository and navigate to the project directory.

   ```bash
    git clone git@github.com:funnyzak/go-gin.git && cd go-gin
   ```
  
2. Copy the `config.example.json` file to `config.json` and update the values.

   ```bash
    cp config.example.json config.json
   ```

3. Run the application.

   ```bash
    go run cmd/main.go
    # or 
    make dev

    # You also specify the config file, e.g. dev, prod, etc.
    go run cmd/main.go -c dev
    ```

### CI/CD

You can click `Use this template` to create a new repository based on this project. and add Secrets Keys: `DOCKER_USERNAME` and `DOCKER_PASSWORD` in the repository settings. And when you push the code, it will automatically build binary and docker image and push to the Docker Hub.

## Structure

```plaintext
├── Dockerfile              // Dockerfile defines how to build a Docker image for the project
├── Makefile                // Contains commands for building, running, testing, etc. the project
├── cmd
│   ├── main.go             // The main entry point for the application
│   └── srv                 // Server controller
├── config.yaml.example     // An example configuration file for the project
├── docker-compose.yml      // Defines services, networks and volumes for docker-compose
├── internal
│   ├── gconfig             // Internal package for configuration
│   └── gogin               // Internal package for the gin framework
├── mappers
│   ├── auth.go             // Data mapper for authentication
│   └── post.go             // Data mapper for posts
├── model
│   ├── auth.go             // Data model for authentication
│   ├── common.go           // Common data models
│   ├── post.go             // Data model for posts
│   └── user.go             // Data model for users
├── pkg
│   ├── logger              // Package for logging
│   ├── mygin               // Custom package for the gin framework
│   └── utils               // Utility functions
├── resource
│   ├── resource.go         // Resource management
│   ├── static              // Static files such as HTML, CSS, JavaScript, images
│   └── template            // Templates for Gin framework
├── script
│   └── build.sh            // A script for building the project
└── service
    └── singleton           // Singleton services for the application
```

## Build

```bash
# Compile multiple platforms architecture (Linux, Windows, MacOS)
make build

# Compile the specified platform architecture
GOOS=linux GOARCH=amd64 go build -o go-gin-linux-amd64 cmd/main.go

# Compile current platform architecture
go build -o go-gin cmd/main.go
```

## Deployment

You can deploy the application using the following ways.

### Docker Deployment

Docker deployment requires the installation of Docker first, and then execute the command.

#### One-click Deployment

Start the service with default configuration parameters, as follows:

```bash
docker run -d \
  --name go-gin \
  --restart always \
  -p 8080:8080 \
  -v ./config.yaml:/app/config.yaml \
  funnyzak/go-gin:latest
```

#### Compose Startup

```bash
# Pull source code
git clone https://go-gin && cd go-gin
# Compose startup, startup parameter configuration can be done by modifying the docker-compose.yml file
docker compose up -d
```

If you need to update the container, you can re-pull the code and build the image in the source code folder, the command is as follows:

```bash
git pull && docker compose up -d --build
```

### Binary Startup

You can pull the source code to compile the binary executable file yourself, or download the binary executable file of the corresponding system architecture from the repository, and then execute the following command to start (note that **go-gin** is the name of the binary executable file, please replace it according to the actual name):

```bash
# Quick start (The config.yaml file needs to be in the same directory as the binary executable file)
./go-gin

# Specify the configuration file. eg. prod, the configuration file is prod.yaml 
./go-gin -c prod

# View help, see available parameters
./go-gin -h
```

**Note:** Please make sure that executable permissions have been set before running. If there are no executable permissions, you can set them through the `chmod +x go-gin` command.

## License

MIT License

[repo-size-image]: https://img.shields.io/github/repo-size/funnyzak/go-gin?style=flat-square&logo=github&logoColor=white&label=size
[down-latest-image]: https://img.shields.io/github/downloads/funnyzak/go-gin/latest/total.svg
[down-total-image]: https://img.shields.io/github/downloads/funnyzak/go-gin/total.svg
[commit-activity-image]: https://img.shields.io/github/commit-activity/m/funnyzak/go-gin?style=flat-square
[last-commit-image]: https://img.shields.io/github/last-commit/funnyzak/go-gin?style=flat-square
[license-image]: https://img.shields.io/github/license/funnyzak/go-gin.svg?style=flat-square
[repository-url]: https://github.com/funnyzak/go-gin
[rle-url]: https://github.com/funnyzak/go-gin/releases/latest
[rle-all-url]: https://github.com/funnyzak/go-gin/releases
[ci-image]: https://img.shields.io/github/workflow/status/funnyzak/go-gin/react-native-android-build-apk
[ci-url]: https://github.com/funnyzak/go-gin/actions
[rle-image]: https://img.shields.io/github/release-date/funnyzak/go-gin.svg?style=flat-square&label=release
[sg-image]: https://img.shields.io/badge/view%20on-Sourcegraph-brightgreen.svg?style=flat-square
[sg-url]: https://sourcegraph.com/github.com/funnyzak/go-gin
[build-status-image]: https://github.com/funnyzak/go-gin/actions/workflows/build.yml/badge.svg
[build-status]: https://github.com/funnyzak/go-gin/actions
[tag-image]: https://img.shields.io/github/tag/funnyzak/go-gin.svg
[docker-image]: https://img.shields.io/docker/pulls/funnyzak/go-gin
[docker-url]: https://hub.docker.com/r/funnyzak/go-gin
