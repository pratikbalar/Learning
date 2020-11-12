
this command needs to be run on node1, which will initiate the cluster
note: make sure u have a fresh installation of k3s

```curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC="server" sh -s - --cluster-init --token secrets``` 

then below command to be run in other two nodes, node2 and node3, which will be added as a servers in this multi server HACluster

```curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC="server" sh -s - --server https://xxx:6443 --token secrets```

or you can just commands defined in documentation


```K3S_TOKEN=SECRET k3s server --cluster-init```


```K3S_TOKEN=SECRET k3s server --server https://<ip or hostname of server1>:6443```

here is the linkto the documentation

```https://rancher.com/docs/k3s/latest/en/installation/ha-embedded/```

note: also you can use --tls-san command with first node1 command execution

read here ```https://rancher.com/docs/k3s/latest/en/installation/ha/```

note: you might endup in CA certificates error and error connection rejected, for that uninstall old k3s and fresh install new k3s and run that ```curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC="server" sh -s - --server https://xxx:6443 --token secrets``` command.


in last you should see ```kubectl get node``` showing all nodes
