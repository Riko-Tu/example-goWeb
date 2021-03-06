## redis

### key常用命令

#### 1. 查看库key值

~~~nosql
keys *0      //查看0号数据库中的key值

dbsize    //查看当前库的key数量
~~~

#### 2.判断key存在

~~~nosql
exists k1        //返回1或0
~~~

#### 3.查看key的类型

~~~nosql
type k1    //返回key的数据类型
~~~

#### 4. 新增KV

~~~nosql
set  k2 ww
~~~

#### 5. 删除k

~~~nosql
del k1   
~~~

#### 6.设置有效时间

~~~sql
expire k1 10   //对已存在k1设置10秒的有效时间
~~~

#### 7. 查看key有效时间

~~~sqll
ttl  k1  //查看k1的剩余有效时间；  -2 表示已过期  -1 永远不过期
~~~

#### 8.切换数据库

~~~sql
select 0   //切换0号库
~~~

### 数据类型

#### 一. string操作

![string类型](.\string类型.jpg)

##### 1. 获取值

~~~sql
get  k1     //获取k1的值
~~~

#####  2. 追加

~~~sql
append    k1   adb   //在k1的值后面拼接adb；如果k1不存在，则变成set k1 adb
~~~

##### 3. 获取value长度

~~~sql
strlen k1    //获取k1 值的长度
~~~

##### 4. 判空设值

~~~sql
setnx  k1  www  //当k1不存在时，才能设置成功
~~~

##### 5. 数值value增减

~~~sql
incr  k1   //当k1的值为int型时，使value加一
decr  k1    //使value减一 


incrby   k1   10   //使value 加10
decrby  k1    20   //使value  减20
~~~

##### 6.批量设置和获值

~~~sql
mset k1 v1 k2 v2 k3 v3   
mget  k1 k2 k3
msetnx k1 v2 k3 v3  //当全部key值都为空时，才成功
~~~

##### 7. 截取vaule

~~~sql
getrange k1 0 3   //获取k1下标0到3的字段，包括3
~~~



##### 8. 更新value下标字符

~~~sql
setrange k1  3 vvvv // 3为起始位置,value长度为结束位置
~~~

##### 9. 设置过期时间

~~~sql
setex k1   20 v22   //设置k1的过期时间为20秒
~~~

##### 10.以新换旧

~~~sql
getset k1   vv  //先获取k1的值，再将vv赋值给k1，获取旧值;设置新值
~~~

#### 二. list操作

![list类型](.\list类型.jpg)

##### 1. 新增

~~~sql
lpush   k1  v1 v2  v3     //从左存储的顺序为 v3  v2  v1  
rpush k2  v1 v2 v3     //从右存储的顺序为  v1 v2 v3
~~~

##### 2. 获取值

~~~sqp
lrange k1 0 1     //结果为 v3 v2 
~~~

##### 3.摘取值

~~~sql
lpop k1   //结果为 v3 ，取出后v3就不在k1中

rpop k2    //结果为v1 ，取出后v1就不在k1中
~~~

##### 4. 摘取转存

~~~sql
rpoplpush k1  k2 // 将k1的右边第一个元素，存在到k2的左边第一个元素
~~~

##### 5. 下标取值

~~~sql
lindex k1 0  //获取k1下标为0的元素
~~~

##### 6. 获取列表长度

~~~sql
llen k1   
~~~

##### 7. 指定元素插入

~~~ 
linsert  k1  before 'v1'  'v2'  //在v1前面插入v2

linsert  k1  after 'v1'  'v2'   //在v1后面插入v2

~~~

##### 8. 指定元素删除

~~~sql
lrem k1 2 'ww'      //从k1左边删除ww两次
~~~

##### 9. 值替换

~~~sql
lset k1 0 '22'    //将k1下标为0的值替换为"22"
~~~



#### 三. set无序集合

- 排重的无序集合

##### ![set集合](F:\goSRC\src\turan\example-goWeb\admin\cache\redisNotes\set集合.jpg)

##### 1. 添加值

~~~sql
sadd k1  v1  v2     //设值k1有v1和v2的集合
~~~

