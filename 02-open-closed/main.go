package main

import "fmt"

// Color type
type Color int

const (
	red Color = iota
	green
	blue
)

// Size type
type Size int

const (
	small Size = iota
	medium
	large
)

// Specification interface
type Specification interface {
	IsSatisfied(p *Product) bool
}

// Product struct
type Product struct {
	name  string
	color Color
	size  Size
}

// ColorSpecification struct
type ColorSpecification struct {
	color Color
}

// IsSatisfied method on ColorSpecification
func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

// SizeSpecification method
type SizeSpecification struct {
	size Size
}

// IsSatisfied method on SizeSpecification
func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

// AndSpecification Composite Specification
type AndSpecification struct {
	first, second Specification
}

// IsSatisfied method on AndSpecification
func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

// BetterFilter struct
type BetterFilter struct {
}

// Filter method
func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}

	return result
}

// // Filter struct
// type Filter struct {
// 	//
// }

// // FilterbyColor method
// func (f *Filter) FilterbyColor(products []Product, color Color) []*Product {
// 	result := make([]*Product, 0)

// 	for i, v := range products {
// 		if v.color == color {
// 			result = append(result, &products[i])
// 		}
// 	}

// 	return result
// }

// // FilterBySize method
// func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
// 	result := make([]*Product, 0)

// 	for i, v := range products {
// 		if v.size == size {
// 			result = append(result, &products[i])
// 		}
// 	}

// 	return result
// }

// // FilterBySizeAndColor method
// func (f *Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {
// 	result := make([]*Product, 0)

// 	for i, v := range products {
// 		if v.size == size && v.color == color {
// 			result = append(result, &products[i])
// 		}
// 	}

// 	return result
// }

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	// Old Way of doing things - Doesn't follow OCP
	//fmt.Printf("Green products: \n")
	// f := Filter{}
	// for _, v := range f.FilterbyColor(products, green) {
	// 	fmt.Printf(" - %s is green\n", v.name)
	// }

	// Using Specification Pattern
	// greenSpec := ColorSpecification{green}
	// fmt.Printf("Green products: \n")

	// f := BetterFilter{}
	// for _, v := range f.Filter(products, greenSpec) {
	// 	fmt.Printf(" - %s is green\n", v.name)
	// }

	// Composite Specification
	greenSpec := ColorSpecification{green}
	largeSpec := SizeSpecification{large}
	lgSpec := AndSpecification{greenSpec, largeSpec}

	fmt.Printf("Large Green products: \n")
	f := BetterFilter{}
	for _, v := range f.Filter(products, lgSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}

}
