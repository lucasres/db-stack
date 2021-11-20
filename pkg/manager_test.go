package pkg

import (
	"testing"
)

func TestManager(t *testing.T) {
	m := NewManager()

	k1 := "key1"
	k2 := "key2"

	k1V1 := "teste de valor um"
	k2V1 := "teste de valor dois"
	k1V2 := "mudei de novo kkkj"
	k1V3 := "novo valor"

	m.Set(k1, k1V1)
	m.Set(k2, k2V1)

	m.Push(k1, k1V2)
	m.Push(k1, k1V3)

	// simulate get
	current, err := m.Get(k1)
	if err != nil {
		t.Fail()
	}
	if current != k1V3 {
		t.Fail()
	}

	current, err = m.Get(k1)
	if err != nil {
		t.Fail()
	}
	if current != k1V3 {
		t.Fail()
	}

	current, err = m.Get(k1)
	if err != nil {
		t.Fail()
	}
	if current != k1V3 {
		t.Fail()
	}

	m.Pop(k1)

	current, err = m.Get(k1)
	if err != nil {
		t.Fail()
	}
	if current != k1V2 {
		t.Fail()
	}

}
