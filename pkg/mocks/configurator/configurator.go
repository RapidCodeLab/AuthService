package mockconfigurator

type Configurator struct {
	httpServerListenNetwork string
	httpServerListenAddr    string
	grpcServerListenNetwork string
	grpcServerListenAddr    string
	grpcUserServiceAddr     string
}

func New() *Configurator {
	return &Configurator{
		httpServerListenNetwork: "tcp4",
		httpServerListenAddr:    ":8080",
		grpcServerListenNetwork: "tcp4",
		grpcServerListenAddr:    ":9090",
		grpcUserServiceAddr:     ":8070",
	}
}

func (c *Configurator) GetHTTPServerListenNetwork() (d string) {
	return c.httpServerListenNetwork
}
func (c *Configurator) GetHTTPServerListenAddr() (d string) {
	return c.httpServerListenAddr
}
func (c *Configurator) GetGRPCServerListenNetwork() (d string) {
	return c.grpcServerListenNetwork
}
func (c *Configurator) GetGRPCServerListenAddr() (d string) {
	return c.grpcServerListenAddr
}

func (c *Configurator) GetGRPCUserServiceAddr() (d string) {
	return c.grpcUserServiceAddr
}
