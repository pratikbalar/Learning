charts:

``` myapp conatines gitea charts```

```consent conatins consent charts```

```hydra contains hydra charts```



Run these commands 

```helm install gitea myapp/```


```helm install consent consent/```


```kubectl apply -f postgres-config.yaml -f postgres-stateful-svc.yaml ``` 

( we are assuming that database(could be anything) is running before we implement the exercise )

Note: in case of other database we weould have to pass --set flag while installing the helm chart, along with the dsn query


```helm install hydra hydra/```


``` kubectl apply -f hydra-migrate-job.yaml ```

( for migration, we are runnning a job )


presuming that you have set hosts in etc/hosts file

now hit the browser with gitea.localhost

```
Initial Configuration

localhost ==> gitea.localhost

3000 ==> 3000

http://localhost:3000 ==> hhtp://gitea.localhost/

install

``` 
now run the two commands from described in commands file (assuming you have installed ory hydra binary on your system)

```


Authsourse Configuration

Name: ORY_Hydra
client: gitea-client
secret: secret
discovery url: http://hydra.default.svc.cluster.local:80/.well-known/openid-configuration








