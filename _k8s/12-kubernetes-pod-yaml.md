---
title: "Kubernetes yaml"
permalink: /k8s/kubernetes-yaml/
excerpt: "Kubernetes yaml"
last_modified_at: 2020-01-29T21:36:11-04:00
categories: kubernetes
redirect_from:
  - /theme-setup/
toc: true
---

# 资源清单定义

## RESTful

**RESTFUL是一种网络应用程序的设计风格和开发方式，基于HTTP，可以使用XML格式定义或JSON格式定义。**
{: .notice}

  - GET,PUT,DELETE,POST, ...
  - 对应 kubectl run,get,edit, ...
  
## 资源：对象
  - 工作负载型 workload: Pod，ReplicaSet,Deployment,StatefulSet,DaemonSet,Job,Cronjob,...
  - 服务发现及均衡：Service，Ingress，...
  - 配置与存储: Volume,CSI   兼容各种第三方存储卷，几十种。
    - ConfigMap,Secret
    - DownwardAPI
  - 集群级资源
    - Namespace，Node，Role，ClusterRole,RoleBinding,ClusterRoleBinding
  - 元数据型资源
    - HPA,PodTemplate,LimitRange

```
kubectl get pods
kubectl get pod podname -o yaml
```

## 创建资源的方法
  - apiserver仅接受JSON格式的资源定义；
  - yaml格式提供配置清单，apiserver可自动将其转化为json格式，然后再提交；

## 大部分资源的配置清单(5个字段 )
  - apiVersion:group/version  （$ kubectl api-version ）
    - $ kubectl api-versions
    - 级别1：内测版
    - 级别2：beta 公测版
    - 级别3：稳定版
  - kind:资源类别
  - metadata：元数据
    - name
    - namespace
    - lables
    - annotations
    - 每个资源的应用PATH  /api/GROUP/VERSION/namespaces/NAMESPACE/TYPE/NAME
  - spec 期望的状态，disired state
  - status 当前状态, current state , 本字段有kubernetes维护

```
kubectl explain pods
kubectl explain pods.<Object>   具有下级字段
kubectl explain pods.metadata
 pods.metadata.finalizers <[]string> 具有字串列表，数组
 pods.metadata.lables <map[string]string> 具有键值对的映射，json数组
 pods.metadata.ownerReference <[]Object> 对象列表
 -required- 必选字段
```
``` 键值对用冒号，列表用横线
mkdir manifests
vim pod-demo.yaml

apiVersin: v1
kind: Pod
metadata:
  name: pod-demo
  namespace: default
  lables:
    app: myapp
    tier: frontend
spec:
  containers:
  - name: myapp
    image: ikubernetes/myapp:v1
    ports:
    - name: http
      containerPort 80
    - name: https
      containerPort 443
  - name: busybox   边车
    image: busybox:latest
    command:
    - "/bin/sh"
    - "-c"
    - "sleep 3600"
  nodeSelector:
    disktype: ssd
```
```
kubectl create -f pod-demo.yaml
kubectl describe pods pod-demo
kubectl logs pod-demo myapp
curl 10.244.2.10
kubectl logs pod-demo myapp
kubectl logs pod-demo busybox   该容器已挂机
kubectl delete pods pod-demo
kubectl create -f pod-demo.yaml
kubectl get pods -w
kubectl exec -it pod-demo -c myapp -- /bin/sh
```

资源的清单格式：

  一级字段 ：apiVersion(group/version),kind,metadata(name,namespace,lables,annotations,....),spec,status(只读)

```
Pod资源
  spec.container <[]object>
  - name <string>
    image <string>
    imagePullPolicy <string>
      Always(注意带库影响)，Never(手动拖镜像)，IfNotPresent
```