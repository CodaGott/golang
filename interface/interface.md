# INTERFACE

Interface is a type in Go which kinda stores collection of methods. These method collections are used to represent certain behavior a type will have.

# Creating an Interface.

Below is how an interface is defined:
type name_of_interface interface{
    method 1
    method 2
    method 3
    ...
}

for example let's define an interface called account with the most basic behavior of an account.

type account interface {
    withdraw()
    deposit()
}

We can create a variable of the above interface by creating a variable name and followed by name of the interface. EXAMPLE BELOW:
acc := account "This is how a variable of interface can be created." The default value of an interface is nill.

# Implemeting an Interface.

Any type that has the methods in our interface is said to implement the interface. So if we define savings struct, which in real world is a type of an account and that struct has the method of our interface, it will be said that the savings is implementing the interface. EXAMPLE:

type savings struct{
    some fields
}

func (s savings) deposit(){
    "This is a deposit into savings"
}

the above savings struct is implementing the deposit interface

func (s savings) withdrawal(){
    "This is a withdrawal from savings"
}

the above savings struct is implementing the withdrawal interface

if we create a variable for the account interface as we did before, we can now assign an instance of the struct we want to the interface.

acc := account

acc = savings{pass the field information's if any}
 acc.deposit() this will give us the deposit method of our savings struct and acc.withdraw() will give us the withdrawal method of the savings struct.

The above operations are possible because the savings struct is implementing both interfaces.

It's important to note that Interface are implemented implicitly because in the above example we didn't have any explicit declaration that the savings type implements the account interface.
In Go you don't have to explicitly say that this type is implementing an interface. A type implements an interface if it implements all the methods of the interface.
A struct must implement all the methods of an interface or an error will be raised.

We can assign the same interface variable name to different struct and their individual implementations or behavior will be called upon execution. This is how Polymorphism is achieved in Go.

Method signature is important when implementing an interface.

We can pass interface type as argument to a function that accepts an argument of the interface.

With our account example we can have a method that calls deposit or savings. And if that method takes our interface as a parameter, we can pass in the instances of the struct type.

type account interface{
    method1
    method2
}

type savings struct(){
    some variables
}
