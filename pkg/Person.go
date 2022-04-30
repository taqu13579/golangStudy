package pkg

import (
	"fmt"
	"strconv"
)

//　構造体
type Person struct {
	// 先頭大文字はパッケージ外からアクセス可、小文字はアクセス不可
	name string
	age  int
	Sex  string
}

func (p *Person) SetPerson(name string, age int, sex string) {
	p.name = name
	p.age = age
	p.Sex = sex
}

func (p *Person) GetPerson() string {
	return "name:" + p.name + ", age:" + strconv.Itoa(p.age) + ", sex:" + p.Sex
}

func (p *Person) Print() {
	fmt.Printf("name=%s, age=%d, sex=%s", p.name, p.age, p.Sex)
}
