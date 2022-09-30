package compactcert

// Code generated by github.com/algorand/msgp DO NOT EDIT.

import (
	"sort"

	"github.com/algorand/msgp/msgp"
)

// The following msgp objects are implemented in this file:
// Cert
//   |-----> (*) MarshalMsg
//   |-----> (*) CanMarshalMsg
//   |-----> (*) UnmarshalMsg
//   |-----> (*) CanUnmarshalMsg
//   |-----> (*) Msgsize
//   |-----> (*) MsgIsZero
//
// CompactOneTimeSignature
//            |-----> (*) MarshalMsg
//            |-----> (*) CanMarshalMsg
//            |-----> (*) UnmarshalMsg
//            |-----> (*) CanUnmarshalMsg
//            |-----> (*) Msgsize
//            |-----> (*) MsgIsZero
//
// Reveal
//    |-----> (*) MarshalMsg
//    |-----> (*) CanMarshalMsg
//    |-----> (*) UnmarshalMsg
//    |-----> (*) CanUnmarshalMsg
//    |-----> (*) Msgsize
//    |-----> (*) MsgIsZero
//
// coinChoice
//      |-----> (*) MarshalMsg
//      |-----> (*) CanMarshalMsg
//      |-----> (*) UnmarshalMsg
//      |-----> (*) CanUnmarshalMsg
//      |-----> (*) Msgsize
//      |-----> (*) MsgIsZero
//
// sigslotCommit
//       |-----> (*) MarshalMsg
//       |-----> (*) CanMarshalMsg
//       |-----> (*) UnmarshalMsg
//       |-----> (*) CanUnmarshalMsg
//       |-----> (*) Msgsize
//       |-----> (*) MsgIsZero
//

// MarshalMsg implements msgp.Marshaler
func (z *Cert) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0003Len := uint32(5)
	var zb0003Mask uint8 /* 6 bits */
	if (*z).PartProofs.MsgIsZero() {
		zb0003Len--
		zb0003Mask |= 0x1
	}
	if (*z).SigProofs.MsgIsZero() {
		zb0003Len--
		zb0003Mask |= 0x2
	}
	if (*z).SigCommit.MsgIsZero() {
		zb0003Len--
		zb0003Mask |= 0x8
	}
	if len((*z).Reveals) == 0 {
		zb0003Len--
		zb0003Mask |= 0x10
	}
	if (*z).SignedWeight == 0 {
		zb0003Len--
		zb0003Mask |= 0x20
	}
	// variable map header, size zb0003Len
	o = append(o, 0x80|uint8(zb0003Len))
	if zb0003Len != 0 {
		if (zb0003Mask & 0x1) == 0 { // if not empty
			// string "P"
			o = append(o, 0xa1, 0x50)
			o = (*z).PartProofs.MarshalMsg(o)
		}
		if (zb0003Mask & 0x2) == 0 { // if not empty
			// string "S"
			o = append(o, 0xa1, 0x53)
			o = (*z).SigProofs.MarshalMsg(o)
		}
		if (zb0003Mask & 0x8) == 0 { // if not empty
			// string "c"
			o = append(o, 0xa1, 0x63)
			o = (*z).SigCommit.MarshalMsg(o)
		}
		if (zb0003Mask & 0x10) == 0 { // if not empty
			// string "r"
			o = append(o, 0xa1, 0x72)
			if (*z).Reveals == nil {
				o = msgp.AppendNil(o)
			} else {
				o = msgp.AppendMapHeader(o, uint32(len((*z).Reveals)))
			}
			zb0001_keys := make([]uint64, 0, len((*z).Reveals))
			for zb0001 := range (*z).Reveals {
				zb0001_keys = append(zb0001_keys, zb0001)
			}
			sort.Sort(SortUint64(zb0001_keys))
			for _, zb0001 := range zb0001_keys {
				zb0002 := (*z).Reveals[zb0001]
				_ = zb0002
				o = msgp.AppendUint64(o, zb0001)
				o = zb0002.MarshalMsg(o)
			}
		}
		if (zb0003Mask & 0x20) == 0 { // if not empty
			// string "w"
			o = append(o, 0xa1, 0x77)
			o = msgp.AppendUint64(o, (*z).SignedWeight)
		}
	}
	return
}

