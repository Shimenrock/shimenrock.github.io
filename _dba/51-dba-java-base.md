---
title: "java基础"
permalink: /dba/java-base/
excerpt: "java基础"
last_modified_at: 2020-05-11T21:36:11-04:00
categories: oracle
redirect_from:
  - /theme-setup/
toc: true
---
 
编程语言类别
- 面对过程：以指令为中心
- 面对对象：以数据为中心

JAVA体系结构
- java编程语言
- java class文件格式
- java api
- java vm

三个技术流派
- J2SE 
- J2EE
- J2ME

Web Container（HTML标签要硬编码在应用程序中）
java先将程序编译成类文件，类文件再加载到jvm
- JDK
- Servlet 
- JSP  java程序翻译成为html，servlet的前端

运行条件1 java虚拟机  2 java运行环境jre  3.大多数用的是jdk 4.j2ee的环境-具备以上环境后允许servlet jsp 等代码
大型商业代码还需要EJB,JMS,JMX，javamail环境
以上整体环境叫做j2ee

BS代码环境，定义servlet规范

Web Container 提供jdk，提供servlet，提供jsp（类库）
Web Container（jsp翻译成servlet代码=jsper） 商业实现
- websphere
- weblogic
- Oc4j
- Glassfish
- JOnAS
- JBOSS
- Geronimo
 
Web Container 开源实现
- Tomcat
- jetty
- resin

JAVA 2 EE APIs
EJB(Enterprise JavaBeans): JAVA 相关的


java ：
- servlet: java 2 ee 的一个特殊类
- jsp ：java 2 ee 的一个特殊类

java 2 ee :java 2 se, servlet ,jsp ,jmx, javamail....

jsp > jasper > servlet > complie > bytecodes > jvm

tomcat : JWS(Sun) + Jserv(ASF)

tomcat : jdk+ tomcat

tomcat :server.xml
