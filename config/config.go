package config

import (
	"github.com/spf13/viper"
)

// AppConfig object with environment configuration information.
type AppConfig struct {
	App struct {
		Name  string
		Debug bool
		Env   string
		Port  string
	}

	Gin struct {
		Mode string
	}

	Database struct {
		Username string
		Password string
		Host     string
		Name     string
		Port     uint32
		SslMode  string
		Timezone string
		Engine   string
	}

	Cors struct {
		Origins []string
	}
}

var cfg *AppConfig

// Config load environments variables.
func Config() *AppConfig {
	if cfg == nil {
		loadConfig()
	}

	return cfg
}

func loadConfig() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// Ignore config file not found, perhaps we will use environment variables.
	_ = viper.ReadInConfig()

	cfg = &AppConfig{}

	// App.
	cfg.App.Name = viper.GetString("APP_NAME")
	cfg.App.Debug = viper.GetBool("APP_DEBUG")
	cfg.App.Env = viper.GetString("APP_ENV")
	cfg.App.Port = viper.GetString("APP_PORT")

	// Gin.
	cfg.Gin.Mode = viper.GetString("GIN_MODE")

	// Database.
	cfg.Database.Username = viper.GetString("DATABASE_USERNAME")
	cfg.Database.Password = viper.GetString("DATABASE_PASSWORD")
	cfg.Database.Host = viper.GetString("DATABASE_HOST")
	cfg.Database.Name = viper.GetString("DATABASE_NAME")
	cfg.Database.Port = viper.GetUint32("DATABASE_PORT")
	cfg.Database.SslMode = viper.GetString("DATABASE_SSLMODE")
	cfg.Database.Timezone = viper.GetString("DATABASE_TIMEZONE")
	cfg.Database.Engine = viper.GetString("DATABASE_ENGINE")

	// CORS.
	cfg.Cors.Origins = viper.GetStringSlice("CORS_ORIGINS")
}
