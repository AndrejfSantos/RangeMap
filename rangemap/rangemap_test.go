package rangemap

/*
go test -v - cover
*/
import "testing"

func TestRangeMapGet_found(t *testing.T) {

	rangeMap := RangeMap[string]{}
	rangeMap.Put(0, 9, "Ones")
	rangeMap.Put(10, 99, "Tens")
	rangeMap.Put(100, 999, "Hundreds")
	rangeMap.Put(1000, 9999, "Thousands")

	for i := 0; i <= 9; i++ {
		helper_TestRangeMapGet_found(t, rangeMap, i, "Ones")
	}

	helper_TestRangeMapGet_found(t, rangeMap, 10, "Tens")
	helper_TestRangeMapGet_found(t, rangeMap, 42, "Tens")
	helper_TestRangeMapGet_found(t, rangeMap, 99, "Tens")

	helper_TestRangeMapGet_found(t, rangeMap, 100, "Hundreds")
	helper_TestRangeMapGet_found(t, rangeMap, 666, "Hundreds")
	helper_TestRangeMapGet_found(t, rangeMap, 999, "Hundreds")

	helper_TestRangeMapGet_found(t, rangeMap, 1000, "Thousands")
	helper_TestRangeMapGet_found(t, rangeMap, 1337, "Thousands")
	helper_TestRangeMapGet_found(t, rangeMap, 9999, "Thousands")
}
func helper_TestRangeMapGet_found(t *testing.T, rangeMap RangeMap[string], key int, wanted string) {
	got, found := rangeMap.Get(key)
	if !found {
		t.Errorf("Get() found not true on %d", key)
	}
	if got == nil {
		t.Errorf("Get() got = nil, want %s", wanted)
	} else if *got != wanted {
		t.Errorf("Get() got = %s, want %s", *got, wanted)
	}
}

func TestRangeMapGet_not_found(t *testing.T) {

	rangeMap := RangeMap[string]{}
	rangeMap.Put(0, 9, "Ones")
	rangeMap.Put(1000, 9999, "Thousands")

	helper_TestRangeMapGet_not_found(t, rangeMap, -1)
	helper_TestRangeMapGet_not_found(t, rangeMap, 10)
	helper_TestRangeMapGet_not_found(t, rangeMap, 999)
	helper_TestRangeMapGet_not_found(t, rangeMap, 10000)
}
func helper_TestRangeMapGet_not_found(t *testing.T, rangeMap RangeMap[string], key int) {
	got, found := rangeMap.Get(key)
	if found {
		t.Errorf("Get() found true on %d", key)
	}
	if got != nil {
		t.Errorf("Get() got = %s, want nil", *got)
	}
}

func TestRangeMapGet_empty(t *testing.T) {

	rangeMap := RangeMap[string]{}
	got, found := rangeMap.Get(0)
	if found {
		t.Error("Get() found true on empty Map")
	}
	if got != nil {
		t.Errorf("Get() got = %s on empty Map, want nil", *got)
	}
}
