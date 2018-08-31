%global KUBE_MAJOR 1
%global KUBE_MINOR 11
%global KUBE_PATCH 0
%global KUBE_VERSION %{KUBE_MAJOR}.%{KUBE_MINOR}.%{KUBE_PATCH}
%global RPM_RELEASE 0
%global ARCH amd64



Name: kubeadm
Version: %{KUBE_VERSION} 
Release: %{RPM_RELEASE}
License: ASL 2.0
Summary: Container Cluster Manager - Kubernetes Cluster Bootstrapping Tool
Requires: kubelet >= 1.8.0
Requires: kubectl >= 1.8.0
Requires: kubernetes-cni >= 0.5.1
Requires: cri-tools >= 1.11.0

URL: https://kubernetes.io

%description
Command-line utility for deploying a Kubernetes cluster.


%prep
rm -rf $RPM_BUILD_DIR/src
zcat $RPM_SOURCE_DIR/src.tar.gz | tar -xv

%build
cd $RPM_BUILD_DIR/src/k8s.io/%{name} 
./make-kubeadm.sh 

%install
mkdir -p %{buildroot}
mkdir -p %{buildroot}%{_sysconfdir}/systemd/system/
mkdir -p %{buildroot}%{_sysconfdir}/systemd/system/kubelet.service.d/
mkdir -p %{buildroot}/usr/bin
cp $RPM_BUILD_DIR/src/k8s.io/%{name}/bin/kubeadm %{buildroot}/usr/bin
cp $RPM_BUILD_DIR/src/k8s.io/%{name}/10-kubeadm.conf %{buildroot}%{_sysconfdir}/systemd/system/kubelet.service.d/ 

%files
/usr/bin/kubeadm
%{_sysconfdir}/systemd/system/kubelet.service.d/10-kubeadm.conf

