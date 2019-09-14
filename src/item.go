package main

import (
	"math/rand"
	"time"
)

type rarity map[string]bool

type item struct {
	name       string
	rarity     rarity
	damage     int
	experience int
}

func newItem(name string) *item {
	var rarity rarity = generateStats()
	return &item{name, rarity, 15, 0}
}

func generateStats() map[string]bool {
	stat := make(map[string]bool)

	rand.Seed(time.Now().UnixNano())

	if rand.Intn(100) > 69 {
		stat["Legendary"] = true
		if rand.Intn(100) > 69 {
			stat["Ultimate"] = true
			if rand.Intn(100) > 84 {
				stat["Ancient"] = true
			}
		}
	}
	return stat
}

func (itm item) getStatedName() (statedName string) {
	for k, v := range itm.rarity {
		if k == "Ancient" && v {
			statedName = "Ancient Ultimate Legendary - " + itm.name
			return
		}

		if k == "Ultimate" && v {
			statedName = "Ultimate " + statedName
		}
		if k == "Legendary" && v {
			statedName += "Legendary - "
		}
	}
	statedName += itm.name
	return
}
