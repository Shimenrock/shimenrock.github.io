---
title: "Elasticsearch 7 安装与配置"
permalink: /elk/elasticsearch-7-setup/
excerpt: "How to quickly install and setup Elasticsearch 7."
last_modified_at: 2020-01-26T21:36:11-04:00
redirect_from:
  - /theme-setup/
toc: true
---

## Elasticsearch 7 安装与配置

### 1. 下载地址 

<u>https://www.elastic.co/cn/downloads/elasticsearch</u>

### 2. win10安装

**配置JDK环境变量**

- JAVA_HOME
  - D:\elasticsearch-7.2.0\jdk\
- Classpath:
  - .;%JAVA_HOME%\lib\dt.jar;%JAVA_HOME%\lib\tools.jar;
- PATH: 
  - ;%JAVA_HOME%\bin;%JAVA_HOME%\jre\bin

确认java版本生效 
```
C:\Windows\system32>java -version
openjdk version "12.0.1" 2019-04-16
OpenJDK Runtime Environment (build 12.0.1+12)
OpenJDK 64-Bit Server VM (build 12.0.1+12, mixed mode, sharing)
```
解压缩Elasticsearch下载的压缩包，执行即可。

### 3.Centos7.6 安装

#### Step 1)  jdk安装
下载  http://openjdk.java.net/
```
# tar -zxvf openjdk-11.0.1_linux-x64_bin.tar.gz
# mv jdk-11.0.1 /usr/local/java/
 
# which java
/usr/bin/java
# ll /usr/bin/java
lrwxrwxrwx. 1 root root 22 9月  20 10:55 /usr/bin/java -> /etc/alternatives/java
# ll /etc/alternatives/java
lrwxrwxrwx. 1 root root 29 9月  20 10:55 /etc/alternatives/java -> /usr/java/jdk-12.0.2/bin/java
# alternatives --install /usr/bin/java java /usr/local/java/jdk-11.0.1/bin/java 3
# alternatives --install /usr/bin/javac javac /usr/local/java/jdk-11.0.1/bin/javac 3
# alternatives --install /usr/bin/jar jar /usr/local/java/jdk-11.0.1/bin/jar 3

# alternatives --config java

共有 2 个提供“java”的程序。

  选项    命令
-----------------------------------------------
*+ 1           /usr/java/jdk-12.0.2/bin/java
   2           /usr/local/java/jdk-11.0.1/bin/java

按 Enter 保留当前选项[+]，或者键入选项编号：2
# java –-version
```


#### Step 2) 安装Elasticsearch

```
# yum install perl-Digest-SHA

# wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-7.3.2-darwin-x86_64.tar.gz

# wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-7.3.2-darwin-x86_64.tar.gz.sha512

# shasum -a 512 -c elasticsearch-7.3.2-darwin-x86_64.tar.gz.sha512 

# tar -xzf elasticsearch-7.3.2-darwin-x86_64.tar.gz

# cd elasticsearch-7.3.2/
```
配置 $ES_HOME
#### Step 3)  配置系统参数
 
```
# vi /etc/sysctl.conf
vm.max_map_count=655360
# sysctl -p

# vi /etc/security/limits.conf
* soft nofile 65536
* hard nofile 131072
* soft nproc 65536
* hard nproc 131072

# vi /etc/security/limits.d/20-nproc.conf
elk  soft  nproc  65536

# useradd elk
# id elk
uid=1000(elk) gid=1000(elk) 组=1000(elk)

# mkdir -pv /elk/data /elk/logs
mkdir: 已创建目录 "/elk"
mkdir: 已创建目录 "/elk/data"
# mkdir -pv /elk/logs
mkdir: 已创建目录 "/elk/logs"

# chown -R elk:elk /elk/data/ /elasticsearch-7.3.2/

# vi /elasticsearch-7.3.2/config/elasticsearch.yml 
cluster.name: my-app
node.name: cka-1
bootstrap.memory_lock: false
bootstrap.system_call_filter: false
path.data: /elk/data
path.logs: /elk/logs
network.host: 0.0.0.0
http.port: 9200
cluster.initial_master_nodes: ["cka-1"]

jvm内存修改

vi /elasticsearch-5.5.2/config/jvm.options

-Xms2g  --》修改 
-Xmx2g  --》修改 

# su - elk
$ vi ~/.bash_profile
export JAVA_HOME=/usr/local/java/jdk-11.0.1/
$ echo $JAVA_HOME
/usr/local/java/jdk-11.0.1/
```
 
