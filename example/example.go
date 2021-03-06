package main

import (
	"encoding/json"
	"fmt"
	"github.com/lforeverhao/hookpanic"
)

type PanicInfo struct {
	Stack, Panic string
}

func main() {
	hookpanic.SetPanicHandler(func(a interface{}, stack []byte) {
		info := PanicInfo{}
		info.Panic = fmt.Sprint(a)
		info.Stack = string(stack)

		bytes, _ := json.Marshal(info)
		fmt.Println(string(bytes))
	})

	panic("shit happens")
}
