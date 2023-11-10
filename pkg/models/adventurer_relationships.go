package models

type AdventurerRelationships struct {
	Parents []Adventurer
	Mentor  *Adventurer
	Friend  *Adventurer
	Enemies []Adventurer
	Allies  []Adventurer
}
