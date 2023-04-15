package fruits

import "strings"

func IsFruitSpecial(fruit string) bool {
	// Checks if a fruit is special
	fruit = strings.ToLower(fruit)
	return fruit == "pineapple" || fruit == "kiwi"
}
