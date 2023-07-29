# MLOPS

![example workflow](https://github.com/MarcelNasser/benchmark-mlops/actions/workflows/docker.yml/badge.svg)  ![example workflow](https://github.com/MarcelNasser/benchmark-mlops/actions/workflows/docker-nfs.yml/badge.svg) 

---
## statement(s)

we investigate the following:

**" ML experiments are hardly reproducible without automatic tracking "**

**" mlflow is the best mlops tool for tracking ML experiments "**

**" mlflow requires a remote deployment to be adopted by ML teams "**

**" but, you don't want to waste days to deploy and test. and pay expensive infrastructure. "**

**" you want to know in 10 minutes, how a remote mlflow will work for you. "**

---
## assumption(s)

bench deployment(s) of open source [mlflow](https://mlflow.org/docs/latest/tracking.html).

bench carried with following remote file systems:

| fs   | branding  |
|------|-----------|
| nfs  | Filestore |
| kv   | AWS S3    |


---

## deployment(s)


| deployment   | filesystem | database | api        | replicas |
|--------------|------------|----------|------------|----------|
| k8s          | S3         | postgres | kubernetes | 1+       |
| local        | local      | postgres | docker     | 1        |
| single-az    | S3         | postgres | docker     | 1        |
| multi-az     | S3         | postgres | docker     | 2        |
| multi-az-nfs | nfs        | postgres | docker     | 2        |


---

### workflow::

to walkthrough deployments, use a workflow like this:

`````
deploy A => test A => destroy A => deploy B => test B => destroy B => (...)
`````
all deployments are served on port 5000 =>> [localhost:5000](http://localhost:5000)

we wrote a basic training [test](src/tests/training.py). 

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
#forward port(s) to localhost
src/shared/minikube.sh open
````

---
### red flags::

**- destroy compose resources between deployment, to avoid collisions...**

```bash
docker compose -f compose_file down
```

**- wrong arguments `mlflow server [args]` can corrupt remote database.**



