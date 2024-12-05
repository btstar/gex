package config

import (
	"github.com/luxun9527/gex/common/pkg/etcd"
	commongorm "github.com/luxun9527/gex/common/pkg/gorm"
	"github.com/luxun9527/gex/common/pkg/pulsar"
	"github.com/luxun9527/gex/common/proto/define"
	logger "github.com/luxun9527/zaplog"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	PulsarConfig     pulsar.PulsarConfig
	LoggerConfig     logger.Config
	GormConf         commongorm.GormConf
	WsConf           zrpc.RpcClientConf
	Symbol           string
	OrderRpcConf     zrpc.RpcClientConf
	RedisConf        redis.RedisConf
	SymbolEtcdConfig etcd.EtcdConfig
	SymbolInfo       *define.SymbolInfo    `json:",optional"`
	EtcdRegisterConf etcd.EtcdRegisterConf `json:",optional"`
}
