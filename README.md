# gopf
gopf

### Framework

Gin  
https://github.com/gin-gonic/gin

## Install dependencies
```bash
    dep ensure -v
```

## To run
```bash
    Rename example.env to .env

    go run main.go
```
## Docker
```bash
    docker image build -t gopf:1.0 .
    docker container run --publish 9000:9000 --name gopf gopf:1.0
    docker container rm --force gopf
```