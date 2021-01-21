# test-api

Test API server

## Docker build

```bash
make build

docker push vincentvega/test-api:dev
docker push vincentvega/test-api:2021-01-21-15-26-59
```

## Run container

```bash
docker run --rm -it \
    -e POSTGRES_USER=cats \
    -e POSTGRES_PASSWORD=12345678 \
    -e POSTGRES_DB=cats \
    -e POSTGRES_HOST=192.168.1.115 \
    -e POSTGRES_PORT=5432 \
    -e GIN_MODE=debug \
    -p 8080:8080 \
    vincentvega/test-api:dev
```

### GIN_MODE

Running in `"debug"` mode by default

Switch to `"release"` mode in production.
