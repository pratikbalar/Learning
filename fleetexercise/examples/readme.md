this repo will create gitea-hydra-nginx full working demo with the help of fleet.

to provide this repo to fleet use below yaml manifest.

```

kind: GitRepo
apiVersion: fleet.cattle.io/v1alpha1
metadata:
  name: simple
  namespace: fleet-local
spec:
  repo: https://github.com/dharmendrakariya/Learning/
  paths:
  - fleetexercise/examples

```

```
kubectl apply -f xyz.yaml
```


ps: fleet doesn't have regex for capital letters, I mean make sure you use small letters only in folder name convention.

otherwise you will see something like this.

time="2020-11-25T06:01:08Z" level=fatal msg="Bundle.fleet.cattle.io \"my-repo-FleetExercise-Examples\" is invalid: metadata.name: 
Invalid value: \"my-repo-FleetExercise-Examples\": a DNS-1123 subdomain must consist of lower case alphanumeric characters, '-' or '.', 
and must start and end with an alphanumeric character
(e.g. 'example.com', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*')"
