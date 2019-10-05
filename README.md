# betty

**WIP: DO NOT USE**

## TODO: 

* Tests
* Improved Docstrings
* Kill containers
* Improve API routing

## Running:

```bash
docker-compose up
```

```bash
docker build -t betty .
docker run --rm -p 7000:7000 -v /var/run/docker.sock:/var/run/docker.sock betty
```