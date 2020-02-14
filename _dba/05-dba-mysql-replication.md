---
title: "MYSQL 主从复制"
permalink: /dba/mysql-master-slave-replication/
excerpt: "MYSQL 主从复制"
last_modified_at: 2020-02-01T21:36:11-04:00
categories: mysql
redirect_from:
  - /theme-setup/
toc: true
---
<!--马哥42课笔记，日期：2月10日-->

# 1.MYSQL的扩展

- MYSQL向上扩展：硬件性能扩展（128G内存，raid10固态硬盘）
- MYSQL横向扩展
  - 复制：每个节点都有相同的数据集

# 2.复制数据库的作用

- 数据分布：
- 负载均衡读操作
- 复制冗余
- 高可用和故障切换
- mysql升级测试

# 3.主从复制原理

- 主节点：
  - dump thread：为每个slave的I/O Thread启动一个dump 线程，用于向其发送binary log events;
- 从节点：
  - I/O Thread :从master请求二进制日志事件，并保存于中继日志中
  - SQL Thread：从中继日志中读取日志事件，在本地完成重放
- 读写分离语句路由
  - r/w spliter
  - 七层负载均衡的反向代理    
- 事务一致
  - 5.6以后，gtid全局事务id

 ## 第三方解决方案       

    MMM过气、MHA、Galera-Cluster 块级别

# 4.主从复制特点
- 异步复制
  - 复制是否成功不需要返回确定
- 主从数据不一致比较常见
  - 网络延迟
  - 服务器性能
  - 事务落到日志上
  - 各种可能性
- 根据业务对数据一致要求综合考虑

# 5.复制架构

- 主从
  - 一主多从
    - 从服务器还可以再有从服务器
  - 一从多主
    - 数据库名不能一样，进行数据汇聚
- 半同步复制
  - Google开发：给一个节点同步复制（带宽足够大，同机架），其他节点异步复制
- 复制过滤器

# 6.二进制日志事件记录格式
    
- STATEMENT（语句模式）
- ROW （行模式）        
- MIXED （混合模式）

# 7.主从配置过程

## 7.1.主节点：

1.启动二级制日志；

```
   [mysqld]
   log_bin=mysql-bin
```

2.为当前节点设置一个全局唯一的ID号；

```
   [mysqld]
   server_id=#
```

3.创建有复制权限的用户账号；

   **REPLICATION SLAVE, REPLICATION CLIENT**

```  
   GRANT REPLCATION SLAVE, REPLICATION CLIENT ON *.* TO 'repluser'@'HOST' IDENTIFIED BY 'replpass';
```

## 7.2.从节点：

1.启动中继日志；

```
   [mysqld]
   relay_log=relay-log
   relay_log_index=relay-log.index
```

2.为当前节点设置一个全局唯一的ID号；
   
## 7.3.启动复制

使用有复制权限的用户账户连接至主服务器，并启动复制线程。

```
    CHANGE MASTER TO MASTER_HOST='host', MASTER_USER='repluser', MASTER_PASSWORD='replpass', MASTER_LOG_FILE='mysql-bin.xxxxx', MASTER_LOG_POS=#;
    mysql> START SLAVE [IO_THREAD|SQL_THREAD];
```

## 7.4.同步修复

如果主节点已经运行一段时间，有大量数据时 

- 通过备份恢复数据至从服务器；
- 复制起始位置为备份时，二进制日志文件及其POS;

## 7.5.操作过程

