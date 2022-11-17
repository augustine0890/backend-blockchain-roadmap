const { ethers } = require('ethers');
const ABI = require('./ABI');

const callReadFunction = async () => {
  // 1. Create provider instance
  const provider = new ethers.providers.JsonRpcProvider('https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161');

  // Create a wallet instance
  const wallet = new ethers.Wallet("YOUR_KEY", provider);

  // 2. Create the contract instance and connect it to provider
  const contract = new ethers.Contract(
    '0xCC8048eF226eb2383B08949F752Cf31932d487cc',
    ABI.humanReadableABI,
    wallet
  );

  // 3. Getting ERC20 token symbol
  const symbol = await contract.symbol();
  console.log('Symbol:', symbol);

  // 4. Getting ERC20 total supply value
  const totalSupply = await contract.totalSupply();
  const formattedTotalSupply = ethers.utils.formatUnits(totalSupply, 18);
  console.log('Total supply:', formattedTotalSupply, symbol);

  // Call mint function from smart contract and mint 10 MOCK token
  try {
    await contract.mint(wallet.address, ethers.utils.parseUnits("20", 18));
    console.log('Mint Success!')
    // 5. Getting an ERC20 balance for specific wallet
    const balance = await contract.balanceOf(wallet.address);
    const formattedBalance = ethers.utils.formatUnits(balance, 18);
    console.log(`Token Balance of ${wallet.address}: ${formattedBalance} ${symbol}`);
  } catch (err) {
    console.log(err)
  }
};
callReadFunction();