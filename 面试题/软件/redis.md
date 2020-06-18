# Redis 基本数据结构

1. String

用作计数器

2. Hash

3. Set

用作交集，差集

4. Zset

5. List

## Redis 雪崩

1. Redis 突然宕机，所有 request 全部进入关系型数据库

2. 使用 Redis 集群，达到高可用

3. Redis 持久化，挂了之后，重启，快速导入 Redis

## Redis 击穿

1. 经常查询不存在的数值。如果 Redis 没有数值，直接进入关系型数据库

2. 使用拦截器，拦截非法参数

3. 如果 Redis 和 关系型数据库都没有这个参数。把空对象放入 Redis. 设置较短的过期时间

## 高并发，多线程同时操作 Redis

1. 有可能需要顺序发生，比如，下单，付款，交钱

2. 使用 zookeeper 分布式锁

## Redis 和 RDS 读写

1. 写数据库之后删除缓存

## Redis 线程模型

1. Redis 单线程监听多个 socket. 将 socket 产生的操作放入队列中等待。