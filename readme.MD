# MLOPS

![example workflow](https://github.com/MarcelNasser/benchmark-mlops/actions/workflows/docker.yml/badge.svg)

### assumption(s)

bench the [deployment](https://mlflow.org/docs/latest/tracking.html) of open source `mlflow`.

bench carried with following remote file systems:

| fs   | branding  |
|------|-----------|
| hdfs | Hadoop    | 
| nfs  | Filestore |
| kv   | AWS S3    |



---
### deployment(s)

to walkthrough deployments, use a workflow like this:

`````
deploy A => test A => destroy A => deploy B => test B => destroy B => (...)
`````
all deployments are served on port 5000 =>> [localhost:5000](http://localhost:5000)

we wrote a basic training [test](src/tests/training.py). add more tests to investigate further the performance. 

````bash
time python -m unittest src/tests/training.py 2>/dev/null
````


deployments(s):

- deploy `dev`:: single instance with local disk

````bash
docker compose up -d
````

- deploy `multi-az`:: replicates a production with artifacts stored in S3

````bash
docker compose -f .deploy/compose/multi-az.yml up -d
````

- deploy `multi-az-nfs`:: replicates a production with artifacts stored in an NFS

````bash
docker compose -f .deploy/compose/nfs.yml up -d
````

- deploy kubernetes [todo]


---
**red flags:** 

**- destroy compose resources between deployment, to avoid collisions...**

```bash
docker compose -f compose_file down
```

**- wrong arguments `mlflow server [args]` can corrupt remote database.**



