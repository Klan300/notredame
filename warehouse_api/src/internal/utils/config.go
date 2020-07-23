package utils

import (
    "io"
    "log"

    "github.com/spf13/viper"
)

type ConfigDefinition struct {
    Source struct {
        Host     string
        Database string
        Username string
        Password string
    }

    Target struct {
        Host string
    }

    Authen AuthenDefinition
    Logging LoggingDefinition
}

type AuthenDefinition struct {
    Usernames []string
    Secret      string
    Expire      string
}

type LoggingDefinition struct {
    Level     string
    Stdout    bool
    Dirname   string
}

var Config *ConfigDefinition = new(ConfigDefinition)

func LoadConfig(filePath string) {

    viper.AutomaticEnv()
    viper.SetConfigFile(filePath)

    err := viper.ReadInConfig()
    if err != nil {
        log.Panicf("[config.go] %v\n", err)
    }

    viper.Unmarshal(&Config)
    LoadLoggingConfig()
}

func (authen *AuthenDefinition) Exists(lookup string) bool {

    for _, username := range authen.Usernames {
        if username == lookup {
            return true
        }
    }
    return false
}

func (logging *LoggingDefinition) Outputs() io.Writer {
    return log.Writer()
}
