# Ethers.js Blockchain Interaction
## Introduction to Ethers.js
- Blockchains contain nodes that share the same state acress all clients, which makes everything accessible as long as you have the access to a node.
- Ethers.js connects to nodes to receive and use the blockchain state for decentralized applications (dApps) purposes.
1. Ethereum Nodes
- A node is any instance of Ethereum client software that is connected to other computers also running Ethereum software, formiing a network.
- A client is an implementation of Ethereum that verifies data against the protocol rules and keeps the network secure.
2. Ethers.js
- Ethers.js is library that helps developers connect with the Ethereum blockchains.
- Ethers.js connects to nodes by using third party APIs, or an injected provider, or other browser wallets.

## Signers and Providers
- A Signer in ethers.js is an abstraction of an Ethereum Account.
- Signers can be used to sign messages and transactions, as well as send signed transactions to the Etherum network to execute state-changing operations.
- A Providers is an abstraction of a connection to the Ethereum network, providing a concise and consistent interface to access standard Ethereum node functionalities.
- Signers is mostly be used when we want to do an operation that changes the state of the blockchain, for example sending ETH, smart contract creation, and calling smart contract write functions.
  - We can also create a signed message with a Signer instance.
- Providers is mostly be used when we want to read data from blockchain, for example getting blockchain data (e.g. block number, transaction data), ether balance for a specific address, and calling read-only functions of smart contracts.
- We needs to pair Signers with Providers to be able to connect to the Ethereum network.

## Web3 Libraries
- Web3.js: JSON-RPC wrapper library for Javascript, developed by the Ethereum Foundation.
- WalletConnect SDK: Web3 library dedicated for mobile interaction, usually paired with Web3.js or Ethers.js, developed by TrustWallet.