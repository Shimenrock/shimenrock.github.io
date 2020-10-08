---
title: "PostgreSql 应用基础"
permalink: /dba/PostgreSql-base/
excerpt: "PostgreSql 应用基础"
last_modified_at: 2020-09-21T21:36:11-04:00
categories: mostgreSql 
redirect_from:
  - /theme-setup/
toc: true
---
<!--
https://gitee.com/komavideo/LearnPostgreSql
-->
# PostgreSql 学习笔记

**官方主页：https://www.postgresql.org/**

**学习环境：**
* Ubuntu Server 16 LTS
* PostgreSql 9.5.x

## 许可
http://www.postgresql.org/about/licence

## 特性矩阵
http://www.postgresql.org/about/featurematrix

## MYSQL对比
http://bbs.chinaunix.net/thread-1688208-1-1.html

## 概述

- ORACLE、MYSQL 索引组织表、 PG 堆表（新数据向后插入，删除时标记为不可用）；没有回滚段。
- PG oracle是进程模式，mysql是线程模式（多路CPU不能利用）

## 1.简易安装方法

版本信息：https://www.postgresql.org/support/versioning/
```
$ apt-cache show postgresql
$ sudo apt-get install postgresql
$ psql --version

$ nmap 127.0.0.1
# yum install nmap
```

官网镜像 https://hub.docker.com/_/postgres
```
# docker pull postgres:9.5
# docker images 
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
postgres            9.5                 8bd68b342799        4 days ago          196MB
# docker run --name postgres -e POSTGRES_PASSWORD=202009 -p 5432:5432 -d postgres:9.5
# docker container ls
# docker stop postgres
# docker container rm postgres 
# mkdir /data
# docker run --name webdb -v /data/postgresql95:/var/lib/postgresql/data -p 5432:5432 -e POSTGRES_USER=dbuser -e POSTGRES_PASSWORD=202009 -d  postgres:9.5
# docker inspect webdb
# docker exec -it webdb /bin/bash
```
如果通过其他方式完成后持久化
```
# docker volume ls
# docker volume create --name -v_webdb_data
# docker run --name webdb -v v_webdb_data:/var/lib/postgresql/data -p 5432:5432 -e POSTGRES_USER=dbuser -e POSTGRES_PASSWORD=202009 -d 
# docker volume create --driver local \
    --opt type=btrfs \
    --opt device=/dev/sda2 \
    foo
# docker volume create --driver local \
    --opt type=nfs \
    --opt o=addr=192.168.1.1,rw \
    --opt device=:/path/to/dir \
    foo
# docker volume inspect v_webdb_data
```
```
客户端安装
```
# yum install postgresql -y
# psql -h localhost -U dbuser
```
## 2.数据库的创建
```
# su postgres
$ psql --version                版本
$ psql -l                       列出数据库
$ createdb komablog
$ psql -l
$ psql komablog
> help                          系统帮助
> \h
> \?                            语言帮助
> \l
> \q
$ psql komablog
> select now();
> select version();
> \q
$ dropdb komablog
$ psql -l
```
## 3.表的创建
```
$ su postgres
$ createdb test
$ psql -l
$ psql test
> create table posts (title varchar(255), content text);
> \dt                           display table
> \d posts                      display posts 表
> alter table posts rename to tests;
> \dt
> drop table tests;
> \dt
> \q
$ nano db.sql
...
create table posts (title varchar(255), content text);
...
$ psql test
> \i db.sql
> \dt
```
## 4.数据类型：

数据类型：
https://www.postgresql.org/docs/9.5/static/datatype.html

## 5.约束

```
create table posts (
    id serial primary key,
    title varchar(255) not null,
    content text check(length(content) > 8),
    is_draft boolean default TRUE,
    is_del boolean default FALSE,
    created_date timestamp default 'now'
);

-- 说明
/*
约束条件：

not null:不能为空
unique:在所有数据中值必须唯一
check:字段设置条件
default:字段默认值
primary key(not null, unique):主键，不能为空，且不能重复
*/

> insert into posts (title, content) values ('', '');
> insert into posts (title, content) values (NULL, '');
> insert into posts (title, content) values ('title1', 'content11');
> select * from posts;
> insert into posts (title, content) values ('title2', 'content22');
> insert into posts (title, content) values ('title3', 'content33');
> select * from posts;
```
## 6.INSERT语句
```
> insert into posts (title, content) values ('', '');
> insert into posts (title, content) values (NULL, '');
> insert into posts (title, content) values ('title1', 'content11');
> select * from posts;
> insert into posts (title, content) values ('title2', 'content22');
> insert into posts (title, content) values ('title3', 'content33');
> select * from posts;
```