package gravatar

import (
	"testing"
)

func TestHash(t *testing.T) {
	if "0bc83cb571cd1c50ba6f3e8a78ef1346" != Hash("MyEmailAddress@example.com ") {
		t.Errorf("incorrect hash")
	}
}
