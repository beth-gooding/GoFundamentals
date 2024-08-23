package payroll

type Manager struct {
	Individual     Employee
	Salary         float64
	CommissionRate float64
}

func (m Manager) Pay() (string, float64) {
	yearlyPay := m.Salary * (1 + m.CommissionRate)
	name := m.Individual.FirstName + " " + m.Individual.LastName

	return name, yearlyPay
}
