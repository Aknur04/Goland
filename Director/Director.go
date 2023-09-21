package Director

type Director struct {
	position string
	salary   int
	address  string
}

func NewDirector() Director {
	return Director{}
}

func (d *Director) GetPosition() string {
	return d.position
}
func (d *Director) GetSalary() int {
	return d.salary
}
func (d *Director) GetAddress() string {
	return d.address
}

func (d *Director) SetPosition(position string) {
	d.position = position
}

func (d *Director) SetSalary(salary int) {
	d.salary = salary
}

func (d *Director) SetAddress(address string) {
	d.address = address
}
