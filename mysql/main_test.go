package mysql_test

import (
	"testing"

	"github.com/debeando/go-common/mysql"
)

func TestParseValue(t *testing.T) {
	if value, ok := mysql.ParseNumberValue("yes"); !ok || value != 1 {
		t.Error("Expected: Found Parse and value = 1.")
	}

	if value, ok := mysql.ParseNumberValue("Yes"); !ok || value != 1 {
		t.Error("Expected: Found Parse and value = 1.")
	}

	if value, ok := mysql.ParseNumberValue("YES"); !ok || value != 1 {
		t.Error("Expected: Found Parse and value = 1.")
	}

	if value, ok := mysql.ParseNumberValue("no"); !ok || value != 0 {
		t.Error("Expected: Found Parse and value = 0.")
	}

	if value, ok := mysql.ParseNumberValue("No"); !ok || value != 0 {
		t.Error("Expected: Found Parse and value = 0.")
	}

	if value, ok := mysql.ParseNumberValue("NO"); !ok || value != 0 {
		t.Error("Expected: Found Parse and value = 0.")
	}

	if value, ok := mysql.ParseNumberValue("ON"); !ok || value != 1 {
		t.Error("Expected: Found Parse and value = 1.")
	}

	if value, ok := mysql.ParseNumberValue("OFF"); !ok || value != 0 {
		t.Error("Expected: Found Parse and value = 0.")
	}

	if value, ok := mysql.ParseNumberValue("true"); ok && value == 0 {
		t.Error("Expected: Imposible Parse.")
	}

	if value, ok := mysql.ParseNumberValue("1234567890"); !ok || value != 1234567890 {
		t.Error("Expected: Found Parse and value = 1234567890.")
	}
}

func TestClearUser(t *testing.T) {
	user := "test[test] @ [127.0.0.1]"
	expected := "test"
	result := mysql.ClearUser(user)

	if result != expected {
		t.Errorf("Expected: '%s', got: '%s'.", expected, result)
	}
}
