package models

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"strings"
)

// Description generates a generic description of a settlement in conversational English.
func (s *Settlement) Description() string {
	rng := rand.New(rand.NewSource(s.Seed))

	if s.description != "" {
		return s.description
	}

	var adjectives = []string{
		"vibrant", "quaint", "picturesque", "historic", "charming", "idyllic",
		"busy", "serene", "mysterious", "sprawling", "cozy", "enchanted",
		"peaceful", "lively", "colorful", "tranquil", "captivating", "majestic",
		"adventurous", "enchanting", "secluded", "exotic", "magical", "whimsical",
		"bustling", "scenic", "romantic", "hidden", "peaceful", "ancient",
		"storybook", "elegant", "mystical", "friendly", "welcoming", "charismatic",
	}

	// Pick random adjectives and traits
	randomAdjective1 := adjectives[rng.Intn(len(adjectives))]
	randomAdjective2 := adjectives[rng.Intn(len(adjectives))]

	// Choose a random facility from the notable facilities
	var notableFacilities []string
	for _, facility := range s.Facilities.List() {
		notableFacilities = append(notableFacilities, facility)
	}

	if len(notableFacilities) == 0 {
		return "This settlement has no notable facilities."
	}
	randomFacility := notableFacilities[rng.Intn(len(notableFacilities))]

	var traitDesc string

	// Generate the description using placeholders
	if len(s.Culture.Traits) == 3 {
		traitDesc = fmt.Sprintf("%s, %s, and %s", s.Culture.Traits[0].Name, s.Culture.Traits[1].Name, s.Culture.Traits[2].Name)
	}

	if len(s.Culture.Traits) == 2 {
		traitDesc = fmt.Sprintf("%s and %s", s.Culture.Traits[0].Name, s.Culture.Traits[1].Name)
	}

	traitDesc = strings.ToLower(traitDesc)
	s.description = fmt.Sprintf("The %s %s of %s, known for its %s %s. The people of %s are known for being %s.",
		randomAdjective1, s.Type, s.Name, randomAdjective2, randomFacility, s.Name, traitDesc)

	return s.description
}

// Description generates a generic description of a settlement in conversational English.
func (s *Settlement) Description2() string {
	rng := rand.New(rand.NewSource(s.Seed))

	if s.description != "" {
		return s.description
	}

	// Define sentence templates with different structures
	sentenceTemplates := []string{
		"The %s %s of %s is known for its %s %s. The people of %s are known for being %s.",
		"%s %s, located in %s, is famous for its %s %s. The residents are renowned for their %s.",
		"%s %s, nestled in %s, boasts a %s %s that is a highlight of the town. Its inhabitants are %s.",
		"%s %s, situated in %s, takes pride in its %s %s. The locals are described as %s.",
		"At %s %s, you'll find a %s %s that defines the town's character. Its populace is %s.",
	}

	var adjectives = []string{
		"vibrant", "quaint", "picturesque", "historic", "charming", "idyllic",
		"busy", "serene", "mysterious", "sprawling", "cozy", "enchanted",
		"peaceful", "lively", "colorful", "tranquil", "captivating", "majestic",
		"adventurous", "enchanting", "secluded", "exotic", "magical", "whimsical",
		"bustling", "scenic", "romantic", "hidden", "peaceful", "ancient",
		"storybook", "elegant", "mystical", "friendly", "welcoming", "charismatic",
	}

	// Pick random adjectives and traits
	randomAdjective1 := adjectives[rng.Intn(len(adjectives))]
	randomAdjective2 := adjectives[rng.Intn(len(adjectives))]

	// Choose a random facility from the notable facilities
	var notableFacilities []string
	for _, facility := range s.Facilities.List() {
		notableFacilities = append(notableFacilities, facility)
	}

	if len(notableFacilities) == 0 {
		return "This settlement has no notable facilities."
	}
	randomFacility := notableFacilities[rng.Intn(len(notableFacilities))]

	// Select a random sentence template
	randomTemplate := sentenceTemplates[rng.Intn(len(sentenceTemplates))]

	var traitDesc string

	// Generate the description using placeholders
	if len(s.Culture.Traits) == 3 {
		traitDesc = fmt.Sprintf("%s, %s, and %s", s.Culture.Traits[0].Name, s.Culture.Traits[1].Name, s.Culture.Traits[2].Name)
	}

	if len(s.Culture.Traits) == 2 {
		traitDesc = fmt.Sprintf("%s and %s", s.Culture.Traits[0].Name, s.Culture.Traits[1].Name)
	}

	traitDesc = strings.ToLower(traitDesc)
	s.description = fmt.Sprintf(randomTemplate, randomAdjective1, s.Type, s.Name, randomAdjective2, randomFacility, s.Name, traitDesc)

	return s.description
}

// GenerateSeedFromString generates an int64 seed from a given string.
func GenerateSeedFromString(input string) int64 {
	// Create a new FNV-1a hash instance
	hash := fnv.New64a()

	// Write the input string to the hash
	_, _ = hash.Write([]byte(input))

	// Get the resulting hash value as an int64 seed
	seed := int64(hash.Sum64())

	return seed
}
