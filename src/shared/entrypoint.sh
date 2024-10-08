#!/usr/bin/env bash

# very critical
## scan env variables
## spin the server in the proper configuration

function error() {
  echo "$1" 2>&1
  exit 1
}

root_dir=$(realpath "$(dirname "$0")")

#exit for dummy
[ -z "$HOST" ] && error "HOST is not provided."
[ -z "$PORT" ] && error "PORT is not provided."

#local db
if [ -z "$DB" ]; then
  echo "x no remote storage.."
  mlflow server  \
    --host "$HOST" --port "$PORT" --gunicorn-opts "--log-level debug" && exit 0 || exit 1

#S3
elif [ -n "$MLFLOW_S3_ENDPOINT_URL" -a -n "$S3" ]; then
   echo "+ remote fs: s3"
   "$root_dir/check_bucket_s3.py" && \
   mlflow server  \
            --host "$HOST" --port "$PORT" \
            --backend-store-uri "$DB" \
            --artifacts-destination "$S3" --gunicorn-opts "--log-level debug" && exit 0 || exit 1

#GCS
elif [ -n "$GCS" ]; then
   echo "+ remote fs: gcs"
   "$root_dir/check_bucket_gcs.py" && \
   mlflow server  \
            --host "$HOST" --port "$PORT" \
            --backend-store-uri "$DB" \
            --artifacts-destination "$GCS" --gunicorn-opts "--log-level debug" && exit 0 || exit 1

#nfs
elif [ -n "$NFS" -a -n "$STORE" ]; then
  echo "+ remote fs: nfs"
  mkdir -p "$STORE";
  mkdir -p /run/sendsigs.omit.d/;
  service rpcbind start && echo "daemon-rpcbind .. ok";
  service nfs-common start && echo "daemon-nfs .. ok";
  mount "$NFS" "$STORE" && echo "mount .. ok";
  echo "" > /data/hello-world || error "nfs mount failed :(";
  mlflow server --backend-store-uri "$DB" \
                --artifacts-destination file://"$STORE" \
                --host "$HOST" --port "$PORT" --gunicorn-opts "--log-level debug" && exit 0 || exit 1

#local fs
else
  echo "x no remote fs"
  mlflow server  \
      --host "$HOST" --port "$PORT" \
      --backend-store-uri "$DB" --gunicorn-opts "--log-level debug" && exit 0 || exit 1
fi