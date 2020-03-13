package main

import "fmt"

type bot interface{
	getGreeting() string
}

type englishBot struct{}
type spainshBot struct{}

func main(){
	eb := englishBot{}
	sb := spainshBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot){
		fmt.Println(b.getGreeting())
	}

// func printGreeting(eb englishBot){
// 	fmt.println(eb.getGreeting())
// }

// func printGreeting(sb spainshBot){
// 	fmt.println(sb.getGreeting())
// }

func (eb englishBot) getGreeting() string {
	return "Hi There!"
}

func (sb spainshBot) getGreeting() string {
	return "Hola!"
}