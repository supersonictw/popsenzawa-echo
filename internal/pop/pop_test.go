package pop

import (
	"bytes"
	"testing"
)

func TestNewPopJSON(t *testing.T) {
	p := NewPop(800, "1.1.1.1", "US")
	if bytes.Compare(p.JSON(), []byte("{\"count\":800,\"address\":\"1.1.1.1\",\"region\":\"US\"}")) != 0 {
		t.Fatalf("unexpected json format:%s", p.JSON())
	}
}
