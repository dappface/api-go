# DAPPFACE API

GraphQL server for DAPPFACE.

## Start Locally

```sh
$ PROJECT_ID=${PROJECT_ID} go run .
```

## Start Docker Container

```sh
$ docker build -t dappface-api-go .
$ docker run -it --rm -p 8080:8080 -e PROJECT_ID=${PROJECT_ID} -e HOME=$HOME -v $HOME:$HOME dappface-api-go
```

## Regenerate when you change graphql schema

```sh
$ go run github.com/99designs/gqlgen -v
```
