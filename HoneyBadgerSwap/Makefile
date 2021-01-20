.PHONY: help

.DEFAULT_GOAL := help

define PRINT_HELP_PYSCRIPT
import re, sys

for line in sys.stdin:
	match = re.match(r'^([a-zA-Z_-]+):.*?## (.*)$$', line)
	if match:
		target, help = match.groups()
		print("%-20s %s" % (target, help))
endef

export PRINT_HELP_PYSCRIPT

help:
	@python -c "$$PRINT_HELP_PYSCRIPT" < $(MAKEFILE_LIST)

clean: clean-mpc-progs
	docker-compose run --rm eth.chain rm -rf /opt/hbswap/poa/data
	#docker-compose run --rm mpc.nodes rm -rf Scripts/hbswap/log/*

clean-mpc-progs:
	docker-compose run --rm mpc.compile rm -rf Programs/Bytecode Programs/Schedules Programs/Public-Input

clean-player-data:
	docker-compose run --rm mpc.trusted.setup rm -rf Player-Data/*

#down: down-init
down:
	docker-compose -f docker-compose.yml -f liquidity-provider.yml -f trader.yml down --remove-orphans

stop:
	docker-compose stop eth.chain mpc.nodes

rm: clean
	#docker-compose rm --stop --force eth.chain contract.deploy contract.deposit
	docker-compose rm --stop --force -v
	docker volume rm \
		hbswap_public-keys \
		hbswap_secrets-p0 \
		hbswap_secrets-p1 \
		hbswap_secrets-p2 \
		hbswap_secrets-p3

init-eth:
	docker-compose up -d eth.chain
	docker-compose up contract.deploy
	docker-compose up pool.init
	docker-compose up client.deposit

simulation: down mpc-keys mpc-compile
	docker-compose up -d eth.chain
	docker-compose up deploy.contract
	docker-compose up -d \
		mpc.node.0 \
		mpc.node.1 \
		mpc.node.2 \
		mpc.node.3
	docker-compose -f docker-compose.yml -f trader.yml up public.deposit
	docker-compose -f docker-compose.yml -f liquidity-provider.yml up init.pool
	docker-compose -f docker-compose.yml -f trader.yml up secret.deposit
	docker-compose -f docker-compose.yml -f trader.yml up trade-1
	docker-compose -f docker-compose.yml -f trader.yml up trade-2
	docker-compose -f docker-compose.yml -f trader.yml up trade-1
	docker-compose -f docker-compose.yml -f trader.yml up trade-2
	docker-compose -f docker-compose.yml -f trader.yml up trade-1
	docker-compose -f docker-compose.yml -f trader.yml up trade-2
	docker-compose -f docker-compose.yml -f trader.yml up trade-1
	docker-compose -f docker-compose.yml -f trader.yml up trade-2
	docker-compose -f docker-compose.yml -f trader.yml up trade-1
	docker-compose -f docker-compose.yml -f trader.yml up trade-2
	docker-compose -f docker-compose.yml -f trader.yml up trade-1
	docker-compose -f docker-compose.yml -f trader.yml up trade-2

start-hbswap:
	docker-compose up -d eth.chain
	docker-compose up deploy.contract
	docker-compose up -d \
		mpc.node.0 \
		mpc.node.1 \
		mpc.node.2 \
		mpc.node.3
	#sh scripts/follow-sim-logs-with-tmux.sh

public-deposit:
	docker-compose -f docker-compose.yml -f trader.yml up public.deposit

init-secrets:
	docker-compose -f docker-compose.yml -f liquidity-provider.yml up init.pool
	docker-compose -f docker-compose.yml -f trader.yml up secret.deposit

trade-1:
	docker-compose -f docker-compose.yml -f trader.yml up trade-1

trade-2:
	docker-compose -f docker-compose.yml -f trader.yml up trade-2

up-eth:
	docker-compose up -d eth.chain

deploy-contract: up-eth
	docker-compose -f docker-compose.yml -f eth.yml up contract.deploy

lp-init-pool:
	docker-compose -f docker-compose.yml -f liquidity-provider.yml up pool.init

trader-deposit:
	docker-compose -f docker-compose.yml -f trader.yml up client.deposit

mpc-keys:
	docker-compose up mpc.trusted.setup

mpc-compile:
	docker-compose up mpc.compile

mpc-init-pool: mpc-keys mpc-compile
	docker-compose up \
		mpc.init.node.0 \
		mpc.init.node.1 \
		mpc.init.node.2 \
		mpc.init.node.3

mpc:
	docker-compose up -d \
		mpc.node.0 \
		mpc.node.1 \
		mpc.node.2 \
		mpc.node.3

run-client:
	docker-compose up client

#up: ## run the example
up: down rm mpc-keys mpc-compile mpc-init-pool ## run the example
	docker-compose up -d eth.chain
	docker-compose up contract.deploy
	docker-compose up contract.deposit
	docker-compose up -d mpc.node.0 mpc.node.1 mpc.node.2 mpc.node.3
	docker-compose up -d client
	sh scripts/follow-sim-logs-with-tmux.sh

down-init:
	docker-compose -f hbswap-init.yml down --remove-orphans

rm-init: clean
	docker-compose -f hbswap-init.yml rm --stop --force mpc.trusted.setup mpc.compile mpc.node.0 mpc.node.1 mpc.node.2 mpc.node.3
	#docker-compose -f hbswap-init.yml rm --stop --force mpc.trusted.setup mpc.compile mpc.allnodes

up-init: down-init rm-init
	docker-compose -f hbswap-init.yml up mpc.trusted.setup mpc.compile
	docker-compose -f hbswap-init.yml up mpc.node.0 mpc.node.1 mpc.node.2 mpc.node.3
	#docker-compose -f hbswap-init.yml up mpc.allnodes
