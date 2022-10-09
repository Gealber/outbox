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

	Cockroach struct {
		DSN string
	}

	Cors struct {
		Origins []string
	}

	GCP struct {
		ProjectID string
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

	// Cockroach.
	cfg.Cockroach.DSN = viper.GetString("COCKROACH_DSN")

	// CORS.
	cfg.Cors.Origins = viper.GetStringSlice("CORS_ORIGINS")

	// GCP.
	cfg.GCP.ProjectID = viper.GetString("GCP_PROJECT_ID")
}
