### video store

This is a sandbox go application created as a minimal project which can be used to experiment with the contract testing tool, [pact](https://docs.pact.io/).

### example requests

Store a video:

```
curl --location --request POST 'localhost:8080/video' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "123",
    "name": "The Matrix",
    "decription": "The most overrated film of all time"
}'
```

Retrieve a video:

```
curl --location --request GET 'localhost:8080/videos?id=123'
```