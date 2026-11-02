package main

func main() {
	cfg := &config{
		commands: getCommands(),
		pokedex:  make(map[string]Pokemon),
	}
	startRepl(cfg)
}