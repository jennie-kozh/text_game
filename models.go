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
