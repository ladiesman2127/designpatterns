package burger

type Burger struct {
	ingredients *[]string
}

func NewBurger(ingredients *[]string) *Burger {
	return &Burger{ingredients: ingredients}
}

func (b *Burger) addIngredient(ing string) {
	*b.ingredients = append(*b.ingredients, ing)
}
