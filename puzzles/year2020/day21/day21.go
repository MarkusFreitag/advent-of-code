package day21

import (
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

var (
	rgxIngredients = regexp.MustCompile(`^(.*?)\(`)
	rgxAllergens   = regexp.MustCompile(`\(contains\s(.*?)\)`)
)

type Food struct {
	Ingredients []string
	Allergens   []string
}

func parseLine(line string) Food {
	f := Food{
		Ingredients: make([]string, 0),
		Allergens:   make([]string, 0),
	}
	matches := rgxIngredients.FindAllStringSubmatch(line, 1)[0]
	f.Ingredients = append(f.Ingredients, strings.Fields(matches[1])...)
	matches = rgxAllergens.FindAllStringSubmatch(line, 1)[0]
	for _, allergen := range strings.Split(matches[1], ",") {
		f.Allergens = append(f.Allergens, strings.TrimSpace(allergen))
	}
	return f
}

func intersection(a, b []string) []string {
	inter := make([]string, 0)
	for _, s := range a {
		if util.StrInSlice(s, b) {
			inter = append(inter, s)
		}
	}
	return inter
}

func Part1(input string) (string, error) {
	foods := make([]Food, 0)
	for _, line := range strings.Split(input, "\n") {
		foods = append(foods, parseLine(line))
	}

	allergens := make(map[string][]string)
	ingredientsCounter := make(map[string]int)

	for _, food := range foods {
		for _, allergen := range food.Allergens {
			if v, ok := allergens[allergen]; ok {
				allergens[allergen] = intersection(v, food.Ingredients)
			} else {
				allergens[allergen] = food.Ingredients
			}
		}
		for _, ingredient := range food.Ingredients {
			v := ingredientsCounter[ingredient]
			ingredientsCounter[ingredient] = v + 1
		}
	}

	for _, ingredients := range allergens {
		for _, ingredient := range ingredients {
			delete(ingredientsCounter, ingredient)
		}
	}

	var sum int
	for _, num := range ingredientsCounter {
		sum += num
	}

	return strconv.Itoa(sum), nil
}

func Part2(input string) (string, error) {
	foods := make([]Food, 0)
	for _, line := range strings.Split(input, "\n") {
		foods = append(foods, parseLine(line))
	}

	allergens := make(map[string][]string)

	for _, food := range foods {
		for _, allergen := range food.Allergens {
			if v, ok := allergens[allergen]; ok {
				allergens[allergen] = intersection(v, food.Ingredients)
			} else {
				allergens[allergen] = food.Ingredients
			}
		}
	}

	for {
		var changed bool
		for allergen, ingredients := range allergens {
			if len(ingredients) > 1 {
				continue
			}
			for a, i := range allergens {
				if allergen == a {
					continue
				}
				for idx, item := range i {
					if item == ingredients[0] {
						allergens[a] = append(allergens[a][:idx], allergens[a][idx+1:]...)
						changed = true
					}
				}
			}
		}
		if !changed {
			break
		}
	}

	keys := make([]string, 0)
	for a := range allergens {
		keys = append(keys, a)
	}
	sort.Strings(keys)

	dangerous := make([]string, len(keys))
	for idx, key := range keys {
		dangerous[idx] = allergens[key][0]
	}

	return strings.Join(dangerous, ","), nil
}
