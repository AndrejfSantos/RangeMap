package rangemap

import "sort"

type intRange struct {
	from int
	to   int
}
type valueIntRange[T any] struct {
	intRange
	value T
}
type RangeMap[T any] struct {
	// keys used for seach on get
	keys []intRange

	// values used to return value on get
	values []T

	// all input added on put
	input []valueIntRange[T]

	// RangeMap needs to rebuild whenever a new input is added,
	// this flag is used to know if the map is ready to be used or needs rebuilding
	ready bool
}

func (m *RangeMap[T]) Get(getValue int) (*T, bool) {
	if !m.ready {
		m.rebuild()
	}

	count := len(m.keys)

	i := sort.Search(count, func(i int) bool {
		return getValue < m.keys[i].from
	})

	i--
	if i >= 0 && i < count && !(getValue > m.keys[i].to) {
		return &m.values[i], true
	}
	return nil, false
}
func (m *RangeMap[T]) Put(from int, to int, value T) {
	m.input = append(m.input,
		valueIntRange[T]{intRange{from, to}, value})
	m.ready = false
}

func (m *RangeMap[T]) rebuild() {
	// sort all input first
	sort.Slice(m.input, func(i, j int) bool {
		return m.input[i].from < m.input[j].from
	})

	size := len(m.input)

	m.keys = make([]intRange, size, size)
	m.values = make([]T, size, size)

	for i, input := range m.input {
		m.keys[i] = intRange{from: input.from, to: input.to}
		m.values[i] = input.value
	}
	m.ready = true
}
