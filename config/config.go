package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	ServiceHost string
	ServicePort int

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase       string
	PostgresUser     string
	PostgresPassword string

	LogLevel string
	HttpPort string

	SmsServiceHost string
	SmsServicePort int

	AuthServiceHost string
	AuthServicePort int

	CorporateServiceHost string
	CorporateServicePort int

	CDN               string
	MinioBucket       string
	CancelledStatusId string

	// Incident Statuses
	IncidentStatusReopened   string
	IncidentStatusClosed     string
	IncidentStatusSolved     string
	IncidentStatusNew        string
	IncidentStatusInProccess string

	SvgateRequestBreak int // minutes

	CryptoKey  string
	EncryptKey string

	// Determines if current node is active
	IsActive                    bool
	DefaultEducationDegreeField string
}

// Load loads environment vars and inflates Config
func Load() Config {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config := Config{}
	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "development"))

	config.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))

	config.ServiceHost = cast.ToString(getOrReturnDefault("SERVICE_HOST", "localhost"))
	config.ServicePort = cast.ToInt(getOrReturnDefault("SERVICE_PORT", 8008))

	config.CorporateServiceHost = cast.ToString(getOrReturnDefault("CORPORATE_SERVICE_HOST", "localhost"))
	config.CorporateServicePort = cast.ToInt(getOrReturnDefault("CORPORATE_SERVICE_PORT", 5003))

	config.SmsServiceHost = cast.ToString(getOrReturnDefault("SMS_SERVICE_HOST", "localhost"))
	config.SmsServicePort = cast.ToInt(getOrReturnDefault("SMS_SERVICE_PORT", 5004))

	config.AuthServiceHost = cast.ToString(getOrReturnDefault("AUTH_SERVICE_HOST", "localhost"))
	config.AuthServicePort = cast.ToInt(getOrReturnDefault("AUTH_SERVICE_PORT", 5000))

	config.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	config.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "m_go_service"))
	config.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "citizenfour"))
	config.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "12321"))

	config.CDN = cast.ToString(getOrReturnDefault("CDN", "https://test.cdn.iman.uz"))
	config.MinioBucket = cast.ToString(getOrReturnDefault("MINIO_BUCKET", "documents"))

	config.SvgateRequestBreak = cast.ToInt(getOrReturnDefault("SVGATE_REQUEST_BREAK", 150))

	config.CancelledStatusId = cast.ToString(getOrReturnDefault("CANCELLED_STATUS_ID", "a3dac56c-9c44-49e6-8db2-7c74a31b6073"))

	config.CryptoKey = cast.ToString(getOrReturnDefault("CRYPTO_KEY", ""))

	config.EncryptKey = cast.ToString(getOrReturnDefault("ENCRYPT_KEY", ""))

	config.IsActive = cast.ToBool(getOrReturnDefault("IS_ACTIVE", true))

	config.DefaultEducationDegreeField = cast.ToString(getOrReturnDefault("EDUCATION_DEGREE", "b2876aff-38b3-11ec-b81a-df71afaf00f4"))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
