# devops-go


```tag
git tag --annotate v1.0.1 --message "init go"
git push origin --tags
```

```run
docker pull ghcr.io/sing3demons/devops-go:c9d81b2300d81778ee24f9835fef6f7e58875368
docker run --name devops-go -d -it -p 2565:2565 -e PORT=2565 ghcr.io/sing3demons/devops-go:c9d81b2300d81778ee24f9835fef6f7e58875368
```

> http://localhost:2565/