package main

var player *Player
var doorOpened bool

func initGame() {
	kitchen := newRoom("кухня", kitchenLook)
	corridor := newRoom("коридор", corridorLook)
	room := newRoom("комната", roomLook)
	street := newRoom("улица", streetLook)

	connectRoom(kitchen, "коридор", corridor)
	connectRoom(corridor, "кухня", kitchen)
	connectRoom(corridor, "комната", room)
	connectRoom(corridor, "улица", street)
	connectRoom(room, "коридор", corridor)
	connectRoom(street, "домой", corridor)

	addItemToRoom(kitchen, "чай", "стол")
	addItemToRoom(room, "ключи", "стол")
	addItemToRoom(room, "конспекты", "стол")
	addItemToRoom(room, "рюкзак", "стул")

	player = newPlayer(kitchen)

	doorOpened = false

	initCommands()
}
