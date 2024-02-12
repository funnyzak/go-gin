package gconfig

const (
	DefaultLogPath        = "logs/log.log"
	DefaultPprofRoutePath = "/debug/pprof"
)

type Config struct {
	Server struct {
		Port uint `mapstructure:"port"`
	} `mapstructure:"server"`
	Site struct {
		Brand       string `mapstructure:"brand"`
		Description string `mapstructure:"description"`
		BaseURL     string `mapstructure:"base_url"`
		CookieName  string `mapstructure:"cookie_name"`
	} `mapstructure:"site"`
	Debug     bool   `mapstructure:"debug"`
	DBPath    string `mapstructure:"db_path"`
	RateLimit struct {
		Max int `mapstructure:"max"`
	} `mapstructure:"rate_limit"`
	Upload struct {
		Dir     string `mapstructure:"dir"`
		MaxSize int    `mapstructure:"max_size"`
	} `mapstructure:"upload"`
	Log struct {
		Level string `mapstructure:"level"`
		Path  string `mapstructure:"path"`
	} `mapstructure:"log"`
	JWT struct {
		AccessSecret           string `mapstructure:"access_secret"`
		RefreshSecret          string `mapstructure:"refresh_secret"`
		TokenExpiration        int    `mapstructure:"token_expiration"`
		RefreshTokenExpiration int    `mapstructure:"refresh_token_expiration"`
	} `mapstructure:"jwt"`
	Users map[string]string `mapstructure:"users"`
}
