#!/bin/sh

echo "**************************************************************"
echo "Starting Integration Tests"
echo "Deploying and building test containers"
docker-compose up --build -d
echo "Sleeping 20 seconds to allow DB population..."
sleep 20
echo "Current activity log:"
docker-compose logs dictio-scrapper
echo "Starting tests execution"
go test -v -tags=integration
echo "Tests execution complete. Tear down test environment..."
docker-compose down
echo "Integration testing finished"
echo "**************************************************************"