package config

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sanitize/sanitize"
	"github.com/johannes-kuhfuss/services_utils/api_error"
	"github.com/johannes-kuhfuss/services_utils/logger"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/microcosm-cc/bluemonday"
)

type AppConfig struct {
	Server struct {
		Host                 string `envconfig:"SERVER_HOST"`
		Port                 string `envconfig:"SERVER_PORT" default:"8080"`
		TlsPort              string `envconfig:"SERVER_TLS_PORT" default:"8443"`
		GracefulShutdownTime int    `envconfig:"GRACEFUL_SHUTDOWN_TIME" default:"10"`
		UseTls               bool   `envconfig:"USE_TLS" default:"false"`
		CertFile             string `envconfig:"CERT_FILE" default:"./cert/cert.pem"`
		KeyFile              string `envconfig:"KEY_FILE" default:"./cert/cert.key"`
	}
	Gin struct {
		Mode         string `envconfig:"GIN_MODE" default:"release"`
		TemplatePath string `envconfig:"TEMPLATE_PATH" default:"./templates/"`
	}
	DeviceDiscovery struct {
		IntervalSec int `envconfig:"DEVICE_DISCOVER_INTERVAL_SEC" default:"5"`
	}
	Misc struct {
		MaxResultLimit int      `envconfig:"MAX_RESULT_LIMIT" default:"100"`
		ApiKeys        []string `envconfig:"API_KEYS"`
	}
	RunTime struct {
		Router      *gin.Engine
		ListenAddr  string
		StartDate   time.Time
		RunDiscover bool
		Sani        *sanitize.Sanitizer
		BmPolicy    *bluemonday.Policy
	}
}

var (
	EnvFile = ".env"
)

func InitConfig(file string, config *AppConfig) api_error.ApiErr {
	logger.Info(fmt.Sprintf("Initalizing configuration from file %v", file))
	loadConfig(file)
	err := envconfig.Process("", config)
	if err != nil {
		return api_error.NewInternalServerError("Could not initalize configuration. Check your environment variables", err)
	}
	logger.Info("Done initalizing configuration")
	return nil
}

func loadConfig(file string) error {
	err := godotenv.Load(file)
	if err != nil {
		logger.Info("Could not open env file. Using Environment variable and defaults")
		return err
	}
	return nil
}
