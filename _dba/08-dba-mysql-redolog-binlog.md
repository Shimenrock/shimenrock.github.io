---
title: "redolog和binlog"
permalink: /dba/mysql-redolog-binlog/
excerpt: "redolog和binlog"
last_modified_at: 2020-02-04T21:36:11-04:00
categories: mysql
redirect_from:
  - /theme-setup/
toc: true
---
<!--
Day29！
学习内容 : binlog和redo log的写入机制
对应篇目：23 | MySQL是怎么保证数据不丢的？
保证 redo log 和 binlog 是完整的，是MySQL保证数据不丢失的关键。今天这篇文章，丁奇会重点与你分析 MySQL 写入 binlog 和 redo log 的流程。同时，你可以结合第2和第15篇文章，对crash-safe 有更清晰的理解。

Day36！
学习内容 :binlog的格式和基本机制
对应篇目：24 | MySQL是怎么保证主备一致的？
在 MySQL 的各种高可用方案上，扮演重要角色的是binlog。binlog有三种格式，分别是statement、row和mixed。理解这三种格式和一些基本机制，是我们理解MySQL读写分离的基础。
-->
# 一、binlog写入机制

事务执行过程中，先把日志写到 binlog cache，事务提交的时候，再把 binlog cache 写到 binlog 文件中。参数 binlog_cache_size 用于控制单个线程内 binlog cache 所占内存的大小。如果超过了这个参数规定的大小，就要暂存到磁盘。每个线程有自己 binlog cache，但是共用同一份 binlog 文件。

1. sync_binlog=0 的时候，表示每次提交事务都只 write，不 fsync；
2. sync_binlog=1 的时候，表示每次提交事务都会执行 fsync；
3. sync_binlog=N(N>1) 的时候，表示每次提交事务都 write，但累积 N 个事务后才 fsync。

出现 IO 瓶颈的场景里，将 sync_binlog 设置成一个比较大的值，可以提升性能。在实际的业务场景中，考虑到丢失日志量的可控性，一般不建议将这个参数设成 0，比较常见的是将其设置为 100~1000 中的某个数值。对应的风险是：如果主机发生异常重启，会丢失最近 N 个事务的 binlog 日志。

# 二、rodo log写入机制

事务在执行过程中，生成的 redo log 是要先写到 redo log buffer 的。redo log buffer 里面的内容，并不每次持久化到磁盘，如果mysql异常重启，这部分事务不提交。

**redolog三种状态**

1. 存在 redo log buffer 中，物理上是在 MySQL 进程内存中；
2. 写到磁盘 (write)，但是没有持久化（fsync)，物理上是在文件系统的 page cache 里面；
3. 持久化到磁盘，对应的是 hard disk。

**InnoDB 提供了 innodb_flush_log_at_trx_commit 参数，它有三种可能取值：**

1. 设置为 0 的时候，表示每次事务提交时都只是把 redo log 留在 redo log buffer 中 ;
2. 设置为 1 的时候，表示每次事务提交时都将 redo log 直接持久化到磁盘；
3. 设置为 2 的时候，表示每次事务提交时都只是把 redo log 写到 page cache。

InnoDB 有一个后台线程，每隔 1 秒，就会把 redo log buffer 中的日志，调用 write 写到文件系统的 page cache，然后调用 fsync 持久化到磁盘。

innodb_flush_log_at_trx_commit 设置成 1，那么 redo log 在 prepare 阶段就要持久化一次，因为有一个崩溃恢复逻辑是要依赖于 prepare 的 redo log，再加上 binlog 来恢复的。

通常我们说 MySQL 的“双 1”配置，指的就是 sync_binlog 和 innodb_flush_log_at_trx_commit 都设置成 1。也就是说，一个事务完整提交前，需要等待两次刷盘，一次是 redo log（prepare 阶段），一次是 binlog。

提升 binlog 组提交的效果:

- binlog_group_commit_sync_delay 参数，表示延迟多少微秒后才调用 fsync;
- binlog_group_commit_sync_no_delay_count 参数，表示累积多少次以后才调用 fsync

设置 binlog_group_commit_sync_delay 和 binlog_group_commit_sync_no_delay_count 参数，减少 binlog 的写盘次数。这个方法是基于“额外的故意等待”来实现的，因此可能会增加语句的响应时间，但没有丢失数据的风险。将 sync_binlog 设置为大于 1 的值（比较常见是 100~1000）。这样做的风险是，主机掉电时会丢 binlog 日志。将 innodb_flush_log_at_trx_commit 设置为 2。这样做的风险是，主机掉电的时候会丢数据。

不建议把 innodb_flush_log_at_trx_commit 设置成 0。因为把这个参数设置成 0，表示 redo log 只保存在内存中，这样的话 MySQL 本身异常重启也会丢数据，风险太大。而 redo log 写到文件系统的 page cache 的速度也是很快的，所以将这个参数设置成 2 跟设置成 0 其实性能差不多，但这样做 MySQL 异常重启时就不会丢数据了.


# 三、主备基本原理

备库为readonly:

1.  备库如果负责查询，设置只读防止误操作。
2.  防止切换逻辑有bug，出现双写，造成主备不一致。
3.  用readonly状态判断节点角色

readonly设置对super权限无效，更新线程拥有超级权限。

# 二、binlog三种格式

- statement
  - 忠实记录语句，但备库执行SQL时，可能因为索引等问题，unsafe
- row
  - 记录真实发生变化行的主键ID，避免主备库不同的问题。
- mixed

1. statement格式binlog可能造成主备不一致
2. row格式占空间，当删掉10万行数据时，10万行的记录都要写入binlog，消耗IO资源。
3. mixed判断SQL语句是否引起主备不一致，有可能用row，否则用statement。

row格式有利于恢复数据，MariaDB的Flashback工具利用此回滚数据。

binlog恢复数据：用mysqlbinlog工具进行解析，把解析的整个结果发给mysql执行。


mysqlbinlog master.000001  --start-position=2738 --stop-position=2973 | mysql -h127.0.0.1 -P13000 -u$user -p$pwd;
将 master.000001 文件里面从第 2738 字节到第 2973 字节中间这段内容解析出来，放到 MySQL 去执行。