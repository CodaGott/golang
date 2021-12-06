package main

import "fmt"

type taxSystem interface {
	calculateTax() int
}

type indianTax struct {
	taxPercentage int
	income        int
}

func (i *indianTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

type singapore struct {
	taxPercentage int
	income        int
}

func (s *singapore) calculateTax() int {
	tax := s.income * s.taxPercentage / 100
	return tax
}

type usaTax struct {
	taxPercentage int
	income        int
}

func (u *usaTax) calculateTax() int {
	tax := u.income * u.taxPercentage / 100
	return tax
}

func main() {
	indianTax := &indianTax{
		taxPercentage: 30,
		income:        1000,
	}

	singaporeTax := &singapore{
		taxPercentage: 10,
		income:        2000,
	}

	usaTax := &usaTax{
		taxPercentage: 40,
		income: 2000,
	}

	taxSystem := []taxSystem{indianTax, singaporeTax, usaTax}

	totalTax := calculateTotalTax(taxSystem)

	fmt.Printf("\nTotal tax is %d\n", totalTax)
}

func calculateTotalTax(taxSystem []taxSystem) int{
	totalTax := 0
	for _, t := range taxSystem {
		totalTax += t.calculateTax()
	}
	return totalTax
}