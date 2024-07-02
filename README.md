## Run the database
```
docker compose up -d
```

## Generate migrations
```
make db-up
```

## Seed the database
```
make db-seed
```

## Run the project
```
make run
```

## Using the endpoints
1. Getting products
```
curl --request GET \
  --url http://localhost:1323/product/1 \
  --header 'User-Agent: insomnia/8.5.0'
```

2. Creating products
```
curl --request POST \
  --url http://localhost:1323/product \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/8.5.0' \
  --data '{
	"title":"test",
	"price":1000,
	"description":"bla"
}'
```

3. Updating products
```
curl --request PUT \
  --url http://localhost:1323/product/1 \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/8.5.0' \
  --data '{
	"title":"edited",
	"price":1000,
	"description":"bla"
}'
```

4. Deleting products
```
curl --request DELETE \
  --url http://localhost:1323/product/1 \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/8.5.0'
```
