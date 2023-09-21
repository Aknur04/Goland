package TheExpert

type TheExpert struct {
	position string
	salary   int
	address  string
}

func NewExpert() TheExpert {
	return TheExpert{}
}

func (t *TheExpert) GetPosition() string {
	return t.position
}
func (t *TheExpert) GetSalary() int {
	return t.salary
}
func (t *TheExpert) GetAddress() string {
	return t.address
}

func (t *TheExpert) SetPosition(position string) {
	t.position = position
}

func (t *TheExpert) SetSalary(salary int) {
	t.salary = salary
}

func (t *TheExpert) SetAddress(address string) {
	t.address = address
}
