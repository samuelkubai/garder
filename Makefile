.PHONY: help
PROJECT_NAME ?= eagle-api
DOCKER_API_DEV_COMPOSE := docker/dev/api/docker-compose.yml
DOCKER_CLIENT_DEV_COMPOSE := docker/dev/client/docker-compose.yml
DOCKER_TEST_COMPOSE_FILE := docker/tests/docker-compose.yml
TARGET_MAX_CHAR_NUM=10
## Show help
help:
	@echo ''
	@echo 'Usage:'
	@echo '${YELLOW} make ${RESET} ${GREEN}<target> [options]${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		message = match(lastLine, /^## (.*)/); \
		if (message) { \
			command = substr($$1, 0, index($$1, ":")-1); \
			message = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} %s\n", command, message; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)
	@echo ''

## Start local development server containers
start:
	${INFO} "Creating postgresql database volume"
	@ docker volume create --name=eagle_data > /dev/null
	@ echo "  "
	@ ${INFO} "Building required docker images"
	@ docker-compose -f $(DOCKER_API_DEV_COMPOSE) build app
	@ ${INFO} "Build Completed successfully"
	@ echo " "
	@ ${INFO} "Starting local development server"
	@ docker-compose -f $(DOCKER_API_DEV_COMPOSE) up

# Start the local API development server container
local-api:
	${INFO} "Creating postgresql database volume"
	@ docker volume create --name=eagle_data > /dev/null
	@ echo "  "
	@ ${INFO} "Building required docker images"
	@ docker-compose -f $(DOCKER_API_DEV_COMPOSE) build app
	@ ${INFO} "Build Completed successfully"
	@ echo " "
	@ ${INFO} "Starting local development server"
	@ docker-compose -f $(DOCKER_API_DEV_COMPOSE) up

# Start the local client development server container
local-client:
	@ ${INFO} "Building required docker images"
	@ docker-compose -f $(DOCKER_CLIENT_DEV_COMPOSE) build web
	@ ${INFO} "Build Completed successfully"
	@ echo " "
	@ ${INFO} "Starting local development server"
	@ docker-compose -f $(DOCKER_CLIENT_DEV_COMPOSE) up web

# Launch storybook to work on the various components.
components:
	@ ${INFO} "Building required docker images"
	@ docker-compose -f $(DOCKER_DEV_COMPOSE_FILE) build storybook
	@ ${INFO} "Build Completed successfully"
	@ echo " "
	@ ${INFO} "Starting storybook environment"
	@ docker-compose -f $(DOCKER_DEV_COMPOSE_FILE) up -d storybook

## Stop local development server containers
stop:
	${INFO} "Stop development server containers"
	@ docker-compose -f $(DOCKER_API_DEV_COMPOSE) down -v
	${INFO} "All containers stopped successfully"

## Run eagle local test suites
test:
	${INFO} "Running tests not yet implemented"

## Remove all development containers and volumes
clean:
	${INFO} "Cleaning your local environment"
	${INFO} "Note all ephemeral volumes will be destroyed"
	@ docker-compose -f $(DOCKER_API_DEV_COMPOSE) down -v
	@ docker-compose -f $(DOCKER_TEST_COMPOSE_FILE) down -v
	@ docker images -q -f label=application=$(PROJECT_NAME) | xargs -I ARGS docker rmi -f ARGS
	${INFO} "Removing dangling images"
	@ docker images -q -f dangling=true -f label=application=$(PROJECT_NAME) | xargs -I ARGS docker rmi -f ARGS
	@ docker system prune
	${INFO} "Clean complete"

## [ service ] Ssh into the API container
api-tunnel:
	${INFO} "Open app container terminal"
	@ docker-compose -f $(DOCKER_API_DEV_COMPOSE) exec app bash

## [ service ] Ssh into the client container
client-tunnel:
	${INFO} "Open the client container terminal"
	@ docker-compose -f $(DOCKER_CLIENT_DEV_COMPOSE) exec web bash

migrate:
	${INFO} "Migration of the database not yet implemented"

seed:
	${INFO} "Seeding of the database not yet implemented"

rollback:
	${INFO} "Rolling back of the database not yet implemented"

# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
NC := "\e[0m"
RESET  := $(shell tput -Txterm sgr0)
# Shell Functions
INFO := @bash -c 'printf "\n"; printf $(YELLOW); echo "===> $$1"; printf $(NC)' SOME_VALUE
SUCCESS := @bash -c 'printf "\n"; printf $(GREEN); echo "===> $$1"; printf $(NC)' SOME_VALUE
