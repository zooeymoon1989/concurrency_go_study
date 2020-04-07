package main

import "fmt"

type celsius float64

func (c celsius) String() string {
	return fmt.Sprintf("%.2f C", c)
}

type temperature =  celsius


func main() {
	c := celsius(10.0)
	fmt.Println(c)

	t := c
	fmt.Println(t)

}

// Func is a very useful thing to document.
//type Func func()
//
//func run(f Func)  {
//	f()
//}
//
//type anotherFunc Func
//
//
//func hello()  {
//	fmt.Println("hello")
//}
//
//func main() {
//	var f anotherFunc = hello
//	fmt.Printf("%T\n",f)
//	run(Func(f))
//}

//type Strings interface {
//	String() string
//}
//
//func main() {
//	var f Strings
//	var s fmt.Stringer
//	fmt.Println(f == s)
//	f = s
//	s = f
//}