**主节点**
```
# vim /etc/mysql/my.cnf

log_bin = /mydata/data/mysql_binlog/mysql-bin
service_id= 1
innodb_file_per_talbe = on
skip_name_resolve = on

# systemctl start mariadb.service

# mysql -uroot -p

mysql> show global variables like '%log%'; 
+--------------------------------------------+-------------------------------------------+
| Variable_name                              | Value                                     |
+--------------------------------------------+-------------------------------------------+
| back_log                                   | 80                                        |
| binlog_cache_size                          | 32768                                     |
| binlog_checksum                            | CRC32                                     |
| binlog_direct_non_transactional_updates    | OFF                                       |
| binlog_error_action                        | ABORT_SERVER                              |
| binlog_format                              | ROW                                       |
| binlog_group_commit_sync_delay             | 0                                         |
| binlog_group_commit_sync_no_delay_count    | 0                                         |
| binlog_gtid_simple_recovery                | ON                                        |
| binlog_max_flush_queue_time                | 0                                         |
| binlog_order_commits                       | ON                                        |
| binlog_row_image                           | FULL                                      |
| binlog_rows_query_log_events               | OFF                                       |
| binlog_stmt_cache_size                     | 32768                                     |
| binlog_transaction_dependency_history_size | 25000                                     |
| binlog_transaction_dependency_tracking     | COMMIT_ORDER                              |
| expire_logs_days                           | 0                                         |
| general_log                                | OFF                                       |
| general_log_file                           | /mydata/data/mysql220.log                 |
| innodb_api_enable_binlog                   | OFF                                       |
| innodb_flush_log_at_timeout                | 1                                         |
| innodb_flush_log_at_trx_commit             | 1                                         |
| innodb_locks_unsafe_for_binlog             | OFF                                       |
| innodb_log_buffer_size                     | 16777216                                  |
| innodb_log_checksums                       | ON                                        |
| innodb_log_compressed_pages                | ON                                        |
| innodb_log_file_size                       | 268435456                                 |
| innodb_log_files_in_group                  | 2                                         |
| innodb_log_group_home_dir                  | ./                                        |
| innodb_log_write_ahead_size                | 8192                                      |
| innodb_max_undo_log_size                   | 1073741824                                |
| innodb_online_alter_log_max_size           | 134217728                                 |
| innodb_undo_log_truncate                   | OFF                                       |
| innodb_undo_logs                           | 128                                       |
| log_bin                                    | ON                                        |
| log_bin_basename                           | /mydata/data/mysql_binlog/mysql-bin       |
| log_bin_index                              | /mydata/data/mysql_binlog/mysql-bin.index |
| log_bin_trust_function_creators            | OFF                                       |
| log_bin_use_v1_row_events                  | OFF                                       |
| log_builtin_as_identified_by_password      | OFF                                       |
| log_error                                  | /var/log/mariadb/mariadb.log              |
| log_error_verbosity                        | 3                                         |
| log_output                                 | FILE                                      |
| log_queries_not_using_indexes              | OFF                                       |
| log_slave_updates                          | OFF                                       |
| log_slow_admin_statements                  | OFF                                       |
| log_slow_slave_statements                  | OFF                                       |
| log_statements_unsafe_for_binlog           | ON                                        |
| log_syslog                                 | OFF                                       |
| log_syslog_facility                        | daemon                                    |
| log_syslog_include_pid                     | ON                                        |
| log_syslog_tag                             |                                           |
| log_throttle_queries_not_using_indexes     | 0                                         |
| log_timestamps                             | UTC                                       |
| log_warnings                               | 2                                         |
| max_binlog_cache_size                      | 18446744073709547520                      |
| max_binlog_size                            | 1073741824                                |
| max_binlog_stmt_cache_size                 | 18446744073709547520                      |
| max_relay_log_size                         | 0                                         |
| relay_log                                  |                                           |
| relay_log_basename                         | /mydata/data/mysql220-relay-bin           |
| relay_log_index                            | /mydata/data/mysql220-relay-bin.index     |
| relay_log_info_file                        | relay-log.info                            |
| relay_log_info_repository                  | FILE                                      |
| relay_log_purge                            | ON                                        |
| relay_log_recovery                         | OFF                                       |
| relay_log_space_limit                      | 0                                         |
| slow_query_log                             | OFF                                       |
| slow_query_log_file                        | /mydata/data/mysql220-slow.log            |
| sql_log_off                                | OFF                                       |
| sync_binlog                                | 1                                         |
| sync_relay_log                             | 10000                                     |
| sync_relay_log_info                        | 10000                                     |
+--------------------------------------------+-------------------------------------------+
73 rows in set (0.00 sec)

mysql> SHOW MASTER LOGS;
+------------------+-----------+
| Log_name         | File_size |
+------------------+-----------+
| mysql-bin.000001 |       143 |
| mysql-bin.000002 |  19203039 |
| mysql-bin.000003 |       177 |
| mysql-bin.000004 |       575 |
| mysql-bin.000005 |       800 |
| mysql-bin.000006 |       154 |
| mysql-bin.000007 |       154 |
| mysql-bin.000008 |       177 |
| mysql-bin.000009 |       154 |
+------------------+-----------+
9 rows in set (0.01 sec)

mysql> show global variables like '%server%';
+---------------------------------+--------------------------------------+
| Variable_name                   | Value                                |
+---------------------------------+--------------------------------------+
| character_set_server            | utf8mb4                              |
| collation_server                | utf8mb4_general_ci                   |
| innodb_ft_server_stopword_table |                                      |
| server_id                       | 1                                    |
| server_id_bits                  | 32                                   |
| server_uuid                     | e18da6ef-4b19-11ea-91e9-000c29daab3a |
+---------------------------------+--------------------------------------+
6 rows in set (0.00 sec)

mysql> show master status;
+------------------+----------+--------------+------------------+-------------------+
| File             | Position | Binlog_Do_DB | Binlog_Ignore_DB | Executed_Gtid_Set |
+------------------+----------+--------------+------------------+-------------------+
| mysql-bin.000009 |      633 |              |                  |                   |
+------------------+----------+--------------+------------------+-------------------+
1 row in set (0.00 sec)

mysql> GRANT REPLICATION SLAVE,REPLICATION CLIENT ON *.* TO 'repluser'@'192.168.%.%' IDENTIFIED BY 'replpass';

mysql> FLUSH PRIVILEGES;
```

