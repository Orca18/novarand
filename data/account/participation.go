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

package account

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Orca18/novarand/config"
	"github.com/Orca18/novarand/crypto"
	"github.com/Orca18/novarand/crypto/merklesignature"
	"github.com/Orca18/novarand/data/basics"
	"github.com/Orca18/novarand/data/transactions"
	"github.com/Orca18/novarand/logging"
	"github.com/Orca18/novarand/protocol"
	"github.com/Orca18/novarand/util/db"
)

// A Participation encapsulates a set of secrets which allows a root to participate in consensus.
// All such accounts are associated with a parent root account via the Address
// (although this parent account may not be resident on this machine).
// Participations are allowed to vote on a user's behalf for some range of rounds.
// After this range, all remaining secrets are destroyed.
// For correctness, all Roots should have no more than one Participation globally active at any time.
// If this condition is violated, the Root may equivocate. (Algorand tolerates a limited fraction of misbehaving accounts.)
//msgp:ignore Participation
/*
Participation은 root가 합의에 참여할 수 있게 하는 비밀 집합(여러종류의 비밀들)을 캡슐화한다.
모든 계정은 주소를 통해 root계정과 연결되어 있다.(부모 계정(root)은 머신상에 존재하지 않는다? 논리적으로만 존재한다는건가?
avm에 올라가지 않는다는건가? 이 상태머신(흠.. 어떤 상태머신이지?)에는 존재하지 않는다. 참여계정만 존재하는건가?)
Participation은 유저가 특정 라운드만큼 투표할 수 있는 권한이다. 모든 라운드가 끝나면 남아있는 비밀(sectret)은 모두 파괴된다.
정확성을 위해 모든 Root는 한번에 하나의 활성화된 참여만 가질 수 있다.
이것이 위배되면 root는 모호해진다(알고랜드는 약각의 오작동은 허용) <= equivocate뜻이 모호해진다여서 이렇게 썼는데 뭐가 모호해진다는건지 모르겠네..
(msgp가 Paticipation을 무시한다? 즉, 동작하지 않는다는건가? 흠.. 모호하다고 하긴했는데 동작을 안하는게 맞겠지?)
*/
type Participation struct {
	Parent basics.Address

	VRF    *crypto.VRFSecrets
	Voting *crypto.OneTimeSignatureSecrets
	// StateProofSecrets is used to sign compact certificates.
	/*
		StateProofSecrets는 컴팩트 인증서에 서명할 때 사용됨
	*/
	StateProofSecrets *merklesignature.Secrets

	// The first and last rounds for which this account is valid, respectively.
	//
	// When lastValid has concluded, this set of secrets is destroyed.
	FirstValid basics.Round
	LastValid  basics.Round

	KeyDilution uint64
}

// ParticipationKeyIdentity is for msgpack encoding the participation data.
// ParticipationKeyIdentity는 participation을 msgpack인코딩한 데이터이다. .
type ParticipationKeyIdentity struct {
	_struct struct{} `codec:",omitempty,omitemptyarray"`

	Parent      basics.Address                  `codec:"addr"`
	VRFSK       crypto.VrfPrivkey               `codec:"vrfsk"`
	VoteID      crypto.OneTimeSignatureVerifier `codec:"vote-id"`
	FirstValid  basics.Round                    `codec:"fv"`
	LastValid   basics.Round                    `codec:"lv"`
	KeyDilution uint64                          `codec:"kd"`
}

// ToBeHashed implements the Hashable interface.
// => Hashable인터페이스를 구현했다.
func (id *ParticipationKeyIdentity) ToBeHashed() (protocol.HashID, []byte) {
	return protocol.ParticipationKeys, protocol.Encode(id)
}

// ID creates a ParticipationID hash from the identity file.
// => ID는 identity file로부터 ParticipationID hash를 생성한다.
func (id ParticipationKeyIdentity) ID() ParticipationID {
	return ParticipationID(crypto.HashObj(&id))
}

// ID computes a ParticipationID.
func (part Participation) ID() ParticipationID {
	idData := ParticipationKeyIdentity{
		Parent:      part.Parent,
		FirstValid:  part.FirstValid,
		LastValid:   part.LastValid,
		KeyDilution: part.KeyDilution,
	}
	if part.VRF != nil {
		copy(idData.VRFSK[:], part.VRF.SK[:])
	}
	if part.Voting != nil {
		copy(idData.VoteID[:], part.Voting.OneTimeSignatureVerifier[:])
	}

	return idData.ID()
}

