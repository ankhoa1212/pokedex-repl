# pokedex-repl
A command-line REPL program using the PokéAPI written in Go.

# Usage

1. Start the REPL environment for the pokedex application.
```
go run .
```
2. Enter commands to interact with the pokedex, such as:
    - ```map```: Show location areas to explore.
    - ```help```: Show available commands and usage information.
    - ```exit```: Close the pokedex.

# Roadmap:
- [x] Use [PokéAPI](https://pokeapi.co/docs/v2) to get data
- [ ] (command) Be able to go to navigate to next/previous map areas
    - [ ] [Convert JSON into Go struct](https://blog.boot.dev/golang/json-golang/?_gl=1*1optvqd*_gcl_au*NjQ2NzA1NjU2LjE3NTI2MDQ1NzE.*_ga*MTc5Nzc5MDkzMi4xNzUyNjA0NTcx*_ga_M7P2PBGN8N*czE3NTI2NTcwNjMkbzQkZzEkdDE3NTI2NTcwNzEkajUyJGwwJGg0NDQ3NjM1#example-unmarshal-json-to-struct-decode)
- [ ] [Caching](https://en.wikipedia.org/wiki/Cache_(computing))
- [ ] (command) Map exploration by listing pokemon at a location
- [ ] (command) Pokemon catching
- [ ] (command) Inspection of specific pokemon
- [ ] (command) Pokedex - list all pokemon caught
