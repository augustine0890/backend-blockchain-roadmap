# Cosmos Academy
- [Cosmos](https://tutorials.cosmos.network/academy/) is a network of interoperable blockchains built on BFT consensus.

## Course Modules
- Module 1: What is Cosmos?
  - Cosmos, its Ecosystem, and ATOM
  - Blockchain Technology and Cosmos
  - The Cosmos Ecosystem
  - Getting ATOM and Staking It
- Module 2: Main Concepts
  - A Blockchain App Architecture
  - Accounts
  - Transactions
  - Messages
  - [Modules](https://tutorials.cosmos.network/academy/2-main-concepts/modules.html)
  - Protobuf
  - Multistore and Keepers
  - BaseApp
  - Queries
  - Events
  - Context
  - Migrations
  - Inter-Blockchain Communication
  - Bridges
- Module 3: Running a Chain
  - Running a Node, API, and CLI
  - [`simapp`](https://tutorials.cosmos.network/academy/3-running-a-chain/node-api-and-cli.html)
- Module 4: My Own Cosmos Chain
- Module 5: What's Next?
  - Continue your Cosmos journey

### Modules
- Most of the work for developers involved in building a Cosmos SDK application consists of building custom modules required by their application that do not exist yet, and integrating them with modules that already exist into one coherent application. Existing modules can come either from the Cosmos SDK itself or from third-party developers. You can download these from an online module repository.

### Migrations
- Upgrading blockchains and blockchain applications is notoriously difficult and risky.
-Upgrading a live chain without software support for upgrades is risky, because all the validators need to pause their state machines at the same block height and apply the upgrade before resuming. If this is not done correctly, there can be state inconsistencies, which are hard to recover from.

## Ignite CLI
- Install Ignite CLI
  - `curl https://get.ignite.com/cli! | bash`
- Verify Ignite CLI version
  - `ignite version`
- Upgrading Ignite CLI
  - Remove all existing Ignite CLI
    1. Press `Ctrl+C` to stop the chain that you started with `ignite chain serve`
    2. Remove the Ignite CLI binary with `rm $(which ignite)`

## Developer Resources
- Cosmos [SDK](https://docs.cosmos.network/)
- [Tendermint](https://docs.tendermint.com/) Core
- Cosmos [Hub](https://hub.cosmos.network/main/hub-overview/overview.html)
- [IBC](https://ibc.cosmos.network/)