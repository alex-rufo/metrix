NAME=metrix
DOCKER_COMPOSE_FILE=build/docker/docker-compose.yml

.PHONY: start
start:
	docker-compose -p $(NAME) -f $(DOCKER_COMPOSE_FILE) up -d

.PHONY: stop
stop:
	docker-compose -p $(NAME) -f $(DOCKER_COMPOSE_FILE) down
