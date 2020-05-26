---
title: "java deserialization"
permalink: /penetrate/java-deserialization/
excerpt: "java deserialization"
published: true
related: true
toc: true
toc_sticky: true
redirect_from:
  - /theme-setup/
categories: 
  - penetrate
---  

## Weblogic

- 基于JAVAEE架构的中间件
- 用于开发、集成、部署和管理大型分布式Web应用、网络应用和数据库应用的Java应用服务器
- 对业内多种标准的全面支持，包括EJB、JSP、JMS、JDBC、XML（标准通用标记语言的子集）和WML，使Web应用系统的实施更为简单。
**功能**
- 集群管理
- 部署基于J2EE 标准编写的服务器JAVA代码，包括servlet,JSP,JavaBean 和EJB。
- 使用J2EE 扩展网络服务集成分布式系统，包括用于数据库连接的JDBC、用于信息传递的JMS、用于网络目录访问的JNDI、用于分布式事务处理的 JTA 和用于电子邮件处理的JavaMail。
- 部署使用远程方法调用（RMI）的纯Java 分布式应用程序。
- 通过使用RMI—IIOP（RMI over Internet Inter-ORB Protocol)协议部署近似CORBA的分布式应用系统。
- 通过使用安全套接层（SSL）和Weblogic的内在支持为用户验证和授权，实现强大的安全性。
- 通过将多个Weblogic服务器组成一个集群提供高可用性、负载均衡和容错能力。
- 利用Java 的多平台能力在Windows NT/2000,Sun Solairs ,HP/UX 和其他Weblogic支持的操作系统上部署Weblogic服务器。
- 在任一平台上，通过使用WebLogic直观的进行基于Web 的管理和监视工具可在网络上轻松管理一个或多个WebLogic服务器。

## Jackson

**Jackson 是一个能够将java对象序列化为JSON字符串，也能够将JSON字符串反序列化为java对象的框架**

无论是序列化还是反序列化，Jackson都提供了三种方式：

1. JSON <--> Java Object 
2. JSON <--> JsonNode Tree（类似于XML的DOM树）
3. JSON <--> Json Stream （这是一个低层次的api，很强大，但是很繁琐）

## Fastjson

**Fastjson 是一个 Java 库，可以将 Java 对象转换为 JSON 格式，当然它也可以将 JSON 字符串转换为 Java 对象。**

- Fastjson 源码地址：https://github.com/alibaba/fastjson
- Fastjson 中文 Wiki：https://github.com/alibaba/fastjson/wiki/Quick-Start-CN