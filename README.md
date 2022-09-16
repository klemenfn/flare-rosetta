# go-flare-rosetta

# System Requirements
- go version >= 1.18.5

# Build
```
./scripts/build.sh
```

# Launch Rosetta

The following scripts use predefined public RPC endpoints for the Flare and Costwo networks, these can be changed by editing the `"rpc_endpoint"` field in `config/flare.json` and `costwo.json`. 

## Flare Mainnet
```
./scripts/rosetta-server.sh flare
```

## Costwo Testnet
```
./scripts/rosetta-server.sh costwo
```

## LocalFlare Local Network
This local network can be run on your machine by following the instructions on the Flare validator codebase repo.
```
./scripts/rosetta-server.sh localflare
```