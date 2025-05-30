# Docker

__SSH__

```bash
vi /etc/ssh/sshd_config
```
```bash
#Port 22
Port 22
PermitRootLogin yes
```
__Static IP__
> GUI MODE
```bash
nmtui
```
```bash
address 192.168.0.000
netmask 255.255.255.0
gateway 192.168.0.1
dns-nameservers 8.8.8.8
```
> CLI MODE
```bash
vi /etc/sysconfig/network-scripts/ifcfg-enp0s3
```
```ini
TYPE=Ethernet
PROXY_METHOD=none
BROWSER_ONLY=no
BOOTPROTO=dhcp
DEFROUTE=yes
IPV4_FAILURE_FATAL=no
IPV6INIT=yes
IPV6_AUTOCONF=yes
IPV6_DEFROUTE=yes
IPV6_FAILURE_FATAL=no
IPV6_ADDR_GEN_MODE=eui64
NAME=enp0s3
UUID=00000000-0000-0000-0000-000000000000
DEVICE=enp0s3
#ONBOOT=no
ONBOOT=yes
IPADDR=192.168.0.000
NETMASK=255.255.255.0
GATEWAY=192.168.0.0
PREFIX=24
DNS1=8.8.8.8
```

__SELINUX disabled__

```bash
vi /etc/selinux/config
```
```ini
#SELINUX=enfonrcing
SELINUX=disabled
```
```bash
reboot
```
## Installation
### 0. If Docker installed, Uninstall Old Docker versions
```bash
dnf remove docker  
            docker-client \ 
            docker-client-latest \  
            docker-common \  
            docker-latest \  
            docker-latest-logrotate \  
            docker-logrotate \  
            docker-engine  
```
> Images, containers, volumes, and networks stored in /var/lib/docker/ aren’t automatically removed when you uninstall Docker.

### 1. Install Docker

https://docs.docker.com/engine/install/centos/  
https://docs.docker.com/engine/install/fedora/

```bash
dnf update
dnf install epel-release
dnf update
dnf install dnf-plugins-core
dnf config-manager	--add-repo https://download.docker.com/linux/centos/docker-ce.repo
dnf install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
systemctl start docker
systemctl status docker
docker run hello-world
```

### 2. Uninstall Docker
```bash
dnf remove docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin docker-ce-rootless-extras
rm -rf /var/lib/docker
rm -rf /var/lib/containerd
```
