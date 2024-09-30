# deployment(s)

---
### warning::

**- destroy resources between deployments, to avoid resources collisions...**

```bash
docker compose -f compose_file down
```

**- if collision anyway, prune manually:**
````bash
# error
[+] Running 1/0
 ✘ Network compose_mlflow-subnet  Error                                                                                                                                                                                                                                                                                                                                                                                      0.0s 
failed to create network compose_mlflow-subnet: Error response from daemon: Pool overlaps with other one on this address space
# fix
$ docker network prune
````

**- wrong arguments `mlflow server [args]` can corrupt remote database. it is the main weakness of mlflow deployments.**

checkout the startup [scripts](../src/shared) to better understand the workaround done.

---
<br>

| deployment      | filesystem | replicas | command line                                             |
|-----------------|------------|----------|----------------------------------------------------------|
| k8s             | S3         | 1+       | src/shared/minikube.sh start                             |
| local           | local      | 1        | docker compose -f .deploy/compose/local.yml up           |
| local-gcs       | gcs        | 1        | docker compose -f .deploy/compose/local-gcs.yml up       |
| single-az       | S3         | 1        | docker compose -f .deploy/compose/single-az.yml up       |
| single-az-auth0 | S3         | 1        | docker compose -f .deploy/compose/single-az-auth0.yml up |
| multi-az        | S3         | 2        | docker compose -f .deploy/compose/multi-az.yml up        |
| multi-az-nfs    | nfs        | 2        | docker compose -f .deploy/compose/multi-az-nfs.yml up    |

<br>

---

### workflow::

to walkthrough deployments, we recommend a workflow like this:

`````
deploy A => test A => destroy A => deploy B => test B => destroy B => (...)
`````
all deployments are served on port 5000 =>> [localhost:5000](http://localhost:5000)

we wrote a basic training [test](../src/tests/training.py). 

add more training tests to investigate further the performance. 

````bash
time python -m unittest src/tests/training.py 2>/dev/null
````

---

### deploy `local`:: 

- *desc*: spin a development mlflow (on local disk).
- *scenarios*:
  1. exploration of `mlflow` features
  2. development new features for `mlflow`

````bash
docker compose -f .deploy/compose/local.yml up -d
````

---

### deploy `local-gcs`:: 

- *desc*: spin a development mlflow connected to remote bucket.
- *scenarios*:
  1. exploration of `mlflow` features
  2. development new features for `mlflow` in google cloud

````bash
GOOGLE_APPLICATION_CREDENTIALS=path_to_service_account_key docker compose -f .deploy/compose/local-gcs.yml up -d
````

---

### deploy `single-az`:: 
- *desc*: spin a development mlflow with S3-like filesystem.
- *scenarios*:
  1. assess the behavior of `mlflow` with remote filesystem
  2. explore network issues of `mlflow`
  3. explore performance issues of `mlflow`

````bash
docker compose -f .deploy/compose/single-az.yml up -d
````

---

### deploy `single-az-auth0`:: 
- *desc*: spin a development mlflow with S3-like filesystem with authentication 2 additional microserveces:
  - authentication service (reverse proxying auth0 + MLFlow UI)
  - logging service (reverse proxying MLFlow run Api + MlFlow artifacts Api)
- *scenarios*:
  1. assess the behavior of `mlflow` with authentication
  2. assess access control of `mlflow` endpoints
  3. assess authentication with [auth0](../.env.tmpl)
  4. assess deployment of auth0 with a reverse [proxy](../src/reverse-proxy) 

````bash
docker compose -f .deploy/compose/single-az-auth0.yml up -d
````

---
### deploy `multi-az`:: 
- *desc*: replicates a production with artifacts stored in S3
- *scenarios*:
  1. assess the behavior of `mlflow` with concurrent access to backend filesystem
  2. explore scalability of `mlflow`
  3. explore resilience of `mlflow`

````bash
docker compose -f .deploy/compose/multi-az.yml up -d
````

---

### deploy `multi-az-nfs`:: 
- *desc*: replicates a production with artifacts stored in a nfs server
- *scenarios*:
  1. assess the behavior of `mlflow` with concurrent access to backend filesystem
  2. explore scalability of `mlflow`
  3. explore resilience of `mlflow`
  4. explore performance of `mlflow`

````bash
docker compose -f .deploy/compose/multi-az.yml up -d
````

---

### deploy `kubernetes`:: 
- *desc*: replicates a production with kubernetes
- *scenarios*:
  1. assess the behavior of `mlflow` in k8s
  2. explore networking of `mlflow` in k8s
  3. explore scalability of `mlflow` in k8s

````bash
#spin minikube
src/shared/minikube.sh start
#deploy services
src/shared/minikube.sh deploy
#forward port(s) to localhost
src/shared/minikube.sh open
````

