# Simple Get Deviation App #
This is repository for Simple HTTP GET golang app that counts standard deviation from random.org integers

> **IMPORTANT:** Because random.org Guidelines suggests not issuing multiple simultaneous requests there is limit for number of requests. This limit can be changed in server/app/config/config.go

## Building Image ##
```bash
docker build -t go-image .
```

## Running Container ##
```bash
docker run -d -p 8080:8080 --name go-server go-image
```

