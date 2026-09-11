#!/bin/sh
set -e

PAYMENT_PORT=${PAYMENT_PORT:-50051}
AUDIT_PORT=${AUDIT_PORT:-50052}

check_port() {
  local port=$1
  for i in $(seq 1 5); do
    if nc -z localhost $port; then
      return 0
    fi
    sleep 1
  done
  return 1
}

if! check_port $PAYMENT_PORT; then
  echo "Payment service is not healthy"
  exit 1
fi

if! check_port $AUDIT_PORT; then
  echo "Audit service is not healthy"
  exit 1
fi

echo "Services are healthy"