func (_ *Cert) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*Cert)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Cert) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0003 int
	var zb0004 bool
	zb0003, zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0003, zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0003 > 0 {
			zb0003--
			bts, err = (*z).SigCommit.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "SigCommit")
				return
			}
		}
		if zb0003 > 0 {
			zb0003--
			(*z).SignedWeight, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "SignedWeight")
				return
			}
		}
		if zb0003 > 0 {
			zb0003--
			bts, err = (*z).SigProofs.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "SigProofs")
				return
			}
		}
		if zb0003 > 0 {
			zb0003--
			bts, err = (*z).PartProofs.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "PartProofs")
				return
			}
		}
		if zb0003 > 0 {
			zb0003--
			var zb0005 int
			var zb0006 bool
			zb0005, zb0006, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Reveals")
				return
			}
			if zb0005 > MaxReveals {
				err = msgp.ErrOverflow(uint64(zb0005), uint64(MaxReveals))
				err = msgp.WrapError(err, "struct-from-array", "Reveals")
				return
			}
			if zb0006 {
				(*z).Reveals = nil
			} else if (*z).Reveals == nil {
				(*z).Reveals = make(map[uint64]Reveal, zb0005)
			}
			for zb0005 > 0 {
				var zb0001 uint64
				var zb0002 Reveal
				zb0005--
				zb0001, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "Reveals")
					return
				}
				bts, err = zb0002.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "Reveals", zb0001)
					return
				}
				(*z).Reveals[zb0001] = zb0002
			}
		}
		if zb0003 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0003)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array")
				return
			}
		}
	} else {
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0004 {
			(*z) = Cert{}
		}
		for zb0003 > 0 {
			zb0003--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "c":
				bts, err = (*z).SigCommit.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "SigCommit")
					return
				}
			case "w":
				(*z).SignedWeight, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "SignedWeight")
					return
				}
			case "S":
				bts, err = (*z).SigProofs.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "SigProofs")
					return
				}
			case "P":
				bts, err = (*z).PartProofs.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "PartProofs")
					return
				}
			case "r":
				var zb0007 int
				var zb0008 bool
				zb0007, zb0008, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Reveals")
					return
				}
				if zb0007 > MaxReveals {
					err = msgp.ErrOverflow(uint64(zb0007), uint64(MaxReveals))
					err = msgp.WrapError(err, "Reveals")
					return
				}
				if zb0008 {
					(*z).Reveals = nil
				} else if (*z).Reveals == nil {
					(*z).Reveals = make(map[uint64]Reveal, zb0007)
				}
				for zb0007 > 0 {
					var zb0001 uint64
					var zb0002 Reveal
					zb0007--
					zb0001, bts, err = msgp.ReadUint64Bytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "Reveals")
						return
					}
					bts, err = zb0002.UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "Reveals", zb0001)
						return
					}
					(*z).Reveals[zb0001] = zb0002
				}
			default:
				err = msgp.ErrNoField(string(field))
				if err != nil {
					err = msgp.WrapError(err)
					return
				}
			}
		}
	}
	o = bts
	return
}

func (_ *Cert) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*Cert)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Cert) Msgsize() (s int) {
	s = 1 + 2 + (*z).SigCommit.Msgsize() + 2 + msgp.Uint64Size + 2 + (*z).SigProofs.Msgsize() + 2 + (*z).PartProofs.Msgsize() + 2 + msgp.MapHeaderSize
	if (*z).Reveals != nil {
		for zb0001, zb0002 := range (*z).Reveals {
			_ = zb0001
			_ = zb0002
			s += 0 + msgp.Uint64Size + zb0002.Msgsize()
		}
	}
	return
}

// MsgIsZero returns whether this is a zero value
func (z *Cert) MsgIsZero() bool {
	return ((*z).SigCommit.MsgIsZero()) && ((*z).SignedWeight == 0) && ((*z).SigProofs.MsgIsZero()) && ((*z).PartProofs.MsgIsZero()) && (len((*z).Reveals) == 0)
}

// MarshalMsg implements msgp.Marshaler
func (z *CompactOneTimeSignature) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(4)
	var zb0001Mask uint8 /* 6 bits */
	if (*z).Signature.MerkleArrayIndex == 0 {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	if (*z).Signature.Proof.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x8
	}
	if (*z).Signature.Signature.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x10
	}
	if (*z).Signature.VerifyingKey.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x20
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "idx"
			o = append(o, 0xa3, 0x69, 0x64, 0x78)
			o = msgp.AppendUint64(o, (*z).Signature.MerkleArrayIndex)
		}
		if (zb0001Mask & 0x8) == 0 { // if not empty
			// string "prf"
			o = append(o, 0xa3, 0x70, 0x72, 0x66)
			o = (*z).Signature.Proof.MarshalMsg(o)
		}
		if (zb0001Mask & 0x10) == 0 { // if not empty
			// string "sig"
			o = append(o, 0xa3, 0x73, 0x69, 0x67)
			o = (*z).Signature.Signature.MarshalMsg(o)
		}
		if (zb0001Mask & 0x20) == 0 { // if not empty
			// string "vkey"
			o = append(o, 0xa4, 0x76, 0x6b, 0x65, 0x79)
			o = (*z).Signature.VerifyingKey.MarshalMsg(o)
		}
	}
	return
}

