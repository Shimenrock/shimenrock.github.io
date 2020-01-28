---
title: "Kubernetes kubectl"
permalink: /k8s/kubernetes-kubectl/
excerpt: "Kubernetes Kubernetes kubectl"
last_modified_at: 2020-01-28T21:36:11-04:00
redirect_from:
  - /theme-setup/
toc: true
---
[Kubernetes 中文文档](http://docs.kubernetes.org.cn/)

[Kubernetes 中文社区文档](https://www.kubernetes.org.cn/%E6%96%87%E6%A1%A3%E4%B8%8B%E8%BD%BD)

[Kubernetes 中文社区教程](https://www.kubernetes.org.cn/course)

# kubectl 命令查看帮助
 
```
    # kubectl
    kubectl controls the Kubernetes cluster manager.

     Find more information at: https://kubernetes.io/docs/reference/kubectl/overview/

    Basic Commands (Beginner):
      create         Create a resource from a file or from stdin.
      expose         使用 replication controller, service, deployment 或者 pod 并暴露它作为一个 新的 Kubernetes Service
      run            在集群中运行一个指定的镜像
      set            为 objects 设置一个指定的特征

    Basic Commands (Intermediate):
      explain        查看资源的文档
      get            显示一个或更多 resources
      edit           在服务器上编辑一个资源
      delete         Delete resources by filenames, stdin, resources and names, or by resources and label selector

    Deploy Commands:
      rollout        Manage the rollout of a resource
      scale          Set a new size for a Deployment, ReplicaSet or Replication Controller
      autoscale      自动调整一个 Deployment, ReplicaSet, 或者 ReplicationController 的副本数量

    Cluster Management Commands:
      certificate    修改 certificate 资源.
      cluster-info   显示集群信息
      top            Display Resource (CPU/Memory/Storage) usage.
      cordon         标记 node 为 unschedulable
      uncordon       标记 node 为 schedulable
      drain          Drain node in preparation for maintenance
      taint          更新一个或者多个 node 上的 taints

    Troubleshooting and Debugging Commands:
      describe       显示一个指定 resource 或者 group 的 resources 详情
      logs           输出容器在 pod 中的日志
      attach         Attach 到一个运行中的 container
      exec           在一个 container 中执行一个命令
      port-forward   Forward one or more local ports to a pod
      proxy          运行一个 proxy 到 Kubernetes API server
      cp             复制 files 和 directories 到 containers 和从容器中复制 files 和 directories.
      auth           Inspect authorization

    Advanced Commands:
      diff           Diff live version against would-be applied version
      apply          通过文件名或标准输入流(stdin)对资源进行配置
      patch          使用 strategic merge patch 更新一个资源的 field(s)
      replace        通过 filename 或者 stdin替换一个资源
      wait           Experimental: Wait for a specific condition on one or many resources.
      convert        在不同的 API versions 转换配置文件
      kustomize      Build a kustomization target from a directory or a remote url.

    Settings Commands:
      label          更新在这个资源上的 labels
      annotate       更新一个资源的注解
      completion     Output shell completion code for the specified shell (bash or zsh)

    Other Commands:
      api-resources  Print the supported API resources on the server
      api-versions   Print the supported API versions on the server, in the form of "group/version"
      config         修改 kubeconfig 文件
      plugin         Provides utilities for interacting with plugins.
      version        输出 client 和 server 的版本信息

    Usage:
      kubectl [flags] [options]

    Use "kubectl <command> --help" for more information about a given command.
    Use "kubectl options" for a list of global command-line options (applies to all commands).
```

# kubectl 查看节点详情
```
    # kubectl describe node k8s-node1-206
    Name:               k8s-node1-206
    Roles:              <none>
    Labels:             beta.kubernetes.io/arch=amd64
                        beta.kubernetes.io/os=linux
                        kubernetes.io/arch=amd64
                        kubernetes.io/hostname=k8s-node1-206
                        kubernetes.io/os=linux
    Annotations:        flannel.alpha.coreos.com/backend-data: {"VtepMAC":"32:9c:aa:25:ca:bc"}
                        flannel.alpha.coreos.com/backend-type: vxlan
                        flannel.alpha.coreos.com/kube-subnet-manager: true
                        flannel.alpha.coreos.com/public-ip: 192.168.11.206
                        kubeadm.alpha.kubernetes.io/cri-socket: /var/run/dockershim.sock
                        node.alpha.kubernetes.io/ttl: 0
                        volumes.kubernetes.io/controller-managed-attach-detach: true
    CreationTimestamp:  Mon, 16 Dec 2019 13:48:44 +0800
    Taints:             node.kubernetes.io/unreachable:NoSchedule
    Unschedulable:      false
    Lease:
      HolderIdentity:  k8s-node1-206
      AcquireTime:     <unset>
      RenewTime:       Mon, 16 Dec 2019 23:58:45 +0800
    Conditions:
      Type             Status    LastHeartbeatTime                 LastTransitionTime                Reason              Message
      ----             ------    -----------------                 ------------------                ------              -------
      MemoryPressure   Unknown   Mon, 16 Dec 2019 23:56:40 +0800   Tue, 17 Dec 2019 09:25:31 +0800   NodeStatusUnknown   Kubelet stopped posting node status.
      DiskPressure     Unknown   Mon, 16 Dec 2019 23:56:40 +0800   Tue, 17 Dec 2019 09:25:31 +0800   NodeStatusUnknown   Kubelet stopped posting node status.
      PIDPressure      Unknown   Mon, 16 Dec 2019 23:56:40 +0800   Tue, 17 Dec 2019 09:25:31 +0800   NodeStatusUnknown   Kubelet stopped posting node status.
      Ready            Unknown   Mon, 16 Dec 2019 23:56:40 +0800   Tue, 17 Dec 2019 09:25:31 +0800   NodeStatusUnknown   Kubelet stopped posting node status.
    Addresses:
      InternalIP:  192.168.11.206
      Hostname:    k8s-node1-206
    Capacity:
      cpu:                2
      ephemeral-storage:  13706Mi
      hugepages-1Gi:      0
      hugepages-2Mi:      0
      memory:             3880732Ki
      pods:               110
    Allocatable:
      cpu:                2
      ephemeral-storage:  12934604369
      hugepages-1Gi:      0
      hugepages-2Mi:      0
      memory:             3778332Ki
      pods:               110
    System Info:
      Machine ID:                 24b72eab25494798a3387d8badc2cf72
      System UUID:                08064D56-CC48-3239-DC66-737A734B6481
      Boot ID:                    f8c3aaa7-604a-453e-a9a3-321ad58c4dfe
      Kernel Version:             3.10.0-957.el7.x86_64
      OS Image:                   CentOS Linux 7 (Core)
      Operating System:           linux
      Architecture:               amd64
      Container Runtime Version:  docker://19.3.5
      Kubelet Version:            v1.17.0
      Kube-Proxy Version:         v1.17.0
    PodCIDR:                      10.244.1.0/24
    PodCIDRs:                     10.244.1.0/24
    Non-terminated Pods:          (2 in total)
      Namespace                   Name                           CPU Requests  CPU Limits  Memory Requests  Memory Limits  AGE
      ---------                   ----                           ------------  ----------  ---------------  -------------  ---
      kube-system                 kube-flannel-ds-amd64-94nmx    100m (5%)     100m (5%)   50Mi (1%)        50Mi (1%)      20h
      kube-system                 kube-proxy-gqk4h               0 (0%)        0 (0%)      0 (0%)           0 (0%)         20h
    Allocated resources:
      (Total limits may be over 100 percent, i.e., overcommitted.)
      Resource           Requests   Limits
      --------           --------   ------
      cpu                100m (5%)  100m (5%)
      memory             50Mi (1%)  50Mi (1%)
      ephemeral-storage  0 (0%)     0 (0%)
    Events:              <none>
```

# kubectl 查看版本
```
    # kubectl version
    Client Version: version.Info{Major:"1", Minor:"17", GitVersion:"v1.17.0", GitCommit:"70132b0f130acc0bed193d9ba59dd186f0e634cf", GitTreeState:"clean", BuildDate:"2019-12-07T21:20:10Z", GoVersion:"go1.13.4", Compiler:"gc", Platform:"linux/amd64"}
    Server Version: version.Info{Major:"1", Minor:"17", GitVersion:"v1.17.0", GitCommit:"70132b0f130acc0bed193d9ba59dd186f0e634cf", GitTreeState:"clean", BuildDate:"2019-12-07T21:12:17Z", GoVersion:"go1.13.4", Compiler:"gc", Platform:"linux/amd64"}
```

# kubectl 查看集群信息
```
    # kubectl cluster-info
    Kubernetes master is running at https://192.168.11.205:6443
    KubeDNS is running at https://192.168.11.205:6443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

    To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
```

# kubectl 运行pod

**注：pod的IP属于cni0**

## kubectl run 查看帮助
```
    # kubectl run --help
    Create and run a particular image, possibly replicated.

     Creates a deployment or job to manage the created container(s).

    Examples:
      # Start a single instance of nginx.
      kubectl run nginx --image=nginx

      # Start a single instance of hazelcast and let the container expose port 5701 .
      kubectl run hazelcast --image=hazelcast --port=5701

      # Start a single instance of hazelcast and set environment variables "DNS_DOMAIN=cluster" and "POD_NAMESPACE=default" in the container.
      kubectl run hazelcast --image=hazelcast --env="DNS_DOMAIN=cluster" --env="POD_NAMESPACE=default"

      # Start a single instance of hazelcast and set labels "app=hazelcast" and "env=prod" in the container.
      kubectl run hazelcast --image=hazelcast --labels="app=hazelcast,env=prod"

      # Start a replicated instance of nginx.
      kubectl run nginx --image=nginx --replicas=5

      # Dry run. Print the corresponding API objects without creating them.
      kubectl run nginx --image=nginx --dry-run

      # Start a single instance of nginx, but overload the spec of the deployment with a partial set of values parsed from JSON.
      kubectl run nginx --image=nginx --overrides='{ "apiVersion": "v1", "spec": { ... } }'

      # Start a pod of busybox and keep it in the foreground, don't restart it if it exits.
      kubectl run -i -t busybox --image=busybox --restart=Never

      # Start the nginx container using the default command, but use custom arguments (arg1 .. argN) for that command.
      kubectl run nginx --image=nginx -- <arg1> <arg2> ... <argN>

      # Start the nginx container using a different command and custom arguments.
      kubectl run nginx --image=nginx --command -- <cmd> <arg1> ... <argN>

      # Start the perl container to compute π to 2000 places and print it out.
      kubectl run pi --image=perl --restart=OnFailure -- perl -Mbignum=bpi -wle 'print bpi(2000)'

      # Start the cron job to compute π to 2000 places and print it out every 5 minutes.
      kubectl run pi --schedule="0/5 * * * ?" --image=perl --restart=OnFailure -- perl -Mbignum=bpi -wle 'print bpi(2000)'

    Options:
          --allow-missing-template-keys=true: If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.
          --attach=false: If true, wait for the Pod to start running, and then attach to the Pod as if 'kubectl attach ...' were called.  Default false, unless '-i/--stdin' is set, in which case the default is true. With '--restart=Never' the exit code of the container process is returned.
          --cascade=true: If true, cascade the deletion of the resources managed by this resource (e.g. Pods created by a ReplicationController).  Default true.
          --command=false: If true and extra arguments are present, use them as the 'command' field in the container, rather than the 'args' field which is the default.
          --dry-run=false: If true, only print the object that would be sent, without sending it.
          --env=[]: Environment variables to set in the container
          --expose=false: If true, a public, external service is created for the container(s) which are run
      -f, --filename=[]: to use to replace the resource.
          --force=false: Only used when grace-period=0. If true, immediately remove resources from API and bypass graceful deletion. Note that immediate deletion of some resources may result in inconsistency or data loss and requires confirmation.
          --generator='': 使用 API generator 的名字, 在 http://kubernetes.io/docs/user-guide/kubectl-conventions/#generators 查看列表.
          --grace-period=-1: Period of time in seconds given to the resource to terminate gracefully. Ignored if negative. Set to 1 for immediate shutdown. Can only be set to 0 when --force is true (force deletion).
          --hostport=-1: The host port mapping for the container port. To demonstrate a single-machine container.
          --image='': 指定容器要运行的镜像.
          --image-pull-policy='': 容器的镜像拉取策略. 如果为空, 这个值将不会 被 client 指定且使用 server 端的默认值
      -k, --kustomize='': Process a kustomization directory. This flag can't be used together with -f or -R.
      -l, --labels='': Comma separated labels to apply to the pod(s). Will override previous values.
          --leave-stdin-open=false: If the pod is started in interactive mode or with stdin, leave stdin open after the first attach completes. By default, stdin will be closed after the first attach completes.
          --limits='': The resource requirement limits for this container.  For example, 'cpu=200m,memory=512Mi'.  Note that server side components may assign limits depending on the server configuration, such as limit ranges.
      -o, --output='': Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-file.
          --overrides='': An inline JSON override for the generated object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field.
          --pod-running-timeout=1m0s: The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running
          --port='': The port that this container exposes.  If --expose is true, this is also the port used by the service that is created.
          --quiet=false: If true, suppress prompt messages.
          --record=false: Record current kubectl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists.
      -R, --recursive=false: Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.
      -r, --replicas=1: Number of replicas to create for this container. Default is 1.
          --requests='': 资源为 container 请求 requests . 例如, 'cpu=100m,memory=256Mi'. 注意服务端组件也许会赋予 requests, 这决定于服务器端配置, 比如 limit ranges.
          --restart='Always': 这个 Pod 的 restart policy.  Legal values [Always, OnFailure, Never]. 如果设置为 'Always' 一个 deployment 被创建, 如果设置为 ’OnFailure' 一个 job 被创建, 如果设置为 'Never', 一个普通的 pod 被创建. 对于后面两个 --replicas 必须为 1.  默认 'Always', 为 CronJobs 设置为 `Never`.
          --rm=false: If true, delete resources created in this command for attached containers.
          --save-config=false: If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.
          --schedule='': A schedule in the Cron format the job should be run with.
          --service-generator='service/v2': 使用 gnerator 的名称创建一个 service.  只有在 --expose 为 true 的时候使用
          --service-overrides='': An inline JSON override for the generated service object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field.  Only used if --expose is true.
          --serviceaccount='': Service account to set in the pod spec
      -i, --stdin=false: Keep stdin open on the container(s) in the pod, even if nothing is attached.
          --template='': Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].
          --timeout=0s: The length of time to wait before giving up on a delete, zero means determine a timeout from the size of the object
      -t, --tty=false: Allocated a TTY for each container in the pod.
          --wait=false: If true, wait for resources to be gone before returning. This waits for finalizers.

    Usage:
      kubectl run NAME --image=image [--env="key=value"] [--port=port] [--replicas=replicas] [--dry-run=bool] [--overrides=inline-json] [--command] -- [COMMAND] [args...] [options]

    Use "kubectl options" for a list of global command-line options (applies to all commands).
```
##  kubectl 运行一个pod
```
    # kubectl run nginx-deploy --image=nginx:1.14-alpine --port=80 --replicas=1 --dry-run=true  干跑模式，并没有运行
    kubectl run --generator=deployment/apps.v1 is DEPRECATED and will be removed in a future version. Use kubectl run --generator=run-pod/v1 or kubectl create instead.
    deployment.apps/nginx-deploy created (dry run)

    # kubectl run nginx-deploy --image=nginx:1.14-alpine --port=80 --replicas=1
    kubectl run --generator=deployment/apps.v1 is DEPRECATED and will be removed in a future version. Use kubectl run --generator=run-pod/v1 or kubectl create instead.
    deployment.apps/nginx-deploy created

    # kubectl get pods
    NAME                            READY   STATUS    RESTARTS   AGE
    nginx-deploy-66ff98548d-xfsd9   1/1     Running   0          10m

    # kubectl get deployment
    NAME           READY   UP-TO-DATE   AVAILABLE   AGE
    nginx-deploy   1/1     1            1           11m

    # kubectl get deployment -o wide
    NAME           READY   UP-TO-DATE   AVAILABLE   AGE   CONTAINERS     IMAGES              SELECTOR
    nginx-deploy   1/1     1            1

            11m   nginx-deploy   nginx:1.14-alpine   run=nginx-deploy

    # kubectl get pods -o wide
    NAME                            READY   STATUS    RESTARTS   AGE   IP           NODE            NOMINATED NODE   READINESS GATES
    nginx-deploy-66ff98548d-xfsd9   1/1     Running   0          11m   10.244.1.2   k8s-node1-206   <none>           <none>

    # curl 10.244.1.2
```

# kubectl删除pod

**自动重建pod，但pod的IP发生变化**
```
    # kubectl get pod
    NAME                            READY   STATUS    RESTARTS   AGE
    nginx-deploy-66ff98548d-xfsd9   1/1     Running   0          18m
    # kubectl delete pods nginx-deploy-66ff98548d-xfsd9
    pod "nginx-deploy-66ff98548d-xfsd9" deleted
    # kubectl get pod
    NAME                            READY   STATUS              RESTARTS   AGE
    nginx-deploy-66ff98548d-v57lq   0/1     ContainerCreating   0          9s
    # kubectl get pod -o wide
    NAME                            READY   STATUS    RESTARTS   AGE   IP           NODE            NOMINATED NODE   READINESS GATES
    nginx-deploy-66ff98548d-v57lq   1/1     Running   0          17s   10.244.2.2   k8s-node2-207   <none>           <none>
```

# kubectl 创建service服务

**给pod一个固定端点，创建service服务，生成IP依然为集群内部IP，解析依靠CoreDNS**
```
    # kubectl expose --help
    Expose a resource as a new Kubernetes service.

     Looks up a deployment, service, replica set, replication controller or pod by name and uses the selector for that
    resource as the selector for a new service on the specified port. A deployment or replica set will be exposed as a
    service only if its selector is convertible to a selector that service supports, i.e. when the selector contains only
    the matchLabels component. Note that if no port is specified via --port and the exposed resource has multiple ports, all
    will be re-used by the new service. Also if no labels are specified, the new service will re-use the labels from the
    resource it exposes.

     Possible resources include (case insensitive):

     pod (po), service (svc), replicationcontroller (rc), deployment (deploy), replicaset (rs)

    Examples:
      # Create a service for a replicated nginx, which serves on port 80 and connects to the containers on port 8000.
      kubectl expose rc nginx --port=80 --target-port=8000

      # Create a service for a replication controller identified by type and name specified in "nginx-controller.yaml",
    which serves on port 80 and connects to the containers on port 8000.
      kubectl expose -f nginx-controller.yaml --port=80 --target-port=8000

      # Create a service for a pod valid-pod, which serves on port 444 with the name "frontend"
      kubectl expose pod valid-pod --port=444 --name=frontend

      # Create a second service based on the above service, exposing the container port 8443 as port 443 with the name
    "nginx-https"
      kubectl expose service nginx --port=443 --target-port=8443 --name=nginx-https

      # Create a service for a replicated streaming application on port 4100 balancing UDP traffic and named 'video-stream'.
      kubectl expose rc streamer --port=4100 --protocol=UDP --name=video-stream

      # Create a service for a replicated nginx using replica set, which serves on port 80 and connects to the containers on
    port 8000.
      kubectl expose rs nginx --port=80 --target-port=8000

      # Create a service for an nginx deployment, which serves on port 80 and connects to the containers on port 8000.
      kubectl expose deployment nginx --port=80 --target-port=8000

    Options:
          --allow-missing-template-keys=true: If true, ignore any errors in templates when a field or map key is missing in
    the template. Only applies to golang and jsonpath output formats.
          --cluster-ip='': ClusterIP to be assigned to the service. Leave empty to auto-allocate, or set to 'None' to create
    a headless service.
          --dry-run=false: If true, only print the object that would be sent, without sending it.
          --external-ip='': Additional external IP address (not managed by Kubernetes) to accept for the service. If this IP
    is routed to a node, the service can be accessed by this IP in addition to its generated service IP.
      -f, --filename=[]: Filename, directory, or URL to files identifying the resource to expose a service
          --generator='service/v2': 使用 generator 的名称. 这里有 2 个 generators: 'service/v1' 和 'service/v2'.
    为一个不同地方是服务端口在 v1 的情况下叫 'default', 如果在 v2 中没有指定名称.
    默认的名称是 'service/v2'.
      -k, --kustomize='': Process the kustomization directory. This flag can't be used together with -f or -R.
      -l, --labels='': Labels to apply to the service created by this call.
          --load-balancer-ip='': IP to assign to the LoadBalancer. If empty, an ephemeral IP will be created and used
    (cloud-provider specific).
          --name='': 名称为最新创建的对象.
      -o, --output='': Output format. One of:
    json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-file.
          --overrides='': An inline JSON override for the generated object. If this is non-empty, it is used to override the
    generated object. Requires that the object supply a valid apiVersion field.
          --port='': 服务的端口应该被指定. 如果没有指定, 从被创建的资源中复制
          --protocol='': 创建 service 的时候伴随着一个网络协议被创建. 默认是 'TCP'.
          --record=false: Record current kubectl command in the resource annotation. If set to false, do not record the
    command. If set to true, record the command. If not set, default to updating the existing annotation value only if one
    already exists.
      -R, --recursive=false: Process the directory used in -f, --filename recursively. Useful when you want to manage
    related manifests organized within the same directory.
          --save-config=false: If true, the configuration of current object will be saved in its annotation. Otherwise, the
    annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.
          --selector='': A label selector to use for this service. Only equality-based selector requirements are supported.
    If empty (the default) infer the selector from the replication controller or replica set.)
          --session-affinity='': If non-empty, set the session affinity for the service to this; legal values: 'None',
    'ClientIP'
          --target-port='': Name or number for the port on the container that the service should direct traffic to.
    Optional.
          --template='': Template string or path to template file to use when -o=go-template, -o=go-template-file. The
    template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].
          --type='': Type for this service: ClusterIP, NodePort, LoadBalancer, or ExternalName. Default is 'ClusterIP'.

    Usage:
      kubectl expose (-f FILENAME | TYPE NAME) [--port=port] [--protocol=TCP|UDP|SCTP] [--target-port=number-or-name]
    [--name=name] [--external-ip=external-ip-of-service] [--type=type] [options]

    Use "kubectl options" for a list of global command-line options (applies to all commands).
```
```
    # kubectl expose deployment nginx-deploy --port=80 --target-port=80 --protocol=TCP
    service/nginx-deploy exposed
    # kubectl get svc
    NAME           TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
    kubernetes     ClusterIP   10.96.0.1      <none>        443/TCP   21h
    nginx-deploy   ClusterIP   10.96.218.77   <none>        80/TCP    21s
    # kubectl get service
    NAME           TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
    kubernetes     ClusterIP   10.96.0.1      <none>        443/TCP   21h
    nginx-deploy   ClusterIP   10.96.218.77   <none>        80/TCP    24s
```
```
    # yum install bind-utils -y
    # kubectl get svc -n kube-system
    NAME       TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)                  AGE
    kube-dns   ClusterIP   10.96.0.10   <none>        53/UDP,53/TCP,9153/TCP   21h
    # kubectl run client --image=busybox --replicas=1 -it --restart=Never     创建pod客户端，进入pod
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
    <!DOCTYPE html>
    <html>
    <head>
    <title>Welcome to nginx!</title>
    <style>
        body {
            width: 35em;
            margin: 0 auto;
            font-family: Tahoma, Verdana, Arial, sans-serif;
        }
    </style>
    </head>
    <body>
    <h1>Welcome to nginx!</h1>
    <p>If you see this page, the nginx web server is successfully installed and
    working. Further configuration is required.</p>

    <p>For online documentation and support please refer to
    <a href="http://nginx.org/">nginx.org</a>.<br/>
    Commercial support is available at
    <a href="http://nginx.com/">nginx.com</a>.</p>

    <p><em>Thank you for using nginx.</em></p>
    </body>
    </html>
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




