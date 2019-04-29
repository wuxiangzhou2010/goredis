# goredis test

A small redis test project

## docker and redis

- [redis5 default config](https://raw.githubusercontent.com/antirez/redis/5.0/redis.conf)

```sh
docker run -p 6379:6379  -v /root/redis.conf:/usr/local/etc/redis/redis.conf --name my-redis-container -d redis # 使用默认选项启动一个docker实例, 使用定义的配置
ocker run -it --name my-redis-cli --link my-redis-container:redis --rm redis redis-cli -h redis -p 6379 ## 使用另外一个docker 示例连接
```

```redis
lpush comments "hello"
lpush comments "world"
lpush comments "testing"
lpush comments "1234"
```

key point:

- net/http
- gorilla mux
- go-redis
- html/template

reference:

- [Write a Web App in Go](https://www.youtube.com/watch?v=rLK7MSt5flE&list=PLmxT2pVYo5LDMV0epL4z4CUbxvIw6umg_&index=5)
