package main

import "fmt"

type IEncoder interface {
	deCode(value string)
	enCode(value string)
}
type XEncoder struct {
}
type EEncoder struct {
}

func (xEncoder XEncoder) enCode(value string) {
	fmt.Println("XEncoder ile encode edildi")
}

func (xEncoder XEncoder) deCode(value string) {
	fmt.Println("XEncoder ile decode edildi")
}
func (eEncoder EEncoder) enCode(value string) {
	fmt.Println("EEncoder ile encode edildi")
}

func (eEncoder EEncoder) deCode(value string) {
	fmt.Println("EEncoder ile decode edildi")
}
func main() {
	var xEncoder *XEncoder
	xEncoder = &XEncoder{}
	xEncoder.enCode("123456")
	xEncoder.deCode("jnoifjmw1234")

	var eEncoder EEncoder
	eEncoder.enCode("1111111")
	eEncoder.deCode("fnkvmafpowk")

	var iEncoder IEncoder
	iEncoder = EEncoder{}
	iEncoder.deCode("iwejiocw22")
	iEncoder.enCode("1111111")
}
