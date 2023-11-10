package world

import (
	"math/rand"
)

// Generate a random fantasy name
func (s *Service) generateNewSettlementName() string {
	var syllables = []string{
		"br", "en", "val", "mir", "step", "pe", "ka", "ro", "th", "ran", "dor", "wyn", "sta", "tu", "sha",
		"al", "un", "la", "er", "zor", "qu", "str", "xi", "tri", "pyr", "ge", "li", "ma", "o", "pe", "or",
		"vel", "gar", "thor", "ark", "ur", "ber", "el", "fin", "va", "leth", "for", "sar", "vyr", "dra", "mol",
		"gen", "syl", "ny", "thu", "grim", "zel", "roth", "yel", "xen", "on", "in", "zer", "kra", "ryn", "an",
		"shad", "thor", "thal", "drak", "shor", "kron", "gha", "zun", "fel", "bal", "cra", "tar", "skal", "dre",
		"tyr", "mur", "wynn", "lor", "har", "bor", "thir", "xar", "vor", "jor", "vald", "yrr", "fyr", "khal",
		"tir", "vorn", "trin", "skyl", "gren", "ul", "zar", "sor", "xir", "fyr", "dal", "garth", "ynd", "ril",
		"ra", "len", "nor", "thra", "mar", "roth", "ral", "ver", "thul", "dral", "nyl", "tril",
	}

	name := ""
	syllableCount := rand.Intn(2) + 2 // Generate 2-3 syllable names

	for i := 0; i < syllableCount; i++ {
		syllable := syllables[rand.Intn(len(syllables))]
		name += syllable
	}

	// Capitalize the first letter
	name = string(name[0]-'a'+'A') + name[1:]

	return name
}
