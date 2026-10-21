package main

import "fmt"

type PolicyError struct {
	Reason string
}

func (e *PolicyError) Error() string {
	return e.Reason
}

func evaluate(ok bool) error {
	if !ok {
		return &PolicyError{Reason: "NOT OK!!!!"}
	}
	return nil
}

func main() {
	err := evaluate(true)
	fmt.Println(err == nil)
	err = evaluate(false)
	fmt.Println(err == nil)
}
