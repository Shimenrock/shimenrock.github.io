---
title:  "modify oracle database sga"
published: true
summary: "modify oracle database sga"
read_time: false
comments: false
related: false
author_profile: false
categories: 
  - dbascript
tags: 
  - sga
  - oracle
---

```

SQL> show parameter spfile

NAME                                 TYPE        VALUE
------------------------------------ ----------- ------------------------------
spfile                               string      +DATAVG/newdb/spfilenewdb.ora
                                                 
//备份参数
SQL> create pfile='/u01/app/oracle/product/11.2.0/dbhome_1/dbs/20151116.ora' from spfile;

File created.

SQL> show parameter memory;

NAME                                 TYPE        VALUE
------------------------------------ ----------- ------------------------------
hi_shared_memory_address             integer     0
memory_max_target                    big integer 0
memory_target                        big integer 0
shared_memory_address                integer     0

SQL> show parameter sga;

NAME                                 TYPE        VALUE
------------------------------------ ----------- ------------------------------
lock_sga                             boolean     FALSE
pre_page_sga                         boolean     FALSE
sga_max_size                         big integer 30G
sga_target                           big integer 30G

SQL> show parameter pga

NAME                                 TYPE        VALUE
------------------------------------ ----------- ------------------------------
pga_aggregate_target                 big integer 9G

//对于linux系统，先检查服务器内存大小和sysctl.conf配置，是否修改合理。

SQL> alter system set sga_max_size=36864M scope=spfile sid=“*”;

SQL> alter system set sga_target=36864M scope=spfile sid=“*”;

SQL> alter system set pga_aggregate_target=12288M scope=spfile sid=“*”;

//重启数据库，测试参数是否生效
SQL> shutdown immediate


```