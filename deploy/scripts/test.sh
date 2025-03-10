#!/bin/bash
lang='50006: 超过最小精度11
100001: 内部错误
100002: 内部错误
100003: 内部错误
100004: 参数错误
100005: 记录未找到
100006: 重复数据
100007: 内部错误
100009: 内部错误
100010: 内部错误
100011: 内部错误
100012: 验证码错误
200001: 用户不存在
200002: 用户余额不足
200003: token验证失败
200004: token到期
200005: 账户密码验证失败1
500001: 订单未找到
500002: 订单已经成交获取已经取消
500003: 市价单不允许手动取消
500004: 订单簿没有买单
500005: 订单簿没有卖单
500006: 超过币种最小精度'

coin1='coinid: 29
coinname: IKUN
prec: 3'

coin2='coinid: 2
coinname: USDT
prec: 5'
symbol='symbolname: IKUN_USDT
symbolid: 6
basecoinname: IKUN
basecoinid: 29
quotecoinname: USDT
quotecoinid: 2
baseCoinPrec: 3
quoteCoinPrec: 5'

docker exec -it etcd /usr/local/bin/etcdctl put language/zh-CN -- "$lang"
docker exec -it etcd /usr/local/bin/etcdctl put Coin/IKUN -- "$coin1"
docker exec -it etcd /usr/local/bin/etcdctl put Coin/USDT -- "$coin2"
docker exec -it etcd /usr/local/bin/etcdctl put Symbol/IKUN_USDT -- "$symbol"