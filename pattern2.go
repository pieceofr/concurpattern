package main

import (
	"fmt"
	"sync"
	"time"
)

// Simple Channel Goroutine Pattern
func Simple() {
	// create new channel of type int
	ch := make(chan int)
	// start new anonymous goroutine
	go func() {
		// send 42 to channel
		ch <- 42
	}()
	// read from channel
	fmt.Println("Simple Result:", <-ch)
}

// Timer Goroutine Pattern
func TimerP() {
	fmt.Println("Timer Pattern Start, tick 5 times per sec")
	for i := 0; i < 5; i++ {
		c := timer(1 * time.Second)
		fmt.Println("Timer Tick", <-c)
	}
}

func timer(d time.Duration) <-chan int {
	c := make(chan int)
	go func() {
		time.Sleep(d)
		c <- 1
	}()
	return c
}

func PingPon() {
	fmt.Println("PingPon Pattern Start, 3 Player play 3 Sec")
	var Ball int
	table := make(chan int)
	go player(table, "1")
	go player(table, "2")
	go player(table, "3")

	table <- Ball
	time.Sleep(3 * time.Second)
	fmt.Println("TimerP Result:", <-table)

}

func player(table chan int, p string) {
	for {
		ball := <-table
		fmt.Println("player ", p, " ball:", ball)
		ball++
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}

func FadeIn() {
	fmt.Println("FadeIn Pattern Start, 1 producer 100ms per input , 2 producer 250ms per input")
	ch := make(chan int)
	out := make(chan int)
	go producer(ch, 100*time.Millisecond, "1")
	go producer(ch, 250*time.Millisecond, "2")
	go reader(out)
	//也可以這樣寫

	for i := 0; i < 30; i++ { //印 30 個就好
		out <- <-ch
		i++
	}

	/*
		for i := range ch {
			fmt.Println("FadeIn:", i)
			out <- i
		}
	*/

}

func producer(ch chan int, d time.Duration, id string) {
	var i int
	if id == "2" {
		i = 10
	}
	for {
		fmt.Println("player:", id, " i:", i)
		ch <- i
		i++
		time.Sleep(d)
	}
}

func reader(out chan int) {
	for x := range out {
		fmt.Println("Reader:", x)
	}
}

func FadeOut() { //Worker Pattern
	fmt.Println("FadeOut/Worker Pattern Start, 3 workers, 14 tasks")
	var wg sync.WaitGroup
	wg.Add(3)
	go pool(&wg, 3, 7)
	wg.Wait()
}

func worker(tasksCh <-chan int, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Println("Start worker:", id)
	for {
		task, ok := <-tasksCh
		//fmt.Println(time.Now(), "--worker:", id, " ok:", ok, " task:", task)
		if !ok {
			return
		}
		//d := time.Duration(task) * time.Millisecond
		d := time.Duration(1) * time.Millisecond
		time.Sleep(d)
		fmt.Println(time.Now(), "worker:", id, "processing task", task)
	}
}

func pool(wg *sync.WaitGroup, workers, tasks int) {
	tasksCh := make(chan int)

	for i := 0; i < workers; i++ {
		go worker(tasksCh, wg, i)
	}
	time.Sleep(100 * time.Millisecond) // 等 worker 全起來, 方便看順序
	for i := 0; i < tasks; i++ {
		tasksCh <- i
	}

	close(tasksCh)
	fmt.Println("pool task and worker assign finished")
}

const (
	WORKERS    = 1
	SUBWORKERS = 3
	TASKS      = 1
	SUBTASKS   = 5
)

func WorkerSubWorker() {
	fmt.Println("SubWorker Pattern Start, 3 workers, 2 subworker , 15 tasks, 10 subtask")
	var wg sync.WaitGroup
	wg.Add(WORKERS)
	tasks := make(chan int)

	for i := 0; i < WORKERS; i++ {
		go worker2(tasks, &wg, i)
	}

	time.Sleep(100 * time.Millisecond) // 等 worker 全起來, 方便看順序

	for i := 0; i < TASKS; i++ {
		tasks <- i
	}
	close(tasks)
	fmt.Println("tasks all finished!")
	wg.Wait()
}

func worker2(tasks <-chan int, wg *sync.WaitGroup, id int) {
	fmt.Println("*Worker2 Up:", id)
	defer wg.Done()
	for {
		task, ok := <-tasks
		fmt.Println("Worker2:", id, " task:", task, "ok", ok)
		if !ok {
			return
		}

		subtasks := make(chan int)
		for i := 0; i < SUBWORKERS; i++ {
			go subworker(subtasks, id, i)
		}
		time.Sleep(100 * time.Millisecond)
		for i := 0; i < SUBTASKS; i++ {
			task1 := task * i
			subtasks <- task1
		}
		close(subtasks)
	}
}

func subworker(subtasks chan int, parent, id int) {
	fmt.Println("subworker Up:", parent, "-", id)
	for {
		task, ok := <-subtasks
		fmt.Println("subworker:", parent, "-", id, " task:", task, "ok", ok)
		if !ok {
			return
		}
		time.Sleep(time.Duration(task) * time.Millisecond)
		fmt.Println(task)
	}
}
