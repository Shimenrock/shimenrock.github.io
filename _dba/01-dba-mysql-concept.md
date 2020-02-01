---
title: "MYSQL 基础概念"
permalink: /dba/mysql-concept/
excerpt: "MYSQL 基础概念"
last_modified_at: 2020-02-01T21:36:11-04:00
categories: mysql
redirect_from:
  - /theme-setup/
toc: true
---
<!--1课笔记，日期：1月20日-->

# 1.基础架构：一条SQL查询语句是如何执行的

<div style='display: none'>
打卡：极客时间小程序-课间-每日打卡-MySQL打卡;
1）每阶段 2 周，每周打卡≥ 3 天，即视为完成该阶段的学习。
2）完成每阶段打卡，且累计打卡≥ 24 次，即视为完成本次打卡计划。
3）每日只计数 1 次打卡，单日内多次打卡视为 1 次。
4) 学习完某部分之后，在小程序上记录学习笔记或者感悟，视为当日打卡成功
</div>

**MySQL可以分为Server层和存储引擎层**
  - Server 层包括连接器、查询缓存、分析器、优化器、执行器等，涵盖 MySQL 的大多数核心服务功能，以及所有的内置函数（如日期、时间、数学和加密函数等），所有跨存储引擎的功能。
  - 存储引擎层负责数据的存储和提取。其架构模式是插件式的，支持 InnoDB、MyISAM、Memory 等多个存储引擎。 MySQL 5.5.5 版本开始，InnoDB成为了默认存储引擎。
    - create table 默认使用InnoDB，否则要指定存储引擎。
    - 不同存储引擎共用一个Server层

## 1.1.连接器
  - 负责跟客户端建立连接，获取权限，维持和管理连接
  - mysql -h$ip -P$port -u$user -p
  - 密码不正确 > Access denied for user 客户端程序结束
  - 密码认证通过 > 连接器查出你拥有权限。
  - 空闲状态  show processlist > Sleep
  - 自动断开 默认8小时 wait_timeout参数控制
  - 长连接：连接成功后，如果客户端持续有请求，则一直使用同一连接。
    - 因为连接过程较复杂，尽量使用长连接。
    - 长连接过多，有时造成内存占用太大，最终连接会被系统强行杀掉。
  - 短连接：每次执行完成很少几次查询后就断开。

  **方案：1.定期断开长连接；2.执行一个较大操作后，执行mysql_reset_connection重新初始化连接资源（过程不需要重连和权限认证）**{: .notice--info}

## 1.2.查询缓存
  - 接到查询请求后，先查询缓存。
  - 之前执行的语句和结果以key-value对形式缓存在内存中
  - 否则继续执行阶段
  - 查询缓存弊大于利
  - 静态表，比如系统配置表适合查询缓存
  - 参数query_cache_type为DEMAND ,默认不使用查询缓存
  - 显式指定 mysql> select SQL_CACHE * FROM T where ID=10;
  - MYSQL 8.0彻底没有缓存

## 1.3.分析器
  - 词法分析
  - 语法分析

## 1.4.优化器
  - 表里有多个索引时，决定使用哪个索引
  - 一个语句有多表关联join，决定各个表连接顺序。

## 1.5.执行器
  - 判断是否执行权限
  - 打开表，根据表的引擎定义，使用引擎接口
  - 慢查询日志 rows_examined 字段：执行过程扫描了多少行。

<!--2课笔记，日期：1月21日-->

# 2.日志系统：一条SQL更新语句是如何执行的？

**回顾：查询语句执行过程：经过连接器、分析器、优化器、执行器等功能模块，最后到达存储引擎。**{: .notice}

```
mysql> create table T(ID int primary key, c int);
mysql> updata T set c=c+1 where ID=2;
```

## 2.1.redolog
  - WAL： Write-Ahead Logging 先写日志，再写磁盘
  - 一条记录更新，InnoDB引擎先将记录写入redo log里，并更新内存；适当时候将这个操作记录更新到磁盘里。
  - 固定大小，比如一组4个文件，每个文件1G
  - crash-safe 保证数据库发生异常重启，之前提交的记录不丢失。
  - innodb_flush_log_at_trx_commit设置1，每次事务都持久化到磁盘

## 2.2.binlog 归档日志
  - Server层：功能层        》binlog
  - 引擎层：负责存储相关事宜 》 redolog
  - sync_binlog 参数设置1，每次事务的binlog都持久化到磁盘

## 2.3.两种日志区别
  - redolog是InnoDB 特有；binlog是server层实现，所有引擎都可以使用
  - redolog是物理日志，记录某个数据页做的修改；binlog是逻辑日志
  - redolog是循环写；binlog追加写入

## 2.4.两阶段提交
  - 两阶段提交是跨系统维持数据逻辑一致性时常用的一个方案。
  - redo log 和binlog都可以用于表示事务的提交状态，而两阶段提交就是让这两个状态保持逻辑上的一致。

