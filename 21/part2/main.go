package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ingredientsPerAllergen := make(map[string]map[string]bool)
	ingredientCounts := make(map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
		chunks := strings.Split(line, "|")
		ingredients := strings.Split(chunks[0], " ")
		allergens := strings.Split(chunks[1], " ")
		fmt.Println(line)
		fmt.Printf("ingredients: %s\n", ingredients)
		fmt.Printf("allergens: %s\n\n", allergens)

		for _, ingredient := range ingredients {
			ingredientCounts[ingredient]++
		}

		for _, allergen := range allergens {
			if ingredientsPerAllergen[allergen] == nil {
				ingredientsPerAllergen[allergen] = make(map[string]bool)

				for _, ingredient := range ingredients {
					ingredientsPerAllergen[allergen][ingredient] = true
				}
			} else {
				// we've already seen this allergen before
				// the responsible ingredient must be in both
				// ingredients lists
				newIngredients := make(map[string]bool)
				for _, ingredient := range ingredients {
					newIngredients[ingredient] = true
				}
				for ingredient := range ingredientsPerAllergen[allergen] {
					if _, isThere := newIngredients[ingredient]; !isThere {
						delete(ingredientsPerAllergen[allergen], ingredient)
					}
				}
			}

		}
	}

	possibleAllergensPerIngredient := make(map[string]map[string]bool)
	for allergen, ingredients := range ingredientsPerAllergen {
		for ingredient := range ingredients {
			if possibleAllergensPerIngredient[ingredient] == nil {
				possibleAllergensPerIngredient[ingredient] = make(map[string]bool)
			}
			possibleAllergensPerIngredient[ingredient][allergen] = true
		}
	}

	sum := 0
	for ingredient, count := range ingredientCounts {
		if _, hasPossibles := possibleAllergensPerIngredient[ingredient]; hasPossibles {
			continue
		}
		fmt.Printf("%s can't be an allergen, shows up in %d recipes\n", ingredient, count)
		sum += count
	}

	fmt.Println(sum)

	for ingredient, possibleAllergens := range possibleAllergensPerIngredient {
		fmt.Printf("%s could countain ", ingredient)
		for allergen := range possibleAllergens {
			fmt.Printf("%s ", allergen)
		}
		fmt.Println()
	}

	knownAllergens := make(map[string]string)
	for i := 0; i < 10; i++ {
		for ingredient, possibleAllergens := range possibleAllergensPerIngredient {
			if len(possibleAllergens) == 1 {
				allergen := ""
				for possible := range possibleAllergens {
					allergen = possible
				}
				knownAllergens[allergen] = ingredient
				fmt.Printf("%s must be %s\n", allergen, ingredient)
			}

			for allergen := range knownAllergens {
				if _, isPresent := possibleAllergens[allergen]; isPresent {
					delete(possibleAllergens, allergen)
				}
			}
		}
		if len(knownAllergens) == len(ingredientsPerAllergen) {
			fmt.Println("DONE")
			break
		}
	}
	// Arrange the ingredients alphabetically by their allergen
	allergens := make([]string, 0)
	for allergen := range knownAllergens {
		allergens = append(allergens, allergen)
	}
	sort.Strings(allergens)

	answer := ""
	for _, allergen := range allergens {
		answer += knownAllergens[allergen] + ","
	}
	fmt.Println(allergens)
	fmt.Println(answer)
}
