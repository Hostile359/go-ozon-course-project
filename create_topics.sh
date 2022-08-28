#!/bin/sh

export KAFKA_CONTAINER="kafka-1"
export KAFKA_HOST="kafka-1:9092"
export KAFKA_TOPICS="/opt/kafka/bin/kafka-topics.sh"
docker exec ${KAFKA_CONTAINER} \
${KAFKA_TOPICS} --bootstrap-server ${KAFKA_HOST} \
             --create --topic add_users \
             --partitions 1 \
             --replication-factor 1

docker exec ${KAFKA_CONTAINER} \
${KAFKA_TOPICS} --bootstrap-server ${KAFKA_HOST} \
             --create --topic update_users \
             --partitions 1 \
             --replication-factor 1

docker exec ${KAFKA_CONTAINER} \
${KAFKA_TOPICS} --bootstrap-server ${KAFKA_HOST} \
             --create --topic delete_users \
             --partitions 1 \
             --replication-factor 1

