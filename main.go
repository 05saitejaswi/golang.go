Basic Go Structure: main
Continuing on with our program, we have:

func main () {
    fmt.Println("Hello World") 
}
There are a few things happening in our main function. We’re introduced to how functions (reusable blocks of code) are defined/created in Go. The basic syntax being:

The func keyword denotes the start of a function declaration.
func is followed by the name of the function.
After name is a pair of parentheses () and a set of curly braces {}.
Inside the function block {}, we indent the code inside using a tab. Then we access the fmt package and call its Println (stands for print line) function to print the message "Hello World" to the terminal.

Now let’s take a second and realize that while we defined our main function, we never explicitly told main to run its block of code. In other programming languages, functions have to be called, i.e. told to run its code. However, a main function is special, a file that has a package main declaration will automatically run the main function!

This nuance becomes more important as we build more complex programs. In the meanwhile, let’s finish up our program!

Instructions
1.
Let’s define our function using the func keyword, followed by the name main, a set of empty parentheses (), and a set of curly braces {}.

2.
Inside the set of curly braces, add a new line and insert:

fmt.Println("Hello World") 
Then make sure the closing brace } is also on its own line.

You can also change the Hello World portion to whatever you want, after all it’s your program!

3.
Time to run main.go! In the command line, type in go run main.go and press enter.

//main.go
package main

import "fmt"

func main() { 
  fmt.Println("Hello World")
}
