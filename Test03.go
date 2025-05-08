package main

import "fmt"

// 选项模式是一种常用的设计模式，允许在创建对象时传递可选参数。
// 这种模式通常用于构建复杂对象，允许用户根据需要选择不同的选项。

type Person struct {
	Name string
	Age  int
}

type PersonOperations func(p *Person)

func withName(name string) PersonOperations {
	return func(p *Person) {
		p.Name = name
	}
}

func withAge(age int) PersonOperations {
	return func(p *Person) {
		p.Age = age
	}
}

func NewPerson(ops ...PersonOperations) *Person {
	p := &Person{}
	for _, op := range ops {
		op(p)
	}
	return p
}

func Test03() {
	Myperson := NewPerson(withName("John"), withAge(30))
	fmt.Println(Myperson.Name) // 输出: John
}
