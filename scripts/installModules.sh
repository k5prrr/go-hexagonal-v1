#!/bin/bash

libraries=(
    "Init App|go mod init App"
    "Gin Server|go get -u github.com/gin-gonic/gin"
    "Postgres|go get -u gorm.io/driver/postgres"
    "Kafka by confluentinc|go get -u github.com/confluentinc/confluent-kafka-go/kafka && go get -u github.com/confluentinc/confluent-kafka-go/v2/kafka"
    #"MySQL|go get -u github.com/go-sql-driver/mysql"
    "BD gorm|go get -u gorm.io/gorm"
    "Token jwt|go get -u github.com/golang-jwt/jwt"
    "Hash bcrypt|go get -u golang.org/x/crypto/bcrypt"
)

# Вывод списка библиотек
echo "Fast install modules"
for index in "${!libraries[@]}"; do
    # Разделяем название и команду
    IFS='|' read -r name command <<< "${libraries[$index]}"
    echo "$index) $name"
done

# Запрос у пользователя
echo "Enter the library numbers to install (separated by a space):"
read -r selected

# Установка выбранных библиотек
for number in $selected; do
    if [[ -n "${libraries[$number]}" ]]; then
        # Разделяем название и команду
        IFS='|' read -r name command <<< "${libraries[$number]}"
        echo "Install $name..."
        # Выполняем команду
        eval "$command"
    else
        echo "Error: the library with the number $number was not found."
    fi
done

chmod -R 755 .
du -sh .

echo "The installation is complete."
