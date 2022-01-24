// Package demo1 contain all generic Funcs
package demo1


// K comparable is predeclared by Go itself, which require the map key to be comparable
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}


// Type constrint
type Number interface {
	float64 | int64  
}

func SumGeneric[K comparable, V Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
