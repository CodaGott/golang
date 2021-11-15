# Slices

Slices are have similarities as arrays but have some improvements. One of those is with arrays we have to specify the number of items we want to pass into it upon creation but sometimes we just might not know how many items we want to be in the array so we need something more dynamic.

Major Difference between Arrays N Slices.

Arrays and Slices have three major characteristics but they work differently:

In ARRAYS.
    1. Pointers = This indicates the index or position of an element.
    2. Length = This indicates the total number of elements in a given array.
    3. Capacity = This indicates the total number of items in a given array just like the length.

In SLICES.
    Slices are just portion of an array, we can have 4 elements in an array but have only 2 elements in a slice.

    1. Pointers = This points to the index of an element within a slice and not the entire arrays.
    2. Length = This is the number of elements in a Slice of array.
    3. Capacity = This is the total number of elements an array can have after the slices.
    Let's say we have a 5 element arrays and within that we have a slice of 3 array elements in-between the first and last element. The capacity of that Array will be 4 because after the last slice there is another room for another element.

To declare a slice we first create an array, pass in the values and then create a slice and pass in the array elements we want to it using the slice operator.
    Syntax:
        var array[5] int = [5]int {1, 2, 4, 3, 5}
        var []slice = array[1:3]
    The above slice is simply saying starting from the array element in index 1 copy the elements up on till the element 3 but don't return the last element. Which is from 2 till 3 but do not include 3. So Only 2 and 4 will be printed if we print the above code.

    So it goes from the index specified on the left till the index specified on the right without it's value.

The numbers on the right and left hand of the slice operator indicates when the slice should start and where it should end.

If we leave the slice operator empty this means we are saying just copy the arrays. So this will start from the beginning to the end of the array and copy all of it's elements.

If we don't have any number on the left hand side of the slice operator we are saying it should start from the very beginning till the specified index on the right.

If we don't have value on the right side it means start from the index on the left and goto the very end of the array.

We can Just declare a slice on it own like the example below:
    var a []int = []int{3, 4, 5, 6, 7}
    This create a slice for us.

To add element to the slice we use the "append()" method and in the method we pass the name of the slice followed by the value we want to pass to our slice comma separated.

Note the append method needs to be assigned to a variable so to append another item into the slice above, we do this example below:

b := append(a, 10), this will add the number 10 to the end of the above arrays and assign the array to the variable b.

The example above copies the value of the array and transfers it to the new variable with the appended value.

This means The previous array still has it's values but the variable with append now has the value of the array plus the appended value.

We can also make slices using the make method, This will create an empty slice for us.
Syntax : variable := make([]dataType, sizeOfSliceElement)

So this "make_slice := make([]int, 5)" will create a slice of 5 elements.
