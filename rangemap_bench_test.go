package rangemap

/*
go test -bench=.

Recorded tests
goos: darwin
goarch: amd64
pkg: rangemap/rangemap
cpu: Intel(R) Core(TM) i5-5257U CPU @ 2.70GHz
BenchmarkSingleKeyMap_Get-4         128684619                9.232 ns/op
BenchmarkUnitsMap_Get-4             79455692                14.98 ns/op
BenchmarkBigMap_Get-4               25379938                45.03 ns/op

Using the single key map as a base line,
4 keys on the map take 50% more time then a single Key,
9001 Keys on the map take 200% more time then 4 keys

*/
import "testing"

// BenchmarkSingleKeyMap_Get - This benchmarks Get on a map with a single key on it
func BenchmarkSingleKeyMap_Get(b *testing.B) {
	rangeMap := RangeMap[string]{}
	rangeMap.Put(0, 9, "Ones")

	for i := 0; i < b.N; i++ {
		_, _ = rangeMap.Get(5)
	}
}

// BenchmarkUnitsMap_Get - This benchmarks Get on a map with the 4 keys used on the unit tests
func BenchmarkUnitsMap_Get(b *testing.B) {
	rangeMap := RangeMap[string]{}
	rangeMap.Put(0, 9, "Ones")
	rangeMap.Put(10, 99, "Tens")
	rangeMap.Put(100, 999, "Hundreds")
	rangeMap.Put(1000, 9999, "Thousands")

	for i := 0; i < b.N; i++ {
		_, _ = rangeMap.Get(5)
	}
}

// BenchmarkBigMap_Get - This benchmarks Get on a map with 9001 keys and is used to test scale
func BenchmarkBigMap_Get(b *testing.B) {
	rangeMap := RangeMap[string]{}

	for i := 0; i < 9000*3+1; i += 3 {
		rangeMap.Put(i, i+1, "T")
	}

	for i := 0; i < b.N; i++ {
		_, _ = rangeMap.Get(5)
	}
}
