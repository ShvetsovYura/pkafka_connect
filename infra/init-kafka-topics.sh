#!/bin/sh

KT="/opt/bitnami/kafka/bin/kafka-topics.sh"

echo "Waiting for kafka..."
"$KT" --bootstrap-server localhost:9092 --list

echo "Creating kafka topics"
"$KT" --bootstrap-server localhost:9092 --create --if-not-exists --topic metrics --replication-factor 2 --partitions 3

echo "Successfully created the following topics:"
"$KT" --bootstrap-server localhost:9092 --list