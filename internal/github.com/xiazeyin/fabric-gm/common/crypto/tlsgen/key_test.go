/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package tlsgen

import (
	"encoding/pem"
	"testing"

	"github.com/stretchr/testify/assert"
	tls "github.com/xiazeyin/gmgo/gmtls"
	"github.com/xiazeyin/gmgo/x509"
)

func TestLoadCert(t *testing.T) {
	pair, err := newCertKeyPair(false, false, nil, nil)
	assert.NoError(t, err)
	assert.NotNil(t, pair)
	tlsCertPair, err := tls.X509KeyPair(pair.Cert, pair.Key)
	assert.NoError(t, err)
	assert.NotNil(t, tlsCertPair)
	block, _ := pem.Decode(pair.Cert)
	cert, err := x509.ParseCertificate(block.Bytes)
	assert.NoError(t, err)
	assert.NotNil(t, cert)
}
