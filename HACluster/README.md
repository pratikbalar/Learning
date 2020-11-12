
this command needs to be run on node1, which will initiate the cluster
note: make sure u have a fresh installation of k3s

```curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC="server" sh -s - --cluster-init --token secrets``` 

then below command to be run in other two nodes, node2 and node3, which will be added as a servers in this multi server HACluster

```curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC="server" sh -s - --server https://xxx:6443 --token secrets```

or you can just commands defined in documentation
