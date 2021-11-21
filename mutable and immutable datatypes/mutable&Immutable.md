# Mutable   &   Immutable Data-Types

Mutable data types are changeable while immutable data-types are not changeable.

Slice is a mutable data-type and this makes it possible that when you assign the value of a slice to a different variable, when ever you make changes in the variable you are assigning the value to, it will affect the original data.

This is so because when you assign a slice value to a different variable, both of these values will start pointing to the same value/data in the memory of the system. Maps works same as a slice.

For arrays, even though it's mutable, "it's value can be changed", it doesn't have the same behavior as maps and slices.

For arrays, when you assign it to a new variable, what it does is it makes a copy of that array and assigns it to the new variable and that means changes made to any of the variable will not affect the other array.
