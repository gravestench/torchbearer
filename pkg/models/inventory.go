package models

type Inventory struct {
	Worn struct {
		Head      *Item
		Neck      *Item
		HandLeft  *Item
		HandRight *Item
		Torso1    *Item
		Torso2    *Item
		Torso3    *Item
		Belt1     *Item
		Belt2     *Item
		Belt3     *Item
		Legs      *Item
		Feet      *Item
	}
	Carried struct {
		HandLeft  *Item
		HandRight *Item
	}
}
