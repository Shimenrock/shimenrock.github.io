---
title: "Implement Node and Pod Affinity/Anti-Affinity in Kubernetes: A Practical Example"
permalink: /k8s/kubernetes-affinity-example/
excerpt: "Implement Node and Pod Affinity/Anti-Affinity in Kubernetes: A Practical Examplet"
last_modified_at: 2020-03-3T21:36:11-04:00
categories: kubernetes
redirect_from:
  - /theme-setup/
toc: true
toc_sticky: true
---

<!--
https://thenewstack.io/implement-node-and-pod-affinity-anti-affinity-in-kubernetes-a-practical-example/
![DIY01](/assets/images/DIY01.jpg)
-->

Implement Node and Pod Affinity/Anti-Affinity in Kubernetes: A Practical Example

在Kubernetes中实现Node和Pod的亲和/反亲和：一个实际的例子

I introduced the concept of node and pod affinity/anti-affinity in last week’s tutorial. We will explore the idea further through a real-world scenario.

在上周的教程中，我介绍了Node和Pod的亲和/反亲和概念。我们将通过实际场景进一步探讨该想法。

Objective

# 目标

We are going to deploy three microservices — MySQL, Redis, and a Python/Flask web app in a four-node Kubernetes cluster. Since one of the nodes is attached to SSD disk, we want to ensure that the MySQL Pod is scheduled on the same node. Redis is used to cache the database queries to accelerate application performance. But no node will run more than one Pod of Redis. Since Redis is utilized as a cache, it doesn’t make sense to run more than one Pod per node. The next goal is to make sure that the web Pod is placed on the same node as the Redis Pod. This will ensure low latency between the web and the cache layer. Even if we scale the number of replicas of the web Pod, it will never get placed on a node that doesn’t have Redis Pod.

我们将在一个四节点的Kubernetes集群中部署三个微服务-MySQL，Redis和一个Python / Flask Web应用程序。由于其中一个Node已挂载了SSD磁盘，因此我们要确保将MySQL Pod调度到同一个Node上。Redis用于缓存数据库查询以提高应用程序性能。但是节点不会运行一个以上的Redis Pod。因为Redis被用作缓存，没必要在每个节点上运行一个以上的Pod。下一个目标是确保Web Pod与Redis Pod在同一节点上。这将确保Web和缓存层之间的低延迟。即使我们调整Web Pod的副本数量，也永远不会将其放置在没有Redis Pod的节点上。

