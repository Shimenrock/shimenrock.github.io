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
<!-- 
重学MySQL打卡行动Day16！

学习内容 : 表锁、行锁和一致性读
对应篇目：
19 | 为什么我只查一行的语句，也执行这么慢？http://gk.link/a/101KL

一般情况下，如果有人跟你说查询性能优化，你首先会想到一些复杂的语句，想到查询需要返回大量的数据。但有些情况下，“查一行”，也会执行得特别慢。

今天，丁奇会用在一个简单的表上执行“查一行”，可能会出现的被锁住和执行慢的例子，和你分析下其中的原理。在分析过程中，会涉及表锁、行锁和一致性读的概念。
-->

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

# 1.普通锁

## 1.1.查询长时间不返回

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
### 1.1.1.等MDL锁

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

### 1.1.2.等flush

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
 
### 1.1.3.等行锁

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

## 1.2.查询慢

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

<!-- 
重学MySQL打卡行动Day17！

学习内容 : 间隙锁和next-key lock	
对应篇目：
20 | 幻读是什么，幻读有什么问题？http://gk.link/a/101Mp

今天这篇文章，会为你讲述关于幻读的两大知识点：

1. 幻读需要注意两点：一是，在“当前读”下才会出现；二是，仅专指“新插入的行”。
2. 引入间隙锁和next-key lock，可以解决幻读问题，但也会带来并发度的问题。
-->



# 2.间隙锁和next-key lock

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

## 2.1.幻读

| | | |
| --- | --- | --- |
| session A t1 | begin; select * from t20 where d=5 for update;| 当前读，加锁，只返回id=5行 |
| session B t2 | update t20 set d=5 where id=0;| |
| session A t3 | select * from t20 where d=5 for update;| 当前读，加锁 |
| session C t4 | insert into t20 values(1,1,5);| |
| session A t5 | select * from t where d=5 for update; commit;| 当前读，加锁 |

performance_schema=on