package gluacrypto_crypto

import (
	"crypto/md5"
	"encoding/hex"

	lua "github.com/yuin/gopher-lua"
)

func md5Fn(L *lua.LState) int {
	h := md5.New()
	s := lua.LVAsString(L.Get(1))
	raw := lua.LVAsBool(L.Get(2))
	_, err := h.Write([]byte(s))
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	var result string
	if !raw {
		result = hex.EncodeToString(h.Sum(nil))
	} else {
		result = string(h.Sum(nil))
	}
	L.Push(lua.LString(result))
	return 1
}
