contractroot=../../brevis-contracts
out=$contractroot/out
cdir=$contractroot/contracts

contract_path_args="$cdir/interfaces/*.sol $cdir/light-client-eth/*.sol $cdir/light-client-eth/common/*.sol $cdir/light-client-eth/interfaces/*.sol $cdir/chunk-sync/*.sol $cdir/verifiers/interfaces/*.sol $cdir/verifiers/*.sol $cdir/verifiers/interfaces/*.sol $cdir/SMT/*.sol $cdir/sdk/core/*.sol $cdir/apps/message-bridge/*.sol $cdir/apps/message-bridge/libraries/*.sol $cdir/apps/message-bridge/interfaces/*.sol $cdir/apps/message-bridge/apps/token-bridge/*.sol $cdir/apps/demo-slot-value/*.sol $cdir/apps/uniswap-v4/uniswap-sum/*.sol $cdir/apps/uniswap-v4/uniswap-v4-hook/*.sol"

echo "run solc"
solc --overwrite --optimize --via-ir --pretty-json \
  --base-path $contractroot \
  --allow-paths $cdir \
  --optimize-runs 800 \
  --combined-json abi,bin \
  -o $contractroot/out \
  '@openzeppelin/'=./node_modules/@openzeppelin/ \
  'solidity-rlp'=./node_modules/solidity-rlp/ \
  'solidity-bytes-utils'=./node_modules/solidity-bytes-utils/ \
  '@uniswap/v4-core/'=./node_modules/@uniswap/v4-core/ \
  $contract_path_args

echo "run abigen"
abigen -combined-json $out/combined.json -pkg eth -out ../common/eth/bindings.go
echo "clean up"
rm -rf $contractroot/out
echo "done"
