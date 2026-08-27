package registry

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdServer struct {
	Host string
	Port int
}

func GetEtcdKv(key string, regServer EtcdServer) ([]byte, error) {

	addr := fmt.Sprintf("%s:%d", regServer.Host, regServer.Port)
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: 5 * time.Second,
		//Username:    "root",
		//Password:    "Password",
	})
	if err != nil {
		fmt.Printf("connect etcd failed: %s", err.Error())
		return nil, err
	}

	resp, err := cli.Get(context.Background(), key)

	if err != nil {
		hlog.Fatalf("etcd kv failed: %s", err.Error())
		return nil, err
	}

	if len(resp.Kvs) == 0 {
		return nil, nil
	}
	return resp.Kvs[0].Value, nil
}
