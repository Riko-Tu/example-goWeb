### 数据查询语言 (DQL)

#### 基础查询

##### 1、查询列表：

- 表中的字段、常量值、表达式、函数

~~~sql
SELECT 查询列表 FORM 表名
~~~

##### 2、常量值:

~~~sql
SELECT 100			字段与值都是常量值
SELECT "JOHON"
SELECT 100*20     //字段为100*20  ；值为结果
~~~

##### 4、表达式:

~~~sql
SELECT VERSION()   //字段为 VERSION()；值为结果
~~~

##### 5、 起别名：

- 别名：出现冲突可以用引号

~~~sql
方式一：            
SELECT 原字段名  AS 新别名 FROM 表名

方式二：
SELECT 字段名 新别名 FROM 表名
~~~

##### 6、字段去重：

~~~sql
SELECT DISTINCT column FROM TABLE
~~~

##### 7、 字符求和：   

~~~sql
SELECT 100+19;             ->  119     int + int ==int
select '123'+90            ->  213		int(string) + int = int
SELECT  "JOHN"+100;        ->  100   //字符串无法转换int， 变成0 + 100 =100
SELECT  null+0；           ->  null  //null拼接任何值都为NULL
~~~

##### 8、字段拼接： concat()

~~~sql
SELECT CONCAT(column 1，column 2) AS column 3  FROM TABLE    //将column 1，column 2的值合并
~~~

##### 9、判断字段NULL

~~~sql
SELECT IFNULL(字段1，返回值)  FROM TABLE    //如果字段1为null就返回指定的值，
~~~

#### 条件查询: 

##### 1、语法

~~~sql
SELECT 查询列表 FROM TABLE WHERE 筛选条件;
~~~

##### 2、 运算符筛选：

-  <   >   =   <>   <=   >=

~~~sql
SELECT 查询列表 FROM TABLE WHERE comlum > 100;
~~~

##### 3、逻辑筛选：

- AND   OR    NOT

~~~sql
SELECT 查询列表 FROM TABLE WHERE conlum1 = 100 AND conlum <> 101;
SELECT 查询列表 FROM TABLE WHERE NOT(conlum1 <= 100) OR conlum > 101;
~~~

##### 4、 模糊筛选：

- LIKE  、  BETWEEN AND  、  IN  、  IS NULL 

~~~txt
LIKE:通常配合通配符一起使用
	通配符：
		% 任意多个字符，包含0个字符
		_ 任意单个字符
	tip：
		转义通配符 \_  转义_为普通字符
		
BETWEEN AND : 包含临界值；值必须从小到大；值必须为同类型的数值；

IN ：包含某个值；值必须与字段值兼容;
~~~

~~~sql
SELECT 查询列表 FROM TABLE WHERE conlum LIKE '_%a\_';
SELECT 查询列表 FROM TABLE WHERE conlum BETWEEN 100 AND 200;
SELECT 查询列表 FROM TABLE WHERE conlum IN ('值1','值2')；
SELECT 查询列表 FROM TABLE WHERE conlum IS NULL ;
SELECT 查询列表 FROM TABLE WHERE conlum IS NOT NULL ;
~~~

#### 排序查询

##### 1、语法：

~~~txt
排序规则：ASC升序    DESC降序
排序列表：
	1，查询列表
	2，别名
	3，函数
	4，表达式
~~~

~~~sql
SELECT 查询列表 FROM TABLE WHERE 筛选条件 ORDER BY 排序列表 排序规则
~~~

##### 2、排序规则

~~~sql
SELECT 查询列表 FROM TABLE WHERE 筛选条件 ORDER BY 排序列表 ASC
SELECT 查询列表 FROM TABLE WHERE 筛选条件 ORDER BY 排序列表 DESC
~~~

##### 3、排序列表

