package main

import "fmt"

type Inter1 interface {
	fun1()
	fun2()
}

type Worker struct {
	Name string
	Age  int
	Inter1
}

func (w Worker) fun1() {
	fmt.Println("Worker fun1")
}

type Salary struct {
	Money int
}

func (s Salary) fun1() {
	fmt.Println("Salary fun1")
}
func (s Salary) fun2() {
	fmt.Println("Salary fun2")
}

func main() {
	s := &Salary{}
	w := Worker{
		Inter1: s,
		Age:    33,
		Name:   "bowen",
	}

	fmt.Println(w.Age)
	fmt.Println(w.Name)
	w.fun1() //调用Work自己fun1
	w.fun2() //调用Salary对象s的fun2
	w.Inter1.fun1()
	w.Inter1.fun2()
	// 无法访问 Money 属性，可以增加方法来实现
}
