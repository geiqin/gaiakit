package registry

import (
	"net"
	"strconv"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hashicorp/consul/api"
)

type ConsulServer struct {
	Host string
	Port int
}

func GetConsulKv(key string, regServer ConsulServer) ([]byte, error) {
	cfg := api.DefaultConfig()
	cfg.Address = net.JoinHostPort(
		regServer.Host,
		strconv.Itoa(regServer.Port))
	consulClient, err := api.NewClient(cfg)

	if err != nil {
		hlog.Fatalf("new consul client failed: %s", err.Error())
		return nil, err
	}
	content, _, err := consulClient.KV().Get(key, nil)
	if err != nil {
		hlog.Fatalf("consul kv failed: %s", err.Error())
		return nil, err
	}
	return content.Value, nil
}
