package strings_test

import (
	"testing"

	"github.com/debeando/go-common/strings"
)

func TestEscape(t *testing.T) {
	expected := "<abc=\\'abc\\'>foo</abc>"
	result := strings.Escape("<abc='abc'>foo</abc>")

	if result != expected {
		t.Error("Expected: <abc=\\'abc\\'>foo</abc>")
	}
}
