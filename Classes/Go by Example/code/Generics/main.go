package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	fmt.Printf("\nMapKeys() type of 'r'...: %T\n", r)
	fmt.Printf("          content of 'r': %+v\n", r)
	for k := range m {
		r = append(r, k)
	}
	fmt.Printf("\nMapKeys() content of 'r': %+v\n", r)
	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	fmt.Printf("\nPush() type of 'v'.....: %T\n", v)
	fmt.Printf("       content of 'v'..: %+v\n", v)
	fmt.Printf("       type of 'lst'...: %T\n", lst)
	fmt.Printf("       content of 'lst': %+v\n", lst)
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
	fmt.Printf("\nPush() content of 'v'..: %+v\n", v)
	fmt.Printf("       content of 'lst': %+v\n", lst)
}

func (lst *List[T]) GetAll() []T {
	fmt.Printf("\nGetAll() type of 'lst'.....: %T\n", lst)
	fmt.Printf("         content of 'lst'..: %+v\n", lst)
	var elems []T
	fmt.Printf("\nGetAll() type of 'elems'...: %T\n", lst)
	fmt.Printf("         content of 'elems': %+v\n", lst)
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	fmt.Printf("\nGetAll() content of 'elems': %+v\n", lst)
	fmt.Printf("         content of 'lst'..: %+v\n", lst)
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	fmt.Printf("main type of 'm'...: %T\n", m)
	fmt.Printf("     content of 'm': %+v\n", m)

	fmt.Println("\nmain keys:", MapKeys(m))

	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("\nmain list:", lst.GetAll())
}
