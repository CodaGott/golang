# Range In Go

Range is a function/keyword that helps with iterating over different items like an array, slices and strings.

when you say range of something it's like saying length of the whole array, slices or string you're looping through.

Syntax:
    var values []int = []int {8, 9, 3, 2, 1, 9}

    when you say:
        for i, element := range values {
            fmt.Println(i, element)
        }
what you are say is for all the element in the values starting from the first to the last loop through them all.
