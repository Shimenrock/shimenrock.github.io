---
title: "Kubernetes Setup"
permalink: /k8s/kubernetes-setup/
excerpt: "Kubernetes Setup"
last_modified_at: 2020-01-28T21:36:11-04:00
redirect_from:
  - /theme-setup/
toc: true
---

# Kubernetes方式1部署

## 一、部署方式说明

### 方式1（复杂，需要很多组CA证书）

  ![setup01](/assets/images/setup01.jpg)

  <font size=3>master节点将 API server，etcd，controller-manager,scheduler安装位系统守护进程，主机节点之上，或yum或编译安装;</font>

  - master(需要高可用)
    - kube-controller-manager
    - kube-scheduler
    - kube-apiserver
    - 第三方存储组件 etcd (保存状态，需要高可用)

  <font size=3>node节点将 kube-proxy，kubelet，docker，flannel，安装在主机节点之上，或yum或编译安装</font>
  - node
    - kubelet
    - kube-proxy
    - docker

### 方式2

  ![setup02](/assets/images/setup02.jpg)

  **master和node都用kubeadm部署**
  - 每个节点都要安装docker
  - 每个节点都要运行kubelet
  - 把第一个节点初始化为master，其余节点初始化为node
  - master的四个组件都运行为pod
  - node都运行kube-proxy为pod
  - master和node都要部署flannel的pod
  - 全部是静态pod，部署k8s管理

## 二、Kubernetes Cluster规划和准备

  **网络规划**

  - 节点网络 172.20.0.0/16
  - pod网络 flannel默认 10.244.x.x/16
  - service网络 10.96.0.0/12

  **IP规划**

  -  master, etcd:192.168.11.205
  -  node1：192.168.11.206
  -  node2： 192.168.11.207

  **环境前提**

  -  1. 基于主机名通信 /etc/hosts;
  -  2. 关闭firewall 和iptables.service;
  -  3. 时间同步;
  -  4. OS:CentOS 7.6.1810 ,Extras仓库中

  **版本选择**

  - 官网：https://github.com/kubernetes/kubernetes/releases
  - 生产环境，不要选择太新版本

  **VPS配置**
  - 最少2核

## 三、准备工作
  2.1 配置所有节点免密登陆
  ```
  ssh-keygen -t rsa -b 2048 -N '' -f ~/.ssh/id_rsa
  ssh-copy-id 192.168.11.205
  ssh-copy-id 192.168.11.206
  ssh-copy-id 192.168.11.207
  ssh 192.168.11.205 date && ssh 192.168.11.206 date && ssh 192.168.11.207 date
  ```
  2.2 编辑ansible主机清单
  ```
  vim /etc/ansible/hosts
  [etcd]
  192.168.11.205 NODE_NAME=etcd1

  [kube-master]
  192.168.11.205 ansible_ssh_port=22 ansible_ssh_user=root

  [kube-node]
  192.168.11.206 ansible_ssh_port=22 ansible_ssh_user=root
  192.168.11.207 ansible_ssh_port=22 ansible_ssh_user=root

  [chrony]
  192.168.11.205

  [all:vars]
  CONTAINER_RUNTIME="docker"
  CLUSTER_NETWORK="flannel"
  PROXY_MODE="ipvs"
  SERVICE_CIDR="10.68.0.0/16"
  CLUSTER_CIDR="172.20.0.0/16"
  NODE_PORT_RANGE="20000-40000"
  CLUSTER_DNS_DOMAIN="cluster.local."
  bin_dir="/opt/kube/bin"
  ca_dir="/etc/kubernetes/ssl"
  base_dir="/etc/ansible"

  ansible all -m ping
  ```
  2.3 推送hosts文件
  ```
  vim /etc/hosts
  127.0.0.1   localhost localhost.localdomain localhost4 localhost4.localdomain4
  ::1         localhost localhost.localdomain localhost6 localhost6.localdomain6
  192.168.11.204 Ansible-203
  192.168.11.205 k8s-master-205
  192.168.11.206 k8s-node1-206
  192.168.11.207 k8s-node2-207

  ansible all -m copy -a "src=/etc/hosts dest=/etc/hosts"
  ```
  2.4 配置基础环境
  ```
  # ansible-playbook 01.prepare.yml
  ```
  2.5 配置时间同步
  ```
  # ansible-playbook 02.chrony.yml
  检查时间同步源 # chronyc sources -v
  检查同步源状态 # chronyc sourcestats -v
  查看配置      # cat /etc/chrony.conf
  ```
  2.6 全部关机
  ```
  ansible all -m shell -a 'init 0'
  ```

