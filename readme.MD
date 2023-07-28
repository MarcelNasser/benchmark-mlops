# MLOPS

![example workflow](https://github.com/MarcelNasser/benchmark-mlops/actions/workflows/docker.yml/badge.svg)

bench the [deployment](https://mlflow.org/docs/latest/tracking.html) of open source `mlflow`.

bench carried with following remote file systems:

| fs   | branding  |
|------|-----------|
| hdfs | Hadoop    | 
| nfs  | Filestore |
| kv   | AWS S3    |

---
**red flags:** 

**- wrong arguments `mlflow server [args]` can corrupt remote database.**



