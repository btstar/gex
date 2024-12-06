package config

import (
	"github.com/luxun9527/gex/common/pkg/etcd"
	logger "github.com/luxun9527/zlog"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	LoggerConfig     logger.Config
	OrderRpcConf     zrpc.RpcClientConf
	MatchRpcConf     zrpc.RpcClientConf
	AccountRpcConf   zrpc.RpcClientConf
	LanguageEtcdConf etcd.EtcdConfig
	SymbolEtcdConfig etcd.EtcdConfig
}
