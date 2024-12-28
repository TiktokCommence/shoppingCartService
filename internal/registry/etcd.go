package registry

import (
	"ShoppingCartService/internal/conf"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func NewRegistrarServer(c *conf.Etcd) *etcd.Registry {
	endpoints := []string{c.Endpoints}
	etcdCfg := clientv3.Config{
		Endpoints:   endpoints,
		Username:    c.Username,
		Password:    c.Password,
		DialTimeout: c.DialTimeout.AsDuration(),
		// DialOptions: []grpc.DialOption{grpc.WithBlock()},
	}

	client, err := clientv3.New(etcdCfg)
	if err != nil {
		panic(err)
	}
	// 创建服务注册 registrar
	registrar := etcd.New(client)
	fmt.Println("registrar successful", registrar)
	return registrar
}
