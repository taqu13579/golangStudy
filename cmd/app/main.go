package main

import (
	"fmt"
	"golangStudy/pkg"
	"time"
)

func main() {
	// 変数宣言
	var int1 int = 123
	var int2 = 456
	int3 := 789
	var str1 string = "aaa"
	var str2 = "bbb"
	str3 := "ccc"
	var (
		str4 string = "ddd"
		str5 string = "eee"
	)
	// 定数
	const foo = 100

	// 文字列出力
	fmt.Print("int1=", int1, "int2=", int2, "int3=", int3, "str1=", str1, "str2=", str2, "str3=", str3, "str4=", str4, "str5=", str5)
	fmt.Println("int1=", int1, "int2=", int2, "int3=", int3, "str1=", str1, "str2=", str2, "str3=", str3, "str4=", str4, "str5=", str5)
	fmt.Printf("int1=%d, int2=%d, int3=%d, str1=%s, str2=%s, str3=%s, str4=%s, str5=%s", int1, int2, int3, str1, str2, str3, str4, str5)

	// 型に別名を付与
	type UtcTime string
	type JstTime string
	// var t1 UtcTime = "00:00:00"
	// var t2 JstTime = "09:00:00"
	// 別の型のため、代入不可
	// t1 = t2

	// 配列
	// パターン1
	var arr1 [2]string
	arr1[0] = "go lang"
	arr1[1] = "java "
	fmt.Println(arr1[0], arr1[1])
	// パターン2
	var arr2 [2]string = [2]string{"go lang", "java"}
	fmt.Println(arr2[0], arr2[1])
	fmt.Println(arr2)
	// パターン3
	arr3 := [...]string{"go lang", "java"}
	fmt.Println(arr3[0], arr3[1])
	fmt.Println(arr3)

	// スライス（可変長配列）
	// パターン1
	slice1 := []string{"go lang", "java"}
	fmt.Println(slice1)

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("%d = %s", i, slice1[i])
	}

	// パターン2
	slice2 := [...]string{"go lang", "java"}
	fmt.Println(slice2)

	// ループ
	for i, v := range slice2 {
		fmt.Printf("%d = %s", i, v)
	}

	// マップ
	map1 := map[string]int{
		"x": 100,
		"y": 100,
	}
	fmt.Println(map1, map1["x"], map1["y"])
	// 要素追加
	map1["z"] = 300
	fmt.Println(map1)
	// 要素削除
	delete(map1, "x")
	fmt.Println(map1)
	// 要素数取得
	fmt.Println(len(map1))
	// mapの要素検証
	_, existZ := map1["z"]
	if existZ {
		fmt.Println("Exist")
	} else {
		fmt.Println("Not Exist")
	}
	// ループ
	for key, value := range map1 {
		fmt.Printf("%s=%d ", key, value)
	}

	// 構造体
	var p1 pkg.Person
	p1.SetPerson("alice", 25, "male")
	p1Info := p1.GetPerson()
	fmt.Println(p1Info)
	// interface
	book := pkg.Book{"booktitle01", 1100}
	var printable pkg.Printable
	printable = book
	print(printable)

	// 変更処理
	go funcA()
	for i := 0; i < 10; i++ {
		fmt.Print("M")
		time.Sleep(time.Second)
	}

}

func print(p pkg.Printable) {
	p.Print()
}

func funcA() {
	for i := 0; i < 10; i++ {
		fmt.Print("A")
		time.Sleep(time.Second)
	}
}
