package main

import (
  "fmt"
//  "os"
)
func main() {
	//if len(os.Args) != 2 {
		//fmt.Printf("first arg %s \n", os.Args[0])
	//	os.Exit(1)
	//}
	//ifmt.Printf("It's over", os.Args[1])
	//_, exists := power("Hankin")
	//if exists == false {
	//	print("handle this error case")
	//}
	//goku := Saiyan{}
	//goku.Name = "Hankin"
	//goku.Power = 9000
	//print(goku.Name, goku.Power)
	goku := Saiyan{"Power", 9000}
	super(goku)
	fmt.Println(goku.Power)
	goku2 := &Saiyan{"Power", 9000}
	super2(goku2)
	fmt.Println(goku2.Power)
	goku3 := &Saiyan{"Power", 9001}
	goku3.super3()
	fmt.Println(goku3.Power)
	goku4 := newSaiyan("Hankin", 9000)
	fmt.Println(goku4.Power)
	gohan := &Saiyan2{
		Name: "Gohan",
		Power: 1000,
		Father: &Saiyan2{
			Name: "Goku",
			Power: 9001,
			Father: nil,
		},
	}
	fmt.Println(gohan.Name, gohan.Power)
	goku5 := new(Saiyan)
	goku5.Name = "goku5"
	goku5.Power = 9001
	goku6 := &Saiyan{
		Name: "goku6",
		Power: 9001,
	}
	fmt.Println(goku5.Name, goku5.Power)
	fmt.Println(goku6.Name, goku6.Power)
}

func log(message string) {
	println(message)
}

func add(a int, b int) int {
	return a+b
}

func power(name string) (int, bool) {
	value := 1000
	exists := false
	return value, exists
}

type Saiyan struct {
	Name	string
	Power	int
}

type Saiyan2 struct {
	Name	string
	Power	int
	Father	*Saiyan2
}

func super(s Saiyan){
    s.Power += 10000
}
func super2(s *Saiyan){
    s.Power += 10000
}

func (s *Saiyan)super3(){
	s.Power += 10000
}

func newSaiyan(name string, power int) Saiyan {
	return Saiyan{
		Name:name,
		Power:power,
	}
}