~~~sql
函数  ： SELECT 查询列表 FROM TABLE WHERE 筛选条件 ORDER BY LENGTH(conlum1) DESC
表达式： SELECT 查询列表 FROM TABLE WHERE 筛选条件 ORDER BY conlum1*12 DESC
别名  ： SELECT 查询列表 AS conlum2 FROM TABLE WHERE 筛选条件 ORDER BY conlum2 DESC
~~~

##### 4、多字段排序

~~~sql
SELECT 查询列表 FROM TABLE WHERE 筛选条件 ORDER BY conlumn 1 ASC, conlumn 2 DESC;
~~~

#### 函数查询

##### 单行函数

- 特点：传入一个值返回一个值;传入多个值返回多个值；

- 单行函数包括：字符函数、数学函数、日期函数、流程控制函数

###### 1、字符函数

~~~sql
 1.length  获取参数值的字节个数 :utf8字符集：中文：三个字节;字母、数字：一个字节
	SELECT  LENGTH(conlum) FROM TABLE WHERE conlum = 3
	
2.concat 拼接字符串
	SELECT CONCAT(conlum1,'_',conlum2) FROM TABLE WHERE conlum3 = %
	
3.UPPER 大写  LOWER 小写	
	SELECT CONCAT(UPPER(conlum1),LOWER(conlum)) FROM TABLE
	
4.SUBSTR 分割	:MYSQL下标从一开始
	SELECT SUBSTR(conlum,3,4) FROM TABLE //从conlum字段的下标为3和下标为4的值截取
	
5.INSTR 返回子串第一次出现的索引，找不到返回0
	SELECT INSTR('123132','13') FROM TABLE  //返回4
	
6.TRIM 去除指定前后字符；默认去空格
	select trim('   123   ') ;   //去除空格返回123 
	SELECT TRIM('a' FROM 'aaaaa123aaaaa');  //去除aaa返回123
	
7.LPAD 左填充
	SELECT LPAD('123',10,'A')  

8.RPAD 右填充
	SELECT RPAD('123',10,'A')

9.REPLACE
	SELECT REPLACE("111111222"，'2','1');   //将2替换成1
~~~

###### 2、数学函数

~~~sql
1.ROUND 四舍五入
	SELECT ROUND(1.2)   //返回1
	SELECT ROUND(1.27,1) //保留小数后一位，返回1.3
	
2.ceil 向上取整，
	SELECT CEIL(1.02)   //返回2
	
3.Floor 向下取整
	SELECT FLOOR(1.02)	//返回1
	
4.truncate 截断
	SELECT TRUNCATE(1.43,1)  //保留小数点后一位
    
5.MOD 取余数
	SELECT MOD(10,3)  //返回1
~~~

###### 3、日期函数

~~~sql
1.NOW 返回当前系统日期+时间
	SELECT NOW();

2.CURDATE()  返回当前系统日期
	SELECT CURDATE()
	
3.CURTIME()  返回当前系统时间
	SELECT CURTIME()
	
4.STR_TO_DATE()  将前端输入的字符串转成数据库可计算的日期格式
	SELECT STR_TO_DATE('2-3 1999','%c-%d %Y');    //返回1999-2-03
	
5.DATE_FORMAT()  将可计算的日期格式转换成字符串
	SELECT DATE_FORMAT('1999-08-23','%y年-%m月-%d日')  //返回99年08月23日
~~~

| 格式符 | 功能                |
| ------ | ------------------- |
| %Y     | 四位年份            |
| %y     | 两位年份            |
| %m     | 月份（01,02...）    |
| %c     | 月份（1,2...）      |
| %d     | 日（01,02）         |
| %H     | 小时（24小时）      |
| %h     | 小时（12小时）      |
| %i     | 分钟（00,01,02...） |
| %s     | 秒钟（00,01...）    |

###### 4、流程控制函数

~~~sql
1，IF  类似IF ELSE
	SELECT IF(10<5,'大','小');

