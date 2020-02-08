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

**实验初始化**
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



 
 

