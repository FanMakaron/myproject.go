package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"time"
)

// базовая функция
func add(x, y int) int {
	return x + y
}

// Возвращение нескольких значений
func swap(x, y string) (string, string) {
	return y, x
}

// "naked" return - возвращение заранее определенных переменных
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var i, j = 1, 2

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	// интересный оператор << и >>
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func printSlice(slice []int, sliceName string) {
	fmt.Printf("%v is %v len=%d cap=%d %v\n", sliceName, reflect.TypeOf(slice).Kind(), len(slice), cap(slice), slice)
}

func arraySliceMap() {
	// Массивы имеют фиксированный размер в статичной памяти
	// объявление массива через var
	var arrayVar [5]int
	arrayVar[0] = 1
	arrayVar[1] = 2
	arrayVar[2] = 3
	fmt.Printf("arrayVar is %v len=%d cap=%d %v\n", reflect.TypeOf(arrayVar).Kind(), len(arrayVar), cap(arrayVar), arrayVar)

	// объявление массива в короткой форме
	arr := [5]int{1, 3}
	fmt.Printf("arr is %v len=%d cap=%d %v\n", reflect.TypeOf(arr).Kind(), len(arr), cap(arr), arr)

	//Объявление слайса
	// диапазон в слайсе включает первый элемент, но доходит до последнего без его включения [0:1] - содержит один нулевой элемент
	// диапазон можно не указывать, [:1] - первый элемент будет нулевой, [1:] - последний элемент будет равен длине массива
	// Длина len() — это количество элементов, которые содержит срез или массив.
	// Емкость cap() — это количество элементов в базовом массиве, считая от ПЕРВОГО элемента В СРЕЗЕ.
	var sliceVar = arrayVar[1:3]
	printSlice(sliceVar, "sliceVar")

	// слайсы - указатели на интервал массива, изменения значений в слайсе приводит к изменению значений в массиве
	sliceVar[1] = 4
	fmt.Printf("arrayVar(updated) is %v len=%d cap=%d %v\n", reflect.TypeOf(arrayVar).Kind(), len(arrayVar), cap(arrayVar), arrayVar)
	// Емкость среза cap() — это количество элементов в базовом массиве, считая от ПЕРВОГО элемента В СРЕЗЕ.
	sliceSmall := arrayVar[2:4]
	printSlice(sliceSmall, "sliceSmall")

	// Нулевое значение среза равно нулю.
	// Нулевой срез имеет длину и емкость 0 и не имеет базового массива.
	var sliceNil []int
	printSlice(sliceNil, "sliceNil")
	if sliceNil == nil {
		fmt.Println("sliceNil is nil!")
	}

	// Срезы можно создавать с помощью make, при этом сначала будет создан массив с указанным размером а затем сразу срез на него
	sliceMake := make([]int, 5)
	// что будет эквивалентно
	sli := []int{0, 0, 0, 0, 0}
	// а еще можно указав емкость третьим аргументом
	sliceMakeCap := make([]int, 5, 6)
	printSlice(sliceMake, "sliceMake")
	printSlice(sli, "sli")
	printSlice(sliceMakeCap, "sliceMakeCap")

	//Slices of slices
	//Slices can contain any type, including other slices.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	fmt.Printf("board is %v len=%d cap=%d %v\n", reflect.TypeOf(board).Kind(), len(board), cap(board), board)

	// Appending
	var sliceAppend []int
	printSlice(sliceAppend, "sliceAppend")
	sliceAppend = append(sliceAppend, 0)
	printSlice(sliceAppend, "sliceAppend(appended)")

	//foreach
	//The first is the index, and the second is a COPY of the element at that index.
	fmt.Println("foreach:")
	for i, v := range sliceVar {
		fmt.Printf("index is %d value is %d\n", i, v)
	}
	//You can skip the index or value by assigning to _.
	//for i, _ := range pow
	//for _, value := range pow

	// Map
	var mapVar = map[string]int{"abc": 2}
	//mapVar = make(map[string]int)
	fmt.Printf("mapVar is %v len=%d %v\n", reflect.TypeOf(mapVar).Kind(), len(mapVar), mapVar)
	mapVar["qwe"] = 1
	fmt.Printf("mapVar is %v len=%d %v\n", reflect.TypeOf(mapVar).Kind(), len(mapVar), mapVar)
	// isset():
	//mapVar["Answer"] = 0
	v, ok := mapVar["Answer"]
	fmt.Println("isset(mapVar[\"Answer\"]):", ok, "Его значение:", v)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

type Vertex struct {
	X, Y float64
}

// Метод привязанный к определенному типу Vertex
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// если убрать указатель, то вызов Scale ничего не изменит
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func methodsInterfaces() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())

	p := &Vertex{4, 3}
	// и так тоже будет работать потому что  method call p.Scale() is interpreted as (*p).Scale().
	p.Scale(3)
	fmt.Println(v.Abs())
}

func main() {
	fmt.Println("----------add, swap, split----------")
	fmt.Println(add(42, 13))
	fmt.Println(swap("hello", "word"))
	fmt.Println(split(25))

	// предопределение переменных в одной строке
	//var c, python, java = true, false, "no!"
	// вместо использования var можно использовать := - но только внутри функции
	fmt.Println("----------предопределение и короткая форма присваивания ----------")
	c, python, java := true, false, "no!"
	fmt.Println(i, j, c, python, java)

	fmt.Println("----------Автоприведение типов констант----------")
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	fmt.Println("----------The init and post statements are optional, если их убрать for будет имитировать while.----------")
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// Бесконечный цикл
	//for {
	//}

	// в IF - условие можно писать без скобок для условий, но тело обязательно {}
	// так же в IF можно объявить новую переменную, но она будет видна в рамках тела условия, в том числе и else
	// else можно опустить, начав с новой строки фигурную скобку ИЛИ его нужно писать на той же строке что и закрывающую фигурную скобку IF
	fmt.Println("----------Эксперименты с IF----------")
	var b float64
	if v := math.Pow(10, 12); v < 15 {
		b = v
	}
	{
		b = 15
	}
	fmt.Println(b)

	// нет break, выходит сразу после первого нахождения
	// Switch without a condition is the same as switch true.
	fmt.Println("----------SWITCH----------")
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// defer - отложить выполнение после возврата из функции
	// вызываются в обратном порядке
	defer fmt.Println("----------this is DEFER! (not sparta)----------")

	//
	fmt.Println("----------Указатели----------")
	var p *int
	i := 42
	p = &i
	fmt.Println(*p) // read i through the pointer p
	*p = 21         // set i through the pointer p
	fmt.Println(i)  // read i through the pointer p

	fmt.Println("----------Структуры----------")
	st := struct {
		X int
		Y int
	}{1, 2}
	fmt.Println(st)
	p1 := &st
	//// При работе со свойствами структуры через ее указатель, не нужно разыменовывать, это делается автоматически
	p1.Y = 3
	fmt.Println(st)

	fmt.Println("----------Array & Slices----------")
	arraySliceMap()

	fmt.Println("----------Замыкания----------")
	pos, neg := adder(), adder()
	for i := 0; i < 3; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	fmt.Println("----------Методы и интерфейсы----------")
	methodsInterfaces()

	fmt.Println("----------Совсем всякий мусор----------")
}