func (_ *CompactOneTimeSignature) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*CompactOneTimeSignature)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CompactOneTimeSignature) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 int
	var zb0002 bool
	zb0001, zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0001, zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).Signature.Signature.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Signature")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			(*z).Signature.MerkleArrayIndex, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "MerkleArrayIndex")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).Signature.Proof.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Proof")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).Signature.VerifyingKey.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "VerifyingKey")
				return
			}
		}
		if zb0001 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0001)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array")
				return
			}
		}
	} else {
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0002 {
			(*z) = CompactOneTimeSignature{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "sig":
				bts, err = (*z).Signature.Signature.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Signature")
					return
				}
			case "idx":
				(*z).Signature.MerkleArrayIndex, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "MerkleArrayIndex")
					return
				}
			case "prf":
				bts, err = (*z).Signature.Proof.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Proof")
					return
				}
			case "vkey":
				bts, err = (*z).Signature.VerifyingKey.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "VerifyingKey")
					return
				}
			default:
				err = msgp.ErrNoField(string(field))
				if err != nil {
					err = msgp.WrapError(err)
					return
				}
			}
		}
	}
	o = bts
	return
}

func (_ *CompactOneTimeSignature) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*CompactOneTimeSignature)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CompactOneTimeSignature) Msgsize() (s int) {
	s = 1 + 4 + (*z).Signature.Signature.Msgsize() + 4 + msgp.Uint64Size + 4 + (*z).Signature.Proof.Msgsize() + 5 + (*z).Signature.VerifyingKey.Msgsize()
	return
}

// MsgIsZero returns whether this is a zero value
func (z *CompactOneTimeSignature) MsgIsZero() bool {
	return ((*z).Signature.Signature.MsgIsZero()) && ((*z).Signature.MerkleArrayIndex == 0) && ((*z).Signature.Proof.MsgIsZero()) && ((*z).Signature.VerifyingKey.MsgIsZero())
}

// MarshalMsg implements msgp.Marshaler
func (z *Reveal) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(2)
	var zb0001Mask uint8 /* 3 bits */
	if (*z).Part.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if ((*z).SigSlot.Sig.MsgIsZero()) && ((*z).SigSlot.L == 0) {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "p"
			o = append(o, 0xa1, 0x70)
			o = (*z).Part.MarshalMsg(o)
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "s"
			o = append(o, 0xa1, 0x73)
			// omitempty: check for empty values
			zb0002Len := uint32(2)
			var zb0002Mask uint8 /* 3 bits */
			if (*z).SigSlot.L == 0 {
				zb0002Len--
				zb0002Mask |= 0x2
			}
			if (*z).SigSlot.Sig.MsgIsZero() {
				zb0002Len--
				zb0002Mask |= 0x4
			}
			// variable map header, size zb0002Len
			o = append(o, 0x80|uint8(zb0002Len))
			if (zb0002Mask & 0x2) == 0 { // if not empty
				// string "l"
				o = append(o, 0xa1, 0x6c)
				o = msgp.AppendUint64(o, (*z).SigSlot.L)
			}
			if (zb0002Mask & 0x4) == 0 { // if not empty
				// string "s"
				o = append(o, 0xa1, 0x73)
				o = (*z).SigSlot.Sig.MarshalMsg(o)
			}
		}
	}
	return
}

