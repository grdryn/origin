// Copyright 2015 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package transport

import (
	"net"
	"net/http"
	"time"
)

// NewTimeoutTransport returns a transport created using the given TLS info.
// If read/write on the created connection blocks longer than its time limit,
// it will return timeout error.
func NewTimeoutTransport(info TLSInfo, rdtimeoutd, wtimeoutd time.Duration) (*http.Transport, error) {
	tr, err := NewTransport(info)
	if err != nil {
		return nil, err
	}
	// the timeouted connection will tiemout soon after it is idle.
	// it should not be put back to http transport as an idle connection for future usage.
	tr.MaxIdleConnsPerHost = -1
	tr.Dial = (&rwTimeoutDialer{
		Dialer: net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		},
		rdtimeoutd: rdtimeoutd,
		wtimeoutd:  wtimeoutd,
	}).Dial
	return tr, nil
}
