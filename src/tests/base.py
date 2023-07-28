import unittest
import io
import boto3
from src.shared.utils import check_bucket


class TestS3(unittest.TestCase):
    bucket_name = "ml-artifacts"
    object_name = "mlflow/hello-world"
    msg = "hello world!"
    client = boto3.client('s3',
                          aws_access_key_id='minio-access-key',
                          aws_secret_access_key='minio-secret-key',
                          endpoint_url='http://localhost:9000', )

    def test_read_write(self):
        check_bucket()
        # write
        self.client.upload_fileobj(io.BytesIO(self.msg.encode()), self.bucket_name, self.object_name)
        # read
        io_byte = io.BytesIO()
        self.client.download_fileobj(self.bucket_name, self.object_name, io_byte)
        # identity
        assert io_byte.getvalue().decode() == self.msg


if __name__ == '__main__':
    unittest.main()
