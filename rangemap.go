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

// RangeMap - A Map that is very useful when its needed to make a map from a key ranging from a set of values
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

// Get - returns the value stored in the map for a key, or nil if no value is present.
// The ok result indicates whether value was found in the map.
func (m *RangeMap[T]) Get(key int) (value *T, ok bool) {
	if !m.ready {
		m.rebuild()
	}

	count := len(m.keys)

	i := sort.Search(count, func(i int) bool {
		return key < m.keys[i].from
	})

	i--
	if i >= 0 && i < count && !(key > m.keys[i].to) {
		return &m.values[i], true
	}
	return nil, false
}

// Put - stores the value for a key that will be within the range of from and to
func (m *RangeMap[T]) Put(from int, to int, value T) {
	m.input = append(m.input, valueIntRange[T]{intRange{from, to}, value})
	m.ready = false
}

// rebuild - for internal use, builds the map from the input values
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
