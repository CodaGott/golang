# STRUCTS AND CUSTOM TYPES

Structs is how class/objects are built in Golang. They are used to create custom types of an object.

To create a struct type in Golang, we start with the keyword "type" followed by name of the object or type you want to create, then the keyword "struct".

For convention, the name of the struct always starts with a capital letter. Example below:
SYNTAX:
type Human struct{
    ==> here we have the fields we want to be type of the struct.
}

To create an object of a struct or type of a struct, you create a variable followed by the name of the point and assign the it with values for the fields in the struct by calling the name of the point again followed by curly braces and the values comma separated.
Example below:
var john Human = Human{values comma separated}

We can have multiple types of a point and a change in one will not cause changes in another.

Let's say we have a point for a bank account, and Mr A and H created an account, when Mr A makes changes in his account it should not affect Mr H's account and vice verse.

You can only make changes to a specific struct without affecting the other type of a struct.

We can have multiple data-types as the values in a struct.

We can also create a type without using the var keyword like we do other variables. Example:

john := Human{== pass in the required values}, this is a valid type.

We create a type and upon creation not know the value we want to pass to the fields of out type, then we can specify the fields or field we want to assign it's values by specifying it's name followed by a colon then the value.
Example:
type Human struct{
    name string
    age string
    phoneNumber string
}
We can decide to create a type of this struct without passing some of it's values by following the syntax below.

john := Human{name: "John Doe"}

When we do this, it adds the known values to the specified names and the default value of data-type to the one fields we didn't specify it's values.

Note we can use a dot "." after the type of a struct to access a specific field of that struct.
we can use john.age to access or assign/change the value to the age in the Human struct.

We can have a type of a struct in another struct. Let's say we have a Human struct and Animal or being more specific Dog struct, we can actually pass the Dog struct as a type in the field of our human struct.

To access a value in an embedded struct, we start with the name of the type that the other struct is inside of then use the dot "." name of the struct inside it another dot "." the name of the field we want to access.

Example: Working with human and dog example, we can do something like this john.dog.name, this will give use access to the name of the dog inside of type of struct john.
