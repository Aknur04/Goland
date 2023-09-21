package Security

type Security struct {
	position string
	salary   int
	address  string
}

func NewSecurity() Security {
	return Security{}
}

func (s *Security) GetPosition() string {
	return s.position
}
func (s *Security) GetSalary() int {
	return s.salary
}
func (s *Security) GetAddress() string {
	return s.address
}

func (s *Security) SetPosition(position string) {
	s.position = position
}

func (s *Security) SetSalary(salary int) {
	s.salary = salary
}

func (s *Security) SetAddress(address string) {
	s.address = address
}
