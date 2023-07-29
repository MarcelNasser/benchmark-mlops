# Docker

<br>

### images::

| image                     | server | size    | 
|---------------------------|--------|---------|
| marcelndeffo/tools:mlflow | mlflow | ~933MB  | 
| marcelndeffo/tools:nfs    | nfs    | ~15.8MB | 

### build:

- marcelndeffo/tools:mlflow

````bash
$ docker compose -f .deploy/compose/local.yml build
````

- marcelndeffo/tools:nfs

````bash
$ docker compose -f .deploy/compose/multi-az-nfs.yml build
````
