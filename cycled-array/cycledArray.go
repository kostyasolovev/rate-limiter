package cycled_array

import (
	"cmp"
	"sync"
)

// CycledArray - для хранения слайсов с ограниченной вместимостью.
// элементы в массиве отсортированы в порядке возрастания
type CycledArray[T cmp.Ordered] struct {
	arr               []T
	nextPos, len, cap int
}

type A[T cmp.Ordered] struct {
	arr CycledArray[T]
	mu  sync.Mutex
}

func (t *CycledArray[T]) Append(obj T) {
	t.arr[t.nextPos] = obj

	var sorted bool
	// проверяем что новый элемент не нарушает сортированности массива
	// check order in right side
	for cur := t.nextPos; cur > 0; cur-- {
		if t.arr[cur-1] < t.arr[cur] {
			t.arr[t.nextPos] = t.arr[cur]
			sorted = true

			break
		}
	}
	// left side if necessary
	for cur := t.len - 1; !sorted && cur > t.nextPos+1; cur-- {
		if t.arr[cur-1] < t.arr[cur] {
			t.arr[t.nextPos] = t.arr[cur]
			break
		}
	}

	if t.nextPos == t.cap-1 {
		t.nextPos = 0
	} else {
		t.nextPos++
	}

	if t.len < t.cap {
		t.len++
	}
}

func (t *CycledArray[T]) Index(from T) int {
	return t.getIndex(from)
}

func (t *CycledArray[T]) PopRange(left, right int) {
	left = min(left, t.nextPos)
}

func (t *CycledArray[T]) getIndex(target T) int {
	left, right := 0, t.len-1
	if t.len == t.cap {
		if t.nextPos > 0 {
			if t.arr[t.len-1] < target {
				if t.arr[0] > target {
					return t.len - 1
				}

				right = t.nextPos - 1
			} else {
				left = t.nextPos - 1
			}
		}
	}

	for left <= right {
		mid := (right + left) / 2
		if t.arr[mid] < target {
			left = mid + 1
		} else if t.arr[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}

	return left
}