## 四、安装步骤

  - 1.etcd cluster，仅master节点；
  - 2.flannel，集群的所有节点;
  - 3.配置K8S的master：仅master节点
    -  kubernetes-master
    -  启动服务 kube-apiserver kube-scheduler  kube-controller-manager
  - 4.配置K8S的各node节点；
    -  kubernetes-node
    -  先设定启动docker服务；启动K8S服务kube-proxy,kubelet
  - 5.kubeadm
    - 1. master,node:安装kubelet，kubeadm,docker
    - 2. master: kubeadm init
    - 3. nodes: kubeadm join

## 五、安装Docker

```
    yum install -y yum-utils device-mapper-persistent-data lvm2
    yum-config-manager --add-repo https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
    yum list docker-ce.x86_64 --showduplicates | sort -r
    yum -y install docker-ce
    service docker start

    docker version
    Client: Docker Engine - Community
    Version:           19.03.5
```

## 六、安装kubeadm

```
    https://opsx.alibaba.com/mirror
    https://mirrors.aliyun.com/kubernetes/yum
    https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64/

    https://developer.aliyun.com/mirror
    https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64/
```
```
    cat <<EOF > /etc/yum.repos.d/kubernetes.repo
    [kubernetes]
    name=Kubernetes
    baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64/
    enabled=1
    gpgcheck=1
    repo_gpgcheck=1
    gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
    EOF
    setenforce 0
    yum install -y kubelet kubeadm kubectl
    systemctl enable kubelet && systemctl start kubelet
    systemctl enable docker && systemctl start docker
    #注：node节点不需要安装kubectl
```

