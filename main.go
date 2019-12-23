package main

import (
	"fmt"
	"reflect"
)

type NumberInterface interface {
	number(a int) (int, error)
}

type FirstStruct struct{}

func (num FirstStruct) number(a int) (int, error) {
	return a + 1, nil
}

type SecondStruct struct{}

func (num SecondStruct) number(b int) (int, error) {
	return b + 1, nil
}

type Number []NumberInterface

func main() {
	numChanel := make(chan int, 2)
	totalChan := make(chan int)
	first := FirstStruct{}
	second := SecondStruct{}
	numList := Number{
		first,
		second,
	}

	for k, v := range numList {
		go func(i int, n NumberInterface) {
			num, _ := n.number(i)
			numChanel <- num
		}(k, v)
	}

	go func() {
		sumTotal := 0
		k := 0
		for i := 0; i < len(numList); i++ {
			select {
			case num := <-numChanel:
				sumTotal += num
				k++
				if k == len(numList) {
					totalChan <- sumTotal
				}
			}
		}
	}()
	total := <-totalChan
	fmt.Println(total)
	fmt.Println(reflect.ValueOf(total).Interface())
	fmt.Println(bar(1))

}

func foo(val int) (int, error) {
	v := val + 2
	return v, nil
}

func bar(v int) (int, error) {
	data, err := foo(v)
	if err != nil {
		fmt.Println(err)
	}
	data++
	return data, nil
}
