const { ethers } = require('ethers')

// const wallet = ethers.Wallet.createRandom()
// console.log(wallet.address)
// console.log(wallet.privateKey)
// console.log(wallet.mnemonic)

const readChainData = async () => {
  // Create a provider instance
  const provider = new ethers.providers.JsonRpcProvider(
    'https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161'
  );

  // Get current block number
  const blockNumber = await provider.getBlockNumber()
  console.log('Current Block Number:', blockNumber);
}
readChainData()