#### 相关报错

- 1. seccomp unavailable 错误

解决方法：elasticsearch.yml 配置
bootstrap.memory_lock: false
bootstrap.system_call_filter: false
- 2. max file descriptors [4096] for elasticsearch process likely too low, increase to at least [65536]

解决方法：修改 /etc/security/limits.conf，配置：
hard nofile 80000
soft nofile 80000
- 3. max virtual memory areas vm.max_map_count [65530] is too low

解决方法：修改 /etc/sysctl.conf，添加 ：
vm.max_map_count = 262144
后 sysctl -p 生效
- 4. the default discovery settings are unsuitable...., last least one of [....] must be configured

解决方法：elasticsearch.yml 开启配置：
node.name: node-1
cluster.initial_master_nodes: ["node-1"]
- 5. 不能用root 运行

 org.elasticsearch.bootstrap.StartupException: java.lang.RuntimeException: can not run elasticsearch as root
- 6. **Failure running machine learning native code. This could be due to running on an unsupported OS or distribution, missing OS libraries, or a problem with the temp directory. To bypass this problem by running Elasticsearch without machine learning functionality set [xpack.ml.enabled: false]**

 增加配置 
禁用X-Pack机器学习功能
 xpack.ml.enabled: flase
xpack.graph.enabled  设置为false禁用X-Pack图形功能。
xpack.ml.enabled  设置为false禁用X-Pack机器学习功能。
xpack.monitoring.enabled 设置为false禁用X-Pack监视功能。
xpack.reporting.enabled 设置为false禁用X-Pack报告功能。
xpack.security.enabled 设置为false禁用X-Pack安全功能。
xpack.watcher.enabled 设置false为禁用观察器。

#### Step 3)  启动

 后台启动 ./elasticsearch -d
 
 启动日志 /elk/logs/my-app.log
 
 yum -y install net-tools
 
 端口查看   netstat -nltp
 
 进程查看 ps -ef | grep elastic
 
 测试 
 
curl -X GET http://localhost:9200/

curl -X GET http://192.168.11.203:9200/

防火墙配置
firewall-cmd --permanent --add-port=9200/tcp
firewall-cmd --reload
### 4. 目录结构
 

| 目录 | 配置文件 | 描述  |
| --- | --- | --- |
| bin  |  | 脚本文件 |
| config | elasticsearch.yml | 集群配置文件 |
| JDK |  |  |
| data | path.data | 数据文件 |
| lib |  | java类库 |
| logs | path.log | 日志文件 |
| modules |  | ES模块 |
| plugins |  | 插件 |



### 5. 安装插件
[elk@CKA1 bin]$ ./elasticsearch-plugin list
[elk@CKA1 bin]$ ./elasticsearch-plugin install analysis-icu

相关报错
```
$ ./elasticsearch-plugin install analysis-icu
Exception in thread "main" java.net.UnknownHostException: artifacts.elastic.co
        at java.base/java.net.AbstractPlainSocketImpl.connect(AbstractPlainSocketImpl.java:220)
        at java.base/java.net.SocksSocketImpl.connect(SocksSocketImpl.java:403)
        
   报错403 访问域名不通
```

### 6. 相关资源
博客 [https://www.elastic.co/cn/blog/](https://www.elastic.co/cn/blog/)

文档 [https://www.elastic.co/guide/index.html](https://www.elastic.co/guide/index.html)

视频 [https://www.elastic.co/cn/videos/](https://www.elastic.co/cn/videos/)

[https://www.elastic.co/guide/en/elastic-stack/current/installing-elastic-stack.html](https://www.elastic.co/guide/en/elastic-stack/current/installing-elastic-stack.html)