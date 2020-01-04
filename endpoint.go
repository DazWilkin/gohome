package gohome

type Endpoint string

func (e Endpoint) String() string {
	return string(e)
}

const (
	Info Endpoint = "setup/eureka_info"
)
