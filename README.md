# GoSpongebob

GoSpongebob is a CLI  implementation of the [Mocking Spongebob](http://knowyourmeme.com/memes/mocking-spongebob) meme written in the Go programming language. 

![Mocking Spongebob](https://i.imgur.com/PVSJvux.jpg)

### Download and Install 

1. Download and configure Go for your operating system. [See Here](https://golang.org/dl/)
2. Run `go get -u github.com/penguingovernor/GoSpongebob`

### Usage 
1. Running `GoSpongebob` will produce an empty screen with a cursor, feel free to type anything you like here!
2. After typing your deseried text press `Ctrl-d` to send the EOF command 
3. Upon seding the EOF command the program should output your text in a Mocking-Spongebob-esque manner! 

### Alternative Usage 
To convert a file to Mocking-Spongebob text run: 
`cat fileName | GoSpongebob`

To save the Mocking-Spongebob text to a file run: 
`cat fileName | GoSpongebob > outputFileName`
