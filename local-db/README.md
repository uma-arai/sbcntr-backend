# Local database for development

```shell
docker compose up -d
```

## Access

Replace `${POSTGRES_USER}` and `${POSTGRES_PASSWORD}` with the values from `docker-compose.yml`.

```shell
POSTGRES_USER=$(yq e ".services.db.environment.POSTGRES_USER" docker-compose.yml)
POSTGRES_PASSWORD=$(yq e ".services.db.environment.POSTGRES_PASSWORD" docker-compose.yml)
docker compose exec db postgres -u ${POSTGRES_USER} -p${POSTGRES_PASSWORD}
```
