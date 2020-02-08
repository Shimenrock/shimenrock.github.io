---
title: "间隙锁和next-key lock"
permalink: /dba/next-key-lock/
excerpt: "间隙锁和next-key lock"
last_modified_at: 2020-02-06T21:36:11-04:00
categories: mysql
redirect_from:
  - /theme-setup/
toc: true
---
<!-- 
重学MySQL打卡行动Day17！

学习内容 : 间隙锁和next-key lock	
对应篇目：
20 | 幻读是什么，幻读有什么问题？http://gk.link/a/101Mp

今天这篇文章，会为你讲述关于幻读的两大知识点：

1. 幻读需要注意两点：一是，在“当前读”下才会出现；二是，仅专指“新插入的行”。
2. 引入间隙锁和next-key lock，可以解决幻读问题，但也会带来并发度的问题。
-->

```
CREATE TABLE `t` (
  `id` int(11) NOT NULL,
  `c` int(11) DEFAULT NULL,
  `d` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `c` (`c`)
) ENGINE=InnoDB;

insert into t values(0,0,0),(5,5,5),(10,10,10),(15,15,15),(20,20,20),(25,25,25);
```
