#!/usr/bin/env bash


COMPOSE_PARALLEL_LIMIT=1 docker compose --file ./.hourglass/docker-compose.yml up
