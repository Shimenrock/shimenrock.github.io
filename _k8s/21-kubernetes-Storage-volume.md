---
title: "Kubernetes Storage"
permalink: /k8s/kubernetes-storage/
excerpt: "Kubernetes lable"
last_modified_at: 2020-01-30T21:36:11-04:00
categories: kubernetes
redirect_from:
  - /theme-setup/
toc: true
---
<!--12课笔记 1小时20分钟-->

**基础架构容器：pause**{: .notice}

# k8s之上可用存储卷
  - emptyDir
    - 临时、缓存、空目录
  - hostPath
    - 主机目录
  - 传统存储设备
    - SAN:iSCSI FC
    - NAS:nfs ,cifs 
  - 分布式存储
    - glusterfs
    - rbd 块级存储
    - cephfs
  - 云存储
    - 亚马逊EBS 弹性块存储
    - Azure Disk

 **查看支持的存储**
```
# kubectl explain pods.spec.volumes
```

**根据性能指标定义存储类**
- Gold Storage Class
- Silver Storage Class
- Bronze Storage Class
  
PVC动态供给

```
# kubectl explain pods.spec.volumes.emptyDir
# kubectl explain pods.spec.containers.volumeMounts
```
**pod-vol-demo.yaml**
```yaml
apiVersion: V1
kind: Pod
metadata:
  name: pod-demo
  namespace: default
  labels:
    app: myapp
    tier: frontend
  annotations:
    magedu.com/create-by: "cluster admin"
spec:
  containers:
  - name: myapp
    image: ikubernetes/myapp:v1
    ports:
    - name: http
      containerPort: 80
    volumeMounts:
    - name: html
      mountPath: /data/web/html/
  - name: busybox
    image: busybox:latest
    imagePullPolicy: IfNotPresent
    volumeMounts:
    - name: html
      mountPath: /data/
    command:
    - "/bin/sh"
    - "-c"
    - "speep 7200"
  volumes:
  - name: html
    emptyDir: {}   #本地磁盘大小不限制
```
```
# kubectl apply -f pod-vol-demo.yaml
# kubectl exec -it pod-demo -c busybox -- /bin/sh
# ls
# mount
# echo $(date) >> /data/index.html

# kubectl exec -it pod-demo -c myapp -- /bin/sh
# cat /date/web/html/index.html

# kubectl delete -f pod-vol-demo.yaml
```

```yaml
apiVersion: V1
kind: Pod
metadata:
  name: pod-demo
  namespace: default
  labels:
    app: myapp
    tier: frontend
  annotations:
    magedu.com/create-by: "cluster admin"
spec:
  containers:
  - name: myapp
    image: ikubernetes/myapp:v1
    imagePullPolicy: IfNotPresent
    ports:
    - name: http
      containerPort: 80
    volumeMounts:
    - name: html
      mountPath: /usr/share/nginx/html/
  - name: busybox
    image: busybox:latest
    imagePullPolicy: IfNotPresent
    volumeMounts:
    - name: html
      mountPath: /data/
    command: ["/bin/sh"]
    arg: ["-c", "while true; do echo $(date) >> /data/index.html; sleep 2; done"]
  volumes:
  - name: html
    emptyDir: {}   #本地磁盘大小不限制
```
```
kubectl app -f pod-vol-demo.yaml
kubectl get pods -o wide
curl IP
```

## gitRepo
  - emptyDir
  - 将git仓库内容克隆到本地，后运行
## hostpath
  - kubectl explain pods.spec.volumes.hostPath.type
  - https://kubernetes.io/docs/concepts/storage/volumes#hostpath

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: pod-vol-hostpath
  namespace: default
spec:
  containers:
  - name: myapp
    image: ikubernetes/myapp:v1
    volumeMounts:
    - name: html
      mountPath: /usr/share/nginx/html/
  volumes:
  - name: html
    hostPath:
      path: /data/pod/volume1
      type: DirectoryOrCreate
```
```
每个节点
mkdir /data/pod/volume1 -p
vim /data/pod/volume1/index.html

kubectl apply -f pod-hostpath-vol.yaml
kubectl get pods -o wide
```

yum -y install nfs-utils
mkdir /data/volumes -pv
vim /etc/exports
/data/volumes 172.20.0.0/16(rw,no_root_squash)
systemctl start nfs
ss -tnl 

mount -t nfs stor01:/data/volumes

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: pod-vol-nfs
  namespace: default
spec:
  containers:
  - name: myapp
    image: ikubernetes/myapp:v1
    volumeMounts:
    - name: html
      mountPath: /usr/share/nginx/html/
  volumes:
  - name: html
    nfs:
      path: /data/volumes
      server: stor01.magedu.com
```








===============================================
ceph  （glusterfs、moosefs、hdfs、fastdfs、等）
  - 文件存储
  - 块设备
  - 对象存储
应用
  - owncloud （php）
  - nextcloud
  - seafile
  - 等
