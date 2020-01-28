---
title: "Kubernetes Concepts"
permalink: /k8s/kubernetes-concepts/
excerpt: "Kubernetes Concepts"
last_modified_at: 2020-01-28T21:36:11-04:00
redirect_from:
  - /theme-setup/
toc: true
---

**CNCF Cloud Native Definition**


- 容器编排  Kubernetes   Helm
- 容器引擎  Containerd   Rocket
- 容器镜像仓库 Notary TUF
- 容器网络  CNI
- 服务网络服务发现  CoreDNS  Linkerd  Envoy
- 容器监控运维  Prometheus   Fluentd   Jaeger  OpenTracing


**Ansible  应用编排工具**

因为Docker，应用变为了容器，Ansible等工具不再使用，需要容器编排工具

- docker compose 单台编排，缺少集群编排》docker swarm
- docker machine 迅速变成docker swarm可管理的工具
- docker三剑客
- MicroServices 微服务

AMP分层架构

单体应用，将程序的所有功能做到一个程序中，不利于扩展

微服务，将程序的每一个服务拆分开，导致微服务之间的调用关系变的极其复杂，调用者，被调用者

目前微服务架构都是构架在容器至上，利用容器技术的优势，迅速找到落地的实现方案

- CI 持续集成
- CD 持续交付  Delivery
- CD 持续部署  Deployment

容器编排工具自身并不能提供DevOps环境，而是需要掌握了容器编排工具之后，把DevOps思想构建和落地到容器编排工具之上。

Kubernetes 2014年发布 起源于Borg  go语言重新构建 1.0版本2015年7月发布

www.github.com/Kubernetes

Openshift  核心 k8s 》k8s的发行版。

特点
  - 自动装箱
  - 自我修复
  - 水平扩展
  - 服务发现和负载均衡
  - 自动发布和回滚
  - 密钥和配置管理

Kubernetes 从运维角度来讲，就是个集群
  - master/node模型
    - master :api Server ,Scheduler,Controller-Manager
    - node : kubelet,容器引擎 docker，
  - Pod  ，Label,Label Selector
    - Lable：key=value
    - Lable Selector
  - Pod:
    - 自主式Pod
    - 控制器管理Pod（不同类型的Pod来运行期待运行的方式）
      - ReplicationController 早期
      - ReplicaSet 新版本副本集
      - Deployment 专门负责管理 无状态  支持二级控制器
      - StatefulSet 有状态副本集
      - DaemonSet 每一个node上运行一个副本
      - Job，Ctonjob  运行作业
HPA : HorizontalPodAutoscaler 水平自动扩展


Pod网络，service网络（虚拟，IPVS规则），Node网络
Node》service》pod

  - 同一个Pod内的多个容器间：lo
  - 各Pod之间的通信 Overlay Network叠加网络
  - Pod与Service之间的通信