2, CASE  等值判断  类似switch case
	SELECT conlum1 
			CASE conlum1 
				WHEN 10 THEN conlum1*1.1 
				WHEN 20 THEN conlum1*1.2 
				ELSE conlum1
			END AS NEWconlum 
	FROM TABLE;
	
3，CASE   区间判断   类似多重IF
	SELECT conlum1 
			CASE 
				WHEN conlum1 >100 THEN  'a'
				WHEN conlum1 >50  THEN  'b'
				ELSE 'C'
			END
	FROM TABLE;
    
~~~

##### 分组函数

- 特点：

  传入一组值返回一个值；只能与GROUP BY 一起用，不能与查询普通字段连用

- 基本分组函数

~~~SQL
1,SUM 求和
	SELECT SUM(conlum1) FROM TABLE
	
2,AVG 求平均
	SELECT AVG(conlum1) FROM TABLE
	
3,MIN 求最小
	SELECT MIN(conlum1) FROM TABLE
	
4,MAX 求最大
	SELECT MAX(conlum1) FROM TABLE
	
5,COUNT 求总条数
	SELECT COUNT(conlum1) FROM TABLE
	SELECT COUNT(*) FROM TABLE    
	
		
6，搭配 DISTINCT
	SELECT COUNT(DISTINCT conlum1) FROM TABLE
		
注意事项：
	支持哪些类型：
		1， SUM 与 AVG  只支持数值型
    	2， MIN MAX COUNT  支持任意类型
    是否忽略null:
    	1，所有分组函数都忽略NULL值
~~~

- 分组查询

~~~sql
1.分组前筛选:筛选原始字段
	SELECT 分组函数,列(要求出现在group by的后面) FROM TABLE 
		WHERE 条件筛选 GROUP BY 分组列表
		
例：	SELECT AVG(conlum1) ,conlum2 FROM TABLE 
		WHERE conlum2 BETWEEN 100 AND 200
        GROUP BY conlum2

		
2.分组后筛选：筛选分组后的结果集
	SELECT 分组函数,分组列 FROM TABLE 
		GROUP BY 分组列表 HAVING 条件筛选
		
例：	SELECT SUM(conlum1) AS 和 ，conlum2 FROM TABLE 
			GROUP BY conlum2 
			HAVING 和 > 100；
		
3.多个分组
	SELECT SUM(conlum1) AS 和 ，conlum2 FROM TABLE 
			GROUP BY conlum2,conlum1
            
4.排序
	SELECT SUM(conlum1) AS 和 ，conlum2 FROM TABLE 
			GROUP BY conlum2,conlum1 ORDER BY conlum2
~~~

#### 连接查询

- 内连接：   

  ~~~txt
  应用：
  	用于查询两个表之间有关联的记录
  ~~~

  

  ~~~txt
  包括：
  	等值连接
  	非等值连接
  	自连接
  ~~~

  ~~~SQL
  语法：
  	SELECT 查询列表 FROM TABLE1 别名1 
  		INNER JOIN TABLE2 别名2
  		ON 连接条件
  ~~~

  - 等值连接：

    条件使用等号

  ~~~sql
  1.  条件筛选
  	SELECT COUNT(*) , TI.conlum, T2.conlum FROM TABLE1 T1 
  	INNER JOIN TABLE2 T2 ON T1.ID = T2.ID
  	WHERE conlum1 = 20
  
  2.  分组+条件筛选+排序
  	SELECT COUNT(*) ,conlum1 FROM TABLE1 T1 
  	INNER JOIN TABLE2 T2 ON T1.ID = T2.ID
  	GROUP BY conlum1 HAVING COUNT(*)>3
  	ORDER BY COUNT(*) DESC
  ~~~

  - 非等值连接:

  ~~~sql
  1.查询T1的字段与T2字段相等的记录：区间
  	SELECT COUNT(*) FROM TABLE1 T1 
  	INNER JOIN TABLE2 T2 
  	ON T1.ID BETWEEN T2.LOWEST_ID AND T2.LOWEST_ID
      
  ~~~

  - 自连接:似乎与等值连接类似

  ~~~sql
  SELECT COUNT(*) , TI.conlum, T2.conlum FROM TABLE1 T1 
  	INNER JOIN TABLE2 T2 ON T1.ID = T2.ID
  	WHERE conlum1 = 20
  ~~~

  

  

