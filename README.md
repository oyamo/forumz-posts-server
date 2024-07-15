# Forum Z Posts Server

## Download

```bash
git clone git@github.com:oyamo/forumz-posts-server.git
cd forumz-posts-server
```

## Build
```shell
docker build -t post-server:1.0.0 .
```

## Liquibase Database Migration
```shell
docker run --rm --network local-sandbox \
    --volume `pwd`/migration:/liquibase/changelog liquibase/liquibase:4.13 \
    --url="jdbc:postgresql://postgres:5432/posts" \
    --changeLogFile=master-changeLog.yml \
    --username=dev \
    --password='Test@12345' \
    --database-changelog-table-name=database_changelog \
    --database-changelog-lock-table-name=database_changelog_lock \
    --driver=org.postgresql.Driver \
    --log-level=info \
    update
```

## Run
```shell
docker run -d \
  --network sandbox \
  -e POST_SERVICE_DATABASE_DSN='postgresql://localhost:5432/posts?user=dev&password=Test@12345' \
  -e POST_SERVICE_P12_CERTIFICATE='./keystore.p12' \
  -e POST_SERVICE_PUBLIC_KEY='public_key.pub' \
  -e POST_SERVICE_CERT_PASSWORD='Testing@123456' \
  -e POST_SERVICE_KAFKA_CONSUMER='' \
  -e POST_SERVICE_KAFKA_PRODUCER='' \
  post-server:1.0.0

```