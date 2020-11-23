fleet: cluster orchestration tool

also it gives feature like flux ( yaml manifests, kustomize configs or helm charts)

single cluster-multi cluster management

upto 1 million clusters

bundle and bundledeployment

bundle consists yaml manigests, helm charts etc fetched by git repo and bundleDeployment is an instance of running (deployed) bundle

GitRepo: Git repositories that are watched by Fleet are represented by the type GitRepo.

multi cluster mechanism will have down stream cluster and also agent running inside them, they will use token to register themselves


============================================================================================================================================

helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo add stable https://charts.helm.sh/stable
helm repo update


helm install my-prometheus-operator prometheus-community/kube-prometheus-stack

kubectl get pod

kubectl port-forward prometheus-my-prometheus-operator-kub-prometheus-0 9090

kubectl port-forward my-prometheus-operator-grafana-7fd6bb6f69-zplrd 3000 

kubectl get secrets

kubectl get secret my-prometheus-operator-grafana -o yaml

echo "YWRtaW4=" | base64 -d 

echo "cHJvbS1vcGVyYXRvcg==" | base64 -d

hits the urls: 

http://localhost:9090
http://localhost:3000
