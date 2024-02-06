package model

const (
	DefaultLogPath        = "logs/log.log"
	DefaultPprofRoutePath = "/debug/pprof"
)

type ErrorResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