![1](https://cdn.thenewstack.io/media/2020/01/6f288b5f-np-aff-1024x447.png)


Setting up a GKE Cluster and Adding an SSD Disk
Let’s launch a GKE cluster, add an SSD persistent disk to one of the nodes, and label the node.

# 配置一个GKE集群并添加一个SSD磁盘
让我们启动一个GKE集群，并且将一个SSD持久化存储添加到其中一个节点，并为该节点增加label标签。

```
gcloud container clusters create "tns" \
	--zone "asia-south1-a" \
	--username "admin" \
	--cluster-version "1.13.11-gke.14" \
	--machine-type "n1-standard-4" \
	--image-type "UBUNTU" \
	--disk-type "pd-ssd" \
	--disk-size "50" \
	--scopes "https://www.googleapis.com/auth/compute","https://www.googleapis.com/auth/devstorage.read_only","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" \
	--num-nodes "4" \
	--enable-stackdriver-kubernetes \
	--network "default" \
	--addons HorizontalPodAutoscaling,HttpLoadBalancing
```

This will result in a 4-node GKE cluster.
创建出一个四节点的GKE集群。

![2](https://cdn.thenewstack.io/media/2020/01/3865a2ed-np-aff-0-1024x286.png)

Let’s create a GCE Persistent Disk and attach it to the first node of the GKE cluster.
让我们创建一个GCE持久化存储并将其附加到GKE群集的第一个节点上。

```
gcloud compute disks create \
 mysql-disk-1 \
 --type pd-ssd \
 --size 20GB \
 --zone asia-south1-a
```
```
gcloud compute instances attach-disk gke-tns-default-pool-b11f5e68-2h4f \
 	--disk mysql-disk-1 \
 	--zone asia-south1-a
```

We need to mount the disk within the node to make it accessible to the applications.
我们需要将这个disk在节点内mount，以使应用程序可以访问它。
```
gcloud compute ssh gke-tns-default-pool-b11f5e68-2h4f \
--zone asia-south1-a
```
Once you SSH into the GKE node, run the below commands to mount the disk.
通过SSH进入GKE节点后，运行以下命令将磁盘mount。
```
sudo mkfs.ext4 -m 0 -F -E lazy_itable_init=0,lazy_journal_init=0,discard /dev/sdb
sudo mkdir -p /mnt/data
sudo mount -o discard,defaults /dev/sdb /mnt/data
sudo chmod a+w /mnt/data
echo UUID=`sudo blkid -s UUID -o value /dev/sdb` /mnt/data ext4 discard,defaults,nofail 0 2 | sudo 
```
Running lsblk command confirms that the disk is mounted at /mnt/data

运行lsblk命令确认磁盘已挂载在 /mnt/data

![3](https://cdn.thenewstack.io/media/2020/01/931257de-np-aff-1-1024x448.png)

Exit the shell and run the below command to label the node as disktype=ssd.

退出并运行以下命令给节点标记为“disktype=ssd”
```
kubectl label node gke-tns-default-pool-b11f5e68-2h4f \
disktype=ssd --overwrite
```
Let’s verify that the node is indeed labeled.

让我们验证该节点是否标记。
```
kubectl get nodes -l disktype=ssd
```

![4](https://cdn.thenewstack.io/media/2020/01/90570f05-np-aff-2-1024x195.png)

Deploying the Database Pod
# 部署数据库容器
Let’s go ahead and deploy a MySQL Pod targeting the above node. Use the below YAML specification to create the database Pod and expose it as a ClusterIP-based Service.

让我们继续在上述节点部署一个MySQL Pod。使用以下YAML文件创建数据库Pod，并
暴露一个基于集群IP的服务。
```
apiVersion: v1
kind: Service
metadata:
  name: mysql
  labels:
    app: mysql
spec:
  ports:
  - port: 3306
    name: mysql
    targetPort: 3306
  selector:
    app: mysql
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mysql
spec:
  selector:
    matchLabels:
      app: mysql
  template:
    metadata:
      labels:
        app: mysql
    spec:
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: disktype
                operator: In
                values:
                - ssd
      containers:
      - image: mysql:5.6
        name: mysql
        env:
        - name: MYSQL_ROOT_PASSWORD
          value: "password"
        ports:
        - containerPort: 3306
          name: mysql
        volumeMounts:
        - name: mysql-persistent-storage
          mountPath: /var/lib/mysql
      volumes:
      - name: mysql-persistent-storage
        hostPath:
          path: /mnt/data
```

There are a few things to note from the above Pod spec. We first implement node affinity by including the below clause in the spec:

在上面的Pod spec定义中需要注意一些事项。我们首先通过以下spec定义实现node的亲和：
```
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: disktype
                operator: In
                values:
                - ssd
```
This will ensure that the Pod is scheduled in the node that has the label disktype=ssd. Since we are sure that it always goes to the same node, we leverage hostPath primitive to create the Persistent Volume. The hostPath primitive has a pointer to the mount point of the SSD disk that we attached in the previous step.

这将确保这个Pod调度到打有“disktype=ssd”标签的Node中。由于我们确信它总是在同一节点，我们利用hostPath创建持久化卷。hostPath存在一个指向我们上一步挂载SSD磁盘的指针。
```
       volumeMounts:
        - name: mysql-persistent-storage
          mountPath: /var/lib/mysql
      volumes:
      - name: mysql-persistent-storage
        hostPath:
          path: /mnt/data
```
Let’s submit the Pod spec to Kubernetes and verify that it is indeed scheduled in the node that matches the label.

我们将Pod定义提交Kubernetes集群，并验证它是否在于标签匹配的节点中进行了调度。

```
kubectl apply -f db.yaml
```
```
kubectl get nodes -l disktype=ssd
```
```
kubectl get pods -o wide
```

![5](https://cdn.thenewstack.io/media/2020/01/96e685f3-np-aff-3-1024x367.png)

It’s evident that the Pod is scheduled in the node that matches the affinity rule.
显然，Pod调度到了匹配亲和规则的的节点上。 

Deploying the Cache Pod
# 部署Cache Pod
It’s time to deploy the Redis Pod that acts as the cache layer. We want to make sure that no two Redis Pods run on the same node. For that, we will define an anti-affinity rule.

现在部署作为换成层的Redis Pod。我们要确保在同一节点上没有两个Redis Pod。为此，我们定义一个反亲和规则。

The below specification creates a Redis Deployment with 3 Pods and exposes them as a ClusterIP.

以下规范将创建一个具有3个Pod的Redis部署，并将为其暴露ClusterIP。
```
apiVersion: v1
kind: Service
metadata:
  name: redis
  labels:
    app: redis
spec:
  ports:
  - port: 6379
    name: redis
    targetPort: 6379
  selector:
    app: redis
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: redis
spec:
  selector:
    matchLabels:
      app: redis
  replicas: 3
  template:
    metadata:
      labels:
        app: redis
    spec:
      affinity:
        podAntiAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
          - labelSelector:
              matchExpressions:
              - key: app
                operator: In
                values:
                - redis
            topologyKey: "kubernetes.io/hostname"
      containers:
      - name: redis-server
        image: redis:3.2-alpine
```
The below clause ensures that a node runs one and only one Redis Pod.

以下子句确保一个节点只运行一个Redis Pod。
```
 affinity:
        podAntiAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
          - labelSelector:
              matchExpressions:
              - key: app
                operator: In
                values:
                - redis
            topologyKey: "kubernetes.io/hostname"
```
Submit the Deployment spec and inspect the distribution of the pods.
提交部署定义并检查pod的分布。
```
kubectl apply -f cache.yaml
```
```
kubectl get pods -l app=redis -o wide
```

![6](https://cdn.thenewstack.io/media/2020/01/156d7348-np-aff-4-1024x367.png)

It’s clear that the Redis Pods have been placed on unique nodes.

很明显Redis Pod在每一个节点部署了一个。

Deploying the Web Pod
# 部署Web Pod
Finally, we want to place a web Pod on the same node as the Redis Pod.
最后，我们想将Web Pod与Redis Pod放在同一节点上。
Submit the Deployment spec to create 3 Pods of the web app and expose them through a Load Balancer.
提交部署定义创建三个Web Pod，并通过一个负载均衡器暴露其服务。

```
apiVersion: v1
kind: Service
metadata:
  name: web
  labels:
    app: web
spec:
  ports:
  - port: 80
    name: redis
    targetPort: 5000
  selector:
    app: web
  type: LoadBalancer    
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: web
spec:
  selector:
    matchLabels:
      app: web
  replicas: 3
  template:
    metadata:
      labels:
        app: web
    spec:
      affinity:
        podAntiAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
          - labelSelector:
              matchExpressions:
              - key: app
                operator: In
                values:
                - web
            topologyKey: "kubernetes.io/hostname"
        podAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
          - labelSelector:
              matchExpressions:
              - key: app
                operator: In
                values:
                - redis
            topologyKey: "kubernetes.io/hostname"
      containers:
      - name: web-app
        image: janakiramm/py-red
        env:       
          - name: "REDIS_HOST"
            value: "redis"
```
```
kubectl apply -f web.yaml
```

The container image used in the web app does nothing but accessing the rows in the database only after checking if they are available in the cache.

web应用程序使用的容器，仅在检查到缓存中的行可用之后，才执行其他操作。

Let’s list all the Pods along with the Node names that they are scheduled in.

我们列出调度后pod在节点分布的结果。
```
kubectl get pods -o wide | awk {'print $1" " $7'} | column -t
```

![7](https://cdn.thenewstack.io/media/2020/01/21962688-np-aff-5-1024x461.png)

We can see that the node gke-tns-default-pool-b11f5e68-2h4f runs three Pods – MySQL, Redis, and Web. The other two nodes run one Pod each for Redis and Web which are co-located for low latency.

我们可以看到，节点gke-tns-default-pool-b11f5e68-2h4f运行了三个Pod： MySQL，Redis和Web。其他两个节点则分别运行Redis和Web各一个Pod，在一起以降低延迟。

Let’s have some fun with the affinity rules. Remember, we are running 4 nodes in the cluster. One of the node is not running any Pod because of the Kubernetes scheduler obeying the rule of co-locating the Web pod and Redis Pod.

让我们来看看相似的亲和规则。记住，我们在集群中运行了4个节点。其中一个节点未运行任何Pod，这是因为Kubernetes调度程序遵循了Web Pod和Redis Pod并置的规则。

What happens when we scale the number of replicas of the Web Pod? Since the anti-affinity rule of Web Deployment imposes a rule that no two Pods of the Web can run on the same node and each Web Pod has to be paired with a Redis Pod, the scheduler wouldn’t be able to place the pod. The new web Pods will be in the pending state forever. This is despite the fact that there is an available node with no Pods running on it.

当我们扩缩容Web Pod的副本数量时会发生什么？由于Web部署的反亲和规则，即两个Web Pod不能在同一节点上运行，并且每个Web Pod必须与Redis Pod配对，因此调度程序不能增加Pod。新的Web Pod将永远处于待处理状态。尽管事实是存在可用节点，并且没有Pod在上面运行。
```
kubectl scale deploy/web --replicas=4
```

![8](https://cdn.thenewstack.io/media/2020/01/e76569db-np-aff-6-1024x547.png)

Remove the anti-affinity rule of the Web Deployment and try scaling the Replica. Now Kubernetes can schedule the Web Pods on any node that has a Redis Pod. This makes the Deployments less restrictive allowing any number of Web Pods to run on any Node provided it runs a Redis Pod.
删除Web部署的反亲和规则，然后尝试扩展副本数量。现在，Kubernetes可以在具有Redis Pod的任何节点上调度Web Pod。这使得部署的限制较少，允许在运行Redis Pod的节点上运行任意数量的Web Pod。

```
kubectl get pods -o wide | awk {'print $1" " $7'} | column -t
````

![9](https://cdn.thenewstack.io/media/2020/01/f49916d9-np-aff-7-1024x498.png)

From the above output, we see that the node gke-tns-default-pool-b11f5e68-cxvw runs two instances of the Web Pod.

从上面的输出中，我们看到节点gke-tns-default-pool-b11f5e68-cxvw运行了两个Web Pod实例。

But, one of the nodes is still lying idle due to the pod affinity/anti-affinity rules. If you want to utilize it, scale the Redis Deployment to run a Pod on the idle node and then scale the Web Deployment to place some Pods on it.

但是，由于Pod的亲和/反亲和规则，其中一个节点仍然处于空闲状态。如果要利用它，请扩容Redis部署，以便Redis在在空闲节点上运行Pod，然后扩容Web部署。

Continuing the theme of co-locating database and cache layers on the same node, in the next part of this series, we will explore the sidecar pattern to deploy low-latency microservices on Kubernetes.

延续在同一节点上的co-locating数据库和缓存层的主题，在本系列的下一部分文章中，我们将探索在Kubernetes上通过边车模式部署低延迟微服务。
