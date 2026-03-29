package main

var player *Player
var doorOpened bool

func initGame() {
	kitchen := newRoom("кухня")
	corridor := newRoom("коридор")
	room := newRoom("комната")
	street := newRoom("улица")

	connectRoom(kitchen, "кухня", corridor)
	connectRoom(corridor, "кухня", kitchen)
	connectRoom(corridor, "комната", room)
	connectRoom(corridor, "улица", street)
	connectRoom(room, "коридор", corridor)
	connectRoom(street, "домой", corridor)

	addItemToRoom(kitchen, "чай", "стол")
	addItemToRoom(room, "ключи", "стол")
	addItemToRoom(room, "конспекты", "стол")
	addItemToRoom(room, "рюкзак", "стул")

	//TODO add look funcs

	player = newPlayer(kitchen)

	doorOpened = false

	initCommands()
}
