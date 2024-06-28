# koivu-api-gateway

Koivu is a lightweight API gateway. It has ratelimiting through Redis, API-key authentication and logging. Koivu is a stateless service and can be run locally, in docker or in kubernetes.

## Get started
Koivu has three configuration files: routes.yaml, api-keys.yaml and redis.yaml.

### Creating simple routes file for koivu-api-gateway.

routes.yaml
```
port: "8080"
routes:
  - name: order
    prefix: /api/v1/order
    destination: http://localhost:8040/v1/order
    authentication: key
    ratelimit:
      requests: 5
      timeframe: second
      type: ip
```

### Secure the order route

api-keys.yaml
```
keys:
  - value: super-secret-password-here
    routes: [order]
```

### Add redis configuration. Request counts for ratelimiter are stored in redis

redis.yaml
```
redis:
  url: localhost:6379
  password: ""
```

### Now run the application
```
cd /src
go run ./main.go
```

### Running koivu in docker
Refactor the destinations of routes to be docker-container names.
```
cd /src
```
Build the image
```
docker build -t koivu-api-gateway .
```
Create a docker network
```
docker network create -d bridge app-network
```
Run the docker application
```
docker run -p 8080:8080 --network app-network koivu-api-gateway:latest
```

Example when running application as a docker container
![image](https://github.com/OnniVirtanen/koivu-api-gateway/assets/116679314/ae0805b0-220a-4a26-9aaa-64eea9a6edef)


Remember that the routed applications should be in the docker network!

### Ratelimiting

Available ratelimiting parameters are the following:

- **requests**: `uint`
- **timeframe**: `second`, `minute`, `hour`, `day`
- **type**: `ip`

### Authentication

Available authentication parameters are the following:

- **authentication**: `key`, `none`

