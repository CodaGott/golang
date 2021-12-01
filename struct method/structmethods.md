# STRUCT METHOD

To create a method for a struct, we start with the func keyword followed by parenthesis and inside of the parenthesis we create a variable for the struct and close the parenthesis. and give name to the method. EXAMPLE:

type Student struct {
 name   string
 grades []int
 age    int
}

To have say a method that set's name to the struct we use the following syntax.

func(student1 Student) setAge(name string){
    student1.name = name
}

It's advised to have pointers on a method that is used to change values of a struct because if we have say name already existing in the student struct if we just pass in a new value without having the method as a pointer it will not change the already existing name. So we want to have the syntax below if the intention is to change an information that is already existing.

func(student1 *Student) setName(name string){
    student1.name = name
}

if the method is basically for maybe calculating or manipulating values then we don't need to add pointers to the method.
