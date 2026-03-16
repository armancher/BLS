package main

import (
	"fmt"
	"testing"

	"github.com/cloudflare/circl/sign/bls"
)

func TestKeyGen(t *testing.T) {
	privKey1, pubKey1, err1 := KeyGen()

	if err1 != nil {
		t.Fatalf("KeyGen returned an error: %v", err1)
	}

	if privKey1 == nil {
		t.Fatal("Private Key is nil")
	}

	if pubKey1 == nil {
		t.Fatal("Public Key is nil")
	}

	// verify that two successive KeyGen calls produce different keys.
	// We only compare public keys because the public key is derived
	// deterministically from the private key via pk = [sk] * g1,
	// so if two public keys are different, the private keys must be
	// different too. Comparing private keys directly is also discouraged
	// for security reasons: the less the private key is manipulated
	// in code, the lower the risk of accidentally exposing it.
	_, pubKey2, err2 := KeyGen()
	if err2 != nil {
		t.Fatalf("Second KeyGen returned an error: %v", err2)
	}

	pub1Bytes, _ := pubKey1.MarshalBinary()
	pub2Bytes, _ := pubKey2.MarshalBinary()

	if string(pub1Bytes) == string(pub2Bytes) {
		t.Fatal("Two KeyGen calls produced the same public key")
	}
}

func TestSignAndVerify(t *testing.T) {
	privKey, pubKey, err := KeyGen()
	if err != nil {
		t.Fatalf("KeyGen returned an error: %v", err)
	}

	msg := []byte("This is a test message")
	sig := Sign(privKey, msg)

	// verify that the signature is valid with the correct key and message
	if !Verify(pubKey, msg, sig) {
		t.Fatal("Verify returned false on a valid signature")
	}

	// verify that the signature is invalid with a wrong message
	if Verify(pubKey, []byte("wrong message"), sig) {
		t.Fatal("Verify returned true on a wrong message")
	}

	// verify that the signature is invalid with a wrong public key
	_, wrongPubKey, err := KeyGen()
	if err != nil {
		t.Fatalf("Second KeyGen returned an error: %v", err)
	}
	if Verify(wrongPubKey, msg, sig) {
		t.Fatal("Verify returned true with a wrong public key")
	}
}

func TestAggregate(t *testing.T) {
	privKey1, pubKey1, err := KeyGen()
	if err != nil {
		t.Fatalf("KeyGen returned an error: %v", err)
	}

	privKey2, pubKey2, err := KeyGen()
	if err != nil {
		t.Fatalf("KeyGen returned an error: %v", err)
	}

	privKey3, pubKey3, err := KeyGen()
	if err != nil {
		t.Fatalf("KeyGen returned an error: %v", err)
	}

	msg1 := []byte("This is the first message")
	msg2 := []byte("This is the second message")
	msg3 := []byte("This is the third message")

	sig1 := Sign(privKey1, msg1)
	sig2 := Sign(privKey2, msg2)
	sig3 := Sign(privKey3, msg3)

	aggSig, err := Aggregate([]bls.Signature{sig1, sig2, sig3})
	if err != nil {
		t.Fatalf("Aggregate returned an error: %v", err)
	}

	pubs := []*bls.PublicKey[bls.G1]{pubKey1, pubKey2, pubKey3}
	msgs := [][]byte{msg1, msg2, msg3}

	if !VerifyAggregate(pubs, msgs, aggSig) {
		t.Fatal("VerifyAggregate returned false on valid aggregated signature")
	}

	wrongMsgs := [][]byte{msg1, []byte("wrong message"), msg3}
	if VerifyAggregate(pubs, wrongMsgs, aggSig) {
		t.Fatal("VerifyAggregate should return false with wrong message")
	}

	_, wrongPubKey, err := KeyGen()
	if err != nil {
		t.Fatalf("KeyGen returned an error: %v", err)
	}
	wrongPubs := []*bls.PublicKey[bls.G1]{wrongPubKey, pubKey2, pubKey3}
	if VerifyAggregate(wrongPubs, msgs, aggSig) {
		t.Fatal("VerifyAggregate should return false with wrong public key")
	}

}

func TestAggregate128(t *testing.T) {
	const n = 128
	privKeys := make([]*bls.PrivateKey[bls.G1], n)
	pubKeys := make([]*bls.PublicKey[bls.G1], n)
	msgs := make([][]byte, n)
	sigs := make([]bls.Signature, n)
	var err error

	for i := 0; i < n; i++ {

		privKeys[i], pubKeys[i], err = KeyGen()
		if err != nil {
			t.Fatalf("KeyGen failed at index %d: %v", i, err)
		}
		msgs[i] = []byte(fmt.Sprintf("This is the message number: %d", i))
		sigs[i] = Sign(privKeys[i], msgs[i])
	}

	aggSig, err := Aggregate(sigs)
	if err != nil {
		t.Fatalf("Aggregate returned an error: %v", err)
	}

	if !VerifyAggregate(pubKeys, msgs, aggSig) {
		t.Fatal("VerifyAggregate returned false on valid aggregated signature")
	}

}

func BenchmarkKeyGen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		KeyGen()
	}
}

func BenchmarkSign(b *testing.B) {
	privKey, _, _ := KeyGen()
	msg := []byte("This is a test message")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Sign(privKey, msg)
	}
}

func BenchmarkVerify(b *testing.B) {
	privKey, pubKey, _ := KeyGen()
	msg := []byte("This is a test message")
	sig := Sign(privKey, msg)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Verify(pubKey, msg, sig)
	}
}

func BenchmarkAggregate(b *testing.B) {
	privKey1, pubKey1, _ := KeyGen()
	privKey2, pubKey2, _ := KeyGen()
	privKey3, pubKey3, _ := KeyGen()
	pubs := []*bls.PublicKey[bls.G1]{pubKey1, pubKey2, pubKey3}

	msg1 := []byte("This is the first message")
	msg2 := []byte("This is the second message")
	msg3 := []byte("This is the third message")
	msgs := [][]byte{msg1, msg2, msg3}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sig1 := Sign(privKey1, msg1)
		sig2 := Sign(privKey2, msg2)
		sig3 := Sign(privKey3, msg3)

		aggSig, err := Aggregate([]bls.Signature{sig1, sig2, sig3})
		if err != nil {
			b.Fatalf("Aggregate returned an error: %v", err)
		}
		VerifyAggregate(pubs, msgs, aggSig)
	}
}

func BenchmarkAggregate128(b *testing.B) {
	const n = 128
	privKeys := make([]*bls.PrivateKey[bls.G1], n)
	pubKeys := make([]*bls.PublicKey[bls.G1], n)
	msgs := make([][]byte, n)
	sigs := make([]bls.Signature, n)
	var err error

	for i := 0; i < n; i++ {

		privKeys[i], pubKeys[i], err = KeyGen()
		if err != nil {
			b.Fatalf("Keygen failed at index %d: %v", i, err)
		}

		msgs[i] = []byte(fmt.Sprintf("This is the message number: %d", i))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			sigs[j] = Sign(privKeys[j], msgs[j])
		}

		aggSig, err := Aggregate(sigs)
		if err != nil {
			b.Fatalf("Aggregate returned an error: %v", err)
		}

		VerifyAggregate(pubKeys, msgs, aggSig)
	}
}
