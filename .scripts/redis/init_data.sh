#!/bin/bash

# Waiting for Redis to be ready
echo "Waiting for Redis to start..."
while ! redis-cli -h redis ping; do
    sleep 1
    echo "Retrying..."
done

echo "Redis is up!"

redis-cli -h redis SET RESET_PASSWORD:T/0je1nPSNM9ZXTONF0KJCdhdMQuqksc+CjN7O7rhcA= '{"userUUID":"b8b13403-5b73-4071-bcae-69690c698ccf","hashEmail":"XviszKHgj0HLJqwP6yzhdhhH5Wt8VKe/SHC4KK2vJNQ=","sessions":[{"uuid":"ec52344d-85d2-4e58-af50-96aed497c999","createdAt":"2023-01-21T16:57:52.036278+07:00"}]}'