package main

import (
	"fmt"
)

func NormalTransform[T1, T2 any](list []T1, transform func(T1) T2) []T2 {
	transformed := make([]T2, len(list))

	for i, t := range list {
		transformed[i] = transform(t)
		fmt.Println("in func - element:", t, "transformed:", transformed[i])
	}

	fmt.Println("transformed:", transformed)

	return transformed
}

func main() {
	list := []int{1, 2, 3, 4, 5}
	doubleFunc := func(i int) int { return i * 2 }

	for i, num := range NormalTransform(list, doubleFunc) {
		fmt.Println("post func - idx:", i, "list:", list[i], "doubleFunced:", num)
	}
}
