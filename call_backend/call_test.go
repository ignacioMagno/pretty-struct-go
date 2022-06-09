package callbackend

import "testing"

func TestCall(t *testing.T) {
	SendStruct(8090, "hola como están")
}
