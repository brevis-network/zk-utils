contractroot=../../brevis-contracts
out=$contractroot/out
contractdir=$contractroot/contracts

contract_path_args="$contractdir/light-client-others/bsc-tendermint/*.sol $contractdir/light-client-others/poa/*.sol"

echo "run solc"
solc --overwrite --optimize --via-ir --pretty-json \
  --base-path $contractroot \
  --allow-paths $contractdir \
  --optimize-runs 800 \
  --combined-json abi,bin \
  -o $contractroot/out \
  '@openzeppelin/'=./node_modules/@openzeppelin/ \
  'solidity-bytes-utils/'=./node_modules/solidity-bytes-utils/ $contract_path_args

echo "run abigen"
abigen -combined-json $out/combined.json -pkg bsc -out ../common/bsc/bindings.go
echo "clean up"
rm -rf $contractroot/out
echo "done"
