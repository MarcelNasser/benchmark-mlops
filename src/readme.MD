# Test Setup

### Testcase(s):

- **test S3**: 
  - desc: check if a bucket 'mlflow' is present in S3 service (minio).
  - scenarios:
    1. the bucket is there:: write 'helloworld'. then, read 'helloword'. exit successfully. 
    2. the bucket is not there:: create the bucket. then write 'helloworld'. then, read 'helloword'. exit successfully.
    3. the S3 service not responding to the localhost port 9000:: crash. (see [setup](#Setup) to fix this)

- **test learning**:
  - desc: log basic training with two methods: autolog and push metric.
  - scenarios:
    1. mlflow is up:: the test run below 10s. then, manual check in mlflow dashboard that artifacts are properly logged.
    2. mlflow is partially broken:: the test run above 10s. warnings could be cryptic client side. investigate server side into 'mlflow' logs. 
    3. mlflow is partially broken:: the test crash. investigate like above. (see [setup](#Setup) to fix this)
    4. mlflow is down:: test will freeze...


run test(s).

````
python -m unittest src/tests/mlflow*.py
````

### Setup

deploy mlflow service stack with guide below. check services are up before starting tests.

---

**use development env of tesselite company**

- create .env

````bash
# create .env file
$ cp .env.tmpl .env
````

- add env url

````cfg
MLFLOW_URL=https://mlops.env1.tesselite.dev  # url here
MLFLOW_EXPERIMENT=test-mlops-env1  # experiment ID
````

---

**deploy services with docker compose:**
````bash
$ docker compose -f .deploy/compose/single-az.yml up -d
````

---
**deploy services with kubernetes:**

- minikube

````bash
#scenario: minikube is not started
$ src/shared/minikube.sh start
#scenario: minikube is started
$ src/shared/minikube.sh deploy
````
then, forward ports:
````bash
$ src/shared/minikube.sh open
````

- external cluster

````bash
#adapt k8s manifests to your cluster
#then deploy
$ kubectl apply -f .deploy/k8s/production
````
then, forward ports:
````bash
#in terminal 1
$ kubectl port-forward deployment/ml 5000:localhost:5000
# in terminal 2
kubectl port-forward deployment/s3 9000:localhost:9000 
````





