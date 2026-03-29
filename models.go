package main

type Room struct {
	name  string
	exits map[string]*Room
	items map[string]*Item
	look  func() string
}

type Item struct {
	name     string
	location string
}

type Player struct {
	location  *Room
	inventory map[string]bool
	hasBag    bool
}

func newRoom(name string) *Room {
	return &Room{
		name:  name,
		exits: make(map[string]*Room),
		items: make(map[string]*Item),
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
