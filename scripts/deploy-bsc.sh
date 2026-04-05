#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT_DIR/contracts"

if [[ ! -f .env ]]; then
  echo "[deploy] contracts/.env not found, copy from contracts/.env.example first"
  exit 1
fi

source .env

forge script script/Deploy.s.sol:DeployScript \
  --rpc-url "$BSC_TESTNET_RPC_URL" \
  --broadcast \
  --verify
