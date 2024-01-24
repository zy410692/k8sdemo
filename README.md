# k8sdemo
用golang写一点k8s的简单小demo

1 获取一些指定name的pod name
2 kubernetes多集群的配置

3 删除namespace下所有资源
kubectl delete namespace default --grace-period=0 --force

4 kuboard这个UI挺好用的 
/1 pod的文件浏览器功能
/2 终端的日志功能--日志下载全部比较麻烦


5 对kuberenetes的nodeport做流量控制
需求：对于kubernetes的nodeport开放的端口，比如kuboard 30080.只允许办公网访问
