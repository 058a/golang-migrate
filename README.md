
```
docker run -v ./migrations:/migrations --network golang-migrate_default migrate/migrate -path=/migrations/ -database postgresql://user:password@openapi-db:5432/openapi?sslmode=disable create -ext sql -dir ./migrations -seq tasks

sudo chown -R k8suser:k8suser migrations
```

```
 docker run --rm -v ./migrations:/migrations --network golang-migrate_default migrate/migrate -path=/migrations -database postgresql://user:password@openapi-db:5432/openapi?sslmode=disable up
```

```
 docker run --rm -v ./migrations:/migrations --network golang-migrate_default migrate/migrate -path=/migrations -database postgresql://user:password@openapi-db:5432/openapi?sslmode=disable down
```