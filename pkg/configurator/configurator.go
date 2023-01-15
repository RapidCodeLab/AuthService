package configurator

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Configurator struct {
	httpServerListenNetwork string `env:"HTTP_SERVER_LISTEN_NETWORK"`
	httpServerListenAddr    string `env:"HTTP_SERVER_LISTEN_ADDR"`
	grpcServerListenNetwork string `env:"GRPC_SERVER_LISTEN_NETWORK"`
	grpcServerListenAddr    string `env:"GPRC_SERVER_LISTEN_ADDR"`
	grpcUserServiceAddr     string `env:"GRPC_USER_SERVICE_ADDR"`
}

func New() (*Configurator, error) {
	cfg := &Configurator{}
	return cfg, readEnv(cfg)
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

func readEnv(cfg any) error {
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		help, err := cleanenv.GetDescription(cfg, nil)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s | %s", help, err.Error())
	}
	return nil
}
