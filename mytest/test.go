package main

import "fmt"

func main()  {
	a := 10
	fmt.Printf("变量的地址:%x\n", &a)

	var b int = 20
	var ip *int
	var ip2 *int

	ip = &b
	fmt.Printf("b 变量的地址:%x\n", &b)

	fmt.Printf("ip 变量的指针地址:%x\n", ip)

	fmt.Printf("ip 变量的值:%d\n", *ip)

	fmt.Printf("ptr 的值为:%x", ip2)
}
