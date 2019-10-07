# Go Simple API

Simple Go API + Docker multistage build (`from scratch`)

[Golang Paris meetup code](https://github.com/GolangParis/dont-panic)

**How to run:**

```
docker build -t gosimpleapi .
docker run --rm -it -p 8080:8080 gosimpleapi
http localhost:8008/speed
```

