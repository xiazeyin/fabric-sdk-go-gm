/*
Copyright IBM Corp All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package operations_test

// http "github.com/xiazeyin/gmgo/gmhttp"只支持tls和x509，不支持gmtls和gmx509
import (
	"io/ioutil"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xiazeyin/fabric-config-gm/healthz"
	"github.com/xiazeyin/fabric-gm/common/crypto/tlsgen"
	http "github.com/xiazeyin/gmgo/gmhttp"
	tls "github.com/xiazeyin/gmgo/gmtls"
	"github.com/xiazeyin/gmgo/x509"
)

func TestOperations(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Operations Suite")
}

func generateCertificates(tempDir string) {
	serverCA, err := tlsgen.NewCA()
	Expect(err).NotTo(HaveOccurred())
	err = ioutil.WriteFile(filepath.Join(tempDir, "server-ca.pem"), serverCA.CertBytes(), 0640)
	Expect(err).NotTo(HaveOccurred())
	serverKeyPair, err := serverCA.NewServerCertKeyPair("127.0.0.1")
	Expect(err).NotTo(HaveOccurred())
	err = ioutil.WriteFile(filepath.Join(tempDir, "server-cert.pem"), serverKeyPair.Cert, 0640)
	Expect(err).NotTo(HaveOccurred())
	err = ioutil.WriteFile(filepath.Join(tempDir, "server-key.pem"), serverKeyPair.Key, 0640)
	Expect(err).NotTo(HaveOccurred())

	clientCA, err := tlsgen.NewCA()
	Expect(err).NotTo(HaveOccurred())
	err = ioutil.WriteFile(filepath.Join(tempDir, "client-ca.pem"), clientCA.CertBytes(), 0640)
	Expect(err).NotTo(HaveOccurred())
	clientKeyPair, err := clientCA.NewClientCertKeyPair()
	Expect(err).NotTo(HaveOccurred())
	err = ioutil.WriteFile(filepath.Join(tempDir, "client-cert.pem"), clientKeyPair.Cert, 0640)
	Expect(err).NotTo(HaveOccurred())
	err = ioutil.WriteFile(filepath.Join(tempDir, "client-key.pem"), clientKeyPair.Key, 0640)
	Expect(err).NotTo(HaveOccurred())
}

func newHTTPClient(tlsDir string, withClientCert bool) *http.Client {
	clientCertPool := x509.NewCertPool()
	caCert, err := ioutil.ReadFile(filepath.Join(tlsDir, "server-ca.pem"))
	Expect(err).NotTo(HaveOccurred())
	clientCertPool.AppendCertsFromPEM(caCert)

	tlsClientConfig := &tls.Config{
		RootCAs: clientCertPool,
	}
	if withClientCert {
		clientCert, err := tls.LoadX509KeyPair(
			filepath.Join(tlsDir, "client-cert.pem"),
			filepath.Join(tlsDir, "client-key.pem"),
		)
		Expect(err).NotTo(HaveOccurred())
		tlsClientConfig.Certificates = []tls.Certificate{clientCert}
	}

	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsClientConfig,
		},
	}
}

//go:generate counterfeiter -o fakes/healthchecker.go -fake-name HealthChecker . healthChecker
type healthChecker interface {
	healthz.HealthChecker
}