// Copyright (c) 2016-present Cloud <cloud@txthinking.com>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of version 3 of the GNU General Public
// License as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package plugin

// ClientAuthman is used to provide extra authentication mechanism.
type ClientAuthman interface {
	// GetToken is used for client to prepare extra data.
	GetToken() ([]byte, error)
}

// ServerAuthman is used to provide extra authentication mechanism
type ServerAuthman interface {
	// VerfiyToken is used for server to verify extra data.
	VerfiyToken(extra []byte, network, dst string) (Internet, error)
}

type Internet interface {
	TCPEgress(int) error
	UDPEgress(int) error
	DNSQuery(string) error
}
