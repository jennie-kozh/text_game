package main

import (
	"sort"
	"strings"
)

func newRoom(name string, look func() string) *Room {
	return &Room{
		name:  name,
		exits: make(map[string]*Room),
		items: make(map[string]*Item),
		look:  look,
	}
}

func connectRoom(from *Room, direction string, to *Room) {
	from.exits[direction] = to
}

func addItemToRoom(room *Room, itemName string, itemLocation string) {
	room.items[itemName] = &Item{
		name:     itemName,
		location: itemLocation,
	}
}

func newPlayer(room *Room) *Player {
	return &Player{
		location:  room,
		inventory: make(map[string]bool),
		hasBag:    false,
	}
}

func kitchenLook() string {
	if player.hasBag && player.inventory["конспекты"] {
		return "ты находишься на кухне, на столе: чай, надо идти в универ. можно пройти - коридор"
	}
	return "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"
}

func corridorLook() string {
	return "ничего интересного. можно пройти - кухня, комната, улица"
}

func roomLook() string {
	room := player.location

	tableItems := []string{}
	hasBag := false

	for _, item := range room.items {
		if item.location == "стол" {
			tableItems = append(tableItems, item.name)
		}
		if item.name == "рюкзак" {
			hasBag = true
		}
	}

	sort.Strings(tableItems)

	var result string

	if len(tableItems) > 0 {
		result = "на столе: " + strings.Join(tableItems, ", ")
	}

	if hasBag {
		if result != "" {
			result += ", на стуле: рюкзак"
		} else {
			result = "на стуле: рюкзак"
		}
	}

	if result == "" {
		result = "пустая комната"
	}

	return result + ". можно пройти - коридор"
}

func streetLook() string {
	return "на улице весна. можно пройти - домой"
}