- 外连接：	

  ~~~txt
  应用：
  	用于查询一个表中有，另一个表中没有的记录
  ~~~

  

  ~~~txt
  特点：
  	1.外连接的查询结果为主表中的所有记录，
  		如果从表中没有和它匹配的，则显示为NULL；
  		如果从表中有和它匹配的，则显示配的值；
  		外连接的查询结果=内连接结果+主表有而从表没有的记录
      2.左外连接，left join 左边的是主表
        右外连接，right join 右边的是主表
      3.左外和右外交换两个表的顺序，可以实现同样的效果
  ~~~

  ~~~txt
  包括：
  	左外连接
  	右外连接
  	全外连接
  ~~~

  - 左外连接

  ~~~sql
  1.左外：left 左边的是主表
  	SELECT T1.ID ,T2.ID FROM TABLE T1 
  	LEFT OUTER JOIN TABLE T2 
  	ON T1.ID =T2.ID    //此时T1表的字段会全部显示，T2,T1ID相等的值匹配，其余展示为空
  	WHERE T2.ID IS NULL				//查询不在主表中的记录
  ~~~

  - 右外连接

  ~~~sql
  1.右外：right 右边的是主表
  	SELECT TI.ID ,T2.ID FORM TABLE T1 
  	RIGHT OUTER JOIN TABLE T2
  	ON T1.ID =T2.ID
  	WHERE T1.ID IS NULL			//查询不在主表中的记录
  
  ~~~

  - 全外连接=内连接结果+左外结果+右外结果

  ~~~sql
  1.全外：没有主从表之分
  	SELECT T1.ID ,T2.ID FROM TABLE T1
  	FULL OUTER JOIN TABLE T2
  	ON T1.ID=T2.ID
  	WHERE T1.ID IS NULL AND T2.ID IS NULL  //查询两个表中相互没有的记录
  ~~~

  

- 交叉连接:笛卡尔积

  ~~~sql
   SELECT COUNT(*) FROM TABLE
   CROSS JOIN  TABLE 
  ~~~

#### 子查询

~~~sql
按子查询出现的位置分类：
	1. SELECT 后面：
			标量子查询
				
	2. FROM  后面：
			表子查询
				
	3. WHERE 或 HAVING 后面：
			标量子查询（单行）必须使用单行操作符： < > <> <= >=
			列子查询(多行)必须使用多行操作符： IN 、 NOT IN 
			行子查询
				
	4. exists  后面：
			表子查询

按结果集的行列数不同：
	标量子查询（结果集只有一行一列）
	列子查询（结果集只有一列多行）
	行子查询（结果集有多行多列）
	表子查询（结果集一般为多行多列）
~~~



- WHERE 或 having后面

  ~~~sql
  标量子查询（单行）：可放 where,having; 子查询返回单行记录
  	1. WHERE
  		SELECT COUNT(*) FROM TABLE 
  		WHERE ID = (SELECT ID FROM TABLE WHERE NAME = 'xx') 
  		AND SAL > (SELECT COUNT(*) FROM TABLE)
  	
  	2. HAVING 
  		SELECT COUNT(*) FROM TABLE WHERE ID = 2
  		GROUP BY dep 
  		HAVING ID = (SELECT ID FROM TABLE WHERE NAME='xx')
  		
  列子查询（多行）： 可放 where
  	1. WHERE 
  		SELECT COUNT(*) FROM TABLE 
  		WHERE NAME IN (SELECT NAME FORM TABLE WHERE ID > 10)   
  		
  ~~~
  
