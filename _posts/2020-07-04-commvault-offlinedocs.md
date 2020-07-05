---
title: "如何安装Commvault离线文档"
published: true
related: true
header:
  teaser: /assets/images/unsplash-image-3.jpg
toc: true
toc_sticky: true
categories: 
  - Commvault
---  

<!--
http://docs.commvault.com/commvault/v11/article?p=49971.htm
-->


# 如何安装Commvault离线文档

Commvault在线文档在国内浏览经常不畅，而且客户现场上网也不是很方便，这里记录了在本地电脑（WIN7）浏览离线文档的方法。

## 1. 下载Commvault离线文档

Commvault官方文档：http://docs.commvault.com/

本次测试的是v11_sp19的离线文档

离线文档下载地址：http://docs.commvault.com/commvault/v11/relatedLinks?action=download&package_name=OfflineDocs.zip


建议使用Free Download Manager下载，有条件最好使用vpn线路加速，迅雷下载会失败。

这里提供了我的云盘下载

链接：https://pan.baidu.com/s/1r5j8pNnpwUDzJ57xxz_Bvg

提取码：uycj


## 2. 下载并装好java sdk

官网要求必须使用大于1.7的jdk，我使用的是1.8的版本，安装以后配置三个系统环境变量。

JAVA_HOME 值"C:\Program Files\Java\jdk1.8.0_161"

Path      在值的后面添加";%JAVA_HOME%\bin"

CLASSPATH  值".;%Java_Home%\bin;%Java_Home%\lib\dt.jar;%Java_Home%\lib\tools.jar"

完成后在命令行里测试
```
D:\Downloads\OfflineDocs\bin> java -version
java version "1.8.0_161"
Java(TM) SE Runtime Environment (build 1.8.0_161-b12)
Java HotSpot(TM) 64-Bit Server VM (build 25.161-b12, mixed mode)
```

## 3. 将离线文档注册为系统服务并运行

将离线文档解压到没有中文的目录，我这里解压到了 “D:\Downloads\OfflineDocs”

运行
```
startService.bat -instance_name Instance001 -service_name "CommvaultOfflineDocs" -service_display_name "Commvault Offline Docs v11" -start_params "start;-port;8080;-cv.solr.jetty.deployall;true;-sysprop.hosted.mode;offline;-sysprop.solr.url;http://localhost:8080/solr"
```
```
C:\Windows\system32>d:

D:\>cd Downloads\OfflineDocs\bin

D:\Downloads\OfflineDocs\bin>startService.bat -instance_name Instance001 -servic
e_name "CommvaultOfflineDocs" -service_display_name "Commvault Offline Docs v11"
 -start_params "start;-port;8080;-cv.solr.jetty.deployall;true;-sysprop.hosted.m
ode;offline;-sysprop.solr.url;http://localhost:8080/solr"
系统找不到指定的路径。
-instance_name = Instance001
SET "instance_name=Instance001"
系统找不到指定的路径。
-service_name = CommvaultOfflineDocs
-service_display_name = Commvault Offline Docs v11
-start_params = start;-port;8080;-cv.solr.jetty.deployall;true;-sysprop.hosted.m
ode;offline;-sysprop.solr.url;http://localhost:8080/solr
SET "start_params=start;-port;8080;-cv.solr.jetty.deployall;true;-sysprop.hosted
.mode;offline;-sysprop.solr.url;http://localhost:8080/solr"
系统找不到指定的路径。
 =
已复制         1 个文件。
created "CommvaultOfflineDocs.exe"
系统找不到指定的路径。
installEmbededJettyService.bat is called with args: -service_name "CommvaultOffl
ineDocs" -instance_name "Instance001" -service_display_name "Commvault Offline D
ocs v11" -service_description "This service belongs to CommVault and is used for
 serving docsearch project." -start_params "start;-port;8080;-cv.solr.jetty.depl
oyall;true;-sysprop.hosted.mode;offline;-sysprop.solr.url;http://localhost:8080/
solr"
系统找不到指定的路径。
-service_name = CommvaultOfflineDocs
-instance_name = Instance001
-service_display_name = Commvault Offline Docs v11
-service_description = This service belongs to CommVault and is used for serving
 docsearch project.
-start_params = start;-port;8080;-cv.solr.jetty.deployall;true;-sysprop.hosted.m
ode;offline;-sysprop.solr.url;http://localhost:8080/solr
 =
C:\Program Files\Java\jdk1.8.0_161 not empty, try to use it.
NEW_PR_JVM: C:\Program Files\Java\jdk1.8.0_161\jre\bin\server\jvm.dll
系统找不到指定的路径。
Installing "CommvaultOfflineDocs(Instance001)"
系统找不到指定的路径。
The service "CommvaultOfflineDocs(Instance001)" has been installed.
已复制         1 个文件。
created "CommvaultOfflineDocs(Instance001)w.exe"
service "CommvaultOfflineDocs(Instance001)" doesn't exist and is installed.
"CommvaultOfflineDocs(Instance001)" is started.
2020-07-05_16:56:49.72 startService completes.
```

看到如下，说明运行成功
```
"CommvaultOfflineDocs(Instance001)" is started.
Fri 03/28/2014_20:13:53.58 startService completes.
```

在windows系统服务里查看，Commvault离线文档已经注册为了服务，平时可以在服务里启停文档

![commvault](/assets/images/2020-07-05-commvault-offinedocs.jpg)


## 4. 访问

访问地址 http://hostname:port/v11_servicepacknumber

这里下载的是v11_sp19版本离线文档，所以访问如下地址：http://localhost:8080/v11_sp19
注意：只能浏览文档，不能搜索，搜索功能需要授权支持。

![commvault2](/assets/images/2020-07-05-commvault-offinedocs2.jpg)

## 5. 卸载服务

```
uninstallEmbededJettyService.bat -instance_name Instance001 -service_name "CommvaultOfflineDocs"
```
