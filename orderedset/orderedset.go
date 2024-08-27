package orderedset

import "sort"

type OrderedSet[T comparable] struct {
	set       map[T]int
	nextIndex int
}

func (o *OrderedSet[T]) Add(elements ...T) {
	if o.set == nil {
		o.set = make(map[T]int)
	}

	for _, element := range elements {
		if !o.Has(element) {
			o.set[element] = o.nextIndex
			o.nextIndex++
		}
	}

}

func (o *OrderedSet[T]) Has(element T) bool {
	_, hasItem := o.set[element]
	return hasItem
}

func (o *OrderedSet[T]) List() []T {
	size := len(o.set)
	list := make([]T, size)
	i := 0

	for el := range o.set {
		list[i] = el
		i++
	}

	sort.Slice(list, func(i, j int) bool {
		return o.GetIndex(list[i]) < o.GetIndex(list[j])
	})

	return list
}

func (o *OrderedSet[T]) GetIndex(el T) int {
	return o.set[el]
}
