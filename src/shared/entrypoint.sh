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

#local disk
if [ -z "$DB" ]; then
  echo "x no remote storage.."
  mlflow server  \
    --host "$HOST" --port "$PORT"
fi

#S3
if [ -n "$MLFLOW_S3_ENDPOINT_URL" -a -n "$S3" ]; then
   echo "+ remote fs: s3"
   "$root_dir/utils.py" && \
   mlflow server  \
            --host "$HOST" --port "$PORT" \
            --backend-store-uri "$DB" \
            --artifacts-destination "$S3" && exit 0 || exit 1
fi

#nfs
if [ -n "$NFS" -a -n "$STORE" ]; then
  echo "+ remote fs: nfs"
  mkdir -p "$STORE";
  mkdir -p /run/sendsigs.omit.d/;
  service rpcbind start && echo "daemon-rpcbind .. ok";
  service nfs-common start && echo "daemon-nfs .. ok";
  mount "$NFS" "$STORE" && echo "mount .. ok";
  echo "" > /data/hello-world || error "nfs mount failed :(";
  mlflow server --backend-store-uri "$DB" \
                --artifacts-destination file://"$STORE" \
                --host "$HOST" --port "$PORT" && exit 0 || exit 1
fi

#finally
echo "x no remote fs"
mlflow server  \
    --host "$HOST" --port "$PORT" \
    --backend-store-uri "$DB"