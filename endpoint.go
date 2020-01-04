package gohome

// Endpoint represents an enum of the REST API paths
type Endpoint string

func (e Endpoint) String() string {
	return string(e)
}

const (
	Info Endpoint = "setup/eureka_info"
)
