package main

import (
	"fmt"
	//"log"
	"math/rand"
	"time"
)

func Common() {
	c := make(chan string)
	go boringCommon("boring", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("you say: %q\n", <-c)
	}
	fmt.Println("Leaving main")
}

func boringCommon(msg string, c chan string) {
	for i := 0; ; i++ {
		//fmt.Println(msg, i)
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}

}

func Generator() {
	c := boringGenerator("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("Leaving Generator!")
}

func boringGenerator(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func Handle() {
	joe := boringGenerator("Joe")
	ann := boringGenerator("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("Leaving Handle!")
}

func FadeInMain() {
	c := FadeInP(boringGenerator("joe"), boringGenerator("ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Leaving FadeInMain")
}

func FadeInP(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}
