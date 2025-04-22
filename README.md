# golf-graphql
A simple graphql implementation of the interface for my golf app


## Example queries

```
mutation cplayer {
  createPlayer(input: { name: "foobar" }) {
    name
  }
}

query getPlayers {
  players {
    name
  }
}
```
