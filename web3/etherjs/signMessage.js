const { ethers } = require("ethers");

const signMessage = async () => {
  // Create a provider instance
  const provider = new ethers.providers.JsonRpcProvider(
    "https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"
  );

  // Create a wallet instance
  const wallet = new ethers.Wallet(
    "YOUR_KEY",
    provider,
  );

  // Sign the message
  const output = await wallet.signMessage("ethers.js is a powerful library")
  console.log("signature of message:", output);
};
signMessage();
