package Employee

type Employee interface {
	GetPosition()
	SetPosition(position string)
	GetSalary()
	SetSalary(salary int)
	GetAddress()
	SetAddress(address string)
}
