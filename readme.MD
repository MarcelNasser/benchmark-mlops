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


see [doc](.deploy/readme.MD) for addition details.

---

## test(s)

doc your test campaign according to your needs.

we wrote basic tests to check the health of deployment.

see the [doc](src/readme.MD) for addition details.



