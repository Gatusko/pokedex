package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/Gatusko/pokedex/internal"
	"os"
	"strings"
	"time"
)

var mapOfCommands = make(map[string]cliCommand)

var currentUrl *string
var previousUrl *string
var secondLineCommand *string

const baseExploreURL = "https://pokeapi.co/api/v2/location-area/"
const basePokemonUrl = "https://pokeapi.co/api/v2/pokemon/"

type urls struct {
	currentUrl string
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func createMap() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get Next 20 Areas",
			callback:    mapArea,
		},
		"bmap": {
			name:        "map",
			description: "Get Previous 20 ",
			callback:    BMapArea,
		},
		"explore": {
			name:        "explore",
			description: "explore an area of pokemons ",
			callback:    ExplorePokemonCommand,
		},
		"catch": {
			name:        "catch",
			description: "catch a pokemon ",
			callback:    CatchPokemon,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect a pokemon ",
			callback:    InspectPokemon,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all your pokemonCatched ",
			callback:    ListAllPokemon,
		},
	}
}

func CatchPokemon() error {
	url := basePokemonUrl + *secondLineCommand
	pokemon, err := client.GetPokemon(url)
	if err != nil {
		errorName := "Pokemon not founnd: " + *secondLineCommand
		return errors.New(errorName)
	}
	fmt.Println("Trying to catch:", *secondLineCommand)
	if client.Pokedex.CatchPokemon(pokemon) {
		fmt.Println("Pokemon catched: ", *secondLineCommand)
		return nil
	} else {
		fmt.Println("Pokemon escaped:", *secondLineCommand)
		return nil
	}
}

func ListAllPokemon() error {
	fmt.Println("Your Pokedex: ")
	client.ListAllPokemon()
	return nil
}

func InspectPokemon() error {
	pokemon, ok := client.Pokedex.GetPokemon(*secondLineCommand)
	if !ok {
		return errors.New("Pokemon not found on pokedex")
	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stats := range pokemon.Stats {
		fmt.Println("-", stats.Stat.Name, ":", stats.BaseStat)
	}
	fmt.Println("Types")
	for _, types := range pokemon.Types {
		fmt.Println("-", types.Type.Name)
	}
	return nil
}
func printAllAreas(areas internal.Areas) {
	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
}

func printAllPokemons(pokemons internal.ExplorePokemon) {
	for _, pokemon := range pokemons.PokemonEncounters {
		fmt.Println(pokemon.Name)
	}
}

func mapArea() error {
	if &currentUrl == nil {
		return errors.New("There is no more areas use mapb for previous Area")
	}
	areas, err := client.GetAreas(*currentUrl)
	if err != nil {
		return errors.New("Errror retrieving the Areas")
	}
	currentUrl = areas.Next
	previousUrl = areas.Previous
	printAllAreas(areas)
	return nil
}

func BMapArea() error {
	if previousUrl == nil {
		return errors.New("There is no more areas use mapb for previous Area")
	}
	areas, err := client.GetAreas(*previousUrl)
	if err != nil {
		return errors.New("Errror retrieving the Areas")
	}
	currentUrl = areas.Next
	previousUrl = areas.Previous
	printAllAreas(areas)
	return nil
}

func ExplorePokemonCommand() error {
	if secondLineCommand == nil {
		return errors.New("Need to select Area to explore Pokemons")
	}
	exploreUrl := baseExploreURL + *secondLineCommand
	exporeArea, err := client.ExplorePokemon(exploreUrl)

	if err != nil {
		return err
	}
	printAllPokemons(exporeArea)
	return nil
}

func commandHelp() error {
	fmt.Println("Commands:")
	for _, commands := range mapOfCommands {
		fmt.Printf("%s : %s\n", commands.name, commands.description)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Exiting Program")
	os.Exit(0)
	return nil
}

var client internal.Client

func main() {
	firstUrl := "https://pokeapi.co/api/v2/location-area/"
	currentUrl = &firstUrl
	fmt.Println("POKEDEX")
	cache := internal.Cache{}

	cache.NewCache(time.Duration(5 * time.Second))
	client.NewClient(cache)
	// init a new Map
	mapOfCommands = createMap()
	// this make a Scan
	scanner := bufio.NewScanner(os.Stdin)
	// for each read line
	fmt.Print("Pokedex>")
	for scanner.Scan() {
		scannedString := scanner.Text()
		splitedCommand := strings.Split(scannedString, " ")
		if len(splitedCommand) > 2 {
			fmt.Println("too many arguments", scanner.Text())
			fmt.Print("Pokedex>")
			continue
		}

		if len(splitedCommand) == 2 {
			secondLineCommand = &splitedCommand[1]
		}

		command, ok := mapOfCommands[splitedCommand[0]]
		if !ok {
			fmt.Println("command doesnt exist :", scanner.Text())
			fmt.Print("Pokedex>")
			continue
		}
		errorCallback := command.callback()
		if errorCallback != nil {
			fmt.Println("Error happened on command :", command.name, errorCallback)
			fmt.Print("Pokedex>")
			continue
		}
		//call the callback

		fmt.Print("Pokedex>")
	}
}
