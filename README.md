# Проект 3-го модуля (Kafka Connect)



Необходимо выполнить: `sudo docker compose up -d` , это запустит:
 - kafka
 - создание топика `metrics`
 - kafka-ui по адресу `http://localhost:8080`
 - prometheus по адресу `http://localhost:9090`
 - webservice обработки метрик по адресу `http://localhost:8001/metrics` (prometheus смотрит на этот endpoint)

Teстовые данные для отправки в топик `metrics` находятся в файле `test_data` (можно отправить через kafka-ui)

Задание 1. Оптимизация параметров для повышения пропускной способности JDBC Source Connector

| Эксперимент | batch.size | linger.ms | compression.type | buffer.memory | Source Record Write Rate (кops/sec) |
|-------------|------------|-----------|------------------|---------------|-------------------------------------|
| 1           | 500        | 1000      | -                | 33554432      | 2,32                                |
| 2           | 50000      | 1000      | -                | 33554432      | 35,6                                |
| 3           | 50000      | 3000      | snappy           | 33554432      | 52,2                                |
| 4           | 50000      | 3000      | gzip             | 33554432      | 38,2                                |
| 5           | 100000     | 3000      | gzip             | 33554432      | 37,3                                | 

Файлы находятся в архиве `pkafka_connect_lesson.zip`

Log-файл по заданию 2 в файле `debezium.log`