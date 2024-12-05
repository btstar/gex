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
	AccountRpcConf   zrpc.RpcClientConf
	OrderRpcConf     zrpc.RpcClientConf
	DtmConf          zrpc.RpcClientConf
	PulsarConfig     pulsar.PulsarConfig
	LoggerConfig     logger.Config
	RedisConf        redis.RedisConf
	GormConf         commongorm.GormConf
	SymbolInfo       *define.SymbolInfo `json:",optional"`
	SnowFlakeWorkID  int64
	WsConf           zrpc.RpcClientConf
	Symbol           string
	SymbolEtcdConfig etcd.EtcdConfig
	EtcdRegisterConf etcd.EtcdRegisterConf `json:",optional"`
}
