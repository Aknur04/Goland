package main

import (
	"Assignment1/Director"
	"Assignment1/MoneyHunter"
	"Assignment1/Security"
	"Assignment1/TheExpert"
	_ "Assignment1/TheExpert"
	"Assignment1/managerr"
	"fmt"
)

func main() {
	sec := Security.NewSecurity()
	sec.SetPosition("security")
	sec.SetSalary(600)
	sec.SetAddress("Arena 5/2")
	fmt.Println(sec.GetPosition(), sec.GetSalary(), sec.GetAddress())

	man := managerr.NewManager()
	man.SetPosition("manager")
	man.SetSalary(500)
	man.SetAddress("Abay 53")
	fmt.Println(man.GetPosition(), man.GetSalary(), man.GetAddress())

	dir := Director.NewDirector()
	dir.SetPosition("director")
	dir.SetSalary(700)
	dir.SetAddress("Bayzakov 185/125")
	fmt.Println(dir.GetPosition(), dir.GetSalary(), dir.GetAddress())

	exp := TheExpert.NewExpert()
	exp.SetPosition("expert")
	exp.SetSalary(900)
	exp.SetAddress("abaya/rozybakieva")
	fmt.Println(exp.GetPosition(), exp.GetSalary(), exp.GetAddress())

	mon := MoneyHunter.NewMoneyHunter()
	mon.SetPosition("money hunter")
	mon.SetSalary(10000)
	mon.SetAddress("Tole bi 59")
	fmt.Println(mon.GetPosition(), mon.GetSalary(), mon.GetAddress())
}