- SELECT后面

  ~~~sql
  标量子查询（单行）：  C
  	1. SELECT  查询每个部门的员工个数
  		SELECT T2.* (SELECT COUNT(*) FROM TABLE T1
                       WHERE T1.ID=T2.ID)
          FROM TABLE T2
  
  ~~~

- FROM后面

  ~~~sql
  表子查询： 子查询充当一张表，该表必须起别名
  	1.FROM 
  		SELECT T2.ID FROM 
  		(SELECT * FROM TABLE GROUP BY ID ) T2
  		WHERE T2.ID = 1
  ~~~

- EXISTS后面

  ~~~sql
  表子查询：  返回 1 或 0 
  	1.exists
  		SELECT EXISTS(SELECT ID FROM TABLE)
  ~~~



#### 分页查询

- 语法

  ~~~sql
  SELECT 查询列表 FROM TABLE 
  [连接查询] [条件查询] [分组查询] [分组后筛选] [排序查询]
  LIMIT offset , size
  
  offset： 起始索引（第一条索引为0）
  size ： 要显示的条目个数
  ~~~

- 查询

  ~~~sql
  显示第11条到第25条
  	SELECT * FROM TABLE LIMIT 10,15;
  
  ~~~


#### 联合查询

- 特点

  ~~~sql
  1.查询结果展示语句1的查询列表
  2.语句1与语句2 的查询列表数必须一致
  3.UNION具有去重效果
  4.union all 不会去重
  ~~~

- 语法

  ~~~sql
  查询语句1 UNION 查询语句2 UNION ...
  ~~~

- 案例

  ~~~sql
  应用场景：需要查询多个表，并且这些表的字段信息一致时
  SELECT id,name,sex FROM TABLE1 WHERE ID =30 
  UNION
  SELECT t_id,t_name,t_sex FROM TABLE2 WHERE NAME LIKE '%A'
  ~~~

  

## 数据操作语言（DML）

### 增加

- 语法

  ~~~sql
  语法一：
  	INSERT INTO TABLE(colnum,...) VALUES(values,...)
  	
  语法二：
  	INSERT INTO TABLE SET colunm1= value1,colunm2 =value2
  
  ~~~

- 特点

  ~~~sql
  1.插入的字段的顺序可以调换
  2.未插入的字段值默认为NULL
  3，可以省略列名，但数值为全表插入
  ~~~

- 案例

  ~~~sql
  语法一：
     a.指定列，插入值	
  	 INSERT INTO user(id,name,sex,date) VALUES(1,'蓝若兮','女',null)
     b.全表插入多行
       INSERT INTO USER VALUES(1,'蓝若兮','女',null),
       						(2,'蓝若兮','女',null),
   						    (3,'蓝若兮','女',null),
    				        	(4,'蓝若兮','女',null);
     C. 使用子查询插入
     	  INSERT INTO TABLE(colunm) SELECT ID FROM TABLE WHERE ID<30
  ~~~

  

### 修改

- 语法

  ~~~sql
  1.修改单表记录
  	UPDATE table set colunm =value where 筛选条件
  	
  2.修改多表记录
  	UPDATE TABLE1 别名 type join TABLE 别名
      ON 连接条件 SET 列=值 ,... WHERE 筛选条件
  ~~~

- 案例

  ~~~sql
  1.单表
  	UPDATE TABLE SET id =2 ，name = '若兮' WHERE NAME LIKE '%A'
  
  2.多表
  
  	
  ~~~

  



### 删除

- 语法

  ~~~sql
  delete
  1.单表
  	delete from table where 筛选条件
  	
  2.多表
  	delete T1，T2 from TABLE T1
  	TYPE JOIN TABLE T2 ON 连接条件 WHERE 筛选条件
  	
  truncate
  1.单表
  	truncate table name；
  ~~~

