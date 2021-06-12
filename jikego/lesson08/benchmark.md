

##### 基本命令：redis-benchmark   [option]   [option value]
- -h host
- -p port
- -s socket  指定连接 socket文件
- -c 并发连接数 默认50
- -n 指定请求数 默认10000
- -d size 以字节形式指定SET/SET值的数据大小 默认2
- -k keep alive 1=keep alive      0=reconnect  默认1
- -r keyspacelen SET/GET/INCR   使用随机key，SADD使用随机值
- -q quit 强制退出Redis。仅显示query/sec值
- --csv 以CSV格式输出
- -l loop 生成循环，永久执行测试
- -t 仅运行以逗号分隔的测试命令列表
- -I Idel model Idle模式。仅打开N个idle连接并等待

##### 示例
> 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
```text
redis-benchmark -t set,get -d 10 -q
    SET: 86132.64 requests per second
    GET: 94607.38 requests per second
redis-benchmark -t set,get -d 20 -q
    SET: 81168.83 requests per second
    GET: 95328.88 requests per second
redis-benchmark -t set,get -d 50 -q
    SET: 67567.57 requests per second
    GET: 84317.03 requests per second
redis-benchmark -t set,get -d 100 -q
    SET: 52854.12 requests per second
    GET: 62227.75 requests per second
redis-benchmark -t set,get -d 200 -q
    SET: 37230.08 requests per second
    GET: 41580.04 requests per second
redis-benchmark -t set,get -d 1000 -q
    SET: 11007.15 requests per second
    GET: 11314.78 requests per second
redis-benchmark -t set,get -d 5000 -q
    SET: 2311.92 requests per second
    GET: 2322.72 requests per second
```

> 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间
- 批量插入数据 `cat data.txt | redis-cli --pipe`
```text

value            redis:key(12)+value
10:
1944184-933104/10000 101
20:
1944008-932896/10000 101
50:
2503912-932704/10000 157
100:
2.75M-932512/10000 195
200:
4089432-932096/10000 315
500:
4089432-932304/10000 611
1000:
12163144-931696/10000 1123
5000:
51995912-931504/10000 5106

```
