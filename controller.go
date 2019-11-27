package main

import (
	"database-controller/pkg/generated/clientset/versioned"
	"flag"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"runtime"
)

var kubeconfig = flag.String("kubeconfig", filepath.Join(home(), ".kube", "config"), "out of cluster")
var master = flag.String("master", "https://192.168.40.170:6443", "out of cluster")

func main() {
	println(os.Getenv("HOME"))
	flag.Parse()
	// 创建客户端
	cfg, err := clientcmd.BuildConfigFromFlags(*master, *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	databaseClient, err := versioned.NewForConfig(cfg)
	if err != nil {
		panic(err.Error())
	}

	list, err := databaseClient.RanchercontrollerV1().Databases("default").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, item := range list.Items {
		logrus.Infof("%s命名空间下的%s元数据为:%s", item.Namespace, item.Name, item.Spec)
	}
}

// 获取home路径
func home() string {
	switch runtime.GOOS {
	case "windows":
		return os.Getenv("USERPROFILE")
	default:
		return os.Getenv("HOME")
	}
}
