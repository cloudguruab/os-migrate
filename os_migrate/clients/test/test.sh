#!/bin/bash

# Set OpenStack environment variables
export OS_AUTH_URL=${OS_AUTH_URL:-"http://127.0.0.1:35357/v2.0/"}
export OS_USERNAME=${OS_USERNAME:-"admin"}
export OS_PASSWORD=${OS_PASSWORD:-"admin"}
export OS_TENANT_NAME=${OS_TENANT_NAME:-"admin"}
export OS_REGION_NAME=${OS_REGION_NAME:-"RegionOne"}

# Build and run the test
cd "$(dirname "$0")" && go build -o test_openstack test_openstack.go

if [ $? -eq 0 ]; then
    ./test_openstack
else
    echo "Build failed"
    exit 1
fi 