## 2.5.双1安全参数
1. redo log 用于保证 crash-safe 能力。innodb_flush_log_at_trx_commit 这个参数设置成 1 的时候，表示每次事务的 redo log 都直接持久化到磁盘。这个参数我建议你设置成 1，这样可以保证 MySQL 异常重启之后数据不丢失
2. sync_binlog 这个参数设置成 1 的时候，表示每次事务的 binlog 都持久化到磁盘。这个参数我也建议你设置成 1，这样可以保证 MySQL 异常重启之后 binlog 不丢失

# 3.事务隔离：为什么你改了我还看不见

<!--3课笔记，日期：1月22日-->

## 3.1.隔离性和隔离级别

  **ACID(Atomicity,Consistency,Isolation,Durability原子性，一致性，隔离性，持久性)**
  - 脏读  dirty read
  - 不可重复读  non-repeatable read
  - 幻读  phantom read

  **隔离与效率**
  - 隔离越严实，效率越低
  - 需要在隔离和效率间寻找平衡点

  **SQL标准事务隔离级别**
  - 读未提交 read uncommitted 一个事务还没有提交，它做的变更就能被别的事务看到。
  - 读提交 read committed  一个事务提交之后，它所做的变更才会被其他事务看到。
  - 可重复读 repeatable read 一个事务执行过程中看到的数据，总是跟这个事务启动时看到数据是一致的。未提交变更对其他事务也是不可见的。
  - 串行化 serializable 对于同一行记录，写会加锁，读会加读锁。出现读写锁冲突的时候，后访问的事务必须等前一个事务执行完成，才能继续执行。
```
create table T(c int) engine=InnoDB
insert into T(c) values(1);
```

| 事务A               | 事务B       |
| ------------------- | ----------- |
| 启动事务查询得到值1 | 启动事务    |
|                     | 查询得到值1 |
|                     | 将1改成2    |
| 查询得到值V1        |             |
|                     | 提交事务B   |
| 查询得到值V2        |             |
| 提交事务A           |             |
| 查询得到值V3        |             |


| 模式     | V1  | V2  | V3  | 视图                      |
| -------- | --- | --- | --- | ------------------------- |
| 读未提交 | 2   | 2   | 2   | 直接返回记录上的值        |
| 读提交   | 1   | 2   | 2   | 每个SQL语句执行时创建视图 |
| 可重复读 | 1   | 1   | 2   | 事务启动时创建视图        |
| 串行化   | 1   | 1   | 2   | 直接用加锁方式并并行访问  |

  **ORACLE默认隔离级别是读提交,迁移ORACLE至MySQL将隔离级别设置为读提交**
  - show variables like 'transaction_isolation'
  - READ-COMMITTED

## 3.2.事务隔离的实现：回滚日志
  - MVCC 多版本并发控制
  - 当系统里没有比这个回滚日志更早的read-view时，删除回滚日志。
  - MySQL 5.5 以前版本，回滚日志和数据字典一起放在ibdata文件，及时长事务提交，回滚段被清理，文件也不会变小，最终导致重建整个库

## 3.3.事务的启动方式
  - 显式启动事务，begin或start transaction  提交commit 回滚 rollback
  - set autocommit=0 自动提交关闭，事务直到commit或rollback
  - 客户端连接架构默认set autocommit=0 导致长事务。
  - 建议 set autocommit=1 通过显式语句方式启动事务。

## 3.4.commit work and chain语法
  - 提交事务并自动启动下一个事务

```
select * from information_schema,innodb_trx where TIME_TO_SEC(timediff(now(),trx_started))>60
查询持续时间超过60s的事务
```

## 3.5.避免长事务对业务影响
  **应用开发端**
  - 确认是否使用了set autocommit=0。检查general_log ,跑一个业务逻辑。目标改成1
  - 确认是否有不必要的只读事务。确认语句有没有必要使用begin/commit。
  - 业务连接数据库，通过SET MAX_EXECUTION_TIME，控制每个语句执行最长时间
  **数据库端**
  - 监控information_schema.Innodb_trx表，设置长事务阀值，超过就报警或着kill;
  - Percona的pt-kill工具
  - 在业务功能测试阶段要求输出所有general_log，分析日志行为提前发现问题；
  - 如果使用mysql5.6以上版本，innodb_undo_tablespaces设置2或更大，如果出现大事务导致回滚段过大，这样设置后清理起来更方便。

# 4.深入浅出索引

<!--4课5课笔记，日期：1月23日-->

## 4.1.索引常见模型
  - 哈希表
    - 键-值 key-value存储数据结构
    - 适用于只有等值查询场景，比如Memcached以及一些NoSQL引擎
  - 有序数组
    - 等值查询和范围查询场景中的性能都优秀
    - 查询效率高，更新数据效率低
    - 只适用于静态存储引擎，比如某市人口信息
  - 搜索树
    - 二叉树时间复杂度O(log(N))
    - N叉树 ，性能优点适配磁盘访问模式，广泛应用在数据库引擎。

