package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func TrimAllStrings(a any) {
	visited := make(map[uintptr]any)

	var trim func(reflect.Value)
	trim = func(v reflect.Value) {
		switch v.Kind() {
		case reflect.Pointer:
			if v.IsNil() {
				return
			}
			k := v.Pointer()
			if _, seen := visited[k]; seen {
				return
			}
			visited[k] = struct{}{}
			trim(v.Elem())
		case reflect.Interface:
			if !v.IsNil() {
				trim(v.Elem())
			}
		case reflect.Struct:
			for _, field := range v.Fields() {
				if field.CanInterface() {
					trim(field)
				}
			}
		case reflect.String:
			if v.CanSet() {
				v.SetString(strings.TrimSpace(v.String()))
			}
		}
	}
	trim(reflect.ValueOf(a))
}

func main() {
	type Person struct {
		Name string
		Age  int
		Next *Person
	}

	a := &Person{
		Name: " name ",
		Age:  20,
		Next: &Person{
			Name: " name2 ",
			Age:  21,
			Next: &Person{
				Name: " name3 ",
				Age:  22,
			},
		},
	}

	TrimAllStrings(&a)

	m, _ := json.Marshal(a)

	fmt.Println(string(m))

	a.Next = a

	TrimAllStrings(&a)

	fmt.Println(a.Next.Next.Name == "name")
}
