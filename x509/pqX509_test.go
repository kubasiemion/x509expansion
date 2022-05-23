package x509

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"testing"
)

var dilprivkey1 = `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`

func getDilTestPrivBytes() []byte {
	b, e := base64.StdEncoding.DecodeString(dilprivkey1)
	if e != nil {
		fmt.Println(e)
	}
	return b
}

var dilithiumTestPubKey = `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`

func getDilTestPubBytes() []byte {
	b, e := base64.StdEncoding.DecodeString(dilithiumTestPubKey)
	if e != nil {
		fmt.Println(e)
	}
	return b
}

func TestX509RSA4DIL(t *testing.T) {

	block, _ := pem.Decode([]byte(selfSignedRootCertRSA4096Pem))
	RootCert, e := ParseCertificate(block.Bytes)
	if e != nil {
		t.Error(e)
		return
	}
	//fmt.Println(hex.EncodeToString(block.Bytes))

	ca := GetTemplate()

	pub := &PQPublicKey{RawBytes: getDilTestPubBytes(), OID: OidDilithiumRawHigh}

	fmt.Println(len(pub.Bytes()))

	block, _ = pem.Decode([]byte(rsa4096privPem))
	rootRSAKey, e = ParsePKCS8PrivateKey(block.Bytes)
	if e != nil {
		t.Error(e)
		return
	}
	certb, e := CreateCertificate(rand.Reader, ca, RootCert, pub, rootRSAKey)
	if e != nil {
		t.Error(e)
		return
	}

	block = &pem.Block{Type: "CERTIFICATE", Bytes: certb}
	certbuf := new(bytes.Buffer)
	e = pem.Encode(certbuf, block)
	if e != nil {
		t.Error(e)
		return
	}
	//fmt.Println(string(certbuf.Bytes()))

	p2, e = ParseCertificate(certb)
	if e != nil {
		t.Error(e)
		return
	}
	cpool := NewCertPool()
	cpool.AddCert(RootCert)
	chain, err := p2.Verify(VerifyOptions{Roots: cpool})
	fmt.Println(err, chain)
}

var rootRSAKey interface{}
var p2 *Certificate // DILITHIUM key certified
func TestDIL4RSA(t *testing.T) {
	fmt.Println("here")

	PrivKey := &PQPrivateKey{Privbytes: getDilTestPrivBytes(), Pubbytes: getDilTestPubBytes(), OID: OidDilithiumRawHigh}

	ca := GetTemplate()

	xc, e := CreateCertificate(rand.Reader, ca, p2, rootRSAKey.(*rsa.PrivateKey).Public(), PrivKey)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(len(xc))
}