## 4.2.InnoDB索引模型
  - 表都根据主机顺序以索引的形式存放--索引组织表
  - B+树索引模型
```
create table T(
  id int primary key,
  k int not null,
  name varchar(16),
  index (k)engine=InnoDB;
  )
```
  - 主键索引（聚簇索引clustered index）和非主键索引（二级索引secondary index）
  - 主键索引叶子节点存整行数据
  - 非主键索引叶子节点内容是主键的值

## 4.3.基于主机索引和普通索引的查询区别
  - 主键查询：搜索ID这棵B+树
  - 普通索引查询：先搜索索引树，得到主键，再在主键索引树搜索一次。这个过程称为回表。

## 4.4.索引维护
  - 页分裂
  - 页合并
  - 自增主键 NOT NULL PRIMARY KEY AUTO_INCREMENT 每次插入新记录，是追加操作，不会触发叶子节点分裂（性能和存储考虑最优）
  - 业务逻辑字段做主键，不容易保证有序插入。
  - **主键长度越小，普通索引的叶子节点越小，普通索引占用空间越小**

## 4.5.知识扩展
重建索引方法1
```
alter table T drop index k;
alter table T add index(k);
```
重建索引方法2
```
alter table T drop primary key;
alter table T add primary key(id);
```

```
create table T(
  ID int primary key,
  k int NOT NULL DEFAULT 0,
  s varchar(16) NOT NULL DEFAULT '',
  index k(k)
  engine=InnoDB;
  )
insert into T values(100,1,'aa'),(200,2,'bb'),(300,3,'cc'),(500,5,'ee'),(600,6,'ff'),(700,7,'gg');
```

## 4.6.覆盖索引
  - select ID from T where k between 3 and 5
  - 减少树的搜索次数，提升查询性能。

## 4.7.联合索引
  - 根据市民身份证查询他的姓名

## 4.8.最前缀原则
  - 第一原则，如果通过调整顺序，可以少维护一个索引，那么这个顺序往往就是需要优先考虑采用的。
  - 考虑空间
  - 索引下推

# 6.全局锁和表锁：给表加个字段怎么有这么多阻碍？

<!--6课7课笔记，日期：1月24日-->

## 6.1.全局锁
  - 整个数据库实例加锁
  - Flush tables with read lock（FTWRL）
  - 阻塞：数据更新语句、数据定义语句、更新类事务提交语句。
  - 典型应用场景，全库逻辑备份
  - mysqlddump 使用参数single-transaction 导数据之前，启动一个事务，确保拿到一致性视图，
  - mysqlddump的一致性读是好，但不是所有引擎支持这个隔离级别，所以需要FTWRL
  - single-transaction方法只适用于所有表使用事务引擎的库，否则备份只能通过FTWRL，这便是InnoDB替代MyISAM愿意之一。

## 6.2.readonly
  - 有些系统中，readonly值被用来做其他逻辑，比如判断一个库是主还是备。
  - 异常处理机制上有差异。
    - FTWRL ,客户端发生异常断开，mysql自动释放全局锁
    - readonly，客户端发生异常，数据库一直保持readonly状态，风险高。

## 6.3.表级锁
  - 表锁
    - lock table...read/write
    - 可用unlock tables主动释放锁
    - 可断开客户端释放锁
    - 即限制其他线程，也限定本线程接下来操作对象
  - 元数据锁 meta data lock (MDL) MySQL5.5以上
    - 不需要显式使用，访问表时自动加上
    - 保证读写正确性

## 6.4.小表加字段
  - 解决长事务，避免一直MDL锁
  - information_schema 库里的innodb_trx表中，查询当前执行中的事务，考虑先暂停DDL或着kill长事务。
  - alter table设定等待时间。
  - ALTER TABLE tbl_name NOWAIT add column
  - ALTER TABLE tbl_name WAIT N add column

## 6.5.两阶段锁
  - 在InnoDB事务中，行锁是在需要时候才加上的，但并不是不需要了就立刻释放，而是要等到事务结束时才释放。这个就是两阶段锁协议。
  - 如果你的事务中需要锁多个行，要把最可能造成锁冲突、最可能影响并发度的锁尽量往后放。

## 6.6.死锁
  - 策略1，直接进入等待，直到超时。超时时间可以通过参数innodb_lock_wait_timeout来设置。
  - 策略2，发起死锁检测，主动回滚死锁链条中的某一个事务，让其他事务得以继续执行。innodb_deadlock_detect设置on，表示开启这个逻辑。
  - innodb_lock_wait_timeout默认50s
  - 并发业务死锁检测要消耗大量CPU资源
    - 临时把死锁检测关掉，业务有损。
    - 控制并发度
    - 中间件控制
    - 一行改成逻辑上的多行来减少锁冲突。
