package calc

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	c := NewCalculator()

	tests := []struct {
		a, b     float64
		operator string
		want     float64
	}{
		{1, 2, "+", 3},
		{5, 3, "-", 2},
		{4, 2, "*", 8},
		{8, 2, "/", 4},
		{5, 0, "/", 0}, // Проверка деления на ноль
	}

	for _, test := range tests {
		got := c.Calculate(test.a, test.b, test.operator)
		if got != test.want {
			t.Errorf("Calculate(%v, %v, %q) = %v; want %v", test.a, test.b, test.operator, got, test.want)
		}
	}
}

func TestDivideByZero(t *testing.T) {
	c := NewCalculator()
	a, b := 5.0, 0.0
	got := c.Calculate(a, b, "/")

	if got != 0 { // Ожидаем, что результат равен 0 при делении на ноль
		t.Errorf("Calculate(%v, %v, \"/\") = %v; want %v", a, b, got, 0)
	}
}
