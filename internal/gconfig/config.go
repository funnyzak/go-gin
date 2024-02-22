package gconfig

import (
	"go-gin/pkg/utils/file"
	"os"
	"path"
)

const (
	DefaultLogPath        = "logs/log.log"
	DefaultPprofRoutePath = "/debug/pprof"
)

type Notification struct {
	Type      string              `mapstructure:"type"`
	Instances []map[string]string `mapstructure:"instances"`
}

type Config struct {
	Server struct {
		Port uint `mapstructure:"port"`
	} `mapstructure:"server"`
	Site struct {
		Brand       string `mapstructure:"brand"`
		Description string `mapstructure:"description"`
		BaseURL     string `mapstructure:"base_url"`
	} `mapstructure:"site"`
	Debug     bool   `mapstructure:"debug"`
	DBPath    string `mapstructure:"db_path"`
	RateLimit struct {
		Max int `mapstructure:"max"`
	} `mapstructure:"rate_limit"`
	EnableCORS             bool `mapstructure:"enable_cors"`
	EnableUserRegistration bool `mapstructure:"enable_user_registration"`
	Upload                 struct {
		Dir              string   `mapstructure:"dir"`
		VirtualPath      string   `mapstructure:"virtual_path"`
		URLPrefix        string   `mapstructure:"url_prefix"`
		MaxSize          int64    `mapstructure:"max_size"`
		KeepOriginalName bool     `mapstructure:"keep_original_name"`
		CreateDateDir    bool     `mapstructure:"create_date_dir"`
		AllowTypes       []string `mapstructure:"allow_types"`
	} `mapstructure:"upload"`
	Log struct {
		Level string `mapstructure:"level"`
		Path  string `mapstructure:"path"`
	} `mapstructure:"log"`
	JWT struct {
		AccessSecret           string `mapstructure:"access_secret"`
		RefreshSecret          string `mapstructure:"refresh_secret"`
		AccessTokenExpiration  int    `mapstructure:"access_token_expiration"`
		RefreshTokenExpiration int    `mapstructure:"refresh_token_expiration"`
		AccessTokenCookieName  string `mapstructure:"access_token_cookie_name"`
		RefreshTokenCookieName string `mapstructure:"refresh_token_cookie_name"`
	} `mapstructure:"jwt"`
	Location      string         `mapstructure:"location"`
	Notifications []Notification `mapstructure:"notifications"`
}

func CreateDefaultConfigFile(config_path string) error {
	err := file.MkdirAllIfNotExists(path.Dir(config_path), os.ModePerm)
	if err != nil {
		return err
	}
	err = file.WriteToFile(config_path, ConfigYamlTemplate, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

const ConfigYamlTemplate = `
server:
  port: 8080 # Server port
site:
  brand: Go-Gin # Site brand
  description: A simple web application using Go and Gin # Site description
  base_url: http://localhost:8080 # Site base URL, used for generating absolute URLs, cant end with /
debug: false # Debug mode, if true, the server will print detailed error messages
log:
  level: debug # debug, info, warn, error, fatal, panic
  path: logs/go-gin.log # Log file path, relative to the project root directory. Or you can use an absolute path, such as /var/www/go-gin.log
db_path: db/go-gin.sqlite # Database path, relative to the project root directory. Or you can use an absolute path, such as /var/www/go-gin.sqlite
rate_limit:
  max: 100 # requests per minute, 0 means no limit
enable_cors: false # Enable CORS
enable_user_registration: true # Enable user registration
upload:
  virtual_path: /upload # Virtual path, used for generating absolute URLs, must start with /, cant end with /
  url_prefix: http://localhost:8080/upload # URL prefix, used for generating absolute URLs, must start with http:// or https:// or /, cant end with /
  dir: upload # Upload directory, relative to the project root directory. Or you can use an absolute path, such as /var/www/upload
  max_size: 10485760 # 10MB, unit: byte
  keep_original_name: false # Keep original file name, if false, the server will generate a random file name
  create_date_dir: true # Create date directory, such as /upload/2021/01/01
  allow_types: # Allowed file types, if empty, all types are allowed
    - image/jpeg
    - image/jpg
    - image/png
    - image/gif
    - image/bmp
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
`
