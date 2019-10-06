# betty

**WIP: DO NOT USE**

## TODO: 

* Tests
* Improved Docstrings
* Kill containers
* Improve API routing
* Kubernetes integration

## Running:

**Run using docker-compose**

```bash
docker-compose up
```

**Run with docker.io image**

```bash
docker run --rm -p 7000:7000 -v /var/run/docker.sock:/var/run/docker.sock --name betty docker.io/teamjorge/betty:latest
```


**Building the image first**

```bash
docker build -t betty .
docker run --rm -p 7000:7000 -v /var/run/docker.sock:/var/run/docker.sock betty
```

## Swagger

Swagger docs can be found at the [root](http://localhost:7000/) of the API