- 特点

  ~~~sq
  1. truncate删除表， 可删除自增长记录，但没有返回值
  2. delete 删除表，无法删除自增长记录，但有返回值
  ~~~

  

## 数据定义语言（DDL）

~~~sql
创建： CREATE
修改： ALTER
删除： DROP
~~~

### 库的管理

#### 创建

- 语法

  ~~~sql
  create database [if not exists] 库名；
  ~~~

- 案例

  ~~~sql
  创建库名 books
  	create database  if not exists books;
  ~~~

#### 修改

- 语法

  ~~~sql
  1.更改字符集
  	ALTER DATABASE BOOKS CHARACTER SET GBK;
  ~~~

#### 删除

- 语法

  ~~~sql
  DROP DATABASE IF EXISTS BOOKS;
  ~~~

  

### 表的管理

#### 创建 

- 语法

  ~~~sql
  CREATE TABLE BOOK(
  	字段 字段值类型【长度，约束】,
      字段 字段值类型【长度，约束】,
      字段 字段值类型【长度，约束】
  )
  ~~~

- 案例

  ~~~sql
  CREATE TABLE processor(
  id INT UNSIGNED NOT NULL AUTO_INCREMENT, 
  NAME VARCHAR(30) NOT NULL,
  `describe` VARCHAR(30) NULL,
  productline VARCHAR(20) NOT NULL,
  applicationtype VARCHAR(20) NOT NULL,
  testtemplate VARCHAR(20) NOT NULL,
  state VARCHAR(200) NULL,
  creator VARCHAR(20) NOT NULL,
  INDEX(productline),
  PRIMARY KEY(id)
  );
  ~~~
  
  

#### 修改

- 语法

  ~~~sql
  alter table 表名 add|change|modify|drop column [列名]
  ~~~

  

- 修改列名

  ~~~sql
  alter table book CHANGE COLUMN [旧列名] [新列名] [新列明类型]
  ~~~

- 修改列的类型，长度，约束

  ~~~sql
  alter table book modify column 列名 列类型
  ~~~

- 修改表名

  ~~~SQL
  alter table book rename to 新表名
  ~~~

- 添加列

  ~~~SQL
  alter table book add column 列名 列的类型
  ~~~

- 删除列

  ~~~sql
  alter table book drop column 列名
  ~~~

  

#### 删除

- 删除表

  ~~~sql
  drop table if exists book 
  ~~~

  

#### 复制

- 只复制表的结构

  ~~~sql
  create table 新表 like 旧表
  ~~~

  

- 复制表结构+全部数据

  ~~~sql
  create table 新表 select * from 旧表
  ~~~

  

- 复制表的部分数据

  ~~~sql
  create table 新表 select * from 旧表 where id <10
  ~~~

- 复制部分结构

  ~~~sql
  只复制id和name的结构
  create table 新表 select id，name from 旧表 where 0 ；
  ~~~

  

### 数据类型

~~~txt
数值型：
		整型
		小数：
				定点数
				浮点数
字符型：
		较短文本：char 、varchar
		较长文本：text 、blob（较长二进制数据）
日期型：				
		
~~~

#### 整型

- 特点

  ~~~TXT
  1.整数类型之间只有存储长度的区别
  2.如果插入的数值超过范围，插入为临界值
  3.不设置长度，有默认长度；
  	a.长度代表该字段的存储空间；
  	b.每个该字段的值都会去申请该字段长度的存储空间;
  ~~~

  

| 整数类型    | 取值范围（字节） |
| ----------- | ---------------- |
| tinyint     | 1                |
| smallint    | 2                |
| mediumint   | 3                |
| int/integer | 4                |
| bigint      | 8                |



#### 小数

- 特点

  ~~~txt
  1.float与double不设置精度时，精度为输入精度
  2.decimale需要给定精度：
  		M:整数部分+小数部分
  		D:小数部分
  3.小数部分过长会四舍五入，小数部分果断会用0填充		
  ~~~

