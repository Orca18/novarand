// Copyright (C) 2019-2022 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package compactcert

import (
	"testing"

	"github.com/Orca18/novarand/crypto"
	"github.com/Orca18/novarand/test/partitiontest"
)

func TestHashCoin(t *testing.T) {
	partitiontest.PartitionTest(t)

	var slots [32]uint64
	var sigcom = make(crypto.GenericDigest, HashSize)
	var partcom = make(crypto.GenericDigest, HashSize)
	var msgHash = make(crypto.GenericDigest, HashSize)

	crypto.RandBytes(sigcom[:])
	crypto.RandBytes(partcom[:])
	crypto.RandBytes(msgHash[:])

	for j := uint64(0); j < 1000; j++ {
		choice := coinChoice{
			J:            j,
			SignedWeight: uint64(len(slots)),
			ProvenWeight: uint64(len(slots)),
			Sigcom:       sigcom,
			Partcom:      partcom,
			MsgHash:      msgHash,
		}

		coin := hashCoin(choice)
		if coin >= uint64(len(slots)) {
			t.Errorf("hashCoin out of bounds")
		}

		slots[coin]++
	}

	for i, count := range slots {
		if count < 3 {
			t.Errorf("slot %d too low: %d", i, count)
		}
		if count > 100 {
			t.Errorf("slot %d too high: %d", i, count)
		}
	}
}

func BenchmarkHashCoin(b *testing.B) {
	var sigcom = make(crypto.GenericDigest, HashSize)
	var partcom = make(crypto.GenericDigest, HashSize)
	var msgHash = make(crypto.GenericDigest, HashSize)

	crypto.RandBytes(sigcom[:])
	crypto.RandBytes(partcom[:])
	crypto.RandBytes(msgHash[:])

	for i := 0; i < b.N; i++ {
		choice := coinChoice{
			J:            uint64(i),
			SignedWeight: 1024,
			ProvenWeight: 1024,
			Sigcom:       sigcom,
			Partcom:      partcom,
			MsgHash:      msgHash,
		}

		hashCoin(choice)
	}
}

func TestNumReveals(t *testing.T) {
	partitiontest.PartitionTest(t)

	billion := uint64(1000 * 1000 * 1000)
	micronova := uint64(1000 * 1000)
	provenWeight := 2 * billion * micronova
	secKQ := uint64(compactCertSecKQForTests)
	bound := uint64(1000)

	for i := uint64(3); i < 10; i++ {
		signedWeight := i * billion * micronova
		n, err := numReveals(signedWeight, provenWeight, secKQ, bound)
		if err != nil {
			t.Error(err)
		}

		if n < 50 || n > 300 {
			t.Errorf("numReveals(%d, %d, %d) = %d looks suspect",
				signedWeight, provenWeight, secKQ, n)
		}
	}
}

func BenchmarkNumReveals(b *testing.B) {
	billion := uint64(1000 * 1000 * 1000)
	micronova := uint64(1000 * 1000)
	provenWeight := 100 * billion * micronova
	signedWeight := 110 * billion * micronova
	secKQ := uint64(compactCertSecKQForTests)
	bound := uint64(1000)

	nr, err := numReveals(signedWeight, provenWeight, secKQ, bound)
	if nr < 900 {
		b.Errorf("numReveals(%d, %d, %d) = %d < 900", signedWeight, provenWeight, secKQ, nr)
	}

	for i := 0; i < b.N; i++ {
		_, err = numReveals(signedWeight, provenWeight, secKQ, bound)
		if err != nil {
			b.Error(err)
		}
	}
}
