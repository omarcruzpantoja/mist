1. add local dev envs
export MIGRATE_VERSION=v4.17.1
export SCYLLA_HOST=channel-service-scylla
export SCYLLA_PORT=9042
export SCYLLA_KEYSPACE=channel_service
export SERVICE_NAME=channel-service
export SERVICE_PORT=3010

2. run `make` to get services running 

3. run `initialize-db` and `migrate` the first time server is running

4. from there on `make`should be used to restart services
( if new migrations are added `migrate` should be used)