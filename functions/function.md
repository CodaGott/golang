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

A function can return more than 1 data-type values. To do this separate the data-type to be returned with a comma and both returning data-type will be in a parenthesis.
Example below:

    syntax: func name_of_function(parameter(s)) (returning_data-type1, returning_data-type2){ return parameters, another_parameters}

Golang has a labeled return function that allows you to label what you are expecting you function to return. This means you can assign a variable of some sort to your return type and afterwards use just a return statement at the end of you method.

If your labeled return type is of same data, you can decide to just specify the return type at the end and separate the variables with comma.

Example below: the example will have a function that returns int data-types

    syntax: func addAndMinusFromNumbers(num1 int, num2 int) (add int, minus int){ 

    add = num1 + num2
    minus = num1 - num2
    return
    }

The above is just to show that this can be achieved using Golang but it's a bad practice of clean code to have your function doing 2 things. A function should only do 1 thing and 1 thing alone.

Golang function has a keyword called defer, which is used to defer an action till the function is done with it's task.

Example below:

    syntax: func addAndMinusFromNumbers(num1 int, num2 int) (add int, minus int){ 
    defer fmt.Println("Done")
    add = num1 + num2
    minus = num1 - num2
    return
    }
