package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "go_default_test",
    srcs = ["proxy_server_test.go"],
    embed = [":go_default_library"],
    deps = [
        "//staging/src/k8s.io/apimachinery/pkg/util/proxy:go_default_library",
        "//staging/src/k8s.io/client-go/rest:go_default_library",
    ],
)

go_library(
    name = "go_default_library",
    srcs = ["proxy_server.go"],
    importpath = "k8s.io/kubeadm/pkg/kubectl/proxy",
    deps = [
        "//pkg/kubectl/util:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/util/net:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/util/proxy:go_default_library",
        "//staging/src/k8s.io/client-go/rest:go_default_library",
        "//staging/src/k8s.io/client-go/transport:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
)
