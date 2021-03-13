# Arrays and Slices

When working with TinyGo be ready to change certain coding conventions that are common place in high level languages, such as Go. Memory ***is*** an object.

It is prefferable to avoid the garbage collector too, as it will slow down your program noticeably. To do this avoid using the `make` built-in.

These are some things to ***avoid*** doing with TinyGo
```go
// TinyGo No-No's
slice := make([]int,8)
for {
    // allocating memory in a for loop.
    // Avoid slices in general.
    anotherSlice := []int{1,2,3}
}
```

### Things that will work most of the time:

* Declaring arrays and passing arrays as slices to functions
    ```go
    array := [...]int{1,2,3}
    // array[:] is the slice of all elements in array
    sliceFunc(array[:])
    // same as
    sliceFunc(array[0:len(array)])
    ```
* Declaring variables outside of for scope and using the 3 statement for to iterate.
    ```go
    var iter = [3]int{1,2,3}
    var j int

    for i:=0; i<len(iter); i++ {
        
        j = mapFunc(i, iter[i])
        
        iter[i] = j*j
    }
    ```