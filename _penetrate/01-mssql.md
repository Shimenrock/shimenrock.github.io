---
title: "mssql"
permalink: /penetrate/mysql/
excerpt: "mssql"
published: true
related: true
toc: true
toc_sticky: true
redirect_from:
  - /theme-setup/
categories: 
  - penetrate
---  

# 一、MSSQL

## 架构

asp+sqlserver

学校、政府、OA、游戏、棋牌、人事考试

## 默认规划

端口： 1433  
管理员： sa  
完整数据库： 数据库文件.mdf+日志文件.ldf

## 系统加固

- sa 要强密码
- 修改sa用户名
- 业务系统要用普通数据库用户，如果使用sa，被入侵后权限过大
- sysadmin角色系统权限，避免给普通用户该权限
 权限划分
 - sa权限，数据库操作，文件管理，命令执行，注册表读取等system
 - db权限，文件管理，数据库操作等users-administrators
 - public权限：数据库操作guest-users

## 爆破工具

- hydra 加载字典

## 数据库导出方法

- 下载数据库文件
- 任务》生成脚本  优点：可跨版本恢复

## 注意
 跨版本数据库文件或脚本，无法恢复到新客户端

## 调用数据库代码
<%
 set conn = server.createobject("adodb.connection")
 conn.open "provider=sqloledb;source=local;uid=sa;pwd=*****;database=database-name"
%>
Web.config

## 注入语句
1. 判断是否有注入，如果页面不正常，说明语句注入到了数据库
and 1=1
and 1=2
/ 
-0

