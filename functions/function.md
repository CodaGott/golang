# Functions in Golang

Functions are ways of writing reuseable codes that can be used in multiple places. To create a function, we start with the keyword "func" followed by the name we want our functions to have, parenthesis, and curl-braces. EXAMPLE below:
    syntax: func print(){} this is how you define a functions without parameters.

To define a function with parameters means to add a variable inside the parenthesis that will take a value when the function is called. To add a parameter to a function, after creating the functions as specified above, inside the parenthesis, add the name of the variable followed by it's data type.
Example below:

syntax: func print(name_of_the_parameter data-type){}

You can have more than 1 parameters in a function, and when the parameters are of the same type, you can omit the data-type of the first parameter.
Example below:

    syntax: func name(first_name, last_name string){}
The above information is saying that we are expecting information of the same data-type so we don't want to repeat the data-type specification.

This works that way because golang is built in a way that it assumes that values coming in is of the same data-type.

when you have parameters of different types be sure to specify the data-types of each parameter.

# How to add  to a function in Golang

To add a return type to a function, we add the data-type the function is expected to return after the closing parenthesis before the curl-braces.
Example below:

    syntax: func name_of_function(parameter(s)) data-type to be returned {}
