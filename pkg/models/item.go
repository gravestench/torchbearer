package models

import (
	"github.com/google/uuid"
)

type TreasureValueType string

const (
	TreasureCopper TreasureValueType = "copper"
	TreasureSilver TreasureValueType = "silver"
	TreasureGold   TreasureValueType = "gold"
)

const (
	RequirementNotPocketed = 0
	RequirementNotPacked   = 0
	RequirementNotCarried  = 0
	RequirementNotBelted   = 0
	RequirementNotWorn     = 0
)

func NewItem() *Item {
	i := &Item{}
	i.ID = uuid.New()
	i.Treasure.Type = TreasureCopper

	return i
}

func NewStorageItem(numSlots int) *Item {
	i := NewItem()
	i.Storage = make([]*Item, numSlots)
	return i
}

type Item struct {
	ID              uuid.UUID
	CurrentLocation uuid.UUID // world, settlement, townsfolk, adventurer, etc
	Name            string
	Description     string
	Treasure        struct {
		Type  TreasureValueType
		Value int
	}
	IsWeapon    bool
	Requirement struct {
		Pack   int
		Carry  int
		Worn   int
		Belt   int
		Pocket int
	}
	Storage []*Item
}

func (i *Item) SetLocation(id uuid.UUID) {
	i.CurrentLocation = id
}