| 浮点型        | 取值范围（占用空间）     |
| ------------- | ------------------------ |
| float         | 4                        |
| double        | 8                        |
| **定点型**    | **取值范围（占用空间）** |
| dec(m,d)      | M+2（与double一致）      |
| decimale(m,d) | M+2（与double一致）      |

- 案例

  ~~~sql
  create table user (
  	id float(5,2),
      name double,
      name_id decimal(5,2)
  )
  ~~~



#### 字符型

- 特点



| 较短文本  | 取值范围   | 特点       | 空间耗费 | 效率 |
| --------- | ---------- | ---------- | -------- | ---- |
| char      | 0~255      | 固定长度   | 比较耗费 | 高   |
| varchar   | 0~6355     | 可变长度   | 比较节省 | 低   |
| enum      | 定义的数值 | 固定数值   | 比较节省 | 高   |
| set       |            |            |          |      |
| binary    |            | 固定二进制 |          |      |
| varbinary |            | 可变二进制 |          |      |

- 案例

  ~~~sql
  create table User(
  	sex enum('1','b'),
      name set('a,b,c,d')
  )
  
  insert into user(sex,name) values('1','a')
  insert into user(sex,name) values('B','a,b,c,d')
  
  ~~~



#### 日期型

- 特点

  ~~~txt
  timestamp（时间戳）受时区影响，空间小，使用多
  datatime 不受时区影响，空间大
  ~~~

| 类型      | 字节 | 最大值              |
| --------- | ---- | ------------------- |
| date      | 4    | 9999-12-31          |
| datetime  | 8    | 9999-12-31 23:59:59 |
| timestamp | 4    | 2038年              |
| time      | 3    | 838:59:59           |
| year      | 1    | 2155                |



### 约束

~~~sql
6种约束：
	a. NOT NULL:非空
		姓名，学号等
	b. DEFAULT : 默认值
		性别
	c. primary key: 主键 
		学号，员工编号
	d. unique：唯一
	
	e. check： 检查约束  (mysql中不支持)
	
	f. FOREIGN KEY:  外键  (mysql中不支持)
	
	
分类：
	列级约束
		支持： primary key,not null ,default,unique
	表级约束
		非空不能用，其他都可以
~~~

- 语法

  ~~~sql
  create table 表名（
  	字段名 字段类型 列级约束，
  	字段名 字段类型，
  	表级约束
  ）
  ~~~

- 案例

  ~~~sql
   create table stuinfo(    
      id int primary key, #主键   
      name varchar(20)  not null, #非空   
      gender char(1) check(gender='男' or gender='女')， #检查   
      seat int unique, #唯一	
      age int default 18, #默认约束 
      majorID int  REFERENCES major(id),# 外键
    )
                        
   create table major(    
       id int primary key  #主键
   )
  ~~~

#### 修改表时添加约束

~~~sql
1.添加非空约束
alter table stuinfo modify column uname varchar(20) not null
2.添加默认约束
alter table stuinfo modify column age int default 18
3.添加主键
	a. 列级约束
		alter table stuinfo modify column id int primary key
	b. 表级约束
		alter table stuinfo add primary key(id)
4.添加唯一
	a. 列级
		alter table stuinfo modify column seat int unique
	b. 表级
    	alter table stuinfo add unique(seat)
5.添加外键
	alter table stuinfo add foreign key(userId) references user(ID)
~~~





### 标识列(自增长列)

- 含义

  系统自动提供一组默认的序列值

---



### 事务控制语言（TCL）

### 支持事务的数据引擎

- 查看当前的数据引擎

  ~~~sql
  
  ~~~

  





#### 事务的特性（ACID）：

- 原子性：
- 一致性：
- 隔离性：
- 持久性：

#### 事务创建：

### 事务的并发问题



#### 事务的隔离级别

