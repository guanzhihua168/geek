# week8 Homework

## 作业内容
1. 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
2. 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

## 测试机器

mac pro 单机

|  硬件   | 参数  |
|  ----  | ----  |
| CPU  | 2.7 GHz Intel Core i5 |
| RAM  | 8 GB |

## 第一题测试
``` 
  redis-benchmark -t set,get -q -d 10
  SET: 46490.00 requests per second
  GET: 46019.32 requests per second
  
  redis-benchmark -t set,get -q -d 20
  SET: 41562.76 requests per second
  GET: 46926.32 requests per second
  
  redis-benchmark -t set,get -q -d 50
  SET: 48520.13 requests per second
  GET: 47596.38 requests per second
  
  redis-benchmark -t set,get -q -d 100
  SET: 46468.40 requests per second
  GET: 45977.01 requests per second
  
  redis-benchmark -t set,get -q -d 200
  SET: 48123.20 requests per second
  GET: 46641.79 requests per second
  
  redis-benchmark -t set,get -q -d 1000
  SET: 46296.29 requests per second
  GET: 43402.78 requests per second
  
  redis-benchmark -t set,get -q -d 5000
  SET: 43365.13 requests per second
  GET: 41528.24 requests per second
```

## 第二题

平均每个key

```
redis-benchmark -t set -r 10000 -d 10000
redis-cli info memory
used_memory_dataset:102857448

```