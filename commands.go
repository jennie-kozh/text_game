package main

import "strings"

var commandHandlers map[string]func(command []string) string

func initCommands() {
	commandHandlers = map[string]func(command []string) string{
		"осмотреться": handleLookCommand,
		"идти":        handleMoveCommand,
		"взять":       handleTakeCommand,
		"надеть":      handleWearCommand,
		"применить":   handleApplyCommand,
	}
}

func handleCommand(cmd string) string {
	parts := strings.Fields(cmd)
	action := parts[0]
	args := parts[1:]

	handler, ok := commandHandlers[action]
	if !ok {
		return "неизвестная команда"
	}
	return handler(args)
}

func handleLookCommand(cmd []string) string {
	if player.location.look == nil {
		return "ничего интересного"
	}
	return player.location.look()
}

func handleMoveCommand(cmd []string) string {
	if len(cmd) < 1 {
		return "куда идти?"
	}

	direction := cmd[0]

	nextRoom, ok := player.location.exits[direction]
	if !ok {
		return "нет пути в " + direction
	}

	if player.location.name == "коридор" && direction == "улица" && !doorOpened {
		return "дверь закрыта"
	}

	player.location = nextRoom

	if player.location.name == "кухня" {
		return "кухня, ничего интересного. можно пройти - коридор"
	}

	if player.location.name == "комната" {
		return "ты в своей комнате. можно пройти - коридор"
	}

	if player.location.look == nil {
		return "ничего интересного"
	}

	return player.location.look()
}

func handleTakeCommand(cmd []string) string {
	if len(cmd) < 1 {
		return "что взять?"
	}

	itemName := cmd[0]

	item, ok := player.location.items[itemName]
	if !ok {
		return "нет такого"
	}

	if !player.hasBag {
		return "некуда класть"
	}

	player.inventory[itemName] = true
	delete(player.location.items, itemName)

	return "предмет добавлен в инвентарь: " + item.name
}

func handleWearCommand(cmd []string) string {
	if len(cmd) < 1 {
		return "нечего надеть"
	}

	itemName := cmd[0]

	_, ok := player.location.items[itemName]
	if !ok {
		return "нет такого"
	}

	if itemName != "рюкзак" {
		return "нельзя надеть"
	}

	player.hasBag = true
	delete(player.location.items, itemName)

	return "вы надели: рюкзак"
}

func handleApplyCommand(cmd []string) string {
	if len(cmd) < 2 {
		return "не к чему применить"
	}

	itemName := cmd[0]
	targetName := cmd[1]

	if !player.inventory[itemName] {
		return "нет предмета в инвентаре - " + itemName
	}

	if itemName == "ключи" && targetName == "дверь" {
		doorOpened = true
		return "дверь открыта"
	}

	return "не к чему применить"
}
