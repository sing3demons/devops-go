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

```docker
docker build -t sing3demons/devopsgo:1.0 .
docker push sing3demons/devopsgo:1.0
```

```kube
kubectl apply -f ./kube-ops
kubectl get all -n go-devops
```

```check
kubectl get po -n go-devops
curl -X GET http://127.0.0.1:80 -v
```

> http://localhost:80

```clean
kubectl delete -f ./kube-ops
```

<hr>
### kube-sql

```build
docker build -t sing3demons/devopsgo:db2.0 .
docker push sing3demons/devopsgo:db2.0
```

```kube
kubectl apply -f ./kube-sql
kubectl get po -n go-devops
```

```clean
kubectl delete -f ./kube-sql/
```

> http://localhost:80
> http://localhost/healthz
> http://localhost/news

```
curl --location --request POST 'http://localhost/news' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "title",
    "content": "content",
    "author": "author"
}'
```

```
curl --request GET 'http://localhost/news'
```

<hr>
