package config

type GlobalConfig struct {
    Host     string `cli:"{\"name\":\"host\",\"usage\":\"\",\"envvar\":\"KUBERNETES_RO_SERVICE_HOST\"}"`
    Port     int    `cli:"{\"name\":\"port\",\"usage\":\"\",\"envvar\":\"KUBERNETES_RO_SERVICE_PORT\"}"`
    LogLevel string `cli:"{\"name\":\"log-level\",\"usage\":\"level which kube2kv should log messages\"}"`
}

func NewGlobalConfig() *GlobalConfig {
    return &GlobalConfig{
        Host: "10.0.36.74",
        Port: 443,
        LogLevel: "",
    }
}
