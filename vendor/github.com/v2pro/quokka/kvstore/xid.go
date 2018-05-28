package kvstore

import "github.com/rs/xid"

var zoneCharToId = map[byte]uint64 {}
var zoneIdToChar = map[uint64]string {}

func init() {
	for i, b := range "0123456789abcdefghijklmnopqrstuv" {
		zoneCharToId[byte(b)] = uint64(i)
		zoneIdToChar[uint64(i)] = string([]byte{byte(b)})
	}
}

// id is 21 bytes
// zoneId should be within [0, 31]
// if one entity type has many entities, we can assign multiple zone to expand capacity
// if many entity types only has small number of entities, we can compact them in same zone
func NewID(zoneId int) string {
	return xid.New().String() + zoneIdToChar[uint64(zoneId)]
}

func lookupZone(b byte) uint64 {
	return zoneCharToId[b] // will default to zone 0
}