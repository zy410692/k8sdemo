kubectl config get-contexts

kubectl --context xxx get pods -n default
通过context进行集群的切换

或者kubectl --kubeconfig=/root/.kube/config
kubectl --kubeconfig=/root/.kube/otherconfig


