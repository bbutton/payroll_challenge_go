package Payroll

import (
	"strconv"
	"testing"
)

func Test_DeveloperPaidCorrecly(t *testing.T) {
	var d Developer
	if d.CalculatePay() != 1000 {
		t.Fatalf("Not paid correctly")
	}
}

func Test_QaTesterPaidCorrectly(t *testing.T) {
	var qa QaTester
	if qa.CalculatePay() != 500 {
		t.Fatalf("Not paid correctly")
	}
}

func Test_ManagerWithNoEmployeesPaidCorrectly(t *testing.T) {
	var m Manager
	if m.CalculatePay() != 300 {
		t.Fatal("Single manager not paid correctly")
	}
}

func Test_ManagerWithSingleEmployeePaidCorrectly(t *testing.T) {
	var m Manager
	var d Developer
	m.AddEmployee(&d)

	pay := m.CalculatePay()

	if pay != 1300 {
		t.Fatalf("Should have been 1300 but was " + strconv.FormatInt(pay, 10))
	}
}
