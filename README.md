# koivu-api-gateway

Koivu works as a reverse proxy and an api gateway. It has ratelimiting, apikey authentication and logging.

## Get started
Koivu has two configuration files. One is main configuration for routes (config.yaml) and other for api keys (keys.yaml). Add a simple route to config.yaml

--config.yaml--
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

Secure the order route with an api key by specifying authentication type to key in config.yaml and adding a key in keys.yaml by route name.

--keys.yaml--
```
keys:
  - value: super-secret-password-here
    routes: [order]
```

Now run the application
cd /src
go run ./main.go
