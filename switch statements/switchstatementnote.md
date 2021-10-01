# Switch statement

Switch statement works like the if statements, they help leave our code more organized and little bit cleaner.

The keyword for starting a switch statement is "switch".

Syntax:
    switch variable {
        case 1:
            some action
        case 2:
            some action
        default:
        some action
    }

    What this does is to check if the variable is equal to all the specified values of the cases in the switch block. And the default is an optional statement that get's executed only if the any case values did not meet any the requirements.

Golang switch statement allows for multiple cases on the same line.