package app

import (
	"os"
	"strconv"
	"sync"

	log "github.com/sirupsen/logrus"
)

var onceConfig sync.Once
var config *Configuration

// Configuration contains config settings read from env variable
type Configuration struct {
	ApplicationPort    string
	Host               string
	Port               string
	Name               string
	User               string
	Password           string
	Type               string
	SSLMode            string
	MaxDBConnections   int
	AcquireConnTimeout int
	MailAPIKey         string
	BaseURL            string
	BasePort           string
	BaseProtocol       string
	S3Region           string
	S3Endpoint         string
	S3AccessKey        string
	S3SecretKey        string
	S3ReportBucket     string
	S3UserImageBucket  string
	S3UserFileBucket   string
	JwtSecret          string
	DatabaseURL        string
}

const (
	// ProductionEnv is the env value to denote Production environment
	ProductionEnv = "production"
)

// GetConfig will return configuration from env variables in singleton pattern
func GetConfig() *Configuration {
	onceConfig.Do(func() {
		config = GetConfiguration()
	})
	return config
}

// GetConfiguration instantiates Configuration
func GetConfiguration() *Configuration {
	mustLoadEnv := func(varName string) string {
		envVar := os.Getenv(varName)
		if envVar == "" {
			log.Fatal("Env var ", varName, " is required but not defined")
		}
		return envVar
	}
	canLoadEnvOrDefault := func(varName string, defaultValue string) string {
		envVar := os.Getenv(varName)
		if envVar == "" {
			log.Info("Using default value ", defaultValue, " for ", varName)
			return defaultValue
		}
		return envVar
	}
	conf := &Configuration{}
	conf.ApplicationPort = canLoadEnvOrDefault("PORT", "5422")
	conf.Host = mustLoadEnv("PGHOST")
	conf.Port = mustLoadEnv("PGPORT")
	conf.Name = mustLoadEnv("PGDATABASE")
	conf.User = mustLoadEnv("PGUSER")
	conf.Password = canLoadEnvOrDefault("PGPASSWORD", "")
	conf.Type = mustLoadEnv("DBTYPE")
	if conf.Type == "sqlite" {
		conf.DatabaseURL = mustLoadEnv("DATABASE_URL")
	}
	conf.SSLMode = mustLoadEnv("PGSSLMODE")
	conf.MaxDBConnections, _ = strconv.Atoi(canLoadEnvOrDefault("PGMAXCONNECTIONS", "20"))
	conf.AcquireConnTimeout, _ = strconv.Atoi(canLoadEnvOrDefault("DB_CONN_ACQUIRE_TIMEOUT", "30"))
	conf.BaseURL = mustLoadEnv("BASE_URL")
	conf.BaseProtocol = mustLoadEnv("BASE_PROTOCOL")
	return conf
}

// Print logs current configuration to stdout
func (c *Configuration) Print() {
	log.Info("Loaded configuration with settings")
	log.Info("Host: ", c.Host)
	log.Info("Port: ", c.Port)
	log.Info("User: ", c.User)
	log.Info("Password: ", len(c.Password), " characters")
	log.Info("Type: ", c.Type)
	log.Info("SSLMode: ", c.SSLMode)
	log.Info("Database name: ", c.Name)
	log.Info("S3 Default Region: ", c.S3Region)
	log.Info("S3 Endpoint: ", c.S3Endpoint)
	log.Info("S3 Reports Bucket: ", c.S3ReportBucket)
	log.Info("S3 User Images Bucket: ", c.S3UserImageBucket)
	log.Info("S3 User Files Bucket: ", c.S3UserFileBucket)
	log.Info("S3 User Access Key: ", len(c.S3AccessKey), "characters")
	log.Info("S3 User Secret Access Key: ", len(c.S3SecretKey), "characters")
	log.Info("Sendgrid API Key: ", len(c.MailAPIKey), " characters")
	log.Info("Mail Base Url: ", c.BaseURL)
	log.Info("Max database connections:", strconv.Itoa(c.MaxDBConnections))
}
