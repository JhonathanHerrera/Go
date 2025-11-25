package concurrent

import (
	"fmt"
	"sync"
	"time"
)

type Student struct {
	studentName  string
	studentGrade string
	studentGPA   float32
}

func probationCheck(student Student) string {

	time.Sleep(1 * time.Second)


	if student.studentGPA < 2.5 {
		fmt.Printf("%s has a BAD GPA of: %.2f\n", student.studentName, student.studentGPA)
	} else {
		fmt.Printf("%s has a GREAT GPA of: %.2f\n", student.studentName, student.studentGPA)
	}

	return "Student has been reviewed"
}

func RunConcurrent() {

	students := []Student{
		{studentName: "Jhonathan", studentGrade: "A", studentGPA: 3.4},
		{studentName: "Jessica", studentGrade: "A", studentGPA: 3.9},
		{studentName: "Mark", studentGrade: "B", studentGPA: 2.01},
		{studentName: "Rob", studentGrade: "C", studentGPA: 3.33},
	}

	start := time.Now()

	wg := sync.WaitGroup{}

	for _, student := range students {
		wg.Add(1)
		go func(s Student) {
			defer wg.Done()
			probationCheck(s)
		}(student)
	}

	wg.Wait()

	fmt.Println("All students are done")

	elapsed := time.Since(start) 
	fmt.Printf("Total execution time: %v\n", elapsed)
}
