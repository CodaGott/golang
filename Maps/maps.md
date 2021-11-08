# Maps

Maps allows use to store a key, value pair. So a key is mapped to a value so to access a value you can use the key.

In array we access based on elements position but with map we access by key.

To define a map we we create a variable followed by the "map" keyword followed by square bracket and the datatype of the key and outside of the bracket the type of values the map is going to hold then assign to this variable map starting with "map" keyword followed by square bracket and the datatype of the key and outside of the bracket the type of values the map is going to hold.

Syntax:
    var map_name map[data_type_for_key]data_type_for_values = map[data_type_for_key]data_type_for_values{
        "key": value, "another_key": value
    }

IMPORTANT TO NOTE:
    You can't have duplicated key but can have duplicated values.

    Map doesn't keep track of the order by which we inserted the values, so it prints or gives everything randomly.

To access a value in a map, we call the variable we used to store the map, followed by a square bracket and inside the bracket call the key of the map.

Syntax:
    map_name["key_of_map_value"]

To change a value in a map, we call the variable we used to store the map, followed by a square bracket and inside the bracket call the key of the map then assign a new value to it.

Syntax:
    map_name["key_of_map_value"] = new_value

To add a value in a map, we call the variable we used to store the map, followed by a square bracket and inside the bracket we create a new key of the map then assign a new value to it. Remember the key must not be the same as any of the previous ones

To delete from a map you just use the delete method and and in side the method you enter the the name of the map you want to delete followed by the key to the value you want to remove comma separated.

Syntax:
    delete(map_name, "key_of_map_value")

To get the number of elements in a map you use the len method.

Syntax:
    len(map_name)

Look up other ways to create maps.
