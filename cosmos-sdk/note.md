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

## [Hello](https://docs.ignite.com/guide/hello.html) - Ignite CLI
- Create blockchain with default directory structure
  - `ignite scaffold chain github.com/augustine/hello`
- The new blockchain imports standard Cosmos SDK modules
  - `staking`: for delegated Proof-of-Stake (PoS) consensus accounts
  - `bank`: for fungible token transfers between accounts
  - `gov`: for on-chain governance
- Learn more about command
  - `ignite scaffold --help`
- Start a blockchain
  - Download dependencies and compiles the source code into a binary called `hellod`
  - `ignite chain serve`
- HTTP API Console: Validator node exposes two API endpoints
  - [http://localhost:26657](http://localhost:26657): low-level Tendermint API
  - [http://localhost:1317](http://localhost:1317): high-level blockchain API

__Say "Hello, Ignite CLI"__
- To get your blockchain to say `Hello! Ignite CLI`:
  - Modify a protocol buffer file
  - Create a keeper query function that return data
  - Register a query function
- Protocol buffer files contain proto rpc calls that define Cosmos SDK queries and message handlers, and proto messages that define Cosmos SDK types.
- The `Keeper` is an abstraction for modifying the state of the blockchain.
- A typical blockchain developer workflow looks something like this:
  - Start with proto files to define Cosmos SDK messages.
  - Define and register queries
  - Define message handler logic
  - Finally, implement the logic of these queries and message handlers in keeper functions

- __Create a query__
  - Create a `posts` query
    - `ignite scaffold query posts --response title,body`
    - The `query` command has created and modified several files:
      ```shell
      modify proto/hello/query.proto
      modify x/hello/client/cli/query.go
      create x/hello/client/cli/query_posts.go
      create x/hello/keeper/grpc_query_posts.go
      ```
  - Updates to the query service
  - Request and response types
- __Posts keeper function__
  - Update keeper function
    - In the `query.proto` file, the response accepts `title` and `body`
    ```go
    func (k Keeper) Posts(goCtx context.Context, req *types.QueryPostsRequest) (*types.QueryPostsResponse, error) {
      //...
    	return &types.QueryPostsResponse{Title: "Hello!", Body: "Ignite CLI"}, nil
    }
    ```
  - Visit the posts endpoint: [http://localhost:1317/augustine/hello/hello/posts](http://localhost:1317/augustine/hello/hello/posts)

- __Register query handlers__
  - Make the required changes to the `x/hello/module.go` file
  - Search for `RegisterGRPCGatewayRoutes`:
    ```go
    // RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module.
    func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
    	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
    }
    ```
  - Start blockchain with reset: `ignite chain serve -r`
- Run the following command and get the JSON
  - `hellod q hello posts`

## Build a [Blog](https://docs.ignite.com/guide/blog/)
- The purpose of this tutorial is to guide you through the implementation of complete feedback loop: submitting data and reading this data back from the blockchain.
- You will learn about:
  - Scaffolding a Cosmos SDK message
  - Defining new types in protocol buffer files
  - Implementing keeper methods to write data to the store
  - Reading data from the store and return it as a result of a query
  - Using the blockchain's CLI to broadcast transactions and query the blockchain

### Create message types
- Create a message type and its handler
  - `ignite scaffold message createPost title body`
### Define messages logic
- Create a variable of type `Post` with title and body from the message
- Append this `Post` to the store
  ```go
  func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
    // Get the context
    ctx := sdk.UnwrapSDKContext(goCtx)
    // Create variable of type Post
    var post = types.Post{
       Creator: msg.Creator,
       Title:   msg.Title,
       Body:    msg.Body,
    }
    // Add a post to the store and get back the ID
    id := k.AppendPost(ctx, post)
    // Return the ID of the post
    return &types.MsgCreatePostResponse{Id: id}, nil
  }
  ```
### Create a post
- Start your chain
  - `ignite chain server`
- Create a post
  - `blogd tx blog create-post foo bar --from alice`
### Display posts
- Scaffold a query
  - `ignite scaffold query posts --response title,body`

## Developer Resources
- Cosmos [SDK](https://docs.cosmos.network/)
- [Tendermint](https://docs.tendermint.com/) Core
- Cosmos [Hub](https://hub.cosmos.network/main/hub-overview/overview.html)
- [IBC](https://ibc.cosmos.network/)