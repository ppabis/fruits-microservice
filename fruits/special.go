package fruits

import "strings"

// Checks if a fruit is special
func IsFruitSpecial(fruit string) bool {
	fruit = strings.ToLower(fruit)
	return fruit == "pineapple" || fruit == "kiwi"
}
