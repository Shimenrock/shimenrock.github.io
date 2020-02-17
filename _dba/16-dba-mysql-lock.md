---
title: "Mysql里的锁"
permalink: /dba/mysql-lock/
excerpt: "Mysql里的锁"
last_modified_at: 2020-02-06T21:36:11-04:00
categories: mysql
redirect_from:
  - /theme-setup/
toc: true
---

# 实验初始化

```
mysql> CREATE TABLE `t` (
  `id` int(11) NOT NULL,
  `c` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;

delimiter ;;
create procedure idata()
begin
  declare i int;
  set i=1;
  while(i<=100000) do
    insert into t values(i,i);
    set i=i+1;
  end while;
end;;
delimiter ;

call idata();
```

# 一、普通锁

<!-- 
Day16: 学习内容 : 表锁、行锁和一致性读
对应篇目：19 | 为什么我只查一行的语句，也执行这么慢？http://gk.link/a/101KL
-->

## 1.查询长时间不返回

```
mysql> select * from t where id=1;

show processlist   # 查询语句处于什么状态

MariaDB [test]> show processlist
    -> ;
+----+------+-----------+------+---------+------+---------------------------------+----------------------------+----------+
| Id | User | Host      | db   | Command | Time | State                           | Info                       | Progress |
+----+------+-----------+------+---------+------+---------------------------------+----------------------------+----------+
| 10 | root | localhost | test | Sleep   |   57 |                                 | NULL                       |    0.000 |
| 12 | root | localhost | test | Query   |   44 | Waiting for table metadata lock | select * from t where id=1 |    0.000 |
| 13 | root | localhost | test | Query   |    0 | NULL                            | show processlist           |    0.000 |
| 14 | root | localhost | test | Sleep   |  154 |                                 | NULL                       |    0.000 |
+----+------+-----------+------+---------+------+---------------------------------+----------------------------+----------+
4 rows in set (0.00 sec)
```
### 1.1.等MDL锁

**MDL:metadata lock，属于表级锁。**

show processlist 显示: Waiting for table metadata lock
这个状态表示的是，现在有一个线程正在表 t 上请求或者持有 MDL 写锁，把 select 语句堵住了

**复现**
|  |  |
| --- | --- |
| sessionA | lock table t write;  |
| sessionB | select * from t where id=1; |

处理方法：找到谁持有 MDL 写锁，然后把它 kill 掉。

MySQL 启动时需要设置 performance_schema=on，相比于设置为 off 会有 10% 左右的性能损失
```
select blocking_pid from sys.schema_table_lock_waits;
```

### 1.2.等flush

mysql对表flush操作一般两种：
```
flush tables t with read lock;    # 只关闭表t
flush tables with read lock;      # 关闭所有打开的表
```
正常两个语句执行都很快，除非它们也被别的线程堵住

**复现**
|  |  |  |
| --- | --- | --- |
| sessionA | select sleep(1) from t; | 1万秒 |
| sessionB | flush tables t; | 等sessionA结束 |
| sessionC | select * from t where id=1; | 被flush命令堵住 |

```
mysql> show processlist;
+----+------+-----------+------+---------+------+-------------------------+----------------------------+
| Id | User | Host      | db   | Command | Time | State                   | Info                       |
+----+------+-----------+------+---------+------+-------------------------+----------------------------+
| 13 | root | localhost | test | Query   |   36 | User sleep              | select sleep(1) from t     |
| 16 | root | localhost | test | Query   |   20 | Waiting for table flush | flush tables t             |
| 17 | root | localhost | test | Query   |    4 | Waiting for table flush | select * from t where id=1 |
| 20 | root | localhost | NULL | Query   |    0 | init                    | show processlist           |
+----+------+-----------+------+---------+------+-------------------------+----------------------------+
4 rows in set (0.00 sec)

mysql> select * from information_schema.processlist where id=17
    -> ;
+----+------+-----------+------+---------+------+-------------------------+----------------------------+
| ID | USER | HOST      | DB   | COMMAND | TIME | STATE                   | INFO                       |
+----+------+-----------+------+---------+------+-------------------------+----------------------------+
| 17 | root | localhost | test | Query   |   26 | Waiting for table flush | select * from t where id=1 |
+----+------+-----------+------+---------+------+-------------------------+----------------------------+
1 row in set (0.00 sec)
```
 
