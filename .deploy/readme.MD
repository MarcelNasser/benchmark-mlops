# deployment(s)

---
### warning::

**- destroy resources between deployments, to avoid collisions...**

```bash
docker compose -f compose_file down
```

**- wrong arguments `mlflow server [args]` can corrupt remote database. it the main weakness of mlflow deployments.**

checkout the startup [scripts](../src/shared) to better understand the workaround done.

---
<br>

| deployment   | filesystem | replicas | command line                                       |
|--------------|------------|----------|----------------------------------------------------|
| k8s          | S3         | 1+       | src/shared/minikube.sh start                       |
| local        | local      | 1        | docker compose -f .deploy/compose/single-az.yml up |
| single-az    | S3         | 1        | docker compose -f .deploy/compose/single-az.yml up |
| multi-az     | S3         | 2        | docker compose -f .deploy/compose/single-az.yml up |
| multi-az-nfs | nfs        | 2        | docker compose -f .deploy/compose/single-az.yml up |

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

