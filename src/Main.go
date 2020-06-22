package main

// var wg = sync.WaitGroup{}

// type Ball struct {
// 	game     string
// 	duration int
// }
// var logCh = make(chan logEntry, 50)
// var doneCh = make(chan gou)

// type logEntry struct {
// 	name string
// }
// type gou struct{}

// func logger() {
// 	for {
// 		select {
// 		case entry := <-logCh:
// 			fmt.Println(entry.name)
// 		case <-doneCh:
// 			break
// 		}
// 	}
// }
func main() {

	// go logger()
	// logCh <- logEntry{name: "Gowtham"}
	// time.Sleep(10 * time.Microsecond)
	// doneCh <- gou{}
	// ch := make(chan int)
	// wg.Add(2)

	// go func(ch <-chan int) {
	// 	// i := <-ch
	// 	for i := range ch {
	// 		fmt.Println(i)
	// 	}
	// 	// ch <- 928
	// 	wg.Done()
	// }(ch)
	// go func(ch chan<- int) {
	// 	ch <- 927
	// 	// fmt.Println(<-ch)
	// 	ch <- 928
	// 	close(ch)
	// 	wg.Done()
	// }(ch)
	// wg.Wait()

	// go func() {
	// 	i := <-ch
	// 	fmt.Println(i)
	// 	wg.Done()
	// }()
	// for m := 0; m < 5; m++ {
	// 	wg.Add(2)
	// 	go func() {
	// 		j := 42
	// 		ch <- j
	// 		j = j + 5
	// 		wg.Done()
	// 	}()
	// }

	// a := 'G'
	// fmt.Println(a)
	// Array

	// myArray := [4]int{1, 2, 3, 4}
	// fmt.Println(myArray)

	// var a [3]int
	// a[1] = 1
	// a[2] = 2
	// a[0] = 0
	// fmt.Println(a)
	// b := a
	// a[0] = 98
	// b[0] = 99
	// fmt.Println(a)
	// fmt.Println(b)
	// c := &a
	// c[0] = 100
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	//slice

	// b := []int{1, 2, 3}
	// fmt.Println(b)
	// fmt.Printf("%v", cap(b))
	// b = append(b, []int{90, 5, 6, 7}...)
	// fmt.Println()
	// fmt.Printf("%v", cap(b))
	// fmt.Println(b)

	//map

	// var k = map[string]int{
	// 	"a": 1,
	// 	"b": 2,
	// 	"c": 3,
	// }
	// fmt.Println(k)
	// fmt.Println(k["b"])
	// k["d"] = 4
	// fmt.Println(k)
	// delete(k, "b")
	// fmt.Println(k)
	// y, ok := k["b"]
	// fmt.Println(y, ok)
	// y1, ok1 := k["a"]
	// fmt.Println(y1, ok1)

	// struct

	// type Employee struct {
	// 	name    string
	// 	age     int
	// 	company []string
	// }

	// emp1 := Employee{
	// 	name: "Gowtham",
	// 	age:  24,
	// 	company: []string{
	// 		"Cts", "EMIS", "Google",
	// 	},
	// }
	// fmt.Println(emp1.company[1])

	// emp2 := struct{ name string }{name: "Gowtham"}
	// fmt.Println(emp2)

	//embedding

	// type Ball struct {
	// 	game     string
	// 	duration int
	// }

	// type CricketBall struct {
	// 	Ball
	// 	place string
	// }
	// o := CricketBall{}
	// fmt.Printf("0 = %+v\n", o)

	// if n := 10; n < 0 {
	// 	fmt.Println("wrong")
	// } else {
	// 	fmt.Println("Cool")
	// }

	// Defer panic and Recover
	// fmt.Println("started")
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println("Error = ", err)
	// 	} else {
	// 		log.Println(err)
	// 	}
	// }()
	// // panic("something wrong")
	// fmt.Println("next")

	//pointers

	// type mystruct struct {
	// 	foo int
	// }
	// var ms *mystruct
	// ms = new(mystruct)
	// ms.foo = 9
	// fmt.Println(ms.foo)

	//functions

	// greet := "hello"
	// name := "Gowtham"
	// callName(&greet, &name)
	// fmt.Println(greet, " ", name)
	// s, err := sum(12.0, 2.0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(s)
	// var f func(int) = func(i int) {
	// 	fmt.Println(i + 10)
	// }
	// f(6)
	// k := animal{name: "KinKong", age: 78}
	// k.desc()
	// var s Shape = Cube{}
	// k := s.(Cube)
	// fmt.Println(k.Area())
	// fmt.Println(k.Volume())
	// fmt.Println("Thread %v\n", runtime.GOMAXPROCS(-1))
	// fmt.Println("started")
	// go firstFun("hello")
	// go secondFun([]int{1, 2, 3, 4, 5})
	// time.Sleep(10 * time.Microsecond)
	// fmt.Println("end line")

}

// func firstFun(k string) {
// 	for _, c := range k {
// 		fmt.Println(c)
// 	}
// }
// func secondFun(j []int) {
// 	for _, c1 := range j {
// 		fmt.Println(c1)
// 	}
// }

// type Shape interface {
// 	Area() float64
// }
// type Object interface {
// 	Volume() float64
// }

// type Cube struct {
// 	side float64
// }

// func (c Cube) Area() float64 {
// 	return 6 * (c.side * c.side)
// }
// func (c Cube) Volume() float64 {
// 	return c.side * c.side * c.side
// }

// type animal struct {
// 	name string
// 	age  int
// }

// func (an animal) desc() {
// 	fmt.Printf("Animal name is %s age is %i", an.name, an.age)
// }

// func callName(g, n *string) {
// 	fmt.Println(*g, " ", *n)
// 	*g="welcome"
// 	fmt.Println(*g, "jhj ", *n)
// }
// func sum(k string, val ...int) (result int) {
// 	result := 0
// 	for _, v := range val {
// 		result = result + v
// 	}
// 	fmt.Println(k, result)
// 	return result

// func sum(a, b float64) (float64, error) {
// 	if b == 0.0 {
// 		return 0.0, fmt.Errorf("0 is not the correct 2nd fn")
// 	}
// 	return a / b, nil
// }
