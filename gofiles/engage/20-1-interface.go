package main

import "fmt"

// defind how to make money and cotrol body of human
type human interface {
	money() int
	body() int
}

type organize struct {
	hand string
	mind string
}

type mind struct {
	iminds int
	level  int
	exe    int
}

var woman, man human

func (or organize) money() int {
	return len(or.hand)*3 + len(or.mind)*400
}
func (in1 mind) body() int {
	in1.iminds = in1.iminds * in1.level
	in1.exe = in1.exe + in1.iminds
	return in1.exe
}
func (or1 organize) body() int {
	return len(or1.hand)*3 + len(or1.mind)*300
}
func (in mind) money() int {
	in.iminds = in.iminds * in.level
	in.exe = in.exe * in.iminds * 2
	//var in1 int
	//in1 = in.exe

	return in.exe
}
func NB(h human) {
	nnb := h.body() + h.money()
	fmt.Println(h)
	fmt.Println("finally score", nnb)
}
func main() {
	fmt.Println("zhanshi interface")
	org := organize{"two hand", "goood man"}
	sd := mind{12, 12, 23}

	NB(org)
	NB(sd)

}
