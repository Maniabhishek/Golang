// What is a map[string]interface{} in Go, and why is it so useful?
// How do we deal with maps of string to interface{} in our programs? What the heck is an interface{}, anyway? Let’s find out.
// Golang ‘map string interface’ example

// When you need to store a collection of arbitrary values of any type, then, identified by strings, a map[string]interface{} or map[string]any is the ideal choice.
// Why is interface{} so useful?
// A variable declared as interface{} can hold a string value, an integer, any kind of struct, a pointer to an os.File, or indeed anything you can think of.
// Suppose we need to write a function that prints out the value passed to it, but we don’t know in advance what type this value would be. This is a job for the empty interface:

//{
//     "name":"John",
//     "age":29,
//     "hobbies":[
//         "martial arts",
//         "breakfast foods",
//         "piano"
//     ]
// }

// map[string]any and arbitrary data

// {
//     "name":"John",
//     "age":29,
//     "hobbies":[
//         "martial arts",
//         "breakfast foods",
//         "piano"
//     ]
// }

// we can define the type for above data , but what this is some alien data and we are not aware of the type of the data then in that case map[string]any or map[string]interface
// is very important for unmarshalling

// p := map[string]any{}
// err := json.Unmarshal(data, &p) check error

// Using map[string]any data
// One thing we can do is use a type switch to do different things depending on the type of the value. Here’s an example:

// for k, v := range p {
//     switch c := v.(type) {
//     case string:
//         fmt.Printf("Item %q is a string, containing %q\n", k, c)
//     case float64:
//         fmt.Printf("item %q is a number, specifically %f\n", k, c)
//     default:
//         fmt.Printf("Not sure what type item %q is, but I think it might be %T\n", k, c)
//     }
// }

// The special syntax switch c := v.(type) tells us that this is a type switch,
// meaning that Go will try to match the type of v to each case in the switch statement.
// For example, the first case will be executed if v is a string:
// In each case, the variable c receives the value of v, but converted to the relevant type. So in the string case, c will be of type string.
// The float64 case will match when v is a float64:

// You might be puzzled that the whole-number value 29 was unmarshaled into a float64, but that’s normal. All JSON numbers are treated as float64 by json.Unmarshal. It’s the most general of Go’s numeric types.
// Finally, if no other case matches, the default case is activated:

// The format specifier %T to fmt.Printf prints the type of its value, which is sometimes handy.
// In this case we can see that the value of "hobbies" is a slice of arbitrary data, which makes sense.

// When to use map[string]any
// whenever the data is coming from outer world for example many apis return data in map format , where you might not in what format the data is

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	foods := map[string]interface{}{
		"bacon": "delicious",
		"eggs": struct {
			source string
			price  float64
		}{"chicken", 1.75},
		"steak": true,
	}

	mbyte, err := json.Marshal(foods)
	if err != nil {
		fmt.Println("error while marshalling", err)
	}

	var unmarshalledResult *map[string]interface{}
	uerr := json.Unmarshal(mbyte, &unmarshalledResult)
	if uerr != nil {
		fmt.Println("error while unmarshalling")
	}

	fmt.Println("-------unmarshalledResult--------", unmarshalledResult)
	for k, v := range *unmarshalledResult {
		switch c := v.(type) {
		case string:
			fmt.Printf("item %q is a string type ", k)
		case float64:
			fmt.Printf("item %q is a number type", k)
			fmt.Printf("item %q is of type interface{} and value %v", k, v)
		default:
			fmt.Printf("item %q could be type of %T, not sure ", k, c)
		}
	}

	fmt.Println(foods["eggs"])
}
