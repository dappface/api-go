<div align="center">
  <img src="https://github.com/dappface/www.dappface.com/raw/master/static/icon-128x128.png" alt="DAPPFACE Logo" />

  <h1>DAPPFACE Go API</h1>

  <strong>
    <p>GraphQL server for DAPPFACE.</p>
  </strong>

  <p>
    <a href="https://github.com/dappface/api-go/actions?workflow=Deploy">
      <img src="https://github.com/dappface/api-go/workflows/Deploy/badge.svg" />
    </a>
  </p>
</div>

## Start Locally

```
PROJECT_ID=${PROJECT_ID} go run .
```

## Start Docker Container

```
docker build -t dappface-api-go .
docker run -it --rm -p 8080:8080 -e PROJECT_ID=${PROJECT_ID} -e HOME=$HOME -v $HOME:$HOME dappface-api-go
```

## Regenerate when you change graphql schema

```sh
go run github.com/99designs/gqlgen -v
```