**从节点**

```
# vim /etc/mysql/my.cnf

relay_log = /mydata/data/mysql_binlog/relay-log
relay-log-index = /mydata/data/mysql_binlog/relay-log.index
server-id=7
innodb__file_per_talbe = on
skip_name_resolve = on

# systemctl start mysqld.service

# mysql -uroot -p

mysql> show global variables like '%log%'; 
+--------------------------------------------+-------------------------------------------+
| Variable_name                              | Value                                     |
+--------------------------------------------+-------------------------------------------+
| back_log                                   | 80                                        |
| binlog_cache_size                          | 32768                                     |
| binlog_checksum                            | CRC32                                     |
| binlog_direct_non_transactional_updates    | OFF                                       |
| binlog_error_action                        | ABORT_SERVER                              |
| binlog_format                              | ROW                                       |
| binlog_group_commit_sync_delay             | 0                                         |
| binlog_group_commit_sync_no_delay_count    | 0                                         |
| binlog_gtid_simple_recovery                | ON                                        |
| binlog_max_flush_queue_time                | 0                                         |
| binlog_order_commits                       | ON                                        |
| binlog_row_image                           | FULL                                      |
| binlog_rows_query_log_events               | OFF                                       |
| binlog_stmt_cache_size                     | 32768                                     |
| binlog_transaction_dependency_history_size | 25000                                     |
| binlog_transaction_dependency_tracking     | COMMIT_ORDER                              |
| expire_logs_days                           | 0                                         |
| general_log                                | OFF                                       |
| general_log_file                           | /mydata/data/mysql-salve.log              |
| innodb_api_enable_binlog                   | OFF                                       |
| innodb_flush_log_at_timeout                | 1                                         |
| innodb_flush_log_at_trx_commit             | 1                                         |
| innodb_locks_unsafe_for_binlog             | OFF                                       |
| innodb_log_buffer_size                     | 16777216                                  |
| innodb_log_checksums                       | ON                                        |
| innodb_log_compressed_pages                | ON                                        |
| innodb_log_file_size                       | 268435456                                 |
| innodb_log_files_in_group                  | 2                                         |
| innodb_log_group_home_dir                  | ./                                        |
| innodb_log_write_ahead_size                | 8192                                      |
| innodb_max_undo_log_size                   | 1073741824                                |
| innodb_online_alter_log_max_size           | 134217728                                 |
| innodb_undo_log_truncate                   | OFF                                       |
| innodb_undo_logs                           | 128                                       |
| log_bin                                    | ON                                        |
| log_bin_basename                           | /mydata/data/mysql_binlog/mysql-bin       |
| log_bin_index                              | /mydata/data/mysql_binlog/mysql-bin.index |
| log_bin_trust_function_creators            | OFF                                       |
| log_bin_use_v1_row_events                  | OFF                                       |
| log_builtin_as_identified_by_password      | OFF                                       |
| log_error                                  | /var/log/mariadb/mariadb.log              |
| log_error_verbosity                        | 3                                         |
| log_output                                 | FILE                                      |
| log_queries_not_using_indexes              | OFF                                       |
| log_slave_updates                          | OFF                                       |
| log_slow_admin_statements                  | OFF                                       |
| log_slow_slave_statements                  | OFF                                       |
| log_statements_unsafe_for_binlog           | ON                                        |
| log_syslog                                 | OFF                                       |
| log_syslog_facility                        | daemon                                    |
| log_syslog_include_pid                     | ON                                        |
| log_syslog_tag                             |                                           |
| log_throttle_queries_not_using_indexes     | 0                                         |
| log_timestamps                             | UTC                                       |
| log_warnings                               | 2                                         |
| max_binlog_cache_size                      | 18446744073709547520                      |
| max_binlog_size                            | 1073741824                                |
| max_binlog_stmt_cache_size                 | 18446744073709547520                      |
| max_relay_log_size                         | 0                                         |
| relay_log                                  | /mydata/data/mysql_binlog/relay-log       |
| relay_log_basename                         | /mydata/data/mysql_binlog/relay-log       |
| relay_log_index                            | /mydata/data/mysql_binlog/relay-log.index |
| relay_log_info_file                        | relay-log.info                            |
| relay_log_info_repository                  | FILE                                      |
| relay_log_purge                            | ON                                        |
| relay_log_recovery                         | OFF                                       |
| relay_log_space_limit                      | 0                                         |
| slow_query_log                             | OFF                                       |
| slow_query_log_file                        | /mydata/data/mysql-salve-slow.log         |
| sql_log_off                                | OFF                                       |
| sync_binlog                                | 1                                         |
| sync_relay_log                             | 10000                                     |
| sync_relay_log_info                        | 10000                                     |
+--------------------------------------------+-------------------------------------------+
73 rows in set (0.00 sec)

mysql> show global variables like '%server%';
+---------------------------------+--------------------------------------+
| Variable_name                   | Value                                |
+---------------------------------+--------------------------------------+
| character_set_server            | utf8mb4                              |
| collation_server                | utf8mb4_general_ci                   |
| innodb_ft_server_stopword_table |                                      |
| server_id                       | 7                                    |
| server_id_bits                  | 32                                   |
| server_uuid                     | 8a1db7ab-4d79-11ea-a163-000c299dde47 |
+---------------------------------+--------------------------------------+
6 rows in set (0.00 sec)

mysql> HELP CHANGE MASTER TO
参数解释

mysql> CHANGE MASTER TO MASTER_HOST='192.168.11.220', MASTER_USER='repluser', MASTER_PASSWORD='replpass', MASTER_LOG_FILE='master-bin.000009', MASTER_LOG_POS=633;
这两个数值，在主节点查询 show master status\G;

mysql> SHOW SLAVE STATUS\G
*************************** 1. row ***************************
               Slave_IO_State: 
                  Master_Host: 192.168.11.200
                  Master_User: repluser
                  Master_Port: 3306
                Connect_Retry: 60
              Master_Log_File: master-bin.000009
          Read_Master_Log_Pos: 633
               Relay_Log_File: relay-log.000001
                Relay_Log_Pos: 4
        Relay_Master_Log_File: master-bin.000009
             Slave_IO_Running: No
            Slave_SQL_Running: No
              Replicate_Do_DB: 
          Replicate_Ignore_DB: 
           Replicate_Do_Table: 
       Replicate_Ignore_Table: 
      Replicate_Wild_Do_Table: 
  Replicate_Wild_Ignore_Table: 
                   Last_Errno: 0
                   Last_Error: 
                 Skip_Counter: 0
          Exec_Master_Log_Pos: 633
              Relay_Log_Space: 154
              Until_Condition: None
               Until_Log_File: 
                Until_Log_Pos: 0
           Master_SSL_Allowed: No
           Master_SSL_CA_File: 
           Master_SSL_CA_Path: 
              Master_SSL_Cert: 
            Master_SSL_Cipher: 
               Master_SSL_Key: 
        Seconds_Behind_Master: NULL
Master_SSL_Verify_Server_Cert: No
                Last_IO_Errno: 0
                Last_IO_Error: 
               Last_SQL_Errno: 0
               Last_SQL_Error: 
  Replicate_Ignore_Server_Ids: 
             Master_Server_Id: 0
                  Master_UUID: 
             Master_Info_File: /mydata/data/master.info
                    SQL_Delay: 0
          SQL_Remaining_Delay: NULL
      Slave_SQL_Running_State: 
           Master_Retry_Count: 86400
                  Master_Bind: 
      Last_IO_Error_Timestamp: 
     Last_SQL_Error_Timestamp: 
               Master_SSL_Crl: 
           Master_SSL_Crlpath: 
           Retrieved_Gtid_Set: 
            Executed_Gtid_Set: 
                Auto_Position: 0
         Replicate_Rewrite_DB: 
                 Channel_Name: 
           Master_TLS_Version: 
1 row in set (0.00 sec)

mysql> HELP START SAVLE

mysql>  START SLAVE;

mysql> SHOW SLAVE STATUS\G
*************************** 1. row ***************************
               Slave_IO_State: Connecting to master
                  Master_Host: 192.168.11.200
                  Master_User: repluser
                  Master_Port: 3306
                Connect_Retry: 60
              Master_Log_File: master-bin.000009
          Read_Master_Log_Pos: 633
               Relay_Log_File: relay-log.000001
                Relay_Log_Pos: 4
        Relay_Master_Log_File: master-bin.000009
             Slave_IO_Running: Connecting
            Slave_SQL_Running: Yes
              Replicate_Do_DB: 
          Replicate_Ignore_DB: 
           Replicate_Do_Table: 
       Replicate_Ignore_Table: 
      Replicate_Wild_Do_Table: 
  Replicate_Wild_Ignore_Table: 
                   Last_Errno: 0
                   Last_Error: 
                 Skip_Counter: 0
          Exec_Master_Log_Pos: 633
              Relay_Log_Space: 154
              Until_Condition: None
               Until_Log_File: 
                Until_Log_Pos: 0
           Master_SSL_Allowed: No
           Master_SSL_CA_File: 
           Master_SSL_CA_Path: 
              Master_SSL_Cert: 
            Master_SSL_Cipher: 
               Master_SSL_Key: 
        Seconds_Behind_Master: NULL
Master_SSL_Verify_Server_Cert: No
                Last_IO_Errno: 0
                Last_IO_Error: 
               Last_SQL_Errno: 0
               Last_SQL_Error: 
  Replicate_Ignore_Server_Ids: 
             Master_Server_Id: 0
                  Master_UUID: 
             Master_Info_File: /mydata/data/master.info
                    SQL_Delay: 0
          SQL_Remaining_Delay: NULL
      Slave_SQL_Running_State: Slave has read all relay log; waiting for more updates
           Master_Retry_Count: 86400
                  Master_Bind: 
      Last_IO_Error_Timestamp: 
     Last_SQL_Error_Timestamp: 
               Master_SSL_Crl: 
           Master_SSL_Crlpath: 
           Retrieved_Gtid_Set: 
            Executed_Gtid_Set: 
                Auto_Position: 0
         Replicate_Rewrite_DB: 
                 Channel_Name: 
           Master_TLS_Version: 
1 row in set (0.00 sec)
```
主节点
```
mysql> CREATE DATABASE mydb;
mysql> show databases;
mysql> show master status;
```
从节点
```
show databases;
SHOW SLAVE STATUS\G
```
排错
```
# 防火墙是否影响了通信
 systemctl status firewalld.service
 systemctl stop firewalld.service
 systemctl disable firewalld.service

# 主节点授权问题
 GRANT REPLICATION SLAVE,REPLICATION CLIENT ON *.* TO 'repluser'@'192.168.%.%' IDENTIFIED BY 'replpass';
 select * from mysql.user where user='repluser'\G
 select host,user,authentication_string from mysql.user;

# 维护
# 主节点日志加一
 flush logs;
# 状态查看
 show slave status\G;
 show master status\G;
# 启停
 mysql>stop slave;
 mysql>start slave;

#格式化
mysql>change master to master_host = '192.168.11.220',
master_user = 'repluser',
master_password = 'replpass',
master_port = 3306,
master_log_file = 'master-bin.000001',
master_log_pos = 120;
```

