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

package protocol

// HashID is a domain separation prefix for an object type that might be hashed
// This ensures, for example, the hash of a transaction will never collide with the hash of a vote
// HashID는 해시될 수 있는 객체 유형에 대한 도메인 분리 접두사입니다.
// 이것은 예를 들어 트랜잭션의 해시가 투표의 해시와 충돌하지 않도록 합니다.
type HashID string

// Hash IDs for specific object types, in lexicographic order.
// Hash IDs must be PREFIX-FREE (no hash ID is a prefix of another).
// 사전순으로 특정 객체 유형에 대한 해시 ID.
// 해시 ID는 PREFIX-FREE여야 합니다(해시 ID가 다른 해시 ID의 접두어가 아님).
const (
	AppIndex HashID = "appID"

	// ARCReserved is used to reserve prefixes starting with `arc` to
	// ARCs-related hashes https://github.com/algorandfoundation/ARCs
	// The prefix for ARC-XXXX should start with:
	// "arcXXXX" (where "XXXX" is the 0-padded number of the ARC)
	// For example ARC-0003 can use any prefix starting with "arc0003"
	// ===================================================================
	// ARCReserved는 'arc' 로 시작하는 접두사를 ARC 관련  해시에 예약하는 데 사용
	// ARC 관련 해시 https://github.com/algorandfoundation/ARCs
	// ARC-XXXX의 접두사는 다음으로 시작해야 합니다.
	// "arcXXXX"(여기서 "XXXX"는 ARC의 0으로 채워진 숫자)
	// 예를 들어 ARC-0003은 "arc0003"으로 시작하는 모든 접두사를 사용할 수 있습니다.
	ARCReserved HashID = "arc"

	AuctionBid        HashID = "aB"
	AuctionDeposit    HashID = "aD"
	AuctionOutcomes   HashID = "aO"
	AuctionParams     HashID = "aP"
	AuctionSettlement HashID = "aS"

	CompactCertCoin HashID = "ccc"
	CompactCertPart HashID = "ccp"
	CompactCertSig  HashID = "ccs"

	AgreementSelector                HashID = "AS"
	BlockHeader                      HashID = "BH"
	BalanceRecord                    HashID = "BR"
	Credential                       HashID = "CR"
	Genesis                          HashID = "GE"
	KeysInMSS                        HashID = "KP"
	MerkleArrayNode                  HashID = "MA"
	MerkleVectorCommitmentBottomLeaf HashID = "MB"
	Message                          HashID = "MX"
	NetPrioResponse                  HashID = "NPR"
	OneTimeSigKey1                   HashID = "OT1"
	OneTimeSigKey2                   HashID = "OT2"
	PaysetFlat                       HashID = "PF"
	Payload                          HashID = "PL"
	Program                          HashID = "Program"
	ProgramData                      HashID = "ProgData"
	ProposerSeed                     HashID = "PS"
	ParticipationKeys                HashID = "PK"
	Seed                             HashID = "SD"
	SpecialAddr                      HashID = "SpecialAddr"
	SignedTxnInBlock                 HashID = "STIB"
	TestHashable                     HashID = "TE"
	TxGroup                          HashID = "TG"
	TxnMerkleLeaf                    HashID = "TL"
	Transaction                      HashID = "TX"
	Vote                             HashID = "VO"
)
