#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"

cleanup() {
  if [[ -n "${BACK_PID:-}" ]]; then
    kill "$BACK_PID" >/dev/null 2>&1 || true
  fi
}

trap cleanup EXIT INT TERM

echo "[up] start backend on :8080"
cd "$ROOT_DIR/backend"
if [[ -f .env ]]; then
  set -a
  source .env
  set +a
else
  echo "[up] warning: backend/.env not found, will rely on shell env"
fi
go run ./cmd/server &
BACK_PID=$!

sleep 1

echo "[up] start frontend on :3000"
cd "$ROOT_DIR/frontend"
npm run dev
