package log

import (
	"bytes"
	"encoding/json"
	"testing"
)

func checkMarshaling(t *testing.T, v Level, expect []byte) {
	data, err := json.Marshal(v)
	if err != nil {
		t.Errorf("Can't marshal %q: %v", v, err)
	}

	if !bytes.Equal(data, expect) {
		t.Errorf("Marshal error, got %q, expected %q", data, expect)
	}
}

func checkUnmarshaling(t *testing.T, data []byte, expect Level) {
	var v Level
	err := json.Unmarshal(data, &v)
	if err != nil {
		t.Errorf("Can't unmarshal %q: %v", data, err)
	}

	if v != expect {
		t.Errorf("Unmarshal error, got %q, expected %q", v, expect)
	}
}

func TestJsonMarshaling(t *testing.T) {
	checkMarshaling(t, INFO, []byte("\"info\""))
	checkMarshaling(t, ERROR, []byte("\"error\""))
	checkMarshaling(t, FATAL, []byte("\"fatal\""))
	checkMarshaling(t, DEBUG, []byte("\"debug\""))
}

func TestJsonUnmarshaling(t *testing.T) {
	checkUnmarshaling(t, []byte("\"info\""), INFO)
	checkUnmarshaling(t, []byte("\"error\""), ERROR)
	checkUnmarshaling(t, []byte("\"vdebug\""), VDEBUG)
	checkUnmarshaling(t, []byte("\"vvdebug\""), VVDEBUG)
}

func TestComplexJson(t *testing.T) {
	type MyType struct {
		L1, L2, L3, L4, L5, L6, L7, L8 Level
	}

	v := MyType{VVVDEBUG, VVDEBUG, VDEBUG, DEBUG, INFO, WARNING, ERROR, FATAL}

	data, err := json.Marshal(&v)
	if err != nil {
		t.Fatalf("can't marshal %q: %v", v, err)
	}

	var v2 MyType
	if err = json.Unmarshal(data, &v2); err != nil {
		t.Fatalf("can't unmarshal %q: %v", data, err)
	}

	if v != v2 {
		t.Errorf("%v != %v", v, v2)
	}
}
