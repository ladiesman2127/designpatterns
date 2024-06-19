package factory

import (
	"github.com/ladiesman2127/designpatterns/burger"
)

type BurgerFactory struct{}

func New() *BurgerFactory {
	return &BurgerFactory{}
}

func (bf *BurgerFactory) createCheeseBurger() *burger.Burger {
	ingredients := []string{"bun", "cheese", "beef-patty", "bun"}
	return burger.NewBurger(&ingredients)
}

func (bf *BurgerFactory) createBigMac() *burger.Burger {
	ingredients := []string{"bun", "beef-patty", "pickles", "cheese", "bun", "onions", "cheese", "bun"}
	return burger.NewBurger(&ingredients)
}

func (bf *BurgerFactory) createDoubleCheeseBurger() *burger.Burger {
	ingredients := []string{"bun", "cheese", "beef-patty", "cheese", "beef-patty", "bun"}
	return burger.NewBurger(&ingredients)
}
