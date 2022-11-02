# Basic Blockchain
## Bitcoin
- Blockchain is about enabling peer-to-peer transactions in a decentralized network
- Validation, Verification, Immutable Recording, and Consensus lead to Trust and Security.

## Blockchain Structure
- A block is composed of a header of information about the block and a set of valid transactions.
- Blockchain was created to support security and trust in a decentralized trustless environment of the cryptocurrency
- UTXO: Unspent Transaction Output. A transaction generates new UTXOs for transferring the amount specified in the input UTXOs.
- Miners are computers that excute operations (software) defined by the blockchain protocol

## Basic Operations
- The algorithm for consensus in the Bitcoin blockchain is Proof of Work (protocol). It gets its name because it involves "work" or computational power to solve the puzzle and to claim the right to form the next block.
- Miners take on added work or computation to verify transactions, broadcast transactions, compete to claim the right to create a block, work on reaching consensus by validating the block, broadcasting the newly created block and confirming transactions.
- Miners are computers that validate and process blockchain transactions and solve the cryptographic puzzle to add new blocks.

# Ethereum Blockchain
## Smart Contracts
- Smart contract is a piece of code deployed in blockchain nodes. It helps creating logic conditions or agreements approved by organizations
- Smart contracts are self-executing lines of code with the terms of an agreement between buyer and seller automatically verified and executed via a computer network
## Ethereum
- The account is the sender and the receiver of a transaction that updates balance as opposed to UTXO methods in Bitcoin.
- Two types of accounts:
  - Externally Owned Accounts (EOAs): which are controlled by private keys. The address of an account is determinted from the public key.
  - Contract Accounts (CA): which are controlled by their contract code and can only be "activated" by an EOA.
    - The address of a contract is determinted at the time the contract is created --> it is derived from the creator address and the number of transactions sent from that address --> called `nonce`.
- Ethereum Virtual Machine (EVM) allows developers to create smart contracts and let's nodes interact with them.
- An Ethereum Node: hold software (including EVM) for
  - Transaction initiation
  - Transaction validation: checking time-stamp, the nonce, sufficient fees.
- The miner nodes in the network receive, verify, gather and execute transactions.
- Gas is what you pay to execute code on the blockchain and to transfer ether to another address. For each instruction on the Ethereum Virtual Machine you pay a certain amount of gas. Some instructions are expensive and some are cheap.
## Ethereum Operations
- Proof-of-Work: miners must compete to solve a difficult puzzle using their computers processing power.
- Proof-of-Stake: instead of requiring mining nodes to run expensive equipment like PoW to discover new blocks, the new PoS system requires users to deposit and lock a specific of coins on the network.
- Stake means how many coins they commited to hold.
## Incentive Model
- Mining process computes gas points required for execution of a transactions
  - Gas limit: the amount of gas points availabe
    - Gas limit = 1,500,000 Units
    - Ether Tx fee = 21,000
    - Gas limit/Tx Fee ~ 70 Txs
  - Gas spent: actual amount spent at completion of block creation.
  - maxPriorityFeePerGas: the maximum amount of gas to be included as a tip to the miner.
# Algorithms and Techniques

# Trust Essentials