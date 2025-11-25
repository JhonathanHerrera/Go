package main

import (
	"fmt"
	"concurrency/sequential"
	"concurrency/concurrent"
)

//go run testMain.go sequential/sequential.go concurrent/concurrent.go


func main() {

	fmt.Println("\n=== Running Sequential ===")
	sequential.RunSequential()

	fmt.Println("\n=== Running Concurrent ===")
	concurrent.RunConcurrent()
}

/*
=== Running Sequential ===
Jhonathan has a GREAT GPA of: 3.40
Jessica has a GREAT GPA of: 3.90
Mark has a BAD GPA of: 2.01
Rob has a GREAT GPA of: 3.33
All students are done
Total execution time: 4.014391916s

=== Running Concurrent ===
Jessica has a GREAT GPA of: 3.90
Rob has a GREAT GPA of: 3.33
Jhonathan has a GREAT GPA of: 3.40
Mark has a BAD GPA of: 2.01
All students are done
Total execution time: 1.007250958s
*/
