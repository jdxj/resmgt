package handler

import (
	"encoding/json"
	"fmt"
	"testing"
)

type People struct {
	Name string
	Age  *int
}

func TestJsonEncode(t *testing.T) {
	p := People{Name: "test"}
	data, _ := json.Marshal(&p)
	fmt.Println("data:", data)

	p2 := People{}
	json.Unmarshal(data, &p2)
	fmt.Println(p2)
}
