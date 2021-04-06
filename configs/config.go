package configs

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Token    Token    `yaml:"token"`
	Logger   Logger   `yaml:"logger"`
}

type Server struct {
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
}

type Database struct {
	Uri string `yaml:"uri"`
}

type Token struct {
	SecretKey string `yaml:"secret_key"`
	PublicKey string `yaml:"public_key"`
	Salt      string `yaml:"salt"`
}

type Logger struct {
	LogLevel string `yaml:"log_level"`
}