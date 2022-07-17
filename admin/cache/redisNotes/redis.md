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



#### 三. set操作







##### 











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

#### 1.开启事物

~~~sql
multi
~~~

#### 2. 预处理命令

~~~sql
set k1 v1
set k2 v2
~~~

#### 3 提交事物

~~~sql
exec
~~~

#### 4.取消事物

~~~sql
discard
~~~





















