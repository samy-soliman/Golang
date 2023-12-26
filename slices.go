/*
Arrays are fixed in size. Once you make an array like [10]int you can't add an 11th element.

A slice is a dynamically-sized, flexible view of the elements of an array.

Slices always have an underlying array, though it isn't always specified explicitly. To explicitly create a slice on top of an array we can do:

primes := [6]int{2, 3, 5, 7, 11, 13}
mySlice := primes[1:4]
// mySlice = {3, 5, 7}

The syntax is:

arrayname[lowIndex:highIndex]
arrayname[lowIndex:]
arrayname[:highIndex]
arrayname[:]

Where lowIndex is inclusive and highIndex is exclusive

Either lowIndex or highIndex or both can be omitted to use the entire array on that side.
*/

// example of a function that returns a slice
/*
func function (s string) ([]string, error) {

}
*/

/*
Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data. Except for items with explicit dimensions such as transformation matrices, most array programming in Go is done with slices rather than simple arrays.

Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array. If a function takes a slice argument, changes it makes to the elements of the slice will be visible to the caller, analogous to passing a pointer to the underlying array. A Read function can therefore accept a slice argument rather than a pointer and a count; the length within the slice sets an upper limit of how much data to read. Here is the signature of the Read() method of the File type in package os:

Referenced from Effective Go

func (f *File) Read(buf []byte) (n int, err error)
*/

// image of slice
/*
Most of the time we don't need to think about the underlying array of a slice. We can create a new slice using the make function:

// func make([]T, len, cap) []T
mySlice := make([]int, 5, 10)

// the capacity argument is usually omitted and defaults to the length
mySlice := make([]int, 5)

Slices created with make will be filled with the zero value of the type.

If we want to create a slice with a specific set of values, we can use a slice literal:

mySlice := []string{"I", "love", "go"}

Note that the array brackets do not have a 3 in them. If they did, you'd have an array instead of a slice.
Length

The length of a slice is simply the number of elements it contains. It is accessed using the built-in len() function:

mySlice := []string{"I", "love", "go"}
fmt.Println(len(mySlice)) // 3

Capacity

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice. It is accessed using the built-in cap() function:

mySlice := []string{"I", "love", "go"}
fmt.Println(cap(mySlice)) // 3

Generally speaking, unless you're hyper-optimizing the memory usage of your program, you don't need to worry about the capacity of a slice because it will automatically grow as needed.
*/

/*
The length of a slice may be changed as long as it still fits within the limits of the underlying array; just assign it to a slice of itself. The capacity of a slice, accessible by the built-in function cap, reports the maximum length the slice may assume. Here is a function to append data to a slice. If the data exceeds the capacity, the slice is reallocated. The resulting slice is returned. The function uses the fact that len and cap are legal when applied to the nil slice, and return 0.

Referenced from Effective Go

func Append(slice, data []byte) []byte {
    l := len(slice)
    if l + len(data) > cap(slice) {  // reallocate
        // Allocate double what's needed, for future growth.
        newSlice := make([]byte, (l+len(data))*2)
        // The copy function is predeclared and works for any slice type.
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:l+len(data)]
    copy(slice[l:], data)
    return slice
}
*/