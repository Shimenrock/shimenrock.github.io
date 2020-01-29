---
title: "Logstash 7 安装与配置"
permalink: /elk/logstash-7-setup/
excerpt: "How to quickly install and setup Logstash 7."
last_modified_at: 2020-01-26T21:36:11-04:00
categories: elk
redirect_from:
  - /theme-setup/
toc: true
---

## 08.Logstash 7 安装与配置
### 1. 下载地址 
[https://www.elastic.co/cn/downloads/logstash](https://www.elastic.co/cn/downloads/logstash)

参考文档：https://www.elastic.co/guide/en/logstash/current/index.html

安装文档：https://www.elastic.co/guide/en/logstash/current/installing-logstash.html#package-repositories

测试数据：https://grouplens.org/datasets/movielens/

### 2.Centos7.6 下安装
```
wget -c https://artifacts.elastic.co/downloads/logstash/logstash-7.4.0.tar.gz
wget -c https://artifacts.elastic.co/downloads/logstash/logstash-7.4.0.tar.gz.sha512
shasum -a 512 logstash-7.4.0.tar.gz 
tar -xzf logstash-7.4.0.tar.gz 
cd logstash-7.4.0/
chown -R elk：elk  logstash-7.4.0/

```

### 3.运行
$ ./logstash -f logstash.conf 
```
input {
  file {
    path => "/tmp/movies.csv"
    start_position => "beginning"
    sincedb_path => "/dev/null"
  }
}
filter {
  csv {
    separator => ","
    columns => ["id","content","genre"]
  }

  mutate {
    split => { "genre" => "|" }
    remove_field => ["path", "host","@timestamp","message"]
  }

  mutate {

    split => ["content", "("]
    add_field => { "title" => "%{[content][0]}"}
    add_field => { "year" => "%{[content][1]}"}
  }

  mutate {
    convert => {
      "year" => "integer"
    }
    strip => ["title"]
    remove_field => ["path", "host","@timestamp","message","content"]
  }

}
output {
   elasticsearch {
     hosts => "http://localhost:9200"
     index => "movies"
     document_id => "%{id}"
   }
  stdout {}
}
```
