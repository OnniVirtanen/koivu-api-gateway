# koivu-api-gateway

Koivu works as a reverse proxy and an api gateway. It has ratelimiting, apikey authentication and logging.

## Get started
Koivu has two configuration files: config.yaml and keys.yaml. Route configuration is in config.yaml. API-key authentication is in keys.yaml.

### Creating simple config file for koivu-api-gateway.

config.yaml
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

### Secure the order route with an api key by specifying authentication type to key in config.yaml and adding a key in keys.yaml by route name.

keys.yaml
```
keys:
  - value: super-secret-password-here
    routes: [order]
```

### Now run the application
```
cd /src
go run ./main.go
```

### Running koivu in docker
Refactor the destination of config.yaml routes to be docker-container names.
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

Remember that the routed applications should be in the docker network!

### Ratelimiting extras
Available ratelimiting parameters are following:
- requests: uint
- timerframe: second, minute, hour, day
- type: ip

### Authentication
Available authentication parameters are following:
- authentication: key, none
