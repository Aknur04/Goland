package MoneyHunter

type MoneyHunter struct {
	position string
	salary   int
	address  string
}

func NewMoneyHunter() MoneyHunter {
	return MoneyHunter{}
}

func (h *MoneyHunter) GetPosition() string {
	return h.position
}
func (h *MoneyHunter) GetSalary() int {
	return h.salary
}
func (h *MoneyHunter) GetAddress() string {
	return h.address
}

func (h *MoneyHunter) SetPosition(position string) {
	h.position = position
}

func (h *MoneyHunter) SetSalary(salary int) {
	h.salary = salary
}

func (h *MoneyHunter) SetAddress(address string) {
	h.address = address
}