// PersistedParticipation encapsulates the static state of the participation
// for a single address at any given moment, while providing the ability
// to handle persistence and deletion of secrets.
//msgp:ignore PersistedParticipation
type PersistedParticipation struct {
	Participation

	Store db.Accessor
}

// ValidInterval returns the first and last rounds for which this participation account is valid.
func (part Participation) ValidInterval() (first, last basics.Round) {
	return part.FirstValid, part.LastValid
}

// Address returns the root account under which this participation account is registered.
func (part Participation) Address() basics.Address {
	return part.Parent
}

// OverlapsInterval returns true if the partkey is valid at all within the range of rounds (inclusive)
func (part Participation) OverlapsInterval(first, last basics.Round) bool {
	if last < first {
		logging.Base().Panicf("Round interval should be ordered (first = %v, last = %v)", first, last)
	}
	if last < part.FirstValid || first > part.LastValid {
		return false
	}
	return true
}

// VRFSecrets returns the VRF secrets associated with this Participation account.
func (part Participation) VRFSecrets() *crypto.VRFSecrets {
	return part.VRF
}

// VotingSecrets returns the voting secrets associated with this Participation account.
func (part Participation) VotingSecrets() *crypto.OneTimeSignatureSecrets {
	return part.Voting
}

// StateProofSigner returns the key used to sign on Compact Certificates.
// might return nil!
func (part Participation) StateProofSigner() *merklesignature.Secrets {
	return part.StateProofSecrets
}

// StateProofVerifier returns the verifier for the StateProof keys.
func (part Participation) StateProofVerifier() *merklesignature.Verifier {
	return part.StateProofSecrets.GetVerifier()
}

// GenerateRegistrationTransaction returns a transaction object for registering a Participation with its parent.
func (part Participation) GenerateRegistrationTransaction(fee basics.MicroNovas, txnFirstValid, txnLastValid basics.Round, leaseBytes [32]byte, includeStateProofKeys bool) transactions.Transaction {
	t := transactions.Transaction{
		Type: protocol.KeyRegistrationTx,
		Header: transactions.Header{
			Sender:     part.Parent,
			Fee:        fee,
			FirstValid: txnFirstValid,
			LastValid:  txnLastValid,
			Lease:      leaseBytes,
		},
		KeyregTxnFields: transactions.KeyregTxnFields{
			VotePK:      part.Voting.OneTimeSignatureVerifier,
			SelectionPK: part.VRF.PK,
		},
	}
	if cert := part.StateProofSigner(); cert != nil {
		if includeStateProofKeys { // TODO: remove this check and parameter after the network had enough time to upgrade
			t.KeyregTxnFields.StateProofPK = *(cert.GetVerifier())
		}
	}
	t.KeyregTxnFields.VoteFirst = part.FirstValid
	t.KeyregTxnFields.VoteLast = part.LastValid
	t.KeyregTxnFields.VoteKeyDilution = part.KeyDilution
	return t
}

// DeleteOldKeys securely deletes ephemeral keys for rounds strictly older than the given round.
func (part PersistedParticipation) DeleteOldKeys(current basics.Round, proto config.ConsensusParams) <-chan error {
	keyDilution := part.KeyDilution
	if keyDilution == 0 {
		keyDilution = proto.DefaultKeyDilution
	}

	part.Voting.DeleteBeforeFineGrained(basics.OneTimeIDForRound(current, keyDilution), keyDilution)

	errorCh := make(chan error, 1)
	deleteOldKeys := func(encodedVotingSecrets []byte) {
		errorCh <- part.Store.Atomic(func(ctx context.Context, tx *sql.Tx) error {
			_, err := tx.Exec("UPDATE ParticipationAccount SET voting=?", encodedVotingSecrets)
			if err != nil {
				return fmt.Errorf("Participation.DeleteOldKeys: failed to update account: %v", err)
			}
			return nil
		})
		close(errorCh)
	}
	voting := part.Voting.Snapshot()
	encodedVotingSecrets := protocol.Encode(&voting)
	go deleteOldKeys(encodedVotingSecrets)
	return errorCh
}

