#!/usr/bin/env python

import os
from logging import getLogger, basicConfig, INFO

import botocore.exceptions
import boto3
from retry import retry

logger = getLogger("datapipelines")
basicConfig(level=INFO, format="%(levelname)s: %(message)s")

bucket_name = os.environ.get("BUCKET", "ml-artifacts")
object_name = "mlflow/hello-world"
aws_access_key_id = os.environ.get("AWS_ACCESS_KEY_ID", 'minio-access-key')
aws_secret_access_key = os.environ.get("AWS_SECRET_ACCESS_KEY", 'minio-secret-key')
endpoint_url = os.environ.get("MLFLOW_S3_ENDPOINT_URL", 'http://localhost:9000')


class Count:
    def __init__(self):
        self.iter = 0

    def incr(self):
        self.iter += 1


count = Count()
client = boto3.client('s3',
                      aws_access_key_id=aws_access_key_id,
                      aws_secret_access_key=aws_secret_access_key,
                      endpoint_url=endpoint_url, )


@retry(botocore.exceptions.EndpointConnectionError, tries=10, delay=10)
def check_bucket():
    count.incr()
    try:
        logger.info(f'creating bucket.. (attempt #{count.iter})')
        client.create_bucket(Bucket=bucket_name)
    except client.exceptions.BucketAlreadyOwnedByYou:
        pass
    except botocore.exceptions:
        logger.info('cannot read bucket')
    except Exception:
        raise
    finally:
        logger.info('yep!')


if __name__ == "__main__":
    logger.info("checking bucket ..")
    check_bucket()
