package Arr

import (
	"fmt"
	"strconv"
)

const capacity = 10

type Arr struct {
	data [capacity]int
	size int
}

// 获取数组的容量
func (arr Arr) getCapacity() int {
	return len(arr.data)
}

// 获取数组中的元素个数
func (arr Arr) getSize() int {
	return arr.size
}

// 返回数组是否为空
func (arr Arr) isEmpty() bool {
	return arr.size == 0
}

//实例化Arr
func newArr() *Arr {
	return &Arr{
		data: [capacity]int{},
		size: 0,
	}
}

// 在index索引的位置插入一个新元素e
func (arr *Arr) add(index int, e int) {
	if arr.size == len(arr.data) {
		return
	}
	if index < 0 || index > arr.size {
		return
	}
	for i := arr.size - 1; i >= index; i-- {
		arr.data[i+1] = arr.data[i]
	}
	arr.data[index] = e
	arr.size++
}

// 向所有元素后添加一个新元素
func (arr *Arr) addLast(e int) {
	arr.add(arr.size, e)
}

// 在所有元素前添加一个新元素
func (arr *Arr) addFirst(e int) {
	arr.add(0, e)
}

// 获取index索引位置的元素
func (arr *Arr) get(index int) int {
	if index < 0 || index >= arr.size {
		return -1
	}
	return arr.data[index]
}

// 获取index索引位置的元素
func (arr *Arr) set(index int, e int) {
	if index < 0 || index >= arr.size {
		return
	}
	arr.data[index] = e
}

// 打印数组
func (arr *Arr) print() {
	s := "Array:size = " + strconv.Itoa(arr.size) + ", capacity = " + strconv.Itoa(len(arr.data))
	fmt.Println(s)
	fmt.Printf("%v\n", arr.data)
}

// 查找数组中是否有元素e
func (arr *Arr) contains(e int) bool {
	for i := 0; i < arr.size; i++ {
		if arr.data[i] == e {
			return true
		}
	}
	return false
}

// 查找数组中元素e所在的索引，如果不存在元素e，则返回-1
func (arr *Arr) find(e int) int {
	for i := 0; i < arr.size; i++ {
		if arr.data[i] == e {
			return i
		}
	}
	return -1
}

// 从数组中删除index位置的元素, 返回删除的元素
func (arr *Arr) remove(index int) int {
	if index < 0 || index >= arr.size {
		return -1
	}
	ret := arr.data[index]
	for i := index + 1; i < arr.size; i++ {
		arr.data[i-1] = arr.data[i]
	}
	arr.size--
	return ret
}

// 从数组中删除第一个元素, 返回删除的元素
func (arr *Arr) removeFirst() int {
	return arr.remove(0)
}

// 从数组中删除最后一个元素, 返回删除的元素
func (arr *Arr) removeLast() int {
	return arr.remove(arr.size - 1)
}

// 从数组中删除元素e
func (arr *Arr) removeElement(e int) {
	index := arr.find(e)
	if index != -1 {
		arr.remove(index)
	}

}
func Run() {
	arr := newArr()
	arr.add(0, 2)
	arr.addFirst(1)
	arr.addLast(3)
	arr.add(1, 5)
	arr.set(0, 7)
	arr.print()
	arr.removeElement(5)
	arr.print()
}