### 1.3.等行锁

| | | |
| --- | --- | --- |
| sessionA | begin;update t set c=c+1 where id=1 | 启动事务，占写锁，不提交 |
| sessionB | select * from t where id=1 lock in share mode; | |

```
mysql> select * from sys.innodb_lock_waits where locked_table=`'test'.'t'`\G
mysql> select * from sys.innodb_lock_waits \G;
*************************** 1. row ***************************
                wait_started: 2020-02-09 19:00:36
                    wait_age: 00:00:36
               wait_age_secs: 36
                locked_table: `test`.`t`
                locked_index: PRIMARY
                 locked_type: RECORD
              waiting_trx_id: 421895922814576
         waiting_trx_started: 2020-02-09 19:00:36
             waiting_trx_age: 00:00:36
     waiting_trx_rows_locked: 1
   waiting_trx_rows_modified: 0
                 waiting_pid: 5
               waiting_query: select * from t where id=1 lock in share mode
             waiting_lock_id: 421895922814576:7:4:2
           waiting_lock_mode: S
             blocking_trx_id: 103191
                blocking_pid: 4
              blocking_query: NULL
            blocking_lock_id: 103191:7:4:2
          blocking_lock_mode: X
        blocking_trx_started: 2020-02-09 18:33:25
            blocking_trx_age: 00:27:47
    blocking_trx_rows_locked: 1
  blocking_trx_rows_modified: 1
     sql_kill_blocking_query: KILL QUERY 4
sql_kill_blocking_connection: KILL 4
1 row in set, 3 warnings (0.01 sec)

ERROR: 
No query specified
```

## 2.查询慢

开启慢查询日志
```
mysql> set global slow_query_log='ON';
mysql> set global slow_query_log_file='/mydata/data/instance-1-slow.log';
mysql> set global long_query_time=0;
mysql> set long_query_time=0;
```
```
# Query_time: 0.007436  Lock_time: 0.000101 Rows_sent: 1  Rows_examined: 50000
SET timestamp=1581271451;
select * from t where c=50000 limit 1;

# Query_time: 0.000231  Lock_time: 0.000119 Rows_sent: 1  Rows_examined: 1
SET timestamp=1581271546;
select * from t where id=1;

# Query_time: 0.000198  Lock_time: 0.000102 Rows_sent: 1  Rows_examined: 1
SET timestamp=1581271620;
select * from t where id=1 lock in share mode;
```
- 带 lock in share mode 的 SQL 语句，是当前读，因此会直接读到 1000001 这个结果，所以速度很快；
- select * from t where id=1 这个语句，是一致性读，因此需要从 1000001 开始，依次执行 undo log，执行了 100 万次以后，才将 1 这个结果返回。


# 二、间隙锁和next-key lock

<!-- 
Day17:学习内容 : 间隙锁和next-key lock	
对应篇目：20 | 幻读是什么，幻读有什么问题？http://gk.link/a/101Mp
1. 幻读需要注意两点：一是，在“当前读”下才会出现；二是，仅专指“新插入的行”。
2. 引入间隙锁和next-key lock，可以解决幻读问题，但也会带来并发度的问题。
-->

实验初始化

```
CREATE TABLE `t20` (
  `id` int(11) NOT NULL,
  `c` int(11) DEFAULT NULL,
  `d` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `c` (`c`)
) ENGINE=InnoDB;

insert into t20 values(0,0,0),(5,5,5),(10,10,10),(15,15,15),(20,20,20),(25,25,25);
```

## 1.幻读

| | | |
| --- | --- | --- |
| session A t1 | begin; select * from t20 where d=5 for update;/*Q1*/| 当前读，加锁，只返回id=5行 |
| session B t2 | update t20 set d=5 where id=0;| |
| session A t3 | select * from t20 where d=5 for update;/*Q2*/| 当前读，加锁 |
| session C t4 | insert into t20 values(1,1,5);| |
| session A t5 | select * from t20 where d=5 for update;/*Q1*/ | 当前读，加锁 |
| session A t6 | commit;| |

注:没有演示出效果

