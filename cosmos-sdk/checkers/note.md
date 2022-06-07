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