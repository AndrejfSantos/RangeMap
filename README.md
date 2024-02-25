# Range map


Very useful when its needed to make a map from a key ranging from a set of values, for example:

| from | to   | value       |
|------|------|-------------|
| 0    | 9    | "Ones"      |
| 10   | 99   | "Tens"      |
| 100  | 999  | "Hundreds"  |
| 1000 | 9999 | "Thousands" |


Code example:

````go
package main

import "github.com/AndrejfSantos/RangeMap/rangemap"

func main()  {

    rangeMap := rangemap.RangeMap[string]{}
    rangeMap.Put(0, 9, "Ones")
    rangeMap.Put(10, 99, "Tens")
    rangeMap.Put(100, 999, "Hundreds")    
    rangeMap.Put(1000, 9999, "Thousands")

    value, found := rangeMap.Get(-1)  // returns nil , false
    value, found = rangeMap.Get(1)    // returns "Ones" , true
    value, found = rangeMap.Get(42)   // returns "Tens" , true
    value, found = rangeMap.Get(666)  // returns "Hundreds" , true
    value, found = rangeMap.Get(1337) // returns "Thousands" , true

}
````