# 8.复制架构中应该注意的问题

## 1.从节点不可写，从节点一旦写入，主从节点不一致

## 2.限制从服务器为只读

- 在从服务器上设置read_only=ON;
- 此限制对拥有SUPER权限的用户均无效；

```
SHOW GLOBEAL VARIABLES LIKE 'read_only'
```
    
- 阻止所有用户：
  
```
mysql> FLUSH TABLES WITH READ LOCK;   #用一个长会话，永不断开连接
```

## 3.如何保证主从复制的事务安全？

- 在master节点启用参数：
  - sync_binlog=ON
- 如果是InnoDB存储引擎
  - innodb_flush_logs_at_trx_commit=ON  事务立即刷入日志
  - innodb_support_xa=ON     分布式事务
- 在slave节点
  - skip_slave_start= ON

- master 节点：
  - sync_master_info
- slave 节点
  - sync_relay_log
  - sync_relay_log_info


# 主主复制
https://v.youku.com/v_show/id_XMzQ4MjY4NzA0NA==.html?spm=a2hzp.8253869.0.0
双主节点，可能造成数据不一致，oracle可以解决这个问题，但mysql不行
https://v.youku.com/v_show/id_XMzQ4MjU4NDA2MA==.html?spm=a2hzp.8253869.0.0

https://www.cnblogs.com/clsn/p/8150036.html