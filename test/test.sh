#!/bin/sh

echo "**************************************************************"
echo "Starting Integration Tests"
echo "Deploying and building test containers"
docker-compose up --build -d
echo "Sleeping 1 minute to allow DB population..."
sleep 60
echo "Current activity log:"
docker-compose logs dictio-scrapper
echo "Starting tests execution"
go test -v -tags=integration

if [ $? -ne "0"]; then
  echo "Tests execution failed"
  exit $?
fi

echo "Tests execution complete. Tear down test environment..."
docker-compose down
docker-compose rm -f
echo "Integration testing finished"
echo "**************************************************************"