#!/usr/bin/env python
import os
import re
from logging import getLogger, basicConfig, INFO

from dotenv import load_dotenv

basicConfig(level=INFO, format="%(levelname)s: %(message)s")
logger = getLogger("mlops")
load_dotenv()


def check(statement):

    def run(*args):
        try:
            out = statement(*args)
            logger.info(f"check {statement.__name__} [{' '.join(args)}] .. succeeded")
            return out
        except Exception as err:
            logger.error(f"check {statement.__name__} .. failed")
            logger.error(err)
            exit(1)

    return run


@check
def envar(var:str) -> str:
    return os.environ[var]

@check
def bucket(url:str):
    name = re.search('gs://(.*)/.*', url, re.IGNORECASE).group(1)
    logger.info(f"fetching bucket .. {name}")
    from google.cloud.storage import client
    logger.info(f"importing storage client ..")
    client = client.Client()
    return client.get_bucket(name)



if __name__ == "__main__":
    creds = envar("GOOGLE_APPLICATION_CREDENTIALS")
    gcs = envar("GCS")
    bucket(gcs)

