# Arrays in Golang

Arrays are simply ordered collection of elements.

Elements are items in array.

Syntax:
    To create an array, first we create a variable, followed by open/close square, brackets [size of elements of the array] , the data-type. Also arrays in GOLANG has a fixed length so upon creation we have to specify the amount of elements it's gonna hold. This doesn't mean we have to fill in all the elements upon creation but it simply means the amount of elements the array can hold can't exceed the specified length.

    Example:
        var arrayName [5] string

Like most programming languages out there the index of array starts at 0. This means the first element of the array has an index (position number) of 0.

To change or assign a value to array at any position of the array, we use the the variable name of the array followed by the square bracket and the index of the position we want to change it's value then assign a value of same data-type to it.

Example:
    var arr [5] int
    arr[0] = 100
The above example gives the first position of our array the element of 100.

To access an element in the array, use the the variable name of the array followed by the square bracket and the index of the position we want to get.

Another way we can define array when we know the values we are expecting is to pass in the value upon creation.

    Syntax:
        arr := [] int {your values}

We can have nested arrays in Golang and nested arrays are simply arrays inside of an array.

When we define a 2D array, the number in the first square bracket is the number of arrays we want to have and the next one is the number of array elements each of the arrays will have.

arr2D := [the number of arrays we are creating][number of elements in the array] = {{}, {}}, example below.

arr2D := [2][2] {{1, 3},{4, 6}}

if you have an array and don't know the number of items you'll be adding to it, you can replace the place you add the number of arrays with triple dots. EXAMPLE BELOW <<<>>>
nameOfArray :=[...] this will make your array open for more inputs.

    Syntax:
        arr := [][][] int {your value(s), {another value(s)}, {another value(s)}}
