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
	return "todo move"
}

func handleTakeCommand(cmd []string) string {
	return "todo take"
}

func handleWearCommand(cmd []string) string {
	return "todo wear"
}

func handleApplyCommand(cmd []string) string {
	return "todo apply"
}
