package main

import "fmt"
import "github.com/yusufpapurcu/wmi"
import "log"

type Win32_Process struct {
	Name string
	ExecutionState uint16
	Status string
}

func main() {
	var dst []Win32_Process
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range dst {
		if v.Status == "" {
			v.Status = "Empty"
		}
		fmt.Println(i, v.Name, v.ExecutionState, v.Status)
	}
}