package config

type BackendConfig interface {
    Type() string
}

func NewBackendConfig(backend string) (bcf BackendConfig) {
    switch backend {
    case "consul":
        bcf = NewConsulBackendConfig()
    case "etcd":
        bcf = NewEtcdBackendConfig()
    default:
        panic("invalid backend, this should never happen!")
    }
    return
}

type ConsulBackendConfig struct {
    Scheme  string   `cli:"{\"name\":\"scheme\"}"`
    Nodes   []string `cli:"{\"name\":\"node\"}"`
    Srv     string   `cli:"{\"name\":\"srv\"}"`
    Cert    string   `cli:"{\"name\":\"client-cert\",\"envvar\":\"CLIENT_CERT\"}"`
    CertKey string   `cli:"{\"name\":\"client-cert-key\",\"envvar\":\"CLIENT_CERT_KEY\"}"`
    CAKeys  string   `cli:"{\"name\":\"client-ca-keys\",\"envvar\":\"CLIENT_CA_KEYS\"}"`
}

func NewConsulBackendConfig() *ConsulBackendConfig {
    return &ConsulBackendConfig{
        Scheme:  "http",
        Nodes:   []string{"127.0.0.1:8500"},
        Srv:     "",
        Cert:    "",
        CertKey: "",
        CAKeys:  "",
    }
}

func (*ConsulBackendConfig) Type() string {
    return "consul"
}

type EtcdBackendConfig struct {
    Nodes   []string `cli:"{\"name\":\"node\",\"envvar\":\"NODES\"}"`
    Srv     string   `cli:"{\"name\":\"srv\"}"`
    Cert    string   `cli:"{\"name\":\"client-cert\",\"envvar\":\"CLIENT_CERT\"}"`
    CertKey string   `cli:"{\"name\":\"client-cert-key\",\"envvar\":\"CLIENT_CERT_KEY\"}"`
    CAKeys  string   `cli:"{\"name\":\"client-ca-keys\",\"envvar\":\"CLIENT_CA_KEYS\"}"`
}

func NewEtcdBackendConfig() *EtcdBackendConfig {
    return &EtcdBackendConfig{
        Nodes:   []string{"127.0.0.1:2379"},
        Srv:     "",
        Cert:    "",
        CertKey: "",
        CAKeys:  "",
    }
}

func (*EtcdBackendConfig) Type() string {
    return "etcd"
}