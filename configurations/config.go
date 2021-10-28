package configurations

// parse config from file or from flags

type ClientConfig struct {
	ClientData string
	Name       string
	Age        uint
}

type ServerConfig struct {
	Port string
	Path string
}

func NewClientConfig() *ClientConfig {
	return &ClientConfig{
		ClientData: "",
		Name:       "",
		Age:        0,
	}
}

func NewServerConfig() *ClientConfig {
	return &ClientConfig{
		ClientData: "",
		Name:       "",
		Age:        0,
	}
}
