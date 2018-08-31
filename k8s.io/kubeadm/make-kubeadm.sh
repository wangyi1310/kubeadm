#! /bin/bash 
cd kubeadm
go build -ldflags " \
-X k8s.io/kubeadm/pkg/version.gitTreeState=dirty \
-X k8s.io/kubeadm/pkg/version.buildDate=2018-08-20T13:18:36Z \
-X k8s.io/kubeadm/pkg/version.gitCommit=31420467ae9503850515786660fa3e93e9d140b2 \
-X k8s.io/kubeadm/vendor/k8s.io/client-go/pkg/version.gitCommit=31420467ae9503850515786660fa3e93e9d140b2 \
-X k8s.io/kubeadm/vendor/k8s.io/client-go/pkg/version.gitTreeState=dirty \
-X k8s.io/kubeadm/pkg/version.gitVersion=v1.13.0-alpha.0.291+31420467ae9503-dirty \
-X k8s.io/kubeadm/vendor/k8s.io/client-go/pkg/version.gitVersion=v1.13.0-alpha.0.291+31420467ae9503-dirty \
-X k8s.io/kubeadm/pkg/version.gitMajor=1 \
-X k8s.io/kubeadm/vendor/k8s.io/client-go/pkg/version.gitMajor=1 \
-X k8s.io/kubeadm/pkg/version.gitMinor=13+ \
-X k8s.io/kubeadm/vendor/k8s.io/client-go/pkg/version.gitMinor=13+"
if [ ! -d ../bin ];then
	mkdir ../bin
fi
mv kubeadm ../bin/
