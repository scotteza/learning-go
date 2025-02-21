package fizzbuzz

import "testing"

func TestNotFizzOrBuzz(t *testing.T) {
	result := FizzBuzz(1)
	if result != "1" {
		t.Errorf("Expected 1, but got %v", result)
	}
}

func TestFizz(t *testing.T) {
	result := FizzBuzz(3)
	if result != "Fizz" {
		t.Errorf("Expected Fizz, but got %v", result)
	}
}

func TestBuzz(t *testing.T) {
	result := FizzBuzz(5)
	if result != "Buzz" {
		t.Errorf("Expected Buzz, but got %v", result)
	}
}

func TestFizzBuzz(t *testing.T) {
	result := FizzBuzz(15)
	if result != "FizzBuzz" {
		t.Errorf("Expected FizzBuzz, but got %v", result)
	}
}
