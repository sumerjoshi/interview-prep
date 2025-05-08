The API server (`server/Server.go`) is a web server.

We use docker and docker-compose to orchestrate the system.

## Running the code

	docker-compose build
	docker-compose up

The code directories are mapped as volumes, so you only need to build once unless  
you add packages to the Pipfile (not necessary to solve the problem, but you can if you  
want).

### Running without Docker

API
```
MODE=api ./run.sh
```

## Querying the API for logs  
```  
curl http://localhost:8080/greeting
```  
