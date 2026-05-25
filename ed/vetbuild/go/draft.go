package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Feito pelo video da aula gravada "ed man s08e01 vector"
type Vector struct {
	data     []int
	size     int
	capacity int
}

// Feito pelo video da aula gravada "ed man s08e01 vector"
func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

// Feito pelo video da aula gravada "ed man s08e01 vector"
func (vec *Vector) String() string {
	return "[" + Join(vec.data[0:vec.size], ", ") + "]"
}

// Feito pelo video da aula gravada "ed man s08e01 vector"
func (vec *Vector) Status() string {
	return fmt.Sprintf("size:%v capacity:%v", vec.size, vec.capacity)
}

/*
// Feito pelo video da aula gravada "ed man s08e01 vector"
func (vec *Vector) PushBack(valor int) {
	vec.data[vec.size] = valor
	vec.size += 1
}
*/
// Feito pelo video da aula gravada "ed man s08e01 vector"
func (vec *Vector) PushBack(valor int) error {
	if vec.size < vec.capacity {
		vec.data[vec.size] = valor
		vec.size += 1
		return nil
	}

	return fmt.Errorf("")
}

// Feito com ajuda do Professor David Sena via Discord.
func (vec *Vector) IndexOf(valor int) int {
	// var a int = 0

	for i := range vec.size {
		if vec.data[i] == valor {
			return i
		}
	}
	return -1
}

// Feito modificando a funcao "IndexOf", que foi feita com a ajuda do Professor David Sena via Discord.
func (vec *Vector) Contains(valor int) bool {

	for i := range vec.size {
		if vec.data[i] == valor {
			return true
		}
	}
	return false

}

// Feito pelo video da aula gravada "edm s05e01 vetbuild"
func (v *Vector) Reserve(capacity int) {
	if capacity < v.size {
		return
	}

	novo := make([]int, capacity)
	for i := range v.size {
		novo[i] = v.data[i]
	}

	v.capacity = capacity
	v.data = novo
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewVector(value)
		case "push":
			// for _, part := range parts[1:] {
			// 	value, _ := strconv.Atoi(part)
			// 	v.PushBack(value)
			// }
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				err := v.PushBack(value)
				if err != nil {
					fmt.Print(err) // antes usava fmt.Println
				}
			}

		case "show":
			fmt.Println(v)
		case "status":
			fmt.Println(v.Status())
		case "pop":
			// err := v.PopBack()
			// if err != nil {
			// 	fmt.Println(err)
			// }
		case "insert":
			// index, _ := strconv.Atoi(parts[1])
			// value, _ := strconv.Atoi(parts[2])
			// err := v.Insert(index, value)
			// if err != nil {
			// 	fmt.Println(err)
			// }
		case "erase":
			// index, _ := strconv.Atoi(parts[1])
			// err := v.Erase(index)
			// if err != nil {
			// 	fmt.Println(err)
			// }
		case "indexOf":
			// value, _ := strconv.Atoi(parts[1])
			// index := v.IndexOf(value)
			// fmt.Println(index)
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			// value, _ := strconv.Atoi(parts[1])
			// if v.Contains(value) {
			// 	fmt.Println("true")
			// } else {
			// 	fmt.Println("false")
			// }
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			// v.Clear()
		case "capacity":
			// fmt.Println(v.Capacity())
		case "get":
			// index, _ := strconv.Atoi(parts[1])
			// value, err := v.At(index)
			// if err != nil {
			// 	fmt.Println(err)
			// } else {
			// 	fmt.Println(value)
			// }
		case "set":
			// index, _ := strconv.Atoi(parts[1])
			// value, _ := strconv.Atoi(parts[2])
			// err := v.Set(index, value)
			// if err != nil {
			// 	fmt.Println(err)
			// }
			//
		case "reserve":
			// newCapacity, _ := strconv.Atoi(parts[1])
			// v.Reserve(newCapacity)
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