##### 2.取值

~~~sql
smembers k1      //取出集合中的所有值
~~~

##### 3.判断值存在

~~~sql
sismember k1 v1        //判断v1是否存在
~~~

##### 4. 返回集合元素个数

~~~sql
scard k1       //返回k1集合中元素的个数
~~~

##### 5.删除指定元素

~~~sql
srem k1 v1      //删除key中的V1值
~~~

##### 6.随机吐值

~~~sql
spop  k1       //吐出的值，不会再存在集合中。k1为空，返回nil
~~~

##### 7.随机取出N个值

~~~sql
srandmember k1  2    //从K1中取值两个值，不会被删除
~~~

##### 8.移动一个值到集合

~~~sql
smove k1 k2   v1  //将k1中的v1值，移动到k2中
~~~

##### 9. 两个集合取交集

~~~sql
sinter k1 k2           
~~~

##### 10.两个集合取并集

~~~sql
sunion k1 k2
~~~

##### 11.两个集合取差集

~~~sql
sdiff  k1 k2
~~~

#### 四.hash表

- 当对象的字段长时，使用hash表

![hash](hash.jpg)



##### 1.添加值

~~~sql
hset user:1 id 1 age 20      //向user：1中存放了字段id和age
hmset user:2 id 2 age 21
~~~

##### 2.获取值

~~~sql
hget user:1 id     //获取user：1中id的值
~~~

##### 3.判断字段存在

~~~sql
hexists user:1 id      //判断user:1中存在id字段，存在返回1，不存在返回0 
~~~

##### 5.获取key的所有字段

~~~sql
hkeys user:1    //返回id 和age
~~~

##### 6.返回key的所有字段值

~~~sql
hvals user：1   //返回1和20
~~~

##### 7.hincrby

~~~sql
hincrby user:1 age 2 //将age字段的值加2
~~~

##### 8.添加k-v

~~~sql
hsetnx user:1 sex 1   //在user：1中添加sex字段值为1，当sex不存在时才成功
~~~

#### 五.zset有序集合

![hash](hash.jpg)

##### 1.添加值

~~~sql
zadd topn  100 java 200 go  //向topn中添加 java和go，评分分别为100和200，排序为从小到大
~~~

##### 2.获取值

~~~sql
zrange topn 0 -1 withscores   //返货topn中所有值和评分
zrange topn 0  1        //返回topn下标为0到1的值
~~~

##### 3.通过评分获取值

~~~sql
zrangebyscore topn 0 300    //获取topn中评分在0到300之间的所有值
~~~

##### 4.降序排序

~~~sql
zrevrangebyscore topn 300 0   //将topn从大到小排序返回,不改变原来topn的排序
~~~

##### 5.增加值的评分

~~~sql
zincrby topn 50 go      //将go的评分增加50
~~~

##### 6.删除集合中的值

~~~ 
zrem topn go   //将topn中的go删除
~~~

##### 7.统计评分之间的个数

~~~sql
zcount topn  0  300  //获取topn中0到300评分中的值个数
~~~

##### 8.返回值的排名

~~~sql
zrank topn go   //返回go在topn的排名
~~~







### 配置文件

 

### 订阅和发布

#### 1. 发布

~~~sql
publish channel hello333
~~~

#### 2.订阅

~~~sql
subscribe channel1
~~~



### 事物

- 介绍：事务就是将多条命令打包成一条命令执行；

  ![事务](事务.jpg)

#### 1.开启事物

- 阶段一：开始组队阶段

~~~sql
multi
~~~

#### 2. 预处理命令

- 组队命令

~~~sql
set k1 v1
set k2 v2
~~~

#### 3 提交事物

- 执行阶段

~~~sql
exec
~~~

#### 4.取消事物

- 在执行之前取消

~~~sql
discard
~~~

#### 5.错误处理

- 如果阶段一出现了错误，执行时会把整个队列取消

![事务错误处理1](事务错误处理1.jpg)



- 如果阶段二出现了错误，执行错误的不成功，执行成功的正常执行

![事务错误处理2](事务错误处理2.jpg)







