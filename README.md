# koivu-api-gateway

Koivu is a lightweight API gateway. By default it has DB-less configuration, ratelimiting, API-key authentication and logging. Koivu is a stateless service and can be run locally, in docker or in kubernetes.

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

Example when running application as a docker container
![image](https://github.com/OnniVirtanen/koivu-api-gateway/assets/116679314/9df233b9-9123-4ef5-8c4e-9488dabcac3b)


Remember that the routed applications should be in the docker network!

### Ratelimiting

Available ratelimiting parameters are the following:

- **requests**: `uint`
- **timeframe**: `second`, `minute`, `hour`, `day`
- **type**: `ip`

### Authentication

Available authentication parameters are the following:

- **authentication**: `key`, `none`

