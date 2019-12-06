#!/usr/bin/env bash

echo "DC"

echo 'server_tasks'
# сборка бинарника server_tasks
cd '/Users/anya/IdeaProjects/tasks'
GOOS=linux GOARCH=amd64 go build -o deploy/app_server_tasks
# создать образ
docker build -f deploy/Dockerfile . -t siannarom/my:server_tasks
rm deploy/app_server_tasks
# локальный запуск
cd "/Users/anya/IdeaProjects/tasks/deploy"
docker-compose up --force-recreate tasks_local

