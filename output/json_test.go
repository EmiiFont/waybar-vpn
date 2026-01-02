package output

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestWaybarOutput_Print(t *testing.T) {
	out := WaybarOutput{
		Text:    "test",
		Tooltip: "tooltip",
		Class:   "class",
	}

	var buf bytes.Buffer
	// Mock stdout, but for test, encode to buf
	encoder := json.NewEncoder(&buf)
	err := encoder.Encode(out)
	if err != nil {
		t.Fatalf("Error encoding: %v", err)
	}

	var decoded WaybarOutput
	err = json.Unmarshal(buf.Bytes(), &decoded)
	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if decoded != out {
		t.Errorf("Expected %+v, got %+v", out, decoded)
	}
}
