// Copyright 2018, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dns

import (
	"net"
)

//
// Sender is interface that for implementing sending DNS message.
//
type Sender interface {
	Send(msg *Message, addr net.Addr) (n int, err error)
}
