Compiling
When we write Go code, we’re writing it to be readable for ourselves and other developers. We’re able to make sense of the code and its intentions. Computers, on the other hand, do NOT understand Go code and thus have no direct idea what our code means/does. What computers do understand are a series of 0’s and 1’s (or binary). To translate Go code to binary, Go has a compiler, a piece of software that converts Go code into a program that the computer understands. This translated code is called an executable or a binary file. We can then run the executable which will do what our Go program was written to do.

To tell the compiler to compile a Go program, we first navigate to our Terminal (on Mac), or Command Prompt (on Windows). Then, type in go build followed by the name of our file and press Enter. If we wanted to run a file called greet.go, the command will look like:

go build greet.go
While nothing obvious is shown after we run our command, if we type in the command ls, we’ll see our original Go program and executable file.
go build greet.go
While nothing obvious is shown after we run our command, if we type in the command ls, we’ll see our original Go program and executable file.

ls
greet     greet.go
To execute the file, we call ./greet

./greet
Note: If the Go compiler finds that our Go code isn’t written correctly, then it will throw an error and our Go program won’t compile. We’ll need to fix the error before the compiler can properly to do its job.

Now let’s go one step further and actually do it!

Instructions
1.
In the file provided (main.go), we have a Go program that prints “Hello World” to the terminal. We’ll get into what all the different parts are later, in this exercise, let’s get the program to execute!

In the terminal, type in go build main.go and press the enter key.

2.
Great, we’ve built our executable file, now it’s time to run it.

In the command line, type in ./main and enter to greet the world!

#hello world program

package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}


Running Files
Great, we were able to compile our program into an executable file that will always print out “Hello World”. If we want our program to run again, we don’t have to compile the program again, we simply run the executable file. Therefore, if we want fast code that users interact with, we’d compile a program once and use the executable file.

But what happens if we ever wanted to change our program? We would have to compile another executable file and then run that file. Imagine having to do that every single time just to check a small change or fix an error! 😱

Thankfully, we have the go run command followed by the name of the Go program. The go run command combines both the compilation and execution of code for us. This allows us to quickly check the output of our code (and for any errors). Unlike go build, go run will NOT create executable file in our current folder.

Instructions
1.
We provided the same Go file as the previous exercise.

This time, in the command line, type in go run main.go and press enter.

#main.go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}



Basic Go Structure: Packages
Now that we understand how to compile and run Go programs, let’s take a detailed look at the structure of Go’s program, specifically its packages:

package main 

import "fmt" 

func main () {
  fmt.Println("Hello World") 
}
Go programs are read from top to down, left to right, so let’s focus on the first line package main. This line is called a package declaration and every Go program starts with one. The package declaration informs the compiler whether to create an executable or library. In contrast to an executable, a library doesn’t outright run/execute code — it is a collection of code that can be used in other programs. Programs that have the package declaration package main (like ours) will create an executable file.

Next is a blank line. Go generally ignores these blank lines, they’re considered whitespace (new lines, spaces, and tabs). While our program doesn’t need the line break, it makes our code easier to read, or increases the program’s readability.

Then we have an import statement, import "fmt". The import keyword allows us to bring in and use code from other packages. We should also note that the package name is enclosed with double quotes ", otherwise our program will get an error and not compile/run.

Packages serve important roles in Go. They group related code together, allow code to be reusable, and make it easier to maintain. We only import the packages we need. In turn, our programs run faster because it’s not bogged down by extra code/packages!

Now that we have a general idea of how a Go program is written. Let’s start to write our own from scratch!

Instructions
1.
In main.go, create the package declaration using package main.

2.
Under the package declaration, import the "fmt" package.

We’ll continue building this program in the next exercise!

Note: Running the program (using the command go run main.go) will result in an error, instead, click the Run button after importing the "fmt" package.


# basiuc prg startswith this 

package main //package declaration every prg start with it
 
 import  "fmt"//The import keyword allows us to bring in and use code from other packages. double qoutes package nm
 
 




