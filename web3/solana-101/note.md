# Rust and Solana
## Create and Mint Tokens
- Wallet public key 
8NGD8ANvgxKmhDQWzdFJER3nPaJDZpojpm2oKovqE7iZ
token 93MxgKeDDoPesxYTA1ow1UVLJzCnxLxH37MJsAHQpzaY

- Set up a new wallet:
  `solana-keygen new`
  (if you’ve already created a wallet, but you want to create a new one, use the `–force switch`)
- Getting the public key of a wallet:
  `solana-keygen pubkey`
- Checking the balance:
  `solana airdrop 1 --url devnet`
  `solana balance --url devnet`
- Checking your balance online: 
  `explorer.solana.com`

  - switch to devnet 
  - check your public key

- Token creation: define the structure and parameters of the token.
- Token minting: instantiate the token and assign the instance(s) to an owner.
- Token transfer: change the owner of the token.

- Installation:
  `cargo install spl-token-cli`
- Creating a token
  `spl-token create-token --url devnet`
  Copy the created token address
  An account in a wallet holds a token. Accounts have to be created within wallets.

  `spl-token create-account <TOKEN_ADDRESS> --url devnet`
  the returned id is the token account address 

- Minting:
  `spl-token mint <TOKEN_ADDRESS> <NUMBER> --url devnet`

- Checking token balance:
  `spl-token balance <TOKEN_ADDRESS> --url devnet`

- Getting the circulating supply:
  `spl-token supply <TOKEN_ADDRESS> --url devnet`

- Renouncing the ability to mint tokens:
  `spl-token authorize <TOKEN_ADDRESS> mint --disable --url devnet`

- Burn tokens (only our own tokens can be burned):
  `spl-token burn <ACCOUNT_ADDRESS> <NUMBER> --url devnet`

- Send tokens to another account
  - Target a Phantom wallet account 
  - Change network to devnet in the settings icon (cog on bottom-right) 
  - Get the wallet address by clicking on the Wallet name
  - Add a new account in the Phantom wallet instance: 
    - Click: Manage token list 
    - + Custom token 
    - then paste the address 
    - Name and symbol can be specified within the wallet 
    - In Manage Tokens, you can enable making the wallet visible 
  `spl-token transfer <SOURCE_TOKEN_ADDRESS> <AMOUNT> <TARGET_TOKEN_ADDRESS> --url devnet --fund-recipient`
without the fund-recipient flag, you’d not be able to add balance to an unfunded address.

## Solana Architecture
- Getting Started with Solana blogpost: https://solana.com/news/getting-started-with-solana-development 
- On-chain code: contains the code describing the smart contract (Rust)
  - The Rust code (or c/cpp) is compiled to bytecode.
  - The compiled bytecode is then deployed to the Solana network.
  - The goal of Solana programs (smart contracts) is to modify the state that lives on the Solana Network (on chain).
  - The Solana Network consists of cloud-hosted validator nodes.
  - State lives in data structures that are called Accounts.
- How can we reach the deployed Smart Contract code? 
  1. CLI solution (Command-Line Interface) 
  2. Through SDKs (Software Development Kit) implementing e.g. JavaScript method calls to abstract API communication.

## Solana Accounts
- Example smart contract: https://github.com/solana-labs/example-helloworld/blob/master/src/program-rust/src/lib.rs 
- Account data structure definition: https://docs.rs/solana-program/1.5.0/solana_program/account_info/struct.AccountInfo.html 
- Account information 
  1. Everything in Solana is an account. 
  2. Programs (smart contracts) are also accounts. 
  3. Smart contracts are accounts with an executable flag.
  4. Account address: public key (32 byte long string). 
  5. Owner of the account is represented by the public key. This owner is an account that has write access to the owned account.
  6. Accounts can store data in form of a byte (u8) array. These account data can be modified by Solana Programs (smart contracts on chain)
  7. Accounts have a Solana token balance measured in Lamports. Lamports are the smallest denomination of Solana tokens (10 to the power of -9, or nano Solana tokens). One Solana is equal to 1 000 000 000 lamports.
  8. Lamports are used to pay rent to keep accounts with non-empty data on chain. 
  9. Executable flag: for smart contracts that are deployed on chain. The bytecode is placed in the data byte array. The executable flag marks if the byte code of the account is executable.

## Logging
- Post your answers here: https://studygroup.moralis.io/t/exercise-logging/45737
- After learning how logging works, perform the following actions: 
  1. Log the value of greeting_account.counter just before and just after the increment operation
  2. Change the increment operation in the following way: 
    1. if the counter is 0, add 1 to it 
    2. otherwise, multiply the counter by 2 
  3. Verify your solution by compiling the code, deploying it, and running it
 
## Instruction Handling
 - Post your answers here: https://studygroup.moralis.io/t/exercise-instruction-handling/45738