func (_ *Reveal) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*Reveal)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Reveal) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 int
	var zb0002 bool
	zb0001, zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0001, zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0001 > 0 {
			zb0001--
			var zb0003 int
			var zb0004 bool
			zb0003, zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
			if _, ok := err.(msgp.TypeError); ok {
				zb0003, zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "SigSlot")
					return
				}
				if zb0003 > 0 {
					zb0003--
					bts, err = (*z).SigSlot.Sig.UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "SigSlot", "struct-from-array", "Sig")
						return
					}
				}
				if zb0003 > 0 {
					zb0003--
					(*z).SigSlot.L, bts, err = msgp.ReadUint64Bytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "SigSlot", "struct-from-array", "L")
						return
					}
				}
				if zb0003 > 0 {
					err = msgp.ErrTooManyArrayFields(zb0003)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "SigSlot", "struct-from-array")
						return
					}
				}
			} else {
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "SigSlot")
					return
				}
				if zb0004 {
					(*z).SigSlot = sigslotCommit{}
				}
				for zb0003 > 0 {
					zb0003--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "SigSlot")
						return
					}
					switch string(field) {
					case "s":
						bts, err = (*z).SigSlot.Sig.UnmarshalMsg(bts)
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "SigSlot", "Sig")
							return
						}
					case "l":
						(*z).SigSlot.L, bts, err = msgp.ReadUint64Bytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "SigSlot", "L")
							return
						}
					default:
						err = msgp.ErrNoField(string(field))
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "SigSlot")
							return
						}
					}
				}
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).Part.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Part")
				return
			}
		}
		if zb0001 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0001)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array")
				return
			}
		}
	} else {
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0002 {
			(*z) = Reveal{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "s":
				var zb0005 int
				var zb0006 bool
				zb0005, zb0006, bts, err = msgp.ReadMapHeaderBytes(bts)
				if _, ok := err.(msgp.TypeError); ok {
					zb0005, zb0006, bts, err = msgp.ReadArrayHeaderBytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "SigSlot")
						return
					}
					if zb0005 > 0 {
						zb0005--
						bts, err = (*z).SigSlot.Sig.UnmarshalMsg(bts)
						if err != nil {
							err = msgp.WrapError(err, "SigSlot", "struct-from-array", "Sig")
							return
						}
					}
					if zb0005 > 0 {
						zb0005--
						(*z).SigSlot.L, bts, err = msgp.ReadUint64Bytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "SigSlot", "struct-from-array", "L")
							return
						}
					}
					if zb0005 > 0 {
						err = msgp.ErrTooManyArrayFields(zb0005)
						if err != nil {
							err = msgp.WrapError(err, "SigSlot", "struct-from-array")
							return
						}
					}
				} else {
					if err != nil {
						err = msgp.WrapError(err, "SigSlot")
						return
					}
					if zb0006 {
						(*z).SigSlot = sigslotCommit{}
					}
					for zb0005 > 0 {
						zb0005--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							err = msgp.WrapError(err, "SigSlot")
							return
						}
						switch string(field) {
						case "s":
							bts, err = (*z).SigSlot.Sig.UnmarshalMsg(bts)
							if err != nil {
								err = msgp.WrapError(err, "SigSlot", "Sig")
								return
							}
						case "l":
							(*z).SigSlot.L, bts, err = msgp.ReadUint64Bytes(bts)
							if err != nil {
								err = msgp.WrapError(err, "SigSlot", "L")
								return
							}
						default:
							err = msgp.ErrNoField(string(field))
							if err != nil {
								err = msgp.WrapError(err, "SigSlot")
								return
							}
						}
					}
				}
			case "p":
				bts, err = (*z).Part.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Part")
					return
				}
			default:
				err = msgp.ErrNoField(string(field))
				if err != nil {
					err = msgp.WrapError(err)
					return
				}
			}
		}
	}
	o = bts
	return
}

func (_ *Reveal) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*Reveal)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Reveal) Msgsize() (s int) {
	s = 1 + 2 + 1 + 2 + (*z).SigSlot.Sig.Msgsize() + 2 + msgp.Uint64Size + 2 + (*z).Part.Msgsize()
	return
}

// MsgIsZero returns whether this is a zero value
func (z *Reveal) MsgIsZero() bool {
	return (((*z).SigSlot.Sig.MsgIsZero()) && ((*z).SigSlot.L == 0)) && ((*z).Part.MsgIsZero())
}

