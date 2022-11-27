# devops-go

```tag
git tag --annotate v1.0.1 --message "init go"
git push origin --tags
```

```run
docker pull ghcr.io/sing3demons/devops-go:c9d81b2300d81778ee24f9835fef6f7e58875368
docker compose up -d
```

> http://localhost:2565/

```down
docker compose down
```

```kube
docker build -t sing3demons/devopsgo:1.0 . 
docker push sing3demons/devopsgo:1.0
kubectl apply -f ./kube-ops
kubectl get all -n go-devops
```

```check
kubectl get po -n go-devops
```

```clean
kubectl delete -f ./kube-ops
```

> http://localhost:80
