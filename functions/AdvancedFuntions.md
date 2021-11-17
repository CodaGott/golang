# Advanced Functions

In golang functions are actually their own data-types, just like float, string, int and others. This means they act similarly to things like other data-types.

We can pass them as variables, we can return a function or just send them around our program the same way we treat other data-types.

Calling a function is when we close the function name with parenthesis. e.g. let's say we have a function that adds numbers "add()" that's an example of a function call.

When we write a function name and not add the parenthesis, we are just making a reference to the function or creating a pointer to it.

If we assign a function reference to a variable we can actually call the variable the same way we could have called the function.
Still on the add function example above, we can decided to make a reference of the function instead by following the example below.

addition := add

To call we can now decide to call the variable as if it's a function. Like so "addition()"

What happens is when addition() is called it calls the function that was assigned to it.

Important to note is we can reference even functions with parameter.

We can create a function in another function in golang. We do this by making a variable and then assigning the variable the func keyword followed by parenthesis, and curly braces.

Even a function inside a function can also have parameters.

We can use a return type for a function inside a function and even assign it a value before calling it.

Example:
func main() {

 test := func(x int) int  {
  return x + 1
 }(8)

 fmt.Println(test)
}

The above example is assigning the test function the value of 8, so can call the function without having to assign any other value to it.
