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

package messagetracer

import (
	"github.com/Orca18/novarand/config"
	"github.com/Orca18/novarand/logging"
)

// MessageTracer interface for configuring trace client and sending trace messages
// 추적 클라이언트 구성 및 추적 메시지 전송을 위한 MessageTracer 인터페이스
type MessageTracer interface {
	// Init configures trace client or returns nil.
	// Init는 추적 클라이언트를 구성하거나 nil을 반환합니다.
	// Caller is expected to check for nil, e.g. `if t != nil {t.HashTrace(...)}`
	// 호출자는 nil을 확인해야 합니다. `if t != nil {t.HashTrace(...)}`
	Init(cfg config.Local) MessageTracer

	// HashTrace submits a trace message to the statistics server.
	// HashTrace는 통계 서버에 추적 메시지를 제출합니다.
	HashTrace(prefix string, data []byte)
}

var implFactory func(logging.Logger) MessageTracer

type nopMessageTracer struct {
}

func (gmt *nopMessageTracer) Init(cfg config.Local) MessageTracer {
	return nil
}
func (gmt *nopMessageTracer) HashTrace(prefix string, data []byte) {
}

var singletonNopMessageTracer nopMessageTracer

// NewTracer constructs a new MessageTracer if that has been compiled in with the build tag `msgtrace`
// NewTracer는 빌드 태그 `msgtrace`로 컴파일된 경우 새 MessageTracer를 생성합니다.
func NewTracer(log logging.Logger) MessageTracer {
	if implFactory != nil {
		log.Info("graphtrace factory enabled")
		return implFactory(log)
	}
	log.Info("graphtrace factory DISabled")
	return &singletonNopMessageTracer
}

// Proposal is a prefix for HashTrace()
// Proposal은 HashTrace()의 접두어입니다.
const Proposal = "prop"