| 隔离级别 | 脏读   | 不可重复读 | 幻读   |
| -------- | ------ | ---------- | ------ |
| 读未提交 | 可能   | 可能       | 可能   |
| 读提交   | 不可能 | 可能       | 可能   |
| 可重复读 | 不可能 | 不可能     | 可能   |
| 串行化   | 不可能 | 不可能     | 不可能 |



- 查看当前的隔离级别

  ~~~sql
  select @@tx_isolation
  ~~~

  

### 视图

#### 1.创建

- 视图：类似封装了sql语句

~~~sql
//语法
 create view [视图名]  as  [查询语句]
~~~

#### 2. 使用

~~~sql
select * from [视图名] where [条件筛选]
~~~

#### 3.修改

~~~sql
alter view [视图名]  as [查询语句]
~~~

#### 4.删除

~~~sql
drop view [视图名]
~~~

#### 5. 查看

~~~sql
show create view [视图名]
~~~

### 系统变量

全局变量 global	:	  针对所有回话连接有效，服务器重启后，变量重置

回话变量 session  : 针对单个连接

#### 1. 查看变量

~~~sql
show global variables；
show session variables;
~~~

####  2.模糊查找

~~~sql
show global variables like 'admin%';
show session variables like 'admin%';
~~~

#### 3.名称查找

~~~sql
select @@global.admin_tls_version ;
select @@admin_port ;
~~~

#### 4.变量赋值

~~~sql
set @@global.admin_tls_version = 0 ;
select @@admin_port =0 ;
~~~

### 自定义变量

- 用户变量 ： 针对当前会话有效，类似系统会话（session）变量

##### 1. 声明 和 更新变量

- 方式一

~~~sql
set @name = 123;     //用户变量
~~~

- 方式二

~~~sql
 // 将学生表的数量复制给studentNumb
select count(id) into @studentNumb from student		 //用户变量
~~~

##### 2. 使用变量

~~~sql
select @name;    //用户变量  
~~~

- 局部变量： 定义它的begin end中有效

~~~sql
create procedure my3(in id int)
begin
    declare result varchar(20) default '';   //局部变量申明
    select s_name into result from			//局部变量赋值
                            Student
                                where s_id=id;
    select result;							//局部变量使用

end;
~~~





### 存储过程

- 适用于增删改查
- 减少连接次数，提高性能

#### 1. 无参过程

- 创建

~~~sql
CREATE PROCEDURE MYP1()
BEGIN
    select * from Student;

end;
~~~

- 使用

~~~sql
call MYP1();
~~~

#### 2. in过程

- 创建

~~~sql
create procedure my3(in id int,in score varchar(20))
begin
    declare result varchar(20) default '';   //局部变量申明
    select s_name into result from			//局部变量赋值
                            Student
                                where s_id=id;
    select result;							//局部变量使用

end;

~~~

- 使用

~~~sql
set @studentNumb =08;  				//声明用户变量

call my3(@studentNumb);	//使用用户变量
~~~

#### 3. out过程

- 创建

~~~sql
create procedure myout1(in id int,out name varchar(20),out sex varchar(20))
begin
    select s_name,s_sex into name,sex    //直接賦值
    from Student where  s_id =id;
end;
~~~

- 使用

~~~sql
call myout1(@studentNumb,@stName,@stSex);   //声明并赋值
select @stName,@stSex;
~~~

#### 4. inout过程

- 创建

~~~sql

create procedure myinout(inout numb int)
begin
    set numb = numb * 2;
end;

~~~

- 使用

~~~sql
set @n = 10;		//声明变量并赋值
call myinout(@n);	//传入皆返回

select @n;			//查看变量值
~~~

#### 5. 删除过程

~~~sql
drop procedure [过程名];      
~~~

#### 6. 查看过程

~~~sql
show create procedure my3;
~~~



### 函数

- 适用于查询
- 有且仅有一个返回值

#### 1. 创建函数



