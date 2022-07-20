### 指针：

![1658295670114](.\指针.jpg)



### 值类型和引用类型

- 引用传递效率高

![值类型和引用类型](.\值类型和引用类型.jpg)

![基本介绍](.\基本介绍.jpg)

### 值传递函数调用过程

![值传递函数调用](.\值传递函数调用.jpg)

![值类型传递函数说明](值类型传递函数说明.jpg)

### 指针传递函数调用过程

![指针传递函数调用](.\指针传递函数调用.jpg)



### 函数类型用法

~~~go
package main
import "fmt"

//函数也是一种数据类型，通过type申明函数类型，只是能为了给函数类型传参
type myf func(int ,int) int

//传入myf栈
func testM(my myf,i1,i2 int)  {
    //执行myf栈
	fmt.Println(my(i1,i2 ))
}
//该方法符合myf类型
func ww(i1 ,i2 int) int {
	return i1+ i2
}
func main() {
	testM(ww,1,2)
}

~~~



### defer

defer栈先进后出，并把当时的值压入栈中

~~~go

func main() {
	v := 2
	v2 := 3
	defer fmt.Println(v)  //4     输出2
	defer fmt.Println(v2) //3     输出3
	v++
	v2++				//执行顺序
	fmt.Println(v)      //1    	  输出3
	fmt.Println(v2)		//2	      输出4
}
~~~



### new

用于分配内部，主要用于分配值类型

![new](.\new.jpg)





### 异常处理机制

defer  + recover 配合使用

~~~go
package main

import "fmt"

func test()  {
	defer func() {
		err := recover()   //负责接收函数内的所有错误，判断是否有收到错误
		if err != nil {
			fmt.Println(111,err)
		}
	}()
	a := 10
	b := 0
	nun := a/b    //此時出現panic，并將错误传给recover
	println(nun)
}

func main() {
	test()
	fmt.Println(1111)  //引发panic后还可以继续执行该代码
}
~~~



### 数组内存分配

![数组内存分步](.\数组内存分步.jpg)

### 切片

#### 截取数据内存分布

![切片内存分布](切片内存分布.jpg)

#### make创建切片内存分布

![make创建切片的内部分布](make创建切片的内部分布.jpg)
