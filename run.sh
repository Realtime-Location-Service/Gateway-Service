#! /bin/bash

PROJECT="gateway_service"

# waiting for consul
until $(curl --output /dev/null --silent --fail http://127.0.0.1:8500/v1/kv); do
    echo 'waiting for consul'
    sleep 5
done

# set consul config key values from example
curl --request PUT --data-binary @config.example.yml http://127.0.0.1:8500/v1/kv/gateway_service
 
# build binary
go build -o $PROJECT .

# set consul env
export GATEWAY_SERVICE_CONSUL_URL="127.0.0.1:8500"
export GATEWAY_SERVICE_CONSUL_PATH=$PROJECT

echo "ENV: GATEWAY_SERVICE_CONSUL_URL =" $GATEWAY_SERVICE_CONSUL_URL
echo "ENV: GATEWAY_SERVICE_CONSUL_PATH =" $GATEWAY_SERVICE_CONSUL_PATH

./$PROJECT serve
