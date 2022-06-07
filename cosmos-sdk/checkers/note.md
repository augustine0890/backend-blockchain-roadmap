# Make a Checkers Blockchain

## Store Object
- The Stored Game Object
- Protobuf Objects
- Query.proto
- Protobuf service interfaces
- Unit test
### Defining the rule set
- Copy the rules file into a `rules` folder inside module
  - `curl https://raw.githubusercontent.com/batkinson/checkers-go/a09daeb1548dd4cc0145d87c8da3ed2ea33a62e3/checkers/checkers.go | sed 's/package checkers/package rules/' > x/checkers/rules/checkers.go`
### The stored game object
- Red player: string, the serialized address
- Black player: string, the serialized address
- Game proper: string, the game as it is serialized by the rules file
- Player to play nextL string
- Call the object that contains the counter `NextGame` and instruct Ignite CLI with `scaffold single`:
  - `ignite scaffold single nextGame idValue:uint --module checkers --no-message`
- You need a map because you're storing games by ID. Instruct Ignite CLI with `scaffold map` using the `StoredGame` name:
  - `ignite scaffold map storedGame game turn red black --module checkers --no-message`
### Additional helper functions
