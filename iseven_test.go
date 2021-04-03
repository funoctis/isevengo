package isevengo

import (
	"testing"
)

func TestIsEven(t *testing.T) {
	result, err := IsEven(10)
	if err != nil {
		t.Errorf("Error Occured: %t\nisEven: %t", err, result)
	}
	if result != true {
		t.Errorf("IsEven(10): Got false, should be true")
	}

	result, err = IsEven(69)
	if err != nil {
		t.Errorf("Error Occured: %t\nisEven: %t", err, result)
	}
	if result != false {
		t.Errorf("IsEven(69): Got true, should be false")
	}

	result, err = IsEven(-1)
	if err == nil {
		t.Errorf("IsEven(-1): Should have received number out of range error")
	}

	result, err = IsEven(9999991)
	if err == nil {
		t.Errorf("IsEven(9999991): Should have received number out of range error")
	}
}

func TestCallIsEvenAPI(t *testing.T) {
	result, err := CallIsEvenAPI(10)
	if err != nil {
		t.Error("Error Occured:\n", err, "\nisEven:\n", result)
	}
}

func TestInvalidNumber(t *testing.T) {
	result, err := CallIsEvenAPI(-1)
	if err != nil {
		t.Error("Error Occured:\n", err, "\nisEven:\n", result)
	}
	if len(result.Error) == 0 {
		t.Error("Error occured on invalid number -1")
	}

	result, err = CallIsEvenAPI(9999991)
	if err != nil {
		t.Error("Error Occured:\n", err, "\nisEven:\n", result)
	}
	if len(result.Error) == 0 {
		t.Error("Error occured on invalid number -1")
	}
}
