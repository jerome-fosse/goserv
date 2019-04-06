# GOSERV

Microservice that expose a REST api to manage music records. This project is written in GO language with a mariadb database. It is my first project written in GO so it's a training project to help me learn and overpower GO.

## API
* POST /artist
* GET  /artist/{id}
* POST /artist/{id}/record
* GET  /artist/{id}/records
* GET  /record/{id}

## Database
GOSERV use mariadb. Creation scripts are in /mariadb. Data is trored in /mariadb-data.

Create and start the database.
```
docker-compose up -d
```

Start the database
```
docker-compose start
```
