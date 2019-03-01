package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	//指针
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	p = &j
	*p = *p / 37
	fmt.Println(j)

	//结构体
	type Vertex struct {
		X int
		Y int
	}
	fmt.Println(Vertex{1, 2})

	//结构体字段
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	//结构体指针
	vp := Vertex{1, 2}
	p2 := &vp
	p2.X = 1e9
	fmt.Println(vp)

	//结构体文法
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		p3 = &Vertex{1, 2}
	)
	fmt.Println(v1, v2, v3, p3)

	//数组
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//切片
	//包括第一个元素，但排除最后一个元素
	var s = primes[1:4]
	fmt.Println("s")
	fmt.Println(s)

	//更改切片的元素会修改其底层数组中对应的元素,
	// 与它共享底层数组的切片都会观测到这些修改
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo", //最后一个也需要逗号
	}
	fmt.Println(names)
	qa := names[0:2]
	qb := names[1:3]
	fmt.Println(qa, qb)

	qb[0] = "XXX"
	fmt.Println(qa, qb)
	fmt.Println(names)

	//切片文法: 类似于没有长度的数组
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	qs := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(qs)

	//切片的默认行为
	//切片下界的默认值为 0，上界则是该切片的长度。
	qm := []int{2, 3, 5, 7, 11, 13}

	qm = qm[1:4]
	fmt.Println(qm)

	qm = qm[:2]
	fmt.Println(qm)

	qm = qm[1:]
	fmt.Println(qm)

	//切片的长度与容量
	//切片的长度就是它所包含的元素个数。
	//切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
	//切片 qc 的长度和容量可通过表达式 len(qc) 和 cap(qc) 来获取。
	qc := []int{2, 3, 5, 7, 11, 13}

	qc = qc[:0] // 截取切片使其长度为 0
	printSlice(qc)

	qc = qc[:4] // 拓展其长度
	printSlice(qc)

	qc = qc[2:] // 舍弃前两个值
	printSlice(qc)

	qc = qc[:4]
	printSlice(qc)

	//nil切片
	//切片的零值是nil
	//nil切片的长度和容量为0且没有底层数组
	var qn []int
	printSlice(qn)
	if qn == nil {
		fmt.Println("nil!")
	}

	//用make创建切片
	ma := make([]int, 5)
	printSlice2("ma", ma)

	mb := make([]int, 0, 5)
	printSlice2("mb", mb)

	mc := mb[:2]
	printSlice2("mc", mc)

	md := mc[2:5]
	printSlice2("md", md)

	//切片的切片
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	} // 创建一个井字板（经典游戏）

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X" // 两个玩家轮流打上 X 和 O

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	//向切片追加元素
	var sa []int
	printSlice3(sa)

	sa = append(sa, 0)
	printSlice3(sa)

	sa = append(sa, 1)
	printSlice3(sa)

	sa = append(sa, 2, 3, 4)
	printSlice3(sa)

	//range
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // == 2**i
		// i 是索引
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}

	//切片练习
	pic := Pic(3, 3)
	fmt.Println(pic)

	//映射
	type Vertex2 struct {
		Lat, Long float64
	}
	var m map[string]Vertex2
	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	//映射的文法
	var m2 = map[string]Vertex2{
		"Bell Labs": Vertex2{
			40.68433, -74.39967,
		},
		"Google": Vertex2{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m2)
	var m3 = map[string]Vertex2{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m3)

	//修改映射
	//插入或修改元素: m[key] = elem
	//获取元素: elem = m[key]
	//删除元素: delete(m, key)
	//通过双赋值检测某个键是否存在: elem, ok = m[key]
	//若key在m中，ok为true;否则,ok为false,若key不在映射中,那么elem是该映射元素类型的零值。
	//当从映射中读取某个不存在的键时，结果是映射的元素类型的零值。
	mm := make(map[string]int)
	mm["Answer"] = 42
	fmt.Println(m["Answer"])
	mm["Answer"] = 48
	fmt.Println(mm["Answer"])
	delete(mm, "Answer")
	fmt.Println(mm["Answer"])
	value, ok := mm["Answer"]
	fmt.Println(value, ok)
	fmt.Println("mmabs")
	fmt.Println(mm["abs"])

	//映射练习
	wm := WordCount("a b s a")
	fmt.Println(wm)
	fmt.Println(wm["a"])

	//函数值
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	//函数的闭包
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

	//闭包练习: 斐波纳契闭包
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap =%d %v\n", s, len(x), cap(x), x)
}

func printSlice3(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for i := range p {
		p[i] = make([]uint8, dx)
	}
	for y, row := range p {
		for x := range row {
			row[x] = uint8(x * y)
		}
	}
	return p
}

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, f := range strings.Fields(s) {
		m[f]++
	}
	return m
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	//函数 adder 返回一个闭包。每个闭包都被绑定在其各自的 sum 变量上。
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	f0, f1, f2 := 0, 1, 0
	index := 0
	return func() int {
		if index == 0 {
			index += 1
			return f0
		} else if index == 1 {
			index += 1
			return f1
		} else {
			f2 = f0 + f1
			f0 = f1
			f1 = f2
			return f2
		}
	}
}
