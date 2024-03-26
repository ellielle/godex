# Godex

Godex is a CLI-based Pokedex using the [PokeAPI](https://pokeapi.co/). It allows you to explore regions of the world and lists which Pokemon are available for capture.

# ⚙️Installation

Clone repo:

```bash
$ git clone git@github.com:ellielle/godex
```

Build and run:

```bash
$ go build -o godex && ./godex
```

Run tests:

```bash
$ go test ./...
```

# Usage

Godex has the following commands available:

```
help                 show the help menu
exit                 exit the program
map                  explore and show 20 regions from the Pokemon API
mapb                 retrace your steps to the previous 20 regions
explore <map-name>   explore a map and see what Pokemon can be caught
catch <pokemon>      attempt to catch a Pokemon and add it to your Pokedex
inspect <pokemon>    inspect a Pokemon you've caught
pokedex              view a list of the Pokemon in your Pokedex
```
