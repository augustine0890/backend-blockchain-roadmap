var Web3 = require('web3');

var web3 = new Web3(
  new Web3.providers.HttpProvider(
    "https://polygon-rpc.com"
  )
);

const address = "0x0C4e510295EAD84BcDb10c0A9f28A22f18c967fB";

web3.eth
    .getBalance(address)
    .then((balance) => 
      console.log(`balance = ${web3.utils.fromWei(balance, "ether")} MATIC`)
    );