## 七、kubeadm init
```
  vim /usr/lib/systemd/system/docker.service
  Environment="HTTPS_PROXY=http://www.ik8s.io:10080"
  Environment="NO_PROXY=127.0.0.0/8,172.20.0.0/16"
  systemctl show docker --property Environment

  systemctl daemon-reload
  systemctl restart docker
  docker info

  cat /proc/sys/net/bridge/bridge-nf-call-iptables
  cat /proc/sys/net/bridge/bridge-nf-call-ip6tables
  确认值是1

  # rpm -ql kubelet
  /etc/kubernetes/manifests         #清单目录
  /etc/sysconfig/kubelet            #配置文件
    KUBELET_EXTRA_ARGS="--fail-swap-on=fals" 忽略swap
  /usr/bin/kubelet
  /usr/lib/systemd/system/kubelet.service

  systemctl start kubelet
  systemctl status kubelet
  tail /var/log/messages
  不要启动，只设置开启自启动
  systemctl enable kubelet
  systemctl enable docker

  kubeadm init --help
  kubeadm init --kubernetes-version=v1.11.1 --pod-network-cidr=10.244.0.0/16 --service-cidr=10.96.0.0/12 --ignore-preflight-errors=Swap

  kubeadm init --kubernetes-version=1.17.0 --pod-network-cidr=10.244.0.0/16 --service-cidr=10.96.0.0/12

  查看进度
  docker image ls
```
``` 日志
  # kubeadm init --kubernetes-version=1.17.0 --pod-network-cidr=10.244.0.0/16 --service-cidr=10.96.0.0/12
  W1216 13:19:53.837158   14623 validation.go:28] Cannot validate kube-proxy config - no validator is available
  W1216 13:19:53.837194   14623 validation.go:28] Cannot validate kubelet config - no validator is available
  [init] Using Kubernetes version: v1.17.0
  [preflight] Running pre-flight checks
          [WARNING IsDockerSystemdCheck]: detected "cgroupfs" as the Docker cgroup driver. The recommended driver is "systemd". Please follow the guide at https://kubernetes.io/docs/setup/cri/
  [preflight] Pulling images required for setting up a Kubernetes cluster
  [preflight] This might take a minute or two, depending on the speed of your internet connection
  [preflight] You can also perform this action in beforehand using 'kubeadm config images pull'
  [kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
  [kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
  [kubelet-start] Starting the kubelet
  [certs] Using certificateDir folder "/etc/kubernetes/pki"
  [certs] Generating "ca" certificate and key
  [certs] Generating "apiserver" certificate and key
  [certs] apiserver serving cert is signed for DNS names [k8s-master-205 kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 192.168.11.205]
  [certs] Generating "apiserver-kubelet-client" certificate and key
  [certs] Generating "front-proxy-ca" certificate and key
  [certs] Generating "front-proxy-client" certificate and key
  [certs] Generating "etcd/ca" certificate and key
  [certs] Generating "etcd/server" certificate and key
  [certs] etcd/server serving cert is signed for DNS names [k8s-master-205 localhost] and IPs [192.168.11.205 127.0.0.1 ::1]
  [certs] Generating "etcd/peer" certificate and key
  [certs] etcd/peer serving cert is signed for DNS names [k8s-master-205 localhost] and IPs [192.168.11.205 127.0.0.1 ::1]
  [certs] Generating "etcd/healthcheck-client" certificate and key
  [certs] Generating "apiserver-etcd-client" certificate and key
  [certs] Generating "sa" key and public key
  [kubeconfig] Using kubeconfig folder "/etc/kubernetes"
  [kubeconfig] Writing "admin.conf" kubeconfig file
  [kubeconfig] Writing "kubelet.conf" kubeconfig file
  [kubeconfig] Writing "controller-manager.conf" kubeconfig file
  [kubeconfig] Writing "scheduler.conf" kubeconfig file
  [control-plane] Using manifest folder "/etc/kubernetes/manifests"
  [control-plane] Creating static Pod manifest for "kube-apiserver"
  [control-plane] Creating static Pod manifest for "kube-controller-manager"
  W1216 13:22:53.091031   14623 manifests.go:214] the default kube-apiserver authorization-mode is "Node,RBAC"; using "Node,RBAC"
  [control-plane] Creating static Pod manifest for "kube-scheduler"
  W1216 13:22:53.091796   14623 manifests.go:214] the default kube-apiserver authorization-mode is "Node,RBAC"; using "Node,RBAC"
  [etcd] Creating static Pod manifest for local etcd in "/etc/kubernetes/manifests"
  [wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory "/etc/kubernetes/manifests". This can take up to 4m0s
  [kubelet-check] Initial timeout of 40s passed.
  [apiclient] All control plane components are healthy after 41.017276 seconds
  [upload-config] Storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
  [kubelet] Creating a ConfigMap "kubelet-config-1.17" in namespace kube-system with the configuration for the kubelets in the cluster
  [upload-certs] Skipping phase. Please see --upload-certs
  [mark-control-plane] Marking the node k8s-master-205 as control-plane by adding the label "node-role.kubernetes.io/master=''"
  [mark-control-plane] Marking the node k8s-master-205 as control-plane by adding the taints [node-role.kubernetes.io/master:NoSchedule]
  [bootstrap-token] Using token: lm27gs.7xyegk1k1kb2y7du
  [bootstrap-token] Configuring bootstrap tokens, cluster-info ConfigMap, RBAC Roles
  [bootstrap-token] configured RBAC rules to allow Node Bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials
  [bootstrap-token] configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Node Bootstrap Token
  [bootstrap-token] configured RBAC rules to allow certificate rotation for all node client certificates in the cluster
  [bootstrap-token] Creating the "cluster-info" ConfigMap in the "kube-public" namespace
  [kubelet-finalize] Updating "/etc/kubernetes/kubelet.conf" to point to a rotatable kubelet client certificate and key
  [addons] Applied essential addon: CoreDNS
  [addons] Applied essential addon: kube-proxy

  Your Kubernetes control-plane has initialized successfully!

  To start using your cluster, you need to run the following as a regular user:

    mkdir -p $HOME/.kube
    sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
    sudo chown $(id -u):$(id -g) $HOME/.kube/config

  You should now deploy a pod network to the cluster.
  Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
    https://kubernetes.io/docs/concepts/cluster-administration/addons/

  Then you can join any number of worker nodes by running the following on each as root:

  kubeadm join 192.168.11.205:6443 --token lm27gs.7xyegk1k1kb2y7du \
      --discovery-token-ca-cert-hash sha256:c8cce0ec69187d13abb828bd723b2d7f67517afdedfc68018cdd2ef6a73bc1e3
```
```
  ss -tnl
  mkdir -p $HOME/.kube
  cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  # kubectl get cs
  NAME                 STATUS    MESSAGE             ERROR
  controller-manager   Healthy   ok
  scheduler            Healthy   ok
  etcd-0               Healthy   {"health":"true"}
  # kubectl get componentstatus
  NAME                 STATUS    MESSAGE             ERROR
  controller-manager   Healthy   ok
  scheduler            Healthy   ok
  etcd-0               Healthy   {"health":"true"}
  # kubectl get node
  NAME             STATUS     ROLES    AGE   VERSION
  k8s-master-205   NotReady   master   15m   v1.17.0

  https://github.com/coreos/flannel

  kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml

  # docker image ls
  REPOSITORY                           TAG                 IMAGE ID            CREATED             SIZE
  k8s.gcr.io/kube-proxy                v1.17.0             7d54289267dc        8 days ago          116MB
  k8s.gcr.io/kube-scheduler            v1.17.0             78c190f736b1        8 days ago          94.4MB
  k8s.gcr.io/kube-apiserver            v1.17.0             0cae8d5cc64c        8 days ago          171MB
  k8s.gcr.io/kube-controller-manager   v1.17.0             5eb3b7486872        8 days ago          161MB
  k8s.gcr.io/coredns                   1.6.5               70f311871ae1        5 weeks ago         41.6MB
  k8s.gcr.io/etcd                      3.4.3-0             303ce5db0e90        7 weeks ago         288MB
  quay.io/coreos/flannel               v0.11.0-amd64       ff281650a721        10 months ago       52.6MB
  k8s.gcr.io/pause                     3.1                 da86e6ba6ca1        24 months ago       742kB
  # kubectl get nodes
  NAME             STATUS   ROLES    AGE   VERSION
  k8s-master-205   Ready    master   19m   v1.17.0
  # kubectl get pods -n kube-system
  NAME                                     READY   STATUS    RESTARTS   AGE
  coredns-6955765f44-4jcc5                 1/1     Running   0          19m
  coredns-6955765f44-hk578                 1/1     Running   0          19m
  etcd-k8s-master-205                      1/1     Running   0          19m
  kube-apiserver-k8s-master-205            1/1     Running   0          19m
  kube-controller-manager-k8s-master-205   1/1     Running   2          19m
  kube-flannel-ds-amd64-9tfnf              1/1     Running   0          2m33s
  kube-proxy-4dlg7                         1/1     Running   0          19m
  kube-scheduler-k8s-master-205            1/1     Running   2          19m
  # kubectl get ns
  NAME              STATUS   AGE
  default           Active   20m
  kube-node-lease   Active   20m
  kube-public       Active   20m
  kube-system       Active   20m
  # kubeadm join 192.168.11.205:6443 --token lm27gs.7xyegk1k1kb2y7du \
      --discovery-token-ca-cert-hash sha256:c8cce0ec69187d13abb828bd723b2d7f67517afdedfc68018cdd2ef6a73bc1e3

  # kubeadm join 192.168.11.205:6443 --token lm27gs.7xyegk1k1kb2y7du \
      >       --discovery-token-ca-cert-hash sha256:c8cce0ec69187d13abb828bd723b2d7f67517afdedfc68018cdd2ef6a73bc1e3
      W1216 13:48:22.748057    7794 join.go:346] [preflight] WARNING: JoinControlPane.controlPlane settings will be ignored when control-plane flag is not set.
      [preflight] Running pre-flight checks
              [WARNING Service-Docker]: docker service is not enabled, please run 'systemctl enable docker.service'
              [WARNING IsDockerSystemdCheck]: detected "cgroupfs" as the Docker cgroup driver. The recommended driver is "systemd". Please follow the guide at https://kubernetes.io/docs/setup/cri/
      [preflight] Reading configuration from the cluster...
      [preflight] FYI: You can look at this config file with 'kubectl -n kube-system get cm kubeadm-config -oyaml'
      [kubelet-start] Downloading configuration for the kubelet from the "kubelet-config-1.17" ConfigMap in the kube-system namespace
      [kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
      [kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
      [kubelet-start] Starting the kubelet
      [kubelet-start] Waiting for the kubelet to perform the TLS Bootstrap...

      This node has joined the cluster:
      * Certificate signing request was sent to apiserver and a response was received.
      * The Kubelet was informed of the new secure connection details.

      Run 'kubectl get nodes' on the control-plane to see this node join the cluster.
```
```
  # docker image ls
  REPOSITORY               TAG                 IMAGE ID            CREATED             SIZE
  k8s.gcr.io/kube-proxy    v1.17.0             7d54289267dc        8 days ago          116MB
  quay.io/coreos/flannel   v0.11.0-amd64       ff281650a721        10 months ago       52.6MB
  k8s.gcr.io/pause         3.1                 da86e6ba6ca1        24 months ago       742kB
  # kubectl get nodes
  NAME             STATUS   ROLES    AGE     VERSION
  k8s-master-205   Ready    master   28m     v1.17.0
  k8s-node1-206    Ready    <none>   2m56s   v1.17.0
  # kubectl get pods -n kube-system -o wide
  NAME                                     READY   STATUS    RESTARTS   AGE     IP               NODE             NOMINATED NODE   READINESS GATES
  coredns-6955765f44-4jcc5                 1/1     Running   0          28m     10.244.0.3       k8s-master-205   <none>           <none>
  coredns-6955765f44-hk578                 1/1     Running   0          28m     10.244.0.2       k8s-master-205   <none>           <none>
  etcd-k8s-master-205                      1/1     Running   0          28m     192.168.11.205   k8s-master-205   <none>           <none>
  kube-apiserver-k8s-master-205            1/1     Running   0          28m     192.168.11.205   k8s-master-205   <none>           <none>
  kube-controller-manager-k8s-master-205   1/1     Running   2          28m     192.168.11.205   k8s-master-205   <none>           <none>
  kube-flannel-ds-amd64-94nmx              1/1     Running   2          3m30s   192.168.11.206   k8s-node1-206    <none>           <none>
  kube-flannel-ds-amd64-9tfnf              1/1     Running   0          11m     192.168.11.205   k8s-master-205   <none>           <none>
  kube-proxy-4dlg7                         1/1     Running   0          28m     192.168.11.205   k8s-master-205   <none>           <none>
  kube-proxy-gqk4h                         1/1     Running   0          3m30s   192.168.11.206   k8s-node1-206    <none>           <none>
  kube-scheduler-k8s-master-205            1/1     Running   2          28m     192.168.11.205   k8s-master-205   <none>           <none>
```


