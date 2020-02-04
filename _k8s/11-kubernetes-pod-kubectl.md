---
title: "Kubernetes kubectl"
permalink: /k8s/kubernetes-kubectl/
excerpt: "Kubernetes Kubernetes kubectl"
last_modified_at: 2020-01-28T21:36:11-04:00
categories: kubernetes
redirect_from:
  - /theme-setup/
toc: true
---
[Kubernetes 中文文档](http://docs.kubernetes.org.cn/)

[Kubernetes 中文社区文档](https://www.kubernetes.org.cn/%E6%96%87%E6%A1%A3%E4%B8%8B%E8%BD%BD)

[Kubernetes 中文社区教程](https://www.kubernetes.org.cn/course)

# 1.kubectl 命令查看帮助
 
```
# kubectl
```

# 2.kubectl 查看节点详情
```
# kubectl get nodes
# kubectl describe node k8s-node1-206
```

# 3.kubectl 查看版本

```
# kubectl version

```

# 4.kubectl 查看集群信息

```
# kubectl cluster-info
```

# 5.kubectl 运行pod

**注：pod的IP属于cni0**

## 5.1.kubectl run 查看帮助
```
# kubectl run --help
```

## 5.2.kubectl 运行一个pod
```
# kubectl run nginx-deploy --image=nginx:1.14-alpine --port=80 --replicas=1 --dry-run=true  干跑模式，并没有运行

# kubectl run nginx-deploy --image=nginx:1.14-alpine --port=80 --replicas=1

# kubectl get pods

# kubectl get deployment

# kubectl get deployment -o wide

# kubectl get pods -o wide

# curl 10.244.1.2
```

## 5.3.kubectl删除pod

```
# kubectl get pod

# kubectl delete pods nginx-deploy-66ff98548d-xfsd9

# kubectl get pod

# kubectl get pod -o wide

正确删除
# kubectl get deployment

# kubectl delete deployment nginx-deploy
```

**自动重建pod，但pod的IP发生变化**

## 5.4.kubectl 创建service服务

**给pod一个固定端点，创建service服务，生成IP依然为集群内部IP，解析依靠CoreDNS**
```
# kubectl expose --help

# kubectl expose deployment nginx-deploy --port=80 --target-port=80 --protocol=TCP

# kubectl get svc

# kubectl get service
```
## 5.5.DNS解析
```
# yum install bind-utils -y

# kubectl get svc -n kube-system
创建pod客户端，进入pod
# kubectl run client --image=busybox --replicas=1 -it --restart=Never     
If you don't see a command prompt, try pressing enter.
/ #
/ # cat /etc/resolv.conf
nameserver 10.96.0.10
search default.svc.cluster.local svc.cluster.local cluster.local
options ndots:5
/ # nslookup nginx-deploy
Server:         10.96.0.10
Address:        10.96.0.10:53

** server can't find nginx-deploy.default.svc.cluster.local: NXDOMAIN

*** Can't find nginx-deploy.svc.cluster.local: No answer
*** Can't find nginx-deploy.cluster.local: No answer
*** Can't find nginx-deploy.default.svc.cluster.local: No answer
*** Can't find nginx-deploy.svc.cluster.local: No answer
*** Can't find nginx-deploy.cluster.local: No answer
/ # wget nginx-deploy
Connecting to nginx-deploy (10.96.218.77:80)
saving to 'index.html'
index.html           100% |*****************************************************************************************************|   612  0:00:00 ETA
'index.html' saved
/ #  wget -O - -q http://nginx-deploy:80
```
```
    # dig -t A nginx-deploy.default.svc.cluster.local @10.96.0.10    在主机上可以解析

    ; <<>> DiG 9.11.4-P2-RedHat-9.11.4-9.P2.el7 <<>> -t A nginx-deploy.default.svc.cluster.local @10.96.0.10
    ;; global options: +cmd
    ;; Got answer:
    ;; WARNING: .local is reserved for Multicast DNS
    ;; You are currently testing what happens when an mDNS query is leaked to DNS
    ;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 57828
    ;; flags: qr aa rd; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1
    ;; WARNING: recursion requested but not available

    ;; OPT PSEUDOSECTION:
    ; EDNS: version: 0, flags:; udp: 4096
    ;; QUESTION SECTION:
    ;nginx-deploy.default.svc.cluster.local.        IN A

    ;; ANSWER SECTION:
    nginx-deploy.default.svc.cluster.local. 30 IN A 10.96.218.77

    ;; Query time: 0 msec
    ;; SERVER: 10.96.0.10#53(10.96.0.10)
    ;; WHEN: 二 12月 17 15:45:56 CST 2019
    ;; MSG SIZE  rcvd: 121
```

# sevice 

**即IPVS规则，关联到pod后端**
```
    # kubectl describe svc nginx
    Name:              nginx-deploy
    Namespace:         default
    Labels:            run=nginx-deploy
    Annotations:       <none>
    Selector:          run=nginx-deploy
    Type:              ClusterIP
    IP:                10.96.218.77
    Port:              <unset>  80/TCP
    TargetPort:        80/TCP
    Endpoints:         10.244.2.2:80
    Session Affinity:  None
    Events:            <none>
    # kubectl get pods --show-labels
    NAME                            READY   STATUS    RESTARTS   AGE     LABELS
    client                          1/1     Running   0          82m     run=client
    nginx-deploy-66ff98548d-v57lq   1/1     Running   0          4h59m   pod-template-hash=66ff98548d,run=nginx-deploy
    # kubectl edit sv nginx-deploy    可编辑IP
    # kubectl edit sv nginx-deploy
    # kubectl expose deployment nginx-deploy --name=nginx
    # kubectl get svc
    # kubectl describe deployment nginx-deploy
    Name:                   nginx-deploy
    Namespace:              default
    CreationTimestamp:      Tue, 17 Dec 2019 10:35:49 +0800
    Labels:                 run=nginx-deploy
    Annotations:            deployment.kubernetes.io/revision: 1
    Selector:               run=nginx-deploy
    Replicas:               1 desired | 1 updated | 1 total | 1 available | 0 unavailable
    StrategyType:           RollingUpdate
    MinReadySeconds:        0
    RollingUpdateStrategy:  25% max unavailable, 25% max surge
    Pod Template:
      Labels:  run=nginx-deploy
      Containers:
       nginx-deploy:
        Image:        nginx:1.14-alpine
        Port:         80/TCP
        Host Port:    0/TCP
        Environment:  <none>
        Mounts:       <none>
      Volumes:        <none>
    Conditions:
      Type           Status  Reason
      ----           ------  ------
      Progressing    True    NewReplicaSetAvailable
      Available      True    MinimumReplicasAvailable
    OldReplicaSets:  <none>
    NewReplicaSet:   nginx-deploy-66ff98548d (1/1 replicas created)
    Events:          <none>
```

# deployment

**控制器副本可动态调整**
```
    # kubectl run myapp --image=ikubernetes/myapp:v1 --replicas=2
    kubectl run --generator=deployment/apps.v1 is DEPRECATED and will be removed in a future version. Use kubectl run --generator=run-pod/v1 or kubectl create instead.
    deployment.apps/myapp created
    # kubectl get deployment -w
    NAME           READY   UP-TO-DATE   AVAILABLE   AGE
    myapp          2/2     2            2           77s
    nginx-deploy   1/1     1            1           5h28m
    # kubectl get pods -o wide
    NAME                            READY   STATUS    RESTARTS   AGE     IP           NODE            NOMINATED NODE   READINESS GATES
    client                          1/1     Running   0          94m     10.244.1.3   k8s-node1-206   <none>           <none>
    myapp-7c468db58f-cx6zk          1/1     Running   0          109s    10.244.2.3   k8s-node2-207   <none>           <none>
    myapp-7c468db58f-w7xrs          1/1     Running   0          109s    10.244.1.4   k8s-node1-206   <none>           <none>
    nginx-deploy-66ff98548d-v57lq   1/1     Running   0          5h10m   10.244.2.2   k8s-node2-207   <none>           <none>

    客户端访问
    / # wget -O - -q 10.244.1.4
    Hello MyApp | Version: v1 | <a href="hostname.html">Pod Name</a>
    / # wget -O - -q 10.244.1.4/hostname.html
    myapp-7c468db58f-w7xrs

    # kubectl expose deployment myapp --name=myapp --port=80
    service/myapp exposed
    # kubectl get svc
    NAME           TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
    kubernetes     ClusterIP   10.96.0.1      <none>        443/TCP   26h
    myapp          ClusterIP   10.96.117.44   <none>        80/TCP    17s
    nginx-deploy   ClusterIP   10.96.218.77   <none>        80/TCP    5h2m
    在两个pod间调度
    / # wget -O - -q 10.244.1.4/hostname.html
    myapp-7c468db58f-w7xrs
    / # wget -O - -q myapp
    Hello MyApp | Version: v1 | <a href="hostname.html">Pod Name</a>
    / # wget -O - -q myapp/hostname.html
    myapp-7c468db58f-w7xrs
    / # wget -O - -q myapp/hostname.html
    myapp-7c468db58f-cx6zk
    / # wget -O - -q myapp/hostname.html
    myapp-7c468db58f-cx6zk
    / # wget -O - -q myapp/hostname.html
    myapp-7c468db58f-cx6zk
    / # wget -O - -q myapp/hostname.html
    myapp-7c468db58f-w7xrs
    / # while true; do wget -O - -q myapp/hostname.html; sleep 1 ; done

    # kubectl scale --help
    Set a new size for a Deployment, ReplicaSet, Replication Controller, or StatefulSet.

     Scale also allows users to specify one or more preconditions for the scale action.

     If --current-replicas or --resource-version is specified, it is validated before the scale is attempted, and it is
    guaranteed that the precondition holds true when the scale is sent to the server.

    Examples:
      # Scale a replicaset named 'foo' to 3.
      kubectl scale --replicas=3 rs/foo

      # Scale a resource identified by type and name specified in "foo.yaml" to 3.
      kubectl scale --replicas=3 -f foo.yaml

      # If the deployment named mysql's current size is 2, scale mysql to 3.
      kubectl scale --current-replicas=2 --replicas=3 deployment/mysql

      # Scale multiple replication controllers.
      kubectl scale --replicas=5 rc/foo rc/bar rc/baz

      # Scale statefulset named 'web' to 3.
      kubectl scale --replicas=3 statefulset/web

    Options:
          --all=false: Select all resources in the namespace of the specified resource types
          --allow-missing-template-keys=true: If true, ignore any errors in templates when a field or map key is missing in
    the template. Only applies to golang and jsonpath output formats.
          --current-replicas=-1: Precondition for current size. Requires that the current size of the resource match this
    value in order to scale.
      -f, --filename=[]: Filename, directory, or URL to files identifying the resource to set a new size
      -k, --kustomize='': Process the kustomization directory. This flag can't be used together with -f or -R.
      -o, --output='': Output format. One of:
    json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-file.
          --record=false: Record current kubectl command in the resource annotation. If set to false, do not record the
    command. If set to true, record the command. If not set, default to updating the existing annotation value only if one
    already exists.
      -R, --recursive=false: Process the directory used in -f, --filename recursively. Useful when you want to manage
    related manifests organized within the same directory.
          --replicas=0: The new desired number of replicas. Required.
          --resource-version='': Precondition for resource version. Requires that the current resource version match this
    value in order to scale.
      -l, --selector='': Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)
          --template='': Template string or path to template file to use when -o=go-template, -o=go-template-file. The
    template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].
          --timeout=0s: The length of time to wait before giving up on a scale operation, zero means don't wait. Any other
    values should contain a corresponding time unit (e.g. 1s, 2m, 3h).

    Usage:
      kubectl scale [--resource-version=version] [--current-replicas=count] --replicas=COUNT (-f FILENAME | TYPE NAME)
    [options]

    Use "kubectl options" for a list of global command-line options (applies to all commands).

    # kubectl scale --replicas=5 deployment myapp
    deployment.apps/myapp scaled
    # kubectl get pods
    NAME                            READY   STATUS    RESTARTS   AGE
    client                          1/1     Running   0          103m
    myapp-7c468db58f-2tgtp          1/1     Running   0          33s
    myapp-7c468db58f-cx6zk          1/1     Running   0          10m
    myapp-7c468db58f-gldfq          1/1     Running   0          33s
    myapp-7c468db58f-w7xrs          1/1     Running   0          10m
    myapp-7c468db58f-xmgnz          1/1     Running   0          33s
    nginx-deploy-66ff98548d-v57lq   1/1     Running   0          5h19m
    # kubectl scale --replicas=3 deployment myapp
    deployment.apps/myapp scaled
    # kubectl get pods
    NAME                            READY   STATUS    RESTARTS   AGE
    client                          1/1     Running   0          103m
    myapp-7c468db58f-cx6zk          1/1     Running   0          10m
    myapp-7c468db58f-w7xrs          1/1     Running   0          10m
    myapp-7c468db58f-xmgnz          1/1     Running   0          48s
    nginx-deploy-66ff98548d-v57lq   1/1     Running   0          5h20m
```

# 动态升级

**(一个一个灰度升级)，回滚**

```
    / # while true; do wget -O - -q myapp; sleep 1 ; done

    # kubectl set image --help
    Update existing container image(s) of resources.

     Possible resources include (case insensitive):

      pod (po), replicationcontroller (rc), deployment (deploy), daemonset (ds), replicaset (rs)

    Examples:
      # Set a deployment's nginx container image to 'nginx:1.9.1', and its busybox container image to 'busybox'.
      kubectl set image deployment/nginx busybox=busybox nginx=nginx:1.9.1

      # Update all deployments' and rc's nginx container's image to 'nginx:1.9.1'
      kubectl set image deployments,rc nginx=nginx:1.9.1 --all

      # Update image of all containers of daemonset abc to 'nginx:1.9.1'
      kubectl set image daemonset abc *=nginx:1.9.1

      # Print result (in yaml format) of updating nginx container image from local file, without hitting the server
      kubectl set image -f path/to/file.yaml nginx=nginx:1.9.1 --local -o yaml

    Options:
          --all=false: Select all resources, including uninitialized ones, in the namespace of the specified resource types
          --allow-missing-template-keys=true: If true, ignore any errors in templates when a field or map key is missing in
    the template. Only applies to golang and jsonpath output formats.
          --dry-run=false: If true, only print the object that would be sent, without sending it.
      -f, --filename=[]: Filename, directory, or URL to files identifying the resource to get from a server.
      -k, --kustomize='': Process the kustomization directory. This flag can't be used together with -f or -R.
          --local=false: If true, set image will NOT contact api-server but run locally.
      -o, --output='': Output format. One of:
    json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-file.
          --record=false: Record current kubectl command in the resource annotation. If set to false, do not record the
    command. If set to true, record the command. If not set, default to updating the existing annotation value only if one
    already exists.
      -R, --recursive=false: Process the directory used in -f, --filename recursively. Useful when you want to manage
    related manifests organized within the same directory.
      -l, --selector='': Selector (label query) to filter on, not including uninitialized ones, supports '=', '==', and
    '!='.(e.g. -l key1=value1,key2=value2)
          --template='': Template string or path to template file to use when -o=go-template, -o=go-template-file. The
    template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].

    Usage:
      kubectl set image (-f FILENAME | TYPE NAME) CONTAINER_NAME_1=CONTAINER_IMAGE_1 ... CONTAINER_NAME_N=CONTAINER_IMAGE_N
    [options]

    Use "kubectl options" for a list of global command-line options (applies to all commands).

    #  kubectl get  pods
    NAME                            READY   STATUS    RESTARTS   AGE
    client                          1/1     Running   0          107m
    myapp-7c468db58f-cx6zk          1/1     Running   0          14m
    myapp-7c468db58f-w7xrs          1/1     Running   0          14m
    myapp-7c468db58f-xmgnz          1/1     Running   0          4m17s
    nginx-deploy-66ff98548d-v57lq   1/1     Running   0          5h23m
    # kubectl set image deployment myapp myapp=ikubernetes/myapp:v2
    deployment.apps/myapp image updated
    # kubectl rollout status deployment myapp
    deployment "myapp" successfully rolled out
    pod已经被全部重新创建
    #  kubectl get  pods
    NAME                            READY   STATUS    RESTARTS   AGE
    client                          1/1     Running   0          111m
    myapp-64758bffd4-bdkqz          1/1     Running   0          2m37s
    myapp-64758bffd4-jvjls          1/1     Running   0          2m54s
    myapp-64758bffd4-nszs5          1/1     Running   0          2m46s
    nginx-deploy-66ff98548d-v57lq   1/1     Running   0          5h27m

    # kubectl rollout --help
    Manage the rollout of a resource.

     Valid resource types include:

      *  deployments
      *  daemonsets
      *  statefulsets

    Examples:
      # Rollback to the previous deployment
      kubectl rollout undo deployment/abc

      # Check the rollout status of a daemonset
      kubectl rollout status daemonset/foo

    Available Commands:
      history     显示 rollout 历史
      pause       标记提供的 resource 为中止状态
      restart     Restart a resource
      resume      继续一个停止的 resource
      status      显示 rollout 的状态
      undo        撤销上一次的 rollout

    Usage:
      kubectl rollout SUBCOMMAND [options]

    Use "kubectl <command> --help" for more information about a given command.
    Use "kubectl options" for a list of global command-line options (applies to all commands).

    # kubectl rollout undo --help
    Rollback to a previous rollout.

    Examples:
      # Rollback to the previous deployment
      kubectl rollout undo deployment/abc

      # Rollback to daemonset revision 3
      kubectl rollout undo daemonset/abc --to-revision=3

      # Rollback to the previous deployment with dry-run
      kubectl rollout undo --dry-run=true deployment/abc

    Options:
          --allow-missing-template-keys=true: If true, ignore any errors in templates when a field or map key is missing in
    the template. Only applies to golang and jsonpath output formats.
          --dry-run=false: If true, only print the object that would be sent, without sending it.
      -f, --filename=[]: Filename, directory, or URL to files identifying the resource to get from a server.
      -k, --kustomize='': Process the kustomization directory. This flag can't be used together with -f or -R.
      -o, --output='': Output format. One of:
    json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-file.
      -R, --recursive=false: Process the directory used in -f, --filename recursively. Useful when you want to manage
    related manifests organized within the same directory.
          --template='': Template string or path to template file to use when -o=go-template, -o=go-template-file. The
    template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].
          --to-revision=0: The revision to rollback to. Default to 0 (last revision).

    Usage:
      kubectl rollout undo (TYPE NAME | TYPE/NAME) [flags] [options]

    Use "kubectl options" for a list of global command-line options (applies to all commands).

    # kubectl rollout undo deployment myapp
    deployment.apps/myapp rolled back
    # kubectl get pods
    NAME                            READY   STATUS              RESTARTS   AGE
    client                          1/1     Running             0          114m
    myapp-64758bffd4-jvjls          1/1     Running             0          5m43s
    myapp-64758bffd4-nszs5          0/1     Terminating         0          5m35s
    myapp-7c468db58f-gvhlp          1/1     Running             0          4s
    myapp-7c468db58f-sqmzh          0/1     ContainerCreating   0          1s
    myapp-7c468db58f-wb2k6          1/1     Running             0          3s
    nginx-deploy-66ff98548d-v57lq   1/1     Running             0          5h30m
    # kubectl get pods
    NAME                            READY   STATUS    RESTARTS   AGE
    client                          1/1     Running   0          114m
    myapp-7c468db58f-gvhlp          1/1     Running   0          19s
    myapp-7c468db58f-sqmzh          1/1     Running   0          16s
    myapp-7c468db58f-wb2k6          1/1     Running   0          18s
    nginx-deploy-66ff98548d-v57lq   1/1     Running   0          5h30m

```

# iptables规则
```
    # iptables --list --verbose --numeric
    Chain INPUT (policy ACCEPT 5283 packets, 849K bytes)
     pkts bytes target     prot opt in     out     source               destination
     6888 1159K KUBE-SERVICES  all  --  *      *       0.0.0.0/0            0.0.0.0/0            ctstate NEW /* kubernetes service portals */
     6888 1159K KUBE-EXTERNAL-SERVICES  all  --  *      *       0.0.0.0/0            0.0.0.0/0            ctstate NEW /* kubernetes externally-visible service portals */
    69946   40M KUBE-FIREWALL  all  --  *      *       0.0.0.0/0            0.0.0.0/0

    Chain FORWARD (policy ACCEPT 0 packets, 0 bytes)
     pkts bytes target     prot opt in     out     source               destination
    15476 1474K KUBE-FORWARD  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes forwarding rules */
     3259  232K KUBE-SERVICES  all  --  *      *       0.0.0.0/0            0.0.0.0/0            ctstate NEW /* kubernetes service portals */
     3259  232K DOCKER-USER  all  --  *      *       0.0.0.0/0            0.0.0.0/0
     3259  232K DOCKER-ISOLATION-STAGE-1  all  --  *      *       0.0.0.0/0            0.0.0.0/0
        0     0 ACCEPT     all  --  *      docker0  0.0.0.0/0            0.0.0.0/0            ctstate RELATED,ESTABLISHED
        0     0 DOCKER     all  --  *      docker0  0.0.0.0/0            0.0.0.0/0
        0     0 ACCEPT     all  --  docker0 !docker0  0.0.0.0/0            0.0.0.0/0
        0     0 ACCEPT     all  --  docker0 docker0  0.0.0.0/0            0.0.0.0/0
     3259  232K ACCEPT     all  --  *      *       10.244.0.0/16        0.0.0.0/0
        0     0 ACCEPT     all  --  *      *       0.0.0.0/0            10.244.0.0/16

    Chain OUTPUT (policy ACCEPT 5222 packets, 492K bytes)
     pkts bytes target     prot opt in     out     source               destination
     5755  670K KUBE-SERVICES  all  --  *      *       0.0.0.0/0            0.0.0.0/0            ctstate NEW /* kubernetes service portals */
    66363 5490K KUBE-FIREWALL  all  --  *      *       0.0.0.0/0            0.0.0.0/0

    Chain DOCKER (1 references)
     pkts bytes target     prot opt in     out     source               destination

    Chain DOCKER-ISOLATION-STAGE-1 (1 references)
     pkts bytes target     prot opt in     out     source               destination
        0     0 DOCKER-ISOLATION-STAGE-2  all  --  docker0 !docker0  0.0.0.0/0            0.0.0.0/0
     3259  232K RETURN     all  --  *      *       0.0.0.0/0            0.0.0.0/0

    Chain DOCKER-ISOLATION-STAGE-2 (1 references)
     pkts bytes target     prot opt in     out     source               destination
        0     0 DROP       all  --  *      docker0  0.0.0.0/0            0.0.0.0/0
        0     0 RETURN     all  --  *      *       0.0.0.0/0            0.0.0.0/0

    Chain DOCKER-USER (1 references)
     pkts bytes target     prot opt in     out     source               destination
     3259  232K RETURN     all  --  *      *       0.0.0.0/0            0.0.0.0/0

    Chain KUBE-EXTERNAL-SERVICES (1 references)
     pkts bytes target     prot opt in     out     source               destination

    Chain KUBE-FIREWALL (2 references)
     pkts bytes target     prot opt in     out     source               destination
        0     0 DROP       all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes firewall for dropping marked packets */ mark match 0x8000/0x8000

    Chain KUBE-FORWARD (1 references)
     pkts bytes target     prot opt in     out     source               destination
        0     0 DROP       all  --  *      *       0.0.0.0/0            0.0.0.0/0            ctstate INVALID
        0     0 ACCEPT     all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes forwarding rules */ mark match 0x4000/0x4000
     4369  447K ACCEPT     all  --  *      *       10.244.0.0/16        0.0.0.0/0            /* kubernetes forwarding conntrack pod source rule */ ctstate RELATED,ESTABLISHED
        0     0 ACCEPT     all  --  *      *       0.0.0.0/0            10.244.0.0/16        /* kubernetes forwarding conntrack pod destination rule */ ctstate RELATED,ESTABLISHED

    Chain KUBE-KUBELET-CANARY (0 references)
     pkts bytes target     prot opt in     out     source               destination

    Chain KUBE-PROXY-CANARY (0 references)
     pkts bytes target     prot opt in     out     source               destination

    Chain KUBE-SERVICES (3 references)
     pkts bytes target     prot opt in     out     source               destination

     # iptables --list --verbose --numeric -t nat
     Chain PREROUTING (policy ACCEPT 1229 packets, 213K bytes)
      pkts bytes target     prot opt in     out     source               destination
      7039  856K KUBE-SERVICES  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes service portals */
      3468  575K DOCKER     all  --  *      *       0.0.0.0/0            0.0.0.0/0            ADDRTYPE match dst-type LOCAL

     Chain INPUT (policy ACCEPT 1229 packets, 213K bytes)
      pkts bytes target     prot opt in     out     source               destination

     Chain OUTPUT (policy ACCEPT 1077 packets, 133K bytes)
      pkts bytes target     prot opt in     out     source               destination
      3413  393K KUBE-SERVICES  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes service portals */
         0     0 DOCKER     all  --  *      *       0.0.0.0/0           !127.0.0.0/8          ADDRTYPE match dst-type LOCAL

     Chain POSTROUTING (policy ACCEPT 2426 packets, 229K bytes)
      pkts bytes target     prot opt in     out     source               destination
      6839  638K KUBE-POSTROUTING  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes postrouting rules */
         0     0 MASQUERADE  all  --  *      !docker0  172.17.0.0/16        0.0.0.0/0
      3424  244K RETURN     all  --  *      *       10.244.0.0/16        10.244.0.0/16
         0     0 MASQUERADE  all  --  *      *       10.244.0.0/16       !224.0.0.0/4
         0     0 RETURN     all  --  *      *      !10.244.0.0/16        10.244.1.0/24
         0     0 MASQUERADE  all  --  *      *      !10.244.0.0/16        10.244.0.0/16

     Chain DOCKER (2 references)
      pkts bytes target     prot opt in     out     source               destination
         0     0 RETURN     all  --  docker0 *       0.0.0.0/0            0.0.0.0/0

     Chain KUBE-KUBELET-CANARY (0 references)
      pkts bytes target     prot opt in     out     source               destination

     Chain KUBE-MARK-DROP (0 references)
      pkts bytes target     prot opt in     out     source               destination
         0     0 MARK       all  --  *      *       0.0.0.0/0            0.0.0.0/0            MARK or 0x8000

     Chain KUBE-MARK-MASQ (17 references)
      pkts bytes target     prot opt in     out     source               destination
         0     0 MARK       all  --  *      *       0.0.0.0/0            0.0.0.0/0            MARK or 0x4000

     Chain KUBE-NODEPORTS (1 references)
      pkts bytes target     prot opt in     out     source               destination

     Chain KUBE-POSTROUTING (1 references)
      pkts bytes target     prot opt in     out     source               destination
         0     0 MASQUERADE  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes service traffic requiring SNAT */ mark match 0x4000/0x4000

     Chain KUBE-PROXY-CANARY (0 references)
      pkts bytes target     prot opt in     out     source               destination

     Chain KUBE-SEP-3TPIUCFAGKLS54CW (1 references)
      pkts bytes target     prot opt in     out     source               destination
         0     0 KUBE-MARK-MASQ  all  --  *      *       10.244.1.9           0.0.0.0/0
       146  8760 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            tcp to:10.244.1.9:80

     Chain KUBE-SEP-4W2OQVTIMNBVJMFT (1 references)
      pkts bytes target     prot opt in     out     source               destination
         0     0 KUBE-MARK-MASQ  all  --  *      *       10.244.2.2           0.0.0.0/0
         0     0 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            tcp to:10.244.2.2:80

     Chain KUBE-SEP-6PHUVOZRUX2TUUW4 (1 references)
      pkts bytes target     prot opt in     out     source               destination
         0     0 KUBE-MARK-MASQ  all  --  *      *       192.168.11.205       0.0.0.0/0
         0     0 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            tcp to:192.168.11.205:6443

     Chain KUBE-SEP-FVQSBIWR5JTECIVC (1 references)
      pkts bytes target     prot opt in     out     source               destination
         0     0 KUBE-MARK-MASQ  all  --  *      *       10.244.0.5           0.0.0.0/0
         0     0 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            tcp to:10.244.0.5:9153

     Chain KUBE-SEP-LASJGFFJP3UOS6RQ (1 references)
      pkts bytes target     prot opt in     out     source               destination
         0     0 KUBE-MARK-MASQ  all  --  *      *       10.244.0.5           0.0.0.0/0
         0     0 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            tcp to:10.244.0.5:53

     Chain KUBE-SEP-LPGSDLJ3FDW46N4W (1 references)
      pkts bytes target     prot opt in     out     source               destination
         0     0 KUBE-MARK-MASQ  all  --  *      *       10.244.0.5           0.0.0.0/0
       447 34419 DNAT       udp  --  *      *       0.0.0.0/0            0.0.0.0/0            udp to:10.244.0.5:53

     Chain KUBE-SEP-PUHFDAMRBZWCPADU (1 references)
      pkts bytes target     prot opt in     out     source               destination
         0     0 KUBE-MARK-MASQ  all  --  *      *       10.244.0.4           0.0.0.0/0
         0     0 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            tcp to:10.244.0.4:9153

     Chain KUBE-SEP-SF3LG62VAE5ALYDV (1 references)
      pkts bytes target     prot opt in     out     source               destination
         0     0 KUBE-MARK-MASQ  all  --  *      *       10.244.0.4           0.0.0.0/0
         0     0 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            tcp to:10.244.0.4:53

     Chain KUBE-SEP-VXUAG2QWVQYXGMBX (1 references)
      pkts bytes target     prot opt in     out     source               destination
         0     0 KUBE-MARK-MASQ  all  --  *      *       10.244.1.8           0.0.0.0/0
       141  8460 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            tcp to:10.244.1.8:80

     Chain KUBE-SEP-WXWGHGKZOCNYRYI7 (1 references)
      pkts bytes target     prot opt in     out     source               destination
         0     0 KUBE-MARK-MASQ  all  --  *      *       10.244.0.4           0.0.0.0/0
       452 34804 DNAT       udp  --  *      *       0.0.0.0/0            0.0.0.0/0            udp to:10.244.0.4:53

     Chain KUBE-SEP-XUXTBSRDRL6YRPAX (1 references)
      pkts bytes target     prot opt in     out     source               destination
         0     0 KUBE-MARK-MASQ  all  --  *      *       10.244.2.7           0.0.0.0/0
       163  9780 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            tcp to:10.244.2.7:80

     Chain KUBE-SERVICES (2 references)
      pkts bytes target     prot opt in     out     source               destination
         0     0 KUBE-MARK-MASQ  tcp  --  *      *      !10.244.0.0/16        10.96.0.10           /* kube-system/kube-dns:dns-tcp cluster IP */ tcp dpt:53
         0     0 KUBE-SVC-ERIFXISQEP7F7OF4  tcp  --  *      *       0.0.0.0/0            10.96.0.10           /* kube-system/kube-dns:dns-tcp cluster IP */ tcp dpt:53
         0     0 KUBE-MARK-MASQ  tcp  --  *      *      !10.244.0.0/16        10.96.0.10           /* kube-system/kube-dns:metrics cluster IP */ tcp dpt:9153
         0     0 KUBE-SVC-JD5MR3NA4I4DYORP  tcp  --  *      *       0.0.0.0/0            10.96.0.10           /* kube-system/kube-dns:metrics cluster IP */ tcp dpt:9153
         0     0 KUBE-MARK-MASQ  tcp  --  *      *      !10.244.0.0/16        10.96.218.77         /* default/nginx-deploy: cluster IP */ tcp dpt:80
         0     0 KUBE-SVC-KPPNLP6EKXQHRN5P  tcp  --  *      *       0.0.0.0/0            10.96.218.77         /* default/nginx-deploy: cluster IP */ tcp dpt:80
         0     0 KUBE-MARK-MASQ  tcp  --  *      *      !10.244.0.0/16        10.96.117.44         /* default/myapp: cluster IP */ tcp dpt:80
       450 27000 KUBE-SVC-NPJI2GAOYBRMPXVD  tcp  --  *      *       0.0.0.0/0            10.96.117.44         /* default/myapp: cluster IP */ tcp dpt:80
         0     0 KUBE-MARK-MASQ  tcp  --  *      *      !10.244.0.0/16        10.96.0.1            /* default/kubernetes:https cluster IP */ tcp dpt:443
         0     0 KUBE-SVC-NPX46M4PTMTKRN6Y  tcp  --  *      *       0.0.0.0/0            10.96.0.1            /* default/kubernetes:https cluster IP */ tcp dpt:443
         0     0 KUBE-MARK-MASQ  udp  --  *      *      !10.244.0.0/16        10.96.0.10           /* kube-system/kube-dns:dns cluster IP */ udp dpt:53
       899 69223 KUBE-SVC-TCOU7JCQXEZGVUNU  udp  --  *      *       0.0.0.0/0            10.96.0.10           /* kube-system/kube-dns:dns cluster IP */ udp dpt:53
      1225  212K KUBE-NODEPORTS  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes service nodeports; NOTE: this must be the last rule in this chain */ ADDRTYPE match dst-type LOCAL

     Chain KUBE-SVC-ERIFXISQEP7F7OF4 (1 references)
      pkts bytes target     prot opt in     out     source               destination
         0     0 KUBE-SEP-SF3LG62VAE5ALYDV  all  --  *      *       0.0.0.0/0            0.0.0.0/0            statistic mode random probability 0.50000000000
         0     0 KUBE-SEP-LASJGFFJP3UOS6RQ  all  --  *      *       0.0.0.0/0            0.0.0.0/0

     Chain KUBE-SVC-JD5MR3NA4I4DYORP (1 references)
      pkts bytes target     prot opt in     out     source               destination
         0     0 KUBE-SEP-PUHFDAMRBZWCPADU  all  --  *      *       0.0.0.0/0            0.0.0.0/0            statistic mode random probability 0.50000000000
         0     0 KUBE-SEP-FVQSBIWR5JTECIVC  all  --  *      *       0.0.0.0/0            0.0.0.0/0

     Chain KUBE-SVC-KPPNLP6EKXQHRN5P (1 references)
      pkts bytes target     prot opt in     out     source               destination
         0     0 KUBE-SEP-4W2OQVTIMNBVJMFT  all  --  *      *       0.0.0.0/0            0.0.0.0/0

     Chain KUBE-SVC-NPJI2GAOYBRMPXVD (1 references)
      pkts bytes target     prot opt in     out     source               destination
       141  8460 KUBE-SEP-VXUAG2QWVQYXGMBX  all  --  *      *       0.0.0.0/0            0.0.0.0/0            statistic mode random probability 0.33333333349
       146  8760 KUBE-SEP-3TPIUCFAGKLS54CW  all  --  *      *       0.0.0.0/0            0.0.0.0/0            statistic mode random probability 0.50000000000
       163  9780 KUBE-SEP-XUXTBSRDRL6YRPAX  all  --  *      *       0.0.0.0/0            0.0.0.0/0

     Chain KUBE-SVC-NPX46M4PTMTKRN6Y (1 references)
      pkts bytes target     prot opt in     out     source               destination
         0     0 KUBE-SEP-6PHUVOZRUX2TUUW4  all  --  *      *       0.0.0.0/0            0.0.0.0/0

     Chain KUBE-SVC-TCOU7JCQXEZGVUNU (1 references)
      pkts bytes target     prot opt in     out     source               destination
       452 34804 KUBE-SEP-WXWGHGKZOCNYRYI7  all  --  *      *       0.0.0.0/0            0.0.0.0/0            statistic mode random probability 0.50000000000
       447 34419 KUBE-SEP-LPGSDLJ3FDW46N4W  all  --  *      *       0.0.0.0/0            0.0.0.0/0
```

# 外部访问
```
# kubectl edit svc myapp
修改type为NodePort
# kubectl get svc
NAME           TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)        AGE
kubernetes     ClusterIP   10.96.0.1      <none>        443/TCP        27h
myapp          NodePort    10.96.117.44   <none>        80:31825/TCP   28m
nginx-deploy   ClusterIP   10.96.218.77   <none>        80/TCP         5h30m

外网浏览器访问
http://192.168.11.206:31825/
http://192.168.11.207:31825/
在集群之外，搭建一个高可用负载均衡器
```




