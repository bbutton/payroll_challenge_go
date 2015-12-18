package main

// Employee is the base interface for things that know how to pay themselves
type Employee interface {
	CalculatePay() int64
}

// Developer defines the pay for a developer
type Developer struct {
}

// CalculatePay for a Developer
func (d *Developer) CalculatePay() int64 {
	return 1000
}

// QaTester defines pay for a qa tester
type QaTester struct {
}

// CalculatePay for a qa tester
func (qa *QaTester) CalculatePay() int64 {
	return 500
}

// Manager can pay all employees
type Manager struct {
	employees []*Employee
}

// NewManager creates a new manager
func NewManager() *Manager {
	manager := new(Manager)
	manager.employees = make([]*Employee, 10)

	return manager
}

// CalculatePay for managers and their reports
func (m *Manager) CalculatePay() int64 {
	var totalPay int64 = 300

	for _, e := range m.employees {
		totalPay = totalPay + (*e).CalculatePay()
	}

	return totalPay
}

// AddEmployee to list of Managers reports
func (m *Manager) AddEmployee(e Employee) {
	m.employees = append(m.employees, &e)
}
