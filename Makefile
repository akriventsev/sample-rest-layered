.DEFAULT_GOAL = help

COMMON_CONFIG = compose/docker-compose.yml
LOCAL_CONFIG = compose/docker-compose.local.yml

ENV_DEPLOY=.env
include ${ENV_DEPLOY}
export

ENV_FILE_APPLICATION = source/.env

COMMITHASH = $$(git rev-parse HEAD)
DATE_BUILD = $$(date +%Y-%m-%d_%H-%M-%S)

---> [ Network ] ---------------------------------------------------------------------->: ## *

create_network: ## Create a network between services
	@docker network create \
		--opt encrypted=true \
		--attachable \
			sample-network

drop_network: ## Drop a network between between services
	@docker network rm sample-network

recreate_network: drop_network create_network ## ReCreate a network between services

---> [ LocalHost ] -------------------------------------------------------------------->: ## *

check_sample_local: ## Check syntax docker-compose file for LocalHost
	@env $$(cat ${ENV_DEPLOY} | grep ^[A-Z0-9] | xargs) \
		docker-compose \
			--file ${COMMON_CONFIG} \
			--file ${LOCAL_CONFIG} \
				config > /dev/null

config_sample_local: ## Show config docker-compose file for LocalHost
	@env $$(cat ${ENV_DEPLOY} | grep ^[A-Z0-9] | xargs) \
		docker-compose \
			--file ${COMMON_CONFIG} \
			--file ${LOCAL_CONFIG} \
				config

build_sample_local: ## Build sample service for LocalHost
	@echo "SAMPLE_COMMITHASH=${COMMITHASH}" >>  ${ENV_FILE_APPLICATION}
	@echo "SAMPLE_DATE_BUILD=${DATE_BUILD}" >> ${ENV_FILE_APPLICATION}
	@env $$(cat ${ENV_DEPLOY} | grep ^[A-Z0-9] | xargs) \
		docker-compose \
			--file ${COMMON_CONFIG} \
			--file ${LOCAL_CONFIG} \
				build

create_sample_local: ## Start sample service on LocalHost
	@env $$(cat ${ENV_DEPLOY} | grep ^[A-Z0-9] | xargs) \
		docker-compose \
			--file ${COMMON_CONFIG} \
			--file ${LOCAL_CONFIG} \
				up \
					--detach \
					--force-recreate

drop_sample_local: ## Stop sample service on LocalHost
	@env $$(cat ${ENV_DEPLOY} | grep ^[A-Z0-9] | xargs) \
		docker-compose \
			--file ${COMMON_CONFIG} \
			--file ${LOCAL_CONFIG} \
				down

recreate_sample_local: drop_sample_local create_sample_local ## ReCreate sample service on LocalHost

logs_sample_local: ## Logs sample service on LocalHost
	@env $$(cat ${ENV_DEPLOY} | grep ^[A-Z0-9] | xargs) \
		docker-compose \
			--file ${COMMON_CONFIG} \
			--file ${LOCAL_CONFIG} \
				logs \
					--follow


help: ## Show help
	@awk	'BEGIN {FS = ":.*?## "} \
			/^[a-z A-Z0-9\[\]\<\>\{\}_-]+:.*?## / \
			{printf "  \033[36m%-33s\033[0m %s\n", $$1, $$2}' \
				$(MAKEFILE_LIST)

clean: ## Clean data
	rm -r ./compose/pg-data