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
