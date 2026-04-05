.PHONY: install contracts-build contracts-test backend-build frontend-build build up deploy-bsc preflight-bsc

install:
	cd backend && go mod tidy
	cd frontend && npm install

contracts-build:
	cd contracts && forge build

contracts-test:
	cd contracts && forge test -q

backend-build:
	cd backend && go build ./...

frontend-build:
	cd frontend && npm run build

build: contracts-build backend-build frontend-build

up:
	./scripts/up.sh

deploy-bsc:
	./scripts/deploy-bsc.sh

preflight-bsc:
	./scripts/preflight-bsc.sh
