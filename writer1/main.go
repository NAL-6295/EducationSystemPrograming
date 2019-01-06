package main

import (
	"fmt"
)

type receiverSampleType struct {
	value int
}

func main() {
	instance := receiverSampleType{}

	instance.value = 1
	instance.notPointerMethod(3)
	fmt.Println(instance) // output : 1
	instance.PointerMethod(3)
	fmt.Println(instance) // output : 3

}

///型名でレシーバーを定義
func (t receiverSampleType) notPointerMethod(value int) {
	t.value = value
}

///ポインタ付きでレシーバーを定義
func (t *receiverSampleType) PointerMethod(value int) {
	t.value = value
}
