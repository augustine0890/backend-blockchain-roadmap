const { ethers } = require("ethers");

const CONTRACT_ADDRESS = '0xCC8048eF226eb2383B08949F752Cf31932d487cc';
const ADDRESS_FROM = '0x009Fdf35d5De7D77909Eb64B50b642A85F447E10';

const eventListener = () => {
  const provider = new ethers.providers.JsonRpcProvider(
    "https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161",
  );

  // const topicSetsFilter = [
  // ethers.utils.id("Transfer(address,address,uint256)"),
  // ];

  // event Transfer(address indexed from, address indexed to, uint256 value)
  const filter = {
    address: CONTRACT_ADDRESS,
    topic: [
      ethers.utils.id("Transfer(address,address,uint256"),
      ADDRESS_FROM,
      // ADDRESS_TO,
    ],
  };

  provider.on(filter, (res) => {
    console.log(res);
  });
};
eventListener();