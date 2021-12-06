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
var acc account "This is how a variable of interface can be created." The default value of an interface is nill.

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

We can pass interface type as argument to a function and when calling the function we can pass a struct as a value to the function.

With our account example we can have a method that that takes our account as a param and such method can be used to access any of the interface methods.

We create a struct variable and pass it to the method when we call it, this will call the withdraw method of that particular struct.

type account interface{
    method1()
    method2()
}

func callWithdraw(acc account){
    acc.withdraw()
}

Passing value to the function above.
func main(){

    savingsAcc := savings{}
    callWithdraw(savingsAcc)

    This will call the withdraw method of the interface.
}

When a struct implements a function using pointer receiver, then a derefrence must be used to access or assign value to the struct.

Note any custom type can implement an interface.

Type Implementing Multiple Interfaces.
A type is said to implement an interface when it defines all it's methods, That same type can also implement other interface(s) by defining it's methods too.

Embedding Interfaces

You can embed an interface inside another one by simply calling or writing it's name in another interface. We can use this to reduce repeating ourselves while writing code. One scenario where this can be used is when method(s) in an interface can be used in another, instead of repeating those method(s) we can simply embed/call another interface that has the needed method(s) and this will give us access to it's method(s). Remember any type that is implementing the interface that has the embed interface must implement the methods in the interface that is embedded in it.

Embedding interface in a struct
Interface can also be embedded in a struct and all it's methods can then be called using the struct.

How these methods will be called will depend upon whether the embedded interface is a named field or an unnamed/anonymous field.

    If the embedded interface is a named field, then interface methods has to be called via the named interface name. Named having a variable for the interface.

    If the embedded interface is unnamed/anonymous field then interface methods can be referred directly or via the interface name. unnamed interface means not having a variable to it.

If the interface is a named field, it can't be called directly. We will need it's interface name to access it. If it is not then we can call it direct.

Also note that while creating the instance of either the struct with the embedded interface, the embedded interface i.e animal is initialized with a type implementing  it.

If we don't initialize the embedded interface animal, then it will be initialized with the zero value of the interface which is nil. Calling it's method(s) will create a panic.
