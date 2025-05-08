package main

// 演示reflect包的使用
// 反射是Go语言中的一个强大特性，允许程序在运行时检查类型和变量的值。

import (
	"fmt"
	"reflect"

)

// 定义一个结构体用于演示
type Person12 struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"gte=0,lte=130"`
	/*
	使用标签（tag）定义了 JSON 序列化规则和验证规则
	json:"name" 指定 JSON 中的字段名
	validate:"required" 表示该字段必填
	validate:"gte=0,lte=130" 表示年龄范围在 0-130
	*/
}

func Test12() {
	// 创建结构体实例
	p := Person12{
		Name: "张三",
		Age:  25,
	}

	// 获取反射类型和值
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	// 打印类型信息
	fmt.Printf("类型名称: %v\n", t.Name())
	fmt.Printf("类型种类: %v\n", t.Kind())
	fmt.Printf("字段数量: %v\n", t.NumField())

	// 遍历结构体字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("\n字段名称: %v\n", field.Name)
		fmt.Printf("字段类型: %v\n", field.Type)
		fmt.Printf("字段标签: %v\n", field.Tag.Get("json"))
		fmt.Printf("字段值: %v\n", value.Interface())
	}

	// 通过反射修改值
	if v.Kind() == reflect.Struct {
		// 获取可修改的值
		newP := reflect.New(t).Elem()
		newP.Set(v)
		// 修改字段值
		if name := newP.FieldByName("Name"); name.IsValid() {
			name.SetString("李四")
		}
		// 获取修改后的实例
		modifiedP := newP.Interface().(Person12)
		fmt.Printf("\n修改后的名字: %v\n", modifiedP.Name)
	}
}
