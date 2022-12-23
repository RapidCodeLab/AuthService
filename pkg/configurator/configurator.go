package configurator

type Configurator struct{}

func New() *Configurator {
	return &Configurator{}
}

func (c *Configurator) GetHTTPServerListenNetwork() (d string) {
	return
}
func (c *Configurator) GetHTTPServerListenAddr() (d string) {
	return
}
func (c *Configurator) GetGRPCServerListenNetwork() (d string) {
	return
}
func (c *Configurator) GetGRPCServerListenAddr() (d string) {
	return
}
