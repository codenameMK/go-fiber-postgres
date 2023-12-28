package providers


import (
	"flag"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type JwtConfig struct {
	Secret string `yaml:"secret" envconfig:"JWT_SECRET_KEY"`
}
type AppConfig struct {
	JwtConfig         JwtConfig         `yaml:"jwt"`
}

var (
	config     *AppConfig
	configPath string
)

func loadConfig() AppConfig {
	f, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var cfg AppConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	err = envconfig.Process("", &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

func GetConfig(path string) (AppConfig, error) {
	if config == nil {
		flag.StringVar(&configPath, "config", path, "Path to config file")
		flag.Parse()
		cnf := loadConfig()
		config = &cnf
	}
	return *config, nil
}