#!/bin/bash

export FLARE_NETWORK=${FLARE_NETWORK:-songbird}
export FLARE_CHAIN=${FLARE_CHAIN:-19}
export FLARE_MODE=${FLARE_MODE:-online}
export FLARE_GENESIS_HASH=${FLARE_GENESIS_HASH:-"0x703ccf62df9fa51265d239a6bb8afb6f2fff2831dac0c2df6208dd254a57585b"}

cat <<EOF > /app/flare-config.json
{
  "network-id": "$FLARE_NETWORK",
  "http-host": "0.0.0.0",
  "api-keystore-enabled": false,
  "api-admin-enabled": false,
  "api-ipcs-enabled": false,
  "api-keystore-enabled": false,
  "db-dir": "/data",
  "chain-config-dir": "/app/configs/chains",
  "network-require-validator-to-connect": true
}
EOF

mkdir -p /app/configs/chains/C

cat <<EOF > /app/configs/chains/C/config.json
{
  "snowman-api-enabled": false,
  "coreth-admin-api-enabled": false,
  "net-api-enabled": true,
  "rpc-gas-cap": 2500000000,
  "rpc-tx-fee-cap": 100,
  "eth-api-enabled": true,
  "personal-api-enabled": false,
  "tx-pool-api-enabled": true,
  "debug-api-enabled": true,
  "web3-api-enabled": true,
  "pruning-enabled": false
}
EOF

cat <<EOF > /app/rosetta-config.json
{
  "mode": "$FLARE_MODE",
  "rpc_endpoint": "http://localhost:9650",
  "listen_addr": "0.0.0.0:8080",
  "network_id": 1,
  "network_name": "$FLARE_NETWORK",
  "chain_id": $FLARE_CHAIN,
  "genesis_block_hash": "$FLARE_GENESIS_HASH"
}
EOF

# Execute a custom command instead of default on
if [ -n "$@" ]; then
  exec $@
fi

exec /app/rosetta-runner \
  -mode $FLARE_MODE \
  -flare-bin /app/flare \
  -flare-config /app/flare-config.json \
  -rosetta-bin /app/rosetta-server \
  -rosetta-config rosetta-config.json