// MarshalMsg implements msgp.Marshaler
func (z *coinChoice) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(6)
	var zb0001Mask uint8 /* 7 bits */
	if (*z).J == 0 {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if (*z).MsgHash.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	if (*z).Partcom.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x8
	}
	if (*z).ProvenWeight == 0 {
		zb0001Len--
		zb0001Mask |= 0x10
	}
	if (*z).Sigcom.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x20
	}
	if (*z).SignedWeight == 0 {
		zb0001Len--
		zb0001Mask |= 0x40
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "j"
			o = append(o, 0xa1, 0x6a)
			o = msgp.AppendUint64(o, (*z).J)
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "msghash"
			o = append(o, 0xa7, 0x6d, 0x73, 0x67, 0x68, 0x61, 0x73, 0x68)
			o = (*z).MsgHash.MarshalMsg(o)
		}
		if (zb0001Mask & 0x8) == 0 { // if not empty
			// string "partcom"
			o = append(o, 0xa7, 0x70, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x6d)
			o = (*z).Partcom.MarshalMsg(o)
		}
		if (zb0001Mask & 0x10) == 0 { // if not empty
			// string "provenweight"
			o = append(o, 0xac, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74)
			o = msgp.AppendUint64(o, (*z).ProvenWeight)
		}
		if (zb0001Mask & 0x20) == 0 { // if not empty
			// string "sigcom"
			o = append(o, 0xa6, 0x73, 0x69, 0x67, 0x63, 0x6f, 0x6d)
			o = (*z).Sigcom.MarshalMsg(o)
		}
		if (zb0001Mask & 0x40) == 0 { // if not empty
			// string "sigweight"
			o = append(o, 0xa9, 0x73, 0x69, 0x67, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74)
			o = msgp.AppendUint64(o, (*z).SignedWeight)
		}
	}
	return
}

func (_ *coinChoice) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*coinChoice)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *coinChoice) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 int
	var zb0002 bool
	zb0001, zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0001, zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0001 > 0 {
			zb0001--
			(*z).J, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "J")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			(*z).SignedWeight, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "SignedWeight")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			(*z).ProvenWeight, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "ProvenWeight")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).Sigcom.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Sigcom")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).Partcom.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Partcom")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).MsgHash.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "MsgHash")
				return
			}
		}
		if zb0001 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0001)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array")
				return
			}
		}
	} else {
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0002 {
			(*z) = coinChoice{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "j":
				(*z).J, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "J")
					return
				}
			case "sigweight":
				(*z).SignedWeight, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "SignedWeight")
					return
				}
			case "provenweight":
				(*z).ProvenWeight, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "ProvenWeight")
					return
				}
			case "sigcom":
				bts, err = (*z).Sigcom.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Sigcom")
					return
				}
			case "partcom":
				bts, err = (*z).Partcom.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Partcom")
					return
				}
			case "msghash":
				bts, err = (*z).MsgHash.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "MsgHash")
					return
				}
			default:
				err = msgp.ErrNoField(string(field))
				if err != nil {
					err = msgp.WrapError(err)
					return
				}
			}
		}
	}
	o = bts
	return
}

func (_ *coinChoice) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*coinChoice)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *coinChoice) Msgsize() (s int) {
	s = 1 + 2 + msgp.Uint64Size + 10 + msgp.Uint64Size + 13 + msgp.Uint64Size + 7 + (*z).Sigcom.Msgsize() + 8 + (*z).Partcom.Msgsize() + 8 + (*z).MsgHash.Msgsize()
	return
}

// MsgIsZero returns whether this is a zero value
func (z *coinChoice) MsgIsZero() bool {
	return ((*z).J == 0) && ((*z).SignedWeight == 0) && ((*z).ProvenWeight == 0) && ((*z).Sigcom.MsgIsZero()) && ((*z).Partcom.MsgIsZero()) && ((*z).MsgHash.MsgIsZero())
}

// MarshalMsg implements msgp.Marshaler
func (z *sigslotCommit) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(2)
	var zb0001Mask uint8 /* 3 bits */
	if (*z).L == 0 {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if (*z).Sig.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "l"
			o = append(o, 0xa1, 0x6c)
			o = msgp.AppendUint64(o, (*z).L)
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "s"
			o = append(o, 0xa1, 0x73)
			o = (*z).Sig.MarshalMsg(o)
		}
	}
	return
}

func (_ *sigslotCommit) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*sigslotCommit)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *sigslotCommit) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 int
	var zb0002 bool
	zb0001, zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0001, zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).Sig.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Sig")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			(*z).L, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "L")
				return
			}
		}
		if zb0001 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0001)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array")
				return
			}
		}
	} else {
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0002 {
			(*z) = sigslotCommit{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "s":
				bts, err = (*z).Sig.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Sig")
					return
				}
			case "l":
				(*z).L, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "L")
					return
				}
			default:
				err = msgp.ErrNoField(string(field))
				if err != nil {
					err = msgp.WrapError(err)
					return
				}
			}
		}
	}
	o = bts
	return
}

func (_ *sigslotCommit) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*sigslotCommit)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *sigslotCommit) Msgsize() (s int) {
	s = 1 + 2 + (*z).Sig.Msgsize() + 2 + msgp.Uint64Size
	return
}

// MsgIsZero returns whether this is a zero value
func (z *sigslotCommit) MsgIsZero() bool {
	return ((*z).Sig.MsgIsZero()) && ((*z).L == 0)
}
