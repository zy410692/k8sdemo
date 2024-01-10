package main

import (
	"context"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"regexp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func hasSubstring(name, sub string) bool {
	match, _ := regexp.MatchString(sub, name)
	return match
}

func main() {
	// 通过RESTConfig构建客户端实例
	kubeconfig := "/Users/zhangyi/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	podClient := clientset.CoreV1().Pods("default")
	podList, _ := podClient.List(context.TODO(), metav1.ListOptions{})
	for _, pod := range podList.Items {
		//fmt.Printf("%s\n", pod.Name)

		//access other pod fields like Namespace, Labels etc
		//fmt.Printf("Namespace: %s\n", pod.Namespace)
		if hasSubstring(pod.Name, "payment") {
			//fmt.Println("Pod name contains substring")
			fmt.Println(pod.Name)
		}
		if hasSubstring(pod.Name, "global") {
			//fmt.Println("Pod name contains substring")
			fmt.Println(pod.Name)
		}
	}
}
