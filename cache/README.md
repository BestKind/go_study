
1. 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
2. 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

默认 50 并发 100000 请求数
```shell
λ bestkind [~] → redis-benchmark -d 10 -t set,get
====== SET ======
  100000 requests completed in 11.16 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

0.00% <= 1 milliseconds

99.12% <= 10 milliseconds

99.95% <= 47 milliseconds
99.95% <= 1005 milliseconds

100.00% <= 1011 milliseconds
8958.97 requests per second

====== GET ======
  100000 requests completed in 9.08 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

0.00% <= 1 milliseconds

99.23% <= 10 milliseconds

100.00% <= 42 milliseconds
11015.64 requests per second


λ bestkind [~] → redis-benchmark -d 10 -t set,get
====== SET ======
  100000 requests completed in 11.16 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

0.00% <= 1 milliseconds

99.12% <= 10 milliseconds

100.00% <= 35 milliseconds
8958.97 requests per second

====== GET ======
  100000 requests completed in 9.08 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

0.00% <= 1 milliseconds

99.23% <= 10 milliseconds

100.00% <= 42 milliseconds
11015.64 requests per second
```

```shell
λ bestkind [~] → redis-benchmark -d 20 -t set,get
====== SET ======
  100000 requests completed in 12.25 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1

0.00% <= 1 milliseconds

94.19% <= 10 milliseconds

100.00% <= 39 milliseconds
8160.60 requests per second

====== GET ======
  100000 requests completed in 11.16 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1

0.00% <= 1 milliseconds

98.15% <= 10 milliseconds

99.95% <= 31 milliseconds
99.95% <= 1001 milliseconds

100.00% <= 1005 milliseconds
8959.77 requests per second


λ bestkind [~] → redis-benchmark -d 20 -t set,get
====== SET ======
  100000 requests completed in 12.25 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1

0.00% <= 1 milliseconds

94.19% <= 10 milliseconds

100.00% <= 39 milliseconds
8160.60 requests per second

====== GET ======
  100000 requests completed in 11.16 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1

0.00% <= 1 milliseconds

98.15% <= 10 milliseconds

100.00% <= 35 milliseconds
8959.77 requests per second
```

```shell
λ bestkind [~] → redis-benchmark -d 50 -t set,get
====== SET ======
  100000 requests completed in 11.15 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1

0.00% <= 1 milliseconds

96.61% <= 10 milliseconds

100.00% <= 31 milliseconds
8967.81 requests per second

====== GET ======
  100000 requests completed in 11.22 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1

0.00% <= 1 milliseconds

96.32% <= 10 milliseconds

100.00% <= 28 milliseconds
8912.66 requests per second
```

```shell
λ bestkind [~] → redis-benchmark -d 100 -t set,get
====== SET ======
  100000 requests completed in 10.53 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1

0.00% <= 1 milliseconds

98.14% <= 10 milliseconds

100.00% <= 22 milliseconds
9498.48 requests per second

====== GET ======
  100000 requests completed in 9.35 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1

0.00% <= 1 milliseconds

99.30% <= 10 milliseconds

100.00% <= 30 milliseconds
10700.91 requests per second
```

```shell
λ bestkind [~] → redis-benchmark -d 200 -t set,get
====== SET ======
  100000 requests completed in 10.52 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1

0.00% <= 1 milliseconds

98.11% <= 10 milliseconds

100.00% <= 28 milliseconds
9504.80 requests per second

====== GET ======
  100000 requests completed in 9.15 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1

0.00% <= 1 milliseconds

99.41% <= 10 milliseconds

100.00% <= 22 milliseconds
10924.19 requests per second
```

```shell
λ bestkind [~] → redis-benchmark -d 1000 -t set,get
====== SET ======
  100000 requests completed in 12.43 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1

0.00% <= 1 milliseconds

94.50% <= 10 milliseconds

100.00% <= 32 milliseconds
8043.11 requests per second

====== GET ======
  100000 requests completed in 9.73 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1

0.00% <= 1 milliseconds

99.25% <= 10 milliseconds

100.00% <= 24 milliseconds
10277.49 requests per second
```

```shell
λ bestkind [~] → redis-benchmark -d 5000 -t set,get
====== SET ======
  100000 requests completed in 16.56 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1

0.00% <= 1 milliseconds

78.40% <= 10 milliseconds

100.00% <= 31 milliseconds
6039.38 requests per second

====== GET ======
  100000 requests completed in 12.44 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1

0.00% <= 1 milliseconds

94.61% <= 10 milliseconds

100.00% <= 38 milliseconds
8037.94 requests per second
```

测试 10 和 20 字节 value 大小时，一次 set 的时间会有阻塞，一次 get 有阻塞
但是再次测试时间又较为正常了
