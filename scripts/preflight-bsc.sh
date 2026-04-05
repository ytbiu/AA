#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"

require_file() {
  local f="$1"
  if [[ ! -f "$f" ]]; then
    echo "[preflight] missing file: $f"
    exit 1
  fi
}

require_cmd() {
  local c="$1"
  if ! command -v "$c" >/dev/null 2>&1; then
    echo "[preflight] command not found: $c"
    exit 1
  fi
}

require_cmd cast
require_file "$ROOT_DIR/contracts/.env"
require_file "$ROOT_DIR/backend/.env"
require_file "$ROOT_DIR/frontend/.env.local"

set -a
source "$ROOT_DIR/contracts/.env"
source "$ROOT_DIR/backend/.env"
source "$ROOT_DIR/frontend/.env.local"
set +a

if [[ -z "${BSC_TESTNET_RPC_URL:-}" ]]; then
  echo "[preflight] BSC_TESTNET_RPC_URL is required"
  exit 1
fi

if [[ -z "${ENTRYPOINT_ADDRESS:-}" ]]; then
  echo "[preflight] ENTRYPOINT_ADDRESS is required in contracts/.env"
  exit 1
fi

if [[ -z "${PAYMASTER_ADDRESS:-}" || -z "${NEXT_PUBLIC_PAYMASTER_ADDRESS:-}" ]]; then
  echo "[preflight] PAYMASTER_ADDRESS / NEXT_PUBLIC_PAYMASTER_ADDRESS missing"
  exit 1
fi

if [[ "$(echo "$PAYMASTER_ADDRESS" | tr '[:upper:]' '[:lower:]')" != "$(echo "$NEXT_PUBLIC_PAYMASTER_ADDRESS" | tr '[:upper:]' '[:lower:]')" ]]; then
  echo "[preflight] paymaster mismatch backend vs frontend"
  echo " backend:  $PAYMASTER_ADDRESS"
  echo " frontend: $NEXT_PUBLIC_PAYMASTER_ADDRESS"
  exit 1
fi

if [[ -z "${NEXT_PUBLIC_USDT_ADDRESS:-}" || -z "${NEXT_PUBLIC_DELEGATE_ACCOUNT_ADDRESS:-}" ]]; then
  echo "[preflight] NEXT_PUBLIC_USDT_ADDRESS / NEXT_PUBLIC_DELEGATE_ACCOUNT_ADDRESS missing"
  exit 1
fi

if [[ -z "${BUNDLER_RPC_URL:-}" ]]; then
  if [[ -z "${RPC_URL:-}" ]]; then
    echo "[preflight] RPC_URL is required when BUNDLER_RPC_URL is empty"
    exit 1
  fi
  if [[ -z "${RELAYER_PRIVATE_KEY:-}" ]]; then
    echo "[preflight] RELAYER_PRIVATE_KEY is required when BUNDLER_RPC_URL is empty"
    exit 1
  fi
fi

CHAIN_ID=$(cast chain-id --rpc-url "$BSC_TESTNET_RPC_URL")
if [[ "$CHAIN_ID" != "97" ]]; then
  echo "[preflight] wrong chain id: $CHAIN_ID (expect 97)"
  exit 1
fi

echo "[preflight] chain id ok: $CHAIN_ID"

code_at() {
  cast code "$1" --rpc-url "$BSC_TESTNET_RPC_URL"
}

check_code() {
  local name="$1"
  local addr="$2"
  local code
  code=$(code_at "$addr")
  if [[ "$code" == "0x" ]]; then
    echo "[preflight] $name has no code: $addr"
    exit 1
  fi
  echo "[preflight] $name code ok: $addr"
}

check_code "EntryPoint" "$ENTRYPOINT_ADDRESS"
check_code "Paymaster" "$PAYMASTER_ADDRESS"
check_code "USDT" "$NEXT_PUBLIC_USDT_ADDRESS"
check_code "DelegateLogic" "$NEXT_PUBLIC_DELEGATE_ACCOUNT_ADDRESS"

BACKEND_SIGNER=$(cast wallet address --private-key "$QUOTE_SIGNER_PRIVATE_KEY")
CHAIN_QUOTE_SIGNER=$(cast call "$PAYMASTER_ADDRESS" "quoteSigner()(address)" --rpc-url "$BSC_TESTNET_RPC_URL")
if [[ "$(echo "$BACKEND_SIGNER" | tr '[:upper:]' '[:lower:]')" != "$(echo "$CHAIN_QUOTE_SIGNER" | tr '[:upper:]' '[:lower:]')" ]]; then
  echo "[preflight] quote signer mismatch"
  echo " backend private key => $BACKEND_SIGNER"
  echo " onchain paymaster   => $CHAIN_QUOTE_SIGNER"
  exit 1
fi

echo "[preflight] quote signer ok: $BACKEND_SIGNER"

ENTRYPOINT_DEPOSIT=$(cast call "$ENTRYPOINT_ADDRESS" "balanceOf(address)(uint256)" "$PAYMASTER_ADDRESS" --rpc-url "$BSC_TESTNET_RPC_URL")
if [[ "$ENTRYPOINT_DEPOSIT" == "0" ]]; then
  echo "[preflight] warning: paymaster entrypoint deposit is 0"
else
  echo "[preflight] paymaster deposit: $ENTRYPOINT_DEPOSIT wei"
fi

if [[ -n "${INITIAL_USER:-}" ]] && [[ "$INITIAL_USER" =~ ^0x[0-9a-fA-F]{40}$ ]]; then
  USER_BAL=$(cast call "$NEXT_PUBLIC_USDT_ADDRESS" "balanceOf(address)(uint256)" "$INITIAL_USER" --rpc-url "$BSC_TESTNET_RPC_URL")
  echo "[preflight] INITIAL_USER usdt balance: $USER_BAL"
fi

if [[ -z "${BUNDLER_RPC_URL:-}" ]]; then
  RELAYER_ADDR=$(cast wallet address --private-key "$RELAYER_PRIVATE_KEY")
  RELAYER_BNB=$(cast balance "$RELAYER_ADDR" --rpc-url "$RPC_URL")
  echo "[preflight] direct mode relayer: $RELAYER_ADDR"
  echo "[preflight] direct mode relayer balance: $RELAYER_BNB wei"
fi

echo "[preflight] all checks passed"
