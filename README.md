# Go Gin

Quick start web application using Go and Gin.

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

## Structure

```plaintext
├── Dockerfile
├── Makefile
├── cmd
├── config.yaml.example
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
├── model
├── pkg
├── resource
└── script
```

## License

MIT License
