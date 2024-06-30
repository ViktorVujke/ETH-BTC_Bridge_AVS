 #!/bin/bash

function create_ts_binding {
    contract_dir=$1
    contract=$2
    binding_dir=$3
    echo "Generating TypeScript bindings for" $contract
    contract_json="$contract_dir/out/${contract}.sol/${contract}.json"
    solc_abi=$(cat ${contract_json} | jq -r '.abi')

    # Ensure the output directory exists
    mkdir -p $binding_dir/${contract}

    # Generate TypeScript bindings using TypeChain
    npx typechain --target ethers-v5 --out-dir $binding_dir/${contract} <(echo $solc_abi)
}

# Clean previous builds and bindings
rm -rf bindings-ts/*
forge clean
forge build

# List your contracts
contracts="IncredibleSquaringServiceManager IncredibleSquaringTaskManager GazelleTaskManager ERC20Mock"

# Generate bindings for each contract
for contract in $contracts; do
    create_ts_binding . $contract ./bindings-ts
done
