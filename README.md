
# Go Gin

Gin gonic starter with zerolog, viper, gorm, jwt, go-cache, rate-limit, cron, notification, etc.

[![Build Status][build-status-image]][build-status]
[![Go Version][go-version-image]](https://github.com/funnyzak/go-gin/blob/main/go.mod)
[![docker][docker-image]][docker-url]
[![license][license-image]][repository-url]
[![GitHub repo size][repo-size-image]][repository-url]
[![release][rle-image]][rle-url]

## Development

If you want to develop with this project, you can follow the steps below.

1. Clone the repository and navigate to the project directory.

   ```bash
    git clone git@github.com:funnyzak/go-gin.git && cd go-gin
   ```
  
2. Copy the `config.example.yaml` file to `config.yaml` and update the values.

   ```bash
    cp config.example.yaml config.yaml
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

You can fork this repository and add Secrets Keys: `DOCKER_USERNAME` and `DOCKER_PASSWORD` in the repository settings. And when you push the code, it will automatically build binary and docker image and push to the Docker Hub.

## Structure

```plaintext
├── Dockerfile              // Dockerfile defines how to build a Docker image for the project
├── Makefile                // Contains commands for building, running, testing, etc. the project
├── cmd
│   ├── main.go             // The main entry point for the application
│   └── srv                 // Server controller
├── config.example.yaml     // An example configuration file for the project
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

## Configuration

The configuration file is in the `config.yaml` file, you can copy the `config.example.yaml` file to `config.yaml` and update the values, the configuration file is as follows:

```yaml
server:
  port: 8080 # Server port
site:
  brand: Go-Gin # Site brand
  description: A simple web application using Go and Gin # Site description
  base_url: http://localhost:8080 # Site base URL, used for generating absolute URLs
debug: false # Debug mode, if true, the server will print detailed error messages
log:
  level: debug # debug, info, warn, error, fatal, panic
  path: logs/go-gin.log # Log file path
db_path: db/go-gin.sqlite # Database path
rate_limit:
  max: 100 # requests per minute
upload:
  dir: upload # Upload directory
  max_size: 10485760 # 10MB
jwt: # JWT settings
  access_secret: qhkxjrRmYcVYKSEobqsvhxhtPVeTWquu # Access token secret
  refresh_secret: qhkxjrRmYcVYKSEobqsvhxhtPV3TWquu # Refresh token secret
  access_token_expiration: 60 # minutes
  refresh_token_expiration: 720 # minutes
  access_token_cookie_name: go-gin-access # Access token cookie name
  refresh_token_cookie_name: go-gin-refresh # Refresh token cookie name
location: Asia/Chongqing # Timezone
notifications: # Notification settings
  - type: apprise # You must install apprise first, more details: https://github.com/caronc/apprise
    instances:
      - url: "apprise-url-1"
      - url: "apprise-url-2"
  - type: dingtalk
    instances:
      - webhook: "dingtalk-webhook-1"
      - webhook: "dingtalk-webhook-2"
  - type: ifttt
    instances:
      - key: "ifttt-key-1"
        event: "event-1"
      - key: "ifttt-key-2"
        event: "event-2"
  - type: smtp
    instances:
      - host: "smtp-host-1"
        port: 587
        username: "user-1"
        password: "password-1"
        from: "from-1"
        to: "to-1"
      - host: "smtp-host-2"
        port: 587
        username: "user-2"
        password: "password-2"
        from: "from-2"
        to: "to-2"
  - type: telegram
    instances:
      - botToken: "telegram-bot-token-1"
        chatID: "chat-id-1"
      - botToken: "telegram-bot-token-2"
        chatID: "chat-id-2"
  - type: wecom
    instances:
      - key: "wecom-key-1"
      - key: "wecom-key-2"
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
git clone git@github.com:funnyzak/go-gin.git && cd go-gin
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

### Run as Service

#### Linux (systemd)

In Linux, services are managed through Systemd. You can use the following commands to install, start, stop, restart, log, and view the status of services, etc.

```bash
bash -c "$(curl -fsSL https://raw.githubusercontent.com/funnyzak/go-gin/main/script/install.sh)"
```

You can also install it manually. The specific steps are as follows:

<details>
  <summary> Click to expand </summary>

1. Download the binary executable file of the corresponding system architecture from the [releases](https://github.com/funnyzak/go-gin/releases) page or [GitHub Actions](https://github.com/funnyzak/go-gin/actions) page, and copy it to the `/opt/go-gin` directory.
2. Grant the executable permission to the file by running the following command:

    ```bash
    sudo chmod +x /opt/go-gin/go-gin
    ```

3. Download [go-gin.service](https://raw.githubusercontent.com/funnyzak/go-gin/main/script/go-gin.service) file to the `/etc/systemd/system` directory.
4. Download [config.example.yaml](https://raw.githubusercontent.com/funnyzak/go-gin/main/config.example.yaml) file to the `/opt/go-gin` directory and rename it to `go-gin.yaml`, and update the values.

Finally, run the following command to start the service:

```bash
sudo systemctl enable go-gin
systemctl start go-gin
```

</details>

#### MacOS (launchd)

Service on MacOS is based on launchd. You can use the following steps to install and start the service.

1. Download the binary executable file of the corresponding system architecture from the [releases](https://github.com/funnyzak/go-gin/releases) page or [GitHub Actions](https://github.com/funnyzak/go-gin/actions) page, and copy it to the `/opt/go-gin` directory.
2. Grant the executable permission to the file by running the following command:

    ```bash
    sudo chmod +x /opt/go-gin/go-gin
    ```

3. Download [com.go-gin.plist](https://raw.githubusercontent.com/funnyzak/go-gin/main/script/com.go-gin.plist) file to the `/Library/LaunchDaemons` directory.
4. Download [config.example.yaml](https://raw.githubusercontent.com/funnyzak/go-gin/main/config.example.yaml) file to the `/opt/go-gin` directory and rename it to `config.yaml`, and update the values.

Finally, run the following command to start the service:

```bash
sudo launchctl load /Library/LaunchDaemons/com.go-gin.plist
sudo launchctl start /Library/LaunchDaemons/com.go-gin.plist
```

#### Windows

Service on Windows can be implemented using Task Scheduler. You can use the following steps to install and start the service.

1. Dwnload the binary executable file of the corresponding system architecture from the [releases](https://github.com/funnyzak/go-gin/releases) page or [GitHub Actions](https://github.com/funnyzak/go-gin/actions) page, and unzip it to the `C:\go-gin` directory.
2. Download [install.ps1](https://raw.githubusercontent.com/funnyzak/go-gin/main/script/install.ps1) file to the `C:\go-gin` directory and rename it to `go-gin.ps1`.
3. Download [config.example.yaml](https://raw.githubusercontent.com/funnyzak/go-gin/main/config.example.yaml) file to the `C:\go-gin` directory and rename it to `config.yaml`, and update the values.

Finally, run the following command to start the service:

```powershell
cd C:\go-gin
.\go-gin.ps1 enable
```

The following are all the commands supported by the script:

```powershell
./go-gin.ps1 enable  # Enable and start the service
./go-gin.ps1 disable # Disable and stop the service
./go-gin.ps1 start   # Start the service
./go-gin.ps1 stop    # Stop the service
./go-gin.ps1 restart # Restart the service
./go-gin.ps1 status  # View the service status
```

## FOSSA Scan

[![FOSSA Status][fossa-image]][fossa-url]

## LICENSE

[MIT License](https://opensource.org/licenses/MIT)

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
[rle-image]: https://img.shields.io/github/release/funnyzak/go-gin.svg?style=smartthings
[sg-image]: https://img.shields.io/badge/view%20on-Sourcegraph-brightgreen.svg?style=flat-square
[sg-url]: https://sourcegraph.com/github.com/funnyzak/go-gin
[build-status-image]: https://github.com/funnyzak/go-gin/actions/workflows/build.yml/badge.svg
[build-status]: https://github.com/funnyzak/go-gin/actions
[tag-image]: https://img.shields.io/github/tag/funnyzak/go-gin.svg
[docker-image]: https://img.shields.io/docker/pulls/funnyzak/go-gin
[docker-url]: https://hub.docker.com/r/funnyzak/go-gin
[fossa-image]: https://app.fossa.com/api/projects/git%2Bgithub.com%2Ffunnyzak%2Fgo-gin.svg?type=large
[fossa-url]: https://app.fossa.com/projects/git%2Bgithub.com%2Ffunnyzak%2Fgo-gin?ref=badge_large
[go-version-image]: https://img.shields.io/github/go-mod/go-version/funnyzak/go-gin?logo=go
