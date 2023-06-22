package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/Gatusko/pokedex/internal"
	"os"
)

var mapOfCommands = make(map[string]cliCommand)

var currentUrl *string
var previousUrl *string

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
	}
}
func printAllAreas(areas internal.Areas) {
	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
}

func mapArea() error {
	if &currentUrl == nil {
		return errors.New("There is no more areas use mapb for previous Area")
	}
	areas, err := internal.GetAreas(*currentUrl)
	if err != nil {
		return errors.New("Errror retrieving the Areas")
	}
	fmt.Print(areas)
	currentUrl = areas.Next
	previousUrl = areas.Previous
	printAllAreas(areas)
	return nil
}

func BMapArea() error {
	if previousUrl == nil {
		return errors.New("There is no more areas use mapb for previous Area")
	}
	areas, err := internal.GetAreas(*previousUrl)
	if err != nil {
		return errors.New("Errror retrieving the Areas")
	}
	currentUrl = areas.Next
	previousUrl = areas.Previous
	printAllAreas(areas)
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

func main() {
	firstUrl := "https://pokeapi.co/api/v2/location-area/"
	currentUrl = &firstUrl
	fmt.Println("POKEDEX")
	// init a new Map
	mapOfCommands = createMap()
	// this make a Scan
	scanner := bufio.NewScanner(os.Stdin)
	// for each read line
	fmt.Print("Pokedex>")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		command, ok := mapOfCommands[scanner.Text()]
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
