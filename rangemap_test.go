package rangemap

/*
go test -v - cover
*/
import "testing"

// Covers GetOrDefault for when key is found
func TestRangeMapGetOrDefault_found(t *testing.T) {
	wanted := "Ones"
	rangeMap := RangeMap[string]{}
	rangeMap.Put(0, 9, wanted)
	got := rangeMap.GetOrDefault(1, "Default")

	if got == "Default" {
		t.Errorf("Get() got = Default, want %s", wanted)
	}
	if got != wanted {
		t.Errorf("Get() got = %s, want %s", got, wanted)
	}
}

// Covers GetOrDefault for when key is not found and default is returned
func TestRangeMapGetOrDefault_not_found(t *testing.T) {
	wanted := "Default"
	rangeMap := RangeMap[string]{}
	got := rangeMap.GetOrDefault(1, wanted)

	if got != wanted {
		t.Errorf("Get() got = %s, want %s", got, wanted)
	}
}

// covers a Get on keys that are found
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

// covers a Get on keys that are not found
func TestRangeMapGet_not_found(t *testing.T) {

	rangeMap := RangeMap[string]{}
	rangeMap.Put(0, 9, "Ones")
	rangeMap.Put(1000, 9999, "Thousands")

	helper_TestRangeMapGet_not_found(t, rangeMap, -1)
	helper_TestRangeMapGet_not_found(t, rangeMap, 10)
	helper_TestRangeMapGet_not_found(t, rangeMap, 999)
	helper_TestRangeMapGet_not_found(t, rangeMap, 10000)
}

// covers a Get on an empty map
func TestRangeMapGet_empty(t *testing.T) {
	rangeMap := RangeMap[string]{}
	helper_TestRangeMapGet_not_found(t, rangeMap, 0)
}

// TestRangeMapGet_flow - A more complete integration test that tests gets and puts together
func TestRangeMapGet_flow(t *testing.T) {

	rangeMap := RangeMap[string]{}
	helper_TestRangeMapGet_not_found(t, rangeMap, -50)
	helper_TestRangeMapGet_not_found(t, rangeMap, -1)
	helper_TestRangeMapGet_not_found(t, rangeMap, 2)
	helper_TestRangeMapGet_not_found(t, rangeMap, 5)
	helper_TestRangeMapGet_not_found(t, rangeMap, 18)

	rangeMap.Put(2, 5, "small")
	helper_TestRangeMapGet_not_found(t, rangeMap, -50)

	for i := -1; i <= 1; i++ {
		helper_TestRangeMapGet_not_found(t, rangeMap, i)
	}

	for i := 2; i <= 5; i++ {
		helper_TestRangeMapGet_found(t, rangeMap, i, "small")
	}

	helper_TestRangeMapGet_not_found(t, rangeMap, 18)

	rangeMap.Put(-60, -40, "negatives")
	helper_TestRangeMapGet_not_found(t, rangeMap, -61)
	helper_TestRangeMapGet_found(t, rangeMap, -60, "negatives")
	helper_TestRangeMapGet_found(t, rangeMap, -50, "negatives")
	helper_TestRangeMapGet_found(t, rangeMap, -40, "negatives")
	helper_TestRangeMapGet_not_found(t, rangeMap, -39)

	for i := -1; i <= 1; i++ {
		helper_TestRangeMapGet_not_found(t, rangeMap, i)
	}

	for i := 2; i <= 5; i++ {
		helper_TestRangeMapGet_found(t, rangeMap, i, "small")
	}

	helper_TestRangeMapGet_not_found(t, rangeMap, 18)

	rangeMap.Put(18, 18, "18")
	helper_TestRangeMapGet_not_found(t, rangeMap, -61)
	helper_TestRangeMapGet_found(t, rangeMap, -60, "negatives")
	helper_TestRangeMapGet_found(t, rangeMap, -50, "negatives")
	helper_TestRangeMapGet_found(t, rangeMap, -40, "negatives")
	helper_TestRangeMapGet_not_found(t, rangeMap, -39)

	for i := -1; i <= 1; i++ {
		helper_TestRangeMapGet_not_found(t, rangeMap, i)
	}

	for i := 2; i <= 5; i++ {
		helper_TestRangeMapGet_found(t, rangeMap, i, "small")
	}

	helper_TestRangeMapGet_not_found(t, rangeMap, 17)
	helper_TestRangeMapGet_found(t, rangeMap, 18, "18")
	helper_TestRangeMapGet_not_found(t, rangeMap, 19)

}

// Helper function for when the key should be found
func helper_TestRangeMapGet_found(t *testing.T, rangeMap RangeMap[string], key int, wanted string) {
	t.Helper()
	got, found := rangeMap.Get(key)
	if !found {
		t.Errorf("Get() found was not true on %d", key)
	}
	if got == nil {
		t.Errorf("Get() got = nil, want %s", wanted)
	} else if *got != wanted {
		t.Errorf("Get() got = %s, want %s", *got, wanted)
	}
}

// Helper function for when the key should not be found
func helper_TestRangeMapGet_not_found(t *testing.T, rangeMap RangeMap[string], key int) {
	t.Helper()
	got, found := rangeMap.Get(key)
	if found {
		t.Errorf("Get() found was true on %d", key)
	}
	if got != nil {
		t.Errorf("Get() got = %s, want nil", *got)
	}
}