1. 在可重复读隔离级别下，普通的查询是快照读，是不会看到别的事务插入的数据的。因此，幻读在“当前读”下才会出现。
2. 上面 session B 的修改结果，被 session A 之后的 select 语句用“当前读”看到，不能称为幻读。幻读仅专指“新插入的行”。

**幻读问题：**

- 语义被破坏，即加锁声明被破坏
- 数据和日志在逻辑上不一致，即数据一致性存在问题。
- 把所有的记录都加上了锁，但没有阻止插入新记录。

## 2.间隙锁

解决幻读：引入间隙锁

间隙锁的引入，可能会导致同样的语句锁住更大的范围，这其实是影响了并发度的

## 3.next-key lock

间隙锁和行锁合称next-key lock，每个next-key lock是前开后闭区间。

# 三、加锁规则

<!-- 
Day18：学习内容 : InnoDB的加锁规则 
对应篇目：
21 | 为什么我只查一行的语句，锁这么多？http://gk.link/a/101Oe
30 | 答疑文章（二）：用动态的观点看加锁 http://gk.link/a/1020p

原则1：加锁的基本单位是next-key lock。
原则2：查找过程中访问到的对象才会加锁。
优化1：索引上的等值查询，给唯一索引加锁的时候，next-key lock退化为行锁。
优化2：索引上的等值查询，向右遍历时且最后一个值不满足等值条件的时候，next-key lock退化为间隙锁。
一个bug：唯一索引上的范围查询会访问到不满足条件的第一个值为止。
-->

# 1.前提

1. MySQL 后面的版本可能会改变加锁策略，限于目前版本 5.x系列<= 5.7.24  8.0系列<= 8.0.13
2. 默认可重复读隔离级别。

# 2.规则

1. 原则 1：加锁的基本单位是 next-key lock。next-key lock 是前开后闭区间。
2. 原则 2：查找过程中访问到的对象才会加锁。
3. 优化 1：索引上的等值查询，给唯一索引加锁的时候，next-key lock 退化为行锁。
4. 优化 2：索引上的等值查询，向右遍历时且最后一个值不满足等值条件的时候，next-key lock 退化为间隙锁。
5. 一个 bug：唯一索引上的范围查询会访问到不满足条件的第一个值为止。

```
CREATE TABLE `t7` (
  `id` int(11) NOT NULL,
  `c` int(11) DEFAULT NULL,
  `d` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `c` (`c`)
) ENGINE=InnoDB;

insert into t7 values(0,0,0),(5,5,5),(10,10,10),(15,15,15),(20,20,20),(25,25,25);
```

### 案例一，等值查询间隙锁

| | | |
| --- | --- | --- |
| sessionA | begin;update t7 set d=d+1 where id=7;| 根据原则1，next-key lock，加锁区间5，10 |
| sessionB | insert into t7 values(8,8,8);(blocked) | |
| sessionC | update t7 set d=d+1 where id=10; | 根据优化2，退化为间隙锁，修改成功 |

### 案例二，非唯一索引等值锁

| | | |
| --- | --- | --- |
| sessionA | select id from t where c=5 lock in share mode; |  |
| sessionB | update t set d=d+1 where id=5; | |
| sessionC | insert into t values(7,7,7); | blocked |

1. 根据原则 1，加锁单位是 next-key lock，因此会给 (0,5]加上 next-key lock。
2. 要注意 c 是普通索引，因此仅访问 c=5 这一条记录是不能马上停下来的，需要向右遍历，查到 c=10 才放弃。根据原则 2，访问到的都要加锁，因此要给 (5,10]加 next-key lock。
3. 但是同时这个符合优化 2：等值判断，向右遍历，最后一个值不满足 c=5 这个等值条件，因此退化成间隙锁 (5,10)。
4. 根据原则 2 ，只有访问到的对象才会加锁，这个查询使用覆盖索引，并不需要访问主键索引，所以主键索引上没有加任何锁，这就是为什么 session B 的 update 语句可以执行完成。

如果你要用 lock in share mode 来给行加读锁避免数据被更新的话，就必须得绕过覆盖索引的优化，在查询字段中加入索引中不存在的字段。比如，将 session A 的查询语句改成 select d from t where c=5 lock in share mode。

