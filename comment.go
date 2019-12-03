package main

import "fmt"

func main() {
	// Are we racing or coding?
  /*
	fmt.Println("Ready")
  fmt.Println("Set")
  */
  fmt.Println("Gooooo!")
}




package main

import "fmt"

func main() {
	// Are we racing or coding?
  /*
	fmt.Println("Ready")
  fmt.Println("Set")
  */
  fmt.Println("Gooooo!")
}


Go Resources
Learning a new language like Go involves learning the accompanying rules and syntax. But, we don’t have to commit everything to memory! It’s ok to search things up, in fact, that’s what all good programmers do!

If you’re ever stuck on something, check out:

Codecademy’s Forums
View questions and answers from learners on this site!
Stack Overflow’s Go questions
A forum of programmers answering programming questions.
Go’s official site
Go to the official site for documentation.
Google
Roll up your sleeves and search it up!
It may help to search Golang instead of Go in certain cases.
In addition to online resources, Go also has it’s own built-in documentation system. To use it, in the command line, use the command go doc followed by a package name. For instance, to find out more information on the fmt package, you can use the command:

go doc fmt 
In the terminal, you’ll see at the top:

package fmt // import "fmt"

Package fmt implements formatted I/O with functions analogous to C's printf
and scanf. The format 'verbs' are derived from C's but are simpler.
...
The information returned actually spans quite a few lines, the example above is truncated. To get more specific information about a function in a package (like fmt‘s Println function) append .Println (or .println, the capitalization of the function doesn’t matter to go doc) to the command:

go doc fmt.Println
The go doc command is also very helpful when you’re interacting with new libraries and packages. Try it out yourself!

package main

import "fmt"

func main() {
	fmt.Println("Using the 'go doc' command is helpful")
  fmt.Println("You can find out more about a package")
  fmt.Println("Or about a function inside the package")
  fmt.Println("Try it out in the terminal!")
}














