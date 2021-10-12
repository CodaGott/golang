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
    Let's say we have a 5 element arrays and within that we have a slice of 3 array elements. The capacity of that Array will be 4 because after the last slice there is another room for another element.