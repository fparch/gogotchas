package main

import "fmt"

type Evaluator interface {
	Evaluate() bool
}

type Policy struct{}

func (p Policy) Evaluate() bool {
	return true
}

func Map[T, U any](items []T, fn func(T) U) []U {
	result := make([]U, 0)
	for _, item := range items {
		result = append(result, fn(item))
	}
	return result
}

type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](values []T) T {
	var sum T
	for _, v := range values {
		sum += v
	}
	return sum
}

const FeeBPS int64 = 750
const BasisPoints int64 = 10_000

type Dollars int64

func (d Dollars) Add(other Dollars) Dollars {
	return d + other
}

func (d Dollars) Sub(other Dollars) Dollars {
	return d - other
}

func (d Dollars) ApplyBPS(bps int64) Dollars {
	// TODO: Perform integer arithmetic and round appropriately.
	return Dollars(
		(int64(d)*bps + BasisPoints/2) / BasisPoints,
	)

}

func (d Dollars) String() string {
	l := d / 100
	r := d % 100
	return fmt.Sprintf("$%d.%02d", l, r)
}

func USD(dollars, cents int64) Dollars {
	return Dollars(dollars*100 + cents)
}

func main() {
	var evaluator Evaluator = Policy{}
	fmt.Printf("Evaluation result: %b", evaluator.Evaluate())

	numbers := []int{1, 2, 3, 4}

	result := Map(numbers, func(n int) string {
		return fmt.Sprintf("ID-%d", n)
	})

	fmt.Println(result)

	names := []string{"Alice", "Bob", "Charlotte"}

	lengths := Map(names, func(s string) int {
		return len(s)
	})

	fmt.Println(lengths)

	numbers2 := []int{10, 20, 30, 40}

	fmt.Println(Sum(numbers2))

	values := []float64{1.5, 2.25, 3.75}

	fmt.Println(Sum(values))

	prices := []Dollars{USD(19, 99), USD(24, 50), USD(5, 51)}

	total := Sum(prices)

	fmt.Println(total)
	fmt.Printf("%T\n", total)

	fmt.Println(USD(10, 50)) // $10.50
	fmt.Println(USD(0, 99))  // $0.99
	fmt.Println(USD(42, 0))  // $42.00

	fmt.Println(USD(10, 0).ApplyBPS(FeeBPS)) // $0.75
	fmt.Println(USD(100, 0).ApplyBPS(250))   // $2.50
	fmt.Println(USD(1, 0).ApplyBPS(FeeBPS))  // $0.08
	fmt.Println(USD(0, 10).ApplyBPS(FeeBPS)) // $0.01

	a := []int{1, 2, 3, 4}
	//b := a[:2]
	b := a[:2:2]

	b = append(b, 99)
	fmt.Println(len(b), cap(b))

}
