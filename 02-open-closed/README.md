# Open-Closed Principle(OCP)
This principle states that types should be open for extension, but closed for modification.


## Analogy

Lets assume we're operating an online store, and we want end users to be able to filter the items by certain criteria ie. by price, size etc

Sw, we have a Product with some description:

```
type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}
```

Lets suppose that our first requirement is to create a Filter type which has the ability to filter by color:

```
type Filter struct {
    //
}

func (f *Filter) FilterbyColor(products []product, color Color) []*Product {
    result := make([]*Product, 0)

    for i, v := range products {
        if v.color == color {
            result = append(result, &products[i])
        }
    }

    return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	fmt.Printf("Green products: \n")

	f := Filter{}
	for _, v := range f.FilterbyColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

}
```

However, imagine we implement this filtering by color and then we also need a filtering by size functionality. This means we have to go back to the Filter type and add another method to it.

```
func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}

	return result
}
```

We now have a new Filter method which filters by size but now we need to implement another method which filters by color and size:

```
func (f *Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {
    result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}
```

What we are doing now is a violation of the Open-Closed principle because as we are going back and modifying/adding additional methoods on the Product, we are sort of interferring something that has already been writtena and tested.

We want to leave the Filter type alone and not going back to keep adding more and more methods. We want to have an extendible setup which is what we get if we use the *Specification Pattern*.

The **Specification Pattern** is an Enterprise Pattern. It has a bunch of interfaces for flexibility.

The first thing we do is to implement a *Specification Interface*:

```
type Specification interface {
    IsSatisfied(p *Product) bool
}
```

The idea behind this interface is that we're testing whether or not a Product(specified above by a pointer) satisfies some criteria. For ex, if we want to check for color, we make a color specification where we specify the color we want to filter on:

```
type ColorSpecification struct {
   color Color
}
```

And then we have a method defined on the ColorSpecification

```
func (c ColorSpecification) IsSatisfied(p *Product) bool {
    return p.color == c.color
}
```

We now need a different filter which we'll likely never modify:

```
type BetterFilter struct {

}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for _, v := range products {
		if spec.isSatisfied(&v) {
			result = append(result, &v)
		}
	}

	return result
}
```

The approach with the *Specification Pattern* gives more flexibility because if we want to add new filters such as filter by type, we only need to make a new specification i.e. `SizeSpecification` and make sure it conforms to the `Specification` interface.

```
type SizeSpecification struct {
   size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
    return p.size == s.size
}
```

But wait, how do would we implement filtering by size and color using the *Specification pattern*. All we have to do in this case, is to make a *Composite Specification* which is also an illustration of the Composite Design Pattern.

A *Composite Specification* just combines two different specifications. So here, we can create a `AndSpecification` or similarly, we can also have a `OrSpecification`:

```
type  AndSpecification struct {
    first, second Specification
}

func (a AndSpecification) IsSatisfied (p *Product) bool {
    return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

```

Now we can test this out in our main funtion like so:

```
func main() {

    apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	fmt.Printf("Large Green products: \n")

	greenSpec := ColorSpecification{green}
    largeSpec := SizeSpecification{large}
    lgSpec := AndSpecification{greenSpec, largeSpec}

    f := BetterFilter{}
	for _, v := range f.Filter(products, lgSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}
}
```


## Recap

Thats pretty much all we have to do for the Open-Close Principle. With types, in this case the `interface` type is open for extension but closed for modification i.e. we're unlikely to ever modify the `Specification` interface. In a similar fashion, we're unlikely to ever modify the `BetterFilter` because there's its very flexible and there's no need for us to ever do so.