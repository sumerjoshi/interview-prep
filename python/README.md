The API server (`api/api.py`) is a web server.

We use docker to orchestrate the system.  You can try running the server and camera:
   
```
docker compose build
docker compose up
```

The API server will start.
### Running the code without Docker

Install the required packages first:
```
pip install -r requirements.txt
```

```
MODE=api ./run.sh
```

## Querying the API for greeting
```
curl http://localhost:8080/greeting
```
