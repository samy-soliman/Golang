/*
When generics were released, a new way of writing interfaces was also released at the same time!

We can now simply list a bunch of types to get a new interface/constraint.

// Ordered is a type constraint that matches any ordered type.
// An ordered type is one that supports the <, <=, >, and >= operators.
type Ordered interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
        ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
        ~float32 | ~float64 |
        ~string
}

*/

/*
Why might you create an interface using a type list?
"You know exactly which types satisfy your interface"
*/