// PersistNewParent writes a new parent address to the partkey database.
func (part PersistedParticipation) PersistNewParent() error {
	return part.Store.Atomic(func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.Exec("UPDATE ParticipationAccount SET parent=?", part.Parent[:])
		return err
	})
}

// FillDBWithParticipationKeys initializes the passed database with participation keys
func FillDBWithParticipationKeys(store db.Accessor, address basics.Address, firstValid, lastValid basics.Round, keyDilution uint64) (part PersistedParticipation, err error) {
	if lastValid < firstValid {
		err = fmt.Errorf("FillDBWithParticipationKeys: firstValid %d is after lastValid %d", firstValid, lastValid)
		return
	}

	// TODO: change to ConsensusCurrentVersion when updated
	interval := config.Consensus[protocol.ConsensusFuture].CompactCertRounds
	maxValidPeriod := config.Consensus[protocol.ConsensusCurrentVersion].MaxKeyregValidPeriod

	if maxValidPeriod != 0 && uint64(lastValid-firstValid) > maxValidPeriod {
		return PersistedParticipation{}, fmt.Errorf("the validity period for mss is too large: the limit is %d", maxValidPeriod)
	}

	// Compute how many distinct participation keys we should generate
	firstID := basics.OneTimeIDForRound(firstValid, keyDilution)
	lastID := basics.OneTimeIDForRound(lastValid, keyDilution)
	numBatches := lastID.Batch - firstID.Batch + 1

	// Generate them
	v := crypto.GenerateOneTimeSignatureSecrets(firstID.Batch, numBatches)

	// Generate a new VRF key, which lives in the participation keys db
	vrf := crypto.GenerateVRFSecrets()

	// Generate a new key which signs the compact certificates
	stateProofSecrets, err := merklesignature.New(uint64(firstValid), uint64(lastValid), interval)
	if err != nil {
		return PersistedParticipation{}, err
	}

	// Construct the Participation containing these keys to be persisted
	part = PersistedParticipation{
		Participation: Participation{
			Parent:            address,
			VRF:               vrf,
			Voting:            v,
			StateProofSecrets: stateProofSecrets,
			FirstValid:        firstValid,
			LastValid:         lastValid,
			KeyDilution:       keyDilution,
		},
		Store: store,
	}
	// Persist the Participation into the database
	err = part.PersistWithSecrets()
	return part, err
}

// PersistWithSecrets writes Participation struct to the database along with all the secrets it contains
func (part PersistedParticipation) PersistWithSecrets() error {
	err := part.Persist()
	if err != nil {
		return err
	}
	return part.StateProofSecrets.Persist(part.Store) // must be called after part.Persist()
}

// Persist writes a Participation out to a database on the disk
func (part PersistedParticipation) Persist() error {
	rawVRF := protocol.Encode(part.VRF)
	voting := part.Voting.Snapshot()
	rawVoting := protocol.Encode(&voting)
	rawStateProof := protocol.Encode(part.StateProofSecrets)

	err := part.Store.Atomic(func(ctx context.Context, tx *sql.Tx) error {
		err := partInstallDatabase(tx)
		if err != nil {
			return fmt.Errorf("failed to install database: %w", err)
		}

		_, err = tx.Exec("INSERT INTO ParticipationAccount (parent, vrf, voting, firstValid, lastValid, keyDilution, stateProof) VALUES (?, ?, ?, ?, ?, ?,?)",
			part.Parent[:], rawVRF, rawVoting, part.FirstValid, part.LastValid, part.KeyDilution, rawStateProof)
		if err != nil {
			return fmt.Errorf("failed to insert account: %w", err)
		}
		return nil
	})

	if err != nil {
		err = fmt.Errorf("PersistedParticipation.Persist: %w", err)
	}
	return err
}

// Migrate is called when loading participation keys.
// Calls through to the migration helper and returns the result.
func Migrate(partDB db.Accessor) error {
	return partDB.Atomic(func(ctx context.Context, tx *sql.Tx) error {
		err := partMigrate(tx)
		if err != nil {
			return err
		}

		return merklesignature.InstallStateProofTable(tx)
	})
}

// Close closes the underlying database handle.
func (part PersistedParticipation) Close() {
	part.Store.Close()
}
