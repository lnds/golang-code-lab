# Hello World

Do you know the classic hello world program?

If you don't, there is a modified version of this in the file hello.go.

You can run this program this way:

    % go run hello.go
    type your name: Eduardo
    
    Hello Eduardo


(If this doesn't work, see the instructions to install Go in your machine in the README.md file at the root of this repo)


## Challenge

Modify the hello world program to receive a list of names (standard input) and say "Hello {name}" for every one.


Time to complete approx: 10 min

## Tips

Instead of a `bufio.Reader` is better to use a `bufio.Scanner` that is more suitable to read text files (and lines).

See official documentation:

    bufio package: https://pkg.go.dev/bufio@go1.17.5

    bufio reader: https://pkg.go.dev/bufio@go1.17.5#Reader

    bufio scanner: https://pkg.go.dev/bufio@go1.17.5#Scanner

