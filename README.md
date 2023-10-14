# Base API Project


# Requirements
- docker
- docker-compose
- golang
- curl

# How to Start

Is recommended to start adding new endpoints to the existing `api.yaml`. Then build the whole flow to insert or publish.  
Then you can remove everything else (the user CRUD example code).


## Run develop environment

Configure env file: copy `env.dist` and name it to `.env`

```shell
make dev
```

Wait for
```
[00] Starting service
[00] INFO[0000] Using Pubsub Emulator: localhost:8085
[00] INFO[0000] server start at PORT 8080
```
Kill the process with ctrl+c

## Debug
```shell
make debug
```

Wait for
```
[00] Got a connection, launched process __debug_bin (pid = 82232).
```
then attach the debugger  
You can kill the process with ctrl+c

## Requests

```bash
curl -v http://localhost:8080/v1/users 
```

```bash
curl --location 'http://localhost:8080/v1/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "jcuzmar@protonmail.cl",
    "name": "juan"
}'
```