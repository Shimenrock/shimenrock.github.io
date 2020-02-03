---
title: "MYSQL 应用基础"
permalink: /dba/mysql-base/
excerpt: "MYSQL 应用集成"
last_modified_at: 2020-02-01T21:36:11-04:00
categories: mysql
redirect_from:
  - /theme-setup/
toc: true
---
<!--马哥22课笔记，日期：2月2日-->

mariadb|mysql
关系模型：二维关系 row,column
数据库管理系统：DBMS
关系：Relational，Rdbms

关系数据库有六种范式，范式越高数据库冗余更小
第一范式基本要求：每一列都是不可分割的原子数据项，即无重复的域，非关系型数据库则无此要求。行的要求
第二范式首先满足第一范式，每个实例或记录可以被唯一区分，即不能有重复的行。列的要求
第三范式必须满足第二范式，一个关系中不包含已在其他关系已包含的非主关键字信息。多张表不能存储相同的非主键字段。表与表之间的关系
一般关系型数据库满足到第三范式，第五范式过于完美，太理想化，完美本身就是不完美。

因为要满足以上范式，关系型数据库经过拆表后，查询时又需要合表，导致影响性能，这是范式要求的必然结果

RDBMS:
  MYSQL:MYSQL,MariaDB，Percona-Server
  PostgreSQL:pgsql --> EnterpriseDB
  Oracle
  MSSQL