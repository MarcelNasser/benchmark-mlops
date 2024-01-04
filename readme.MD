# MLOPS

![example workflow](https://github.com/MarcelNasser/benchmark-mlops/actions/workflows/docker.yml/badge.svg)


## statement(s)

we investigate the following:

**" ML experiments are hardly reproducible without automatic tracking "**

**" mlflow is the best mlops tool for tracking ML experiments "**

**" mlflow requires a remote deployment to be adopted by ML teams "**



## assumption(s)

your beloved team lack observability (a-k-a tracking/logging) on ML experiments. 

you desire to investigate thoroughly `mlflow` as the tool of choice, the official documentation of [project](https://mlflow.org/docs/latest/tracking.html).

the documentation does not demonstrate the tool can really reach the observability promise. As it is not clear, `mlflow` can be deployed remotely and matches production-like standards. 

bench yourself the deployment in 10 minutes, and assess all of that.

we run deployment of `mlflow` on those remote file systems:

| fs   | branding  |
|------|-----------|
| nfs  | Filestore |
| kv   | AWS S3    |




## deployment(s)


| deployment   | filesystem | database | api        | replicas |
|--------------|------------|----------|------------|----------|
| k8s          | S3         | postgres | kubernetes | 1+       |
| local        | local      | postgres | docker     | 1        |
| single-az    | S3         | postgres | docker     | 1        |
| multi-az     | S3         | postgres | docker     | 2        |
| multi-az-nfs | nfs        | postgres | docker     | 2        |


see deployment [doc](.deploy/readme.MD) for addition details.

Basically, do this:
````bash
docker compose -f .deploy/compose/single-az.yml up -d
````



## test(s)

build your test campaign according to your team needs.

we wrote basic tests to check the health of the deployment.

see the test [doc](src/readme.MD) for addition details.


- Basically, add those 3 lines at entry of your training script, **after deploying the server**:

````python
import mlflow

# set server URL
mlflow.set_tracking_uri("http://localhost:5000")

# start logging
mlflow.autolog()
````


- Quick the training like:


````
python -m unittest src/tests/*py
````

- Follow-up in the U.I.: http://localhost:5000




