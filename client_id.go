package go_librespot

import "encoding/hex"

var ClientId = []byte{0x0, 0xdd, 0x5c, 0xb4, 0x48, 0xe5, 0x4b, 0x23, 0x90, 0xa9, 0xf8, 0xc5, 0xab, 0x79, 0xf2, 0x88}
var ClientIdHex = hex.EncodeToString(ClientId)
