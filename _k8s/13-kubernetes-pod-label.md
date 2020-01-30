---
title: "Kubernetes lable"
permalink: /k8s/kubernetes-lable/
excerpt: "Kubernetes lable"
last_modified_at: 2020-01-29T21:36:11-04:00
categories: kubernetes
redirect_from:
  - /theme-setup/
toc: true
---

修改镜像中的默认应用
  command,args
    https://kubernetes.io/docs/tasks/inject-data-application/define-commmand-argument-container/
标签：(键值键名最小63个字符)
  key=value
    key：字母、数字、_、-、
    value: 可以为空，只能字母或着数字开头及结尾，中间可使用

标签选择器：
  等值关系： =,==,!=
  集合关系：
    KEY in (VALUE1,VALUE2,...)
    KEY notin (VALUE1,VALUE2,...)
    KEY
    !KEY

许多资源支持内嵌字段定义其使用的标签选择器：
  matchLables:直接给定键值
  matchExpressions:基于给定的表达式来定义使用标签选择器，{key:"KEY" ,operator:"OPERATOR", values:[VAL1,VAL2,...]}
    操作符：
      In, NotIn:values 字段的值必须为非空列表；
      Exists，NotExists：values字段的值必须为空列表；

nodeSelector <map[string]string>
  节点标签选择器。
nodeName <string>

annotations
  与label 不同的地方在于，它不能用于挑选资源对象，仅用于为对象提供“元数据”








kubectl get pods
kubectl get pods --show-labels
kubectl get pods -L app  显示拥有app标签的
kubectl get pods -l app
kubectl get pods -l app --show-labels
kubectl get pods -L app,run


kubectl label pods pod-demo release=canary
kubectl get pods -l app --show-labels
kubectl label pods pod-demo release=stable
kubectl label pods pod-demo release=stable --overwrite

kubectl get pods -l release,app
kubectl get pods -l release=stable,app=myapp
kubectl get pods -l release!=canary
kubectl get pods -l "release in (canary,beta,alpha)"
kubectl get pods -l "release notion (canary,beta,alpha)"

kubectl get nodes --show-labels
kubectl label nodes node01.magedu.com disktype=ssd

Pod的生命周期：
  状态：Pending，Running ,Failed,Succeeded,Unknown

  创建Pod
  Pod生命周期中的重要行为
    初始化容器
    容器探测
      liveness
      readiness

restartPolicy
  Always,OnFailure,Never,Default to Always
![podlive](../pic/podlive.jpg)
