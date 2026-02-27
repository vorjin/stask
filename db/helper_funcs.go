// Package db contains with operations with db and other numeric functions
package db

import (
	"encoding/binary"
)

func uToB(id uint64) []byte {
	idBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(idBytes, id)
	return idBytes
}
