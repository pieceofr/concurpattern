# Golang Concurrency Pattern 
This code is base on "Visualizing Concurrency in Go" article. 
The code is slitly changed in order to printout in windows and have a better understanding in concurrency process.

### Reference 
[Visualizing Concurrency in Go](http://divan.github.io/posts/go_concurrency_visualize/)

### Concurrency Pattern Listed
1:Simple  : A simple channel and goroutine pattern
2:Timer   : Timer ticks every 1 sec
3:PingPon : Players recieve and send to each other through a channel
4:FadeIn  : Different goroutines send their request to a task handler to do jobs
5:Worker  : Many tasks has assigned to a task handler and handler assign jobs to  several workers
6:WorkerSubWorker :Many tasks has assigned to a task handler and handler assign jobs to  several workers and each worker assign sub-jobs to subworkers

### To Demostrate
In windows:
go build
goroutine.exe -pattern=1

pattern argument can be follow
1:Simple 2:Timer 3:PingPon 4:FadeIn 5:Worker 6:WorkerSubWorker

### Notes
pattern2.go is based on the reference
pattern1.go is some code from Gophercon (it does not in demostrations)