# Kubernetes方式2部署

**｛10-12课｝**

参考 https://www.kubernetes.org.cn/5462.html
* * *


- Kubernetes项目发布初期：部署依靠脚本
- 大厂使用SaltStack、Ansible运维工具自动化部署
- 2017年社区发布独立部署工具 [Kubeadm](https://github.com/kubernetes/kubeadm)

kubelet : 操作容器运行核心组件，在配置容器网络、管理容器数据卷时需要直接操作宿主机，不推荐将该组件容器化，设计上是独立组件

方案：把Kubelet直接运行在宿主机上，使用容器部署其他组件。即运行kubeadm需要在宿主机手动安装kubeadm、kubelet、kubectl

kubeadm init 工作流程
    1.Preflight Check  工作预检查
    2.生成k8s对外提供服务所需各种证书和对应目录。
       k8s对外服务器必须通过HTTPS访问kube-apiserver等各种证书
       所有证书在Master节点的/etc/kubernetes/pki/ca.{crt,key}
    3.证书生成后，生成其他组件访问kube-apiserver的配置文件
       /etc/kubernetes/*.conf
    4.为Master组件生成Pod配置文件
       kube-apiserver   kube-controller-manager  kube-scheduler
       /etc/kubernetes/manifests
       Master容器启动后，kubeadm会检查localhost:/healthz 是否健康
    5.kubeadm为集群生成一个bootstrap token，让节点通过token加进来
    6. kubeadm将证书master节点信息通过configmap方式保存在etcd中，供后续节点使用
    7.最后一步，安装默认插件 kube-proxy  DNS两个容器

Kubeadm部署参数
kubeadm init --config kubeadm.yaml
生成实例代码

>**kubeadm 不能用于生产环境：目前缺少一键部署高可用的k8s集群，即etcd、master应该是多点
>生产环境使用kops或着saltstack部署工具**



## 一、准备工作

- 1. 满足Docker安装要求；
- 2. 机器之间网络互通；
- 3. 有外网访问权限，可以拉取镜像；
- 4. 能够访问gcr.io  quay.io 两个docker registry；
- 5. 单机2核 8G内存；
- 6. 30G以上硬盘。

## 二、操作系统配置

- 1.操作系统版本
```
$ uname -a
Linux k8s-1 4.15.0-64-generic #73-Ubuntu SMP Thu Sep 12 13:16:13 UTC 2019 x86_64 x86_64 x86_64 GNU/Linux
baiy@k8s-1:/etc/apt/sources.list.d$ cat /etc/issue
Ubuntu 18.04.3 LTS \n \l
```
- 2.配置root登陆

```
$ sudo passwd root
Enter new UNIX password:
Retype new UNIX password:
passwd: password updated successfully
$ su  root
```

- 3.安装必要软件

```
# apt-get update && apt-get install -y apt-transport-https curl vim ca-certificates software-properties-common
```

- 4.更换dns管理服务

```
# apt-get install -y unbound
# systemctl stop systemd-resolved
# systemctl disable systemd-resolved
# mv /etc/resolv.conf /etc/resolvconf.bk
# vim /etc/NetworkManager/NetworkManager.conf
 # 在[main]下面添加
 dns=unbound
#  reboot   重启生效
————————————————
tips: 系统自带的systemd-resolved服务会将/etc/resolv.conf软链接到/run/systemd/resolv/stub-resolv.conf，并在里面写入localloop地址。而coredns会读取/etc/resolv.conf文件中的dns服务器地址，如果读到的是localloop，那么coredns会启动失败。当然有很多种方法来解决这个问题，这里采用禁用systemd-resolved，更换为unbound来管理dns服务
```

- 5.修改主机名

```
hostnamectl set-hostname 主机名
```

- 6.配置包转发

```
# vim /etc/sysctl.conf
net.ipv4.ip_forward = 1
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
# modprobe br_netfilter
# sysctl -p              生效
# vim /etc/init.d/load_br_netfilter.sh
  #!/bin/bash
  ### BEGIN INIT INFO
  # Provides:       svnd.sh
  # Required-Start: $local_fs $remote_fs $network $syslog
  # Required-Stop:  $local_fs $remote_fs $network $syslog
  # Default-Start:  2 3 4 5
  # Default-Stop:   0 1 6
  # Short-Description: starts the svnd.sh daemon
  # Description:       starts svnd.sh using start-stop-daemon
  ### END INIT INFO
  sudo modprobe br_netfilter

# chmod 775 /etc/init.d/load_br_netfilter.sh
# update-rc.d load_br_netfilter.sh defaults 90
# 如果要取消开机自动载入模块
# update-rc.d -f load_br_netfilter.sh remove
```

- 7.关闭swap

```
# swapoff -a   暂时关闭
# free -m       查看
# vi /etc/fstab    注释swap行
```

- 8.关闭防火墙

```
# ufw disable
```

- 9.关闭selinux

```

```

## 三、安装Docker和kubeadm

- 安装kubeadm
```
# curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
$ cat <<EOF > /etc/apt/sources.list.d/kubernetes.list
deb http://apt.kubernetes.io/ kubernetes-xenial main
EOF
$ apt-get update
$ apt-get install -y docker.io kubeadm
```
通过国内源安装
```
# curl https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | apt-key add -
# cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
deb https://mirrors.aliyun.com/kubernetes/apt/kubernetes-xenial main
EOF
# apt-get update
# apt-get install -y kubelet kubeadm kubectl
kubeadm安装会自动安装 kubelet  kubectl  kubernetes-cni
安装指定版本：
# apt-get install kubeadm=1.11.0-00 kubectl=1.11.0-00 kubelet=1.11.0-00
# apt-get install -y docker.io    不使用社区版，社区版往往没有通过K8S项目验证
```
启动kubelet.service
```
# systemctl enable kubelet && systemctl start kubelet
```


ls -ltr /etc/kubernetes/manifests/yaml文件列表，每个文件都会写着镜像的地址和版本

## 四、部署Kubernetes Master

老版本yaml文件
```
apiVersion: kubeadm.k8s.io/v1alpha1
kind: MasterConfiguration
controllerManagerExtraArgs:
         horizontal-pod-autoscaler-use-rest-clients: "true"
         horizontal-pod-autoscaler-sync-period: "10s"
         node-monitor-grace-period: "10s"
apiServerExtraArgs:
         runtime-config: "api/all=true"
kubernetesVersion: "stable-1.11"
```
报错
```
your configuration file uses an old API spec: "kubeadm.k8s.io/v1alpha1". Please use kubeadm v1.11 instead and run 'kubeadm config migrate --old-config old.yaml --new-config new.yaml', which will write the new, similar spec using a newer API version.
```
提示可以通过以下命令，转成新版本配置文件
kubeadm config migrate --old-config old.yaml --new-config new.yaml
但是最新版本kubeadm对老配置文件识别很困难，比如默认安装的是1.16

检查当前版本
```
# kubeadm version
kubeadm version: &version.Info{Major:"1", Minor:"16", GitVersion:"v1.16.0", GitCommit:"2bd9643cee5b3b3a5ecbd3af49d09018f0773c77", GitTreeState:"clean", BuildDate:"2019-09-18T14:34:01Z", GoVersion:"go1.12.9", Compiler:"gc", Platform:"linux/amd64"}
```
卸载重装方法
```
# apt remove kubelet kubectl kubeadm
# apt install kubelet=1.12.0-00
# apt install kubectl=1.11.0-00
# apt install kubeadm=1.11.0-00
```

https://cr.console.aliyun.com/cn-hangzhou/images
支付宝登陆
Registry登陆密码 AZbn-567


注册个账号，在家找容器云服务，这里就可以找到K8S所有的镜像
在里找容器云服务
不仅有这，还可以给你的docker拉镜像 加速


[k8s]
name=k8s
baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg

学习kubeadm的github文档

https://juejin.im/post/5d7fb46d5188253264365dcf

## 五、部署容器网络插件

## 六、部署Kubernetes Worker

## 七、部署Dashboard 可视化插件

## 八、部署容器存储插件
