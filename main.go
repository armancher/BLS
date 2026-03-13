package main

import (
	"fmt"

	"github.com/cloudflare/circl/sign/bls"
)

func main() {
	fmt.Println("=== BLS Signature Demo ===")
	fmt.Println()

	// ── 1. Key Generation 
	fmt.Println("--- Key Generation ---")
	privKey1, pubKey1, err := KeyGen()
	if err != nil {
		fmt.Println("Error KeyGen 1:", err)
		return
	}
	privKey2, pubKey2, err := KeyGen()
	if err != nil {
		fmt.Println("Error KeyGen 2:", err)
		return
	}
	privKey3, pubKey3, err := KeyGen()
	if err != nil {
		fmt.Println("Error KeyGen 3:", err)
		return
	}

	pub1Bytes, _ := pubKey1.MarshalBinary()
	pub2Bytes, _ := pubKey2.MarshalBinary()
	pub3Bytes, _ := pubKey3.MarshalBinary()
	fmt.Printf("Public key 1 (%d bytes): %x\n", len(pub1Bytes), pub1Bytes)
	fmt.Printf("Public key 2 (%d bytes): %x\n", len(pub2Bytes), pub2Bytes)
	fmt.Printf("Public key 3 (%d bytes): %x\n", len(pub3Bytes), pub3Bytes)
	fmt.Println()

	// ── 2. Signing 
	fmt.Println("--- Signing ---")
	msg1 := []byte("message from signer 1")
	msg2 := []byte("message from signer 2")
	msg3 := []byte("message from signer 3")

	sig1 := Sign(privKey1, msg1)
	sig2 := Sign(privKey2, msg2)
	sig3 := Sign(privKey3, msg3)

	fmt.Printf("Signature 1 (%d bytes): %x\n", len(sig1), sig1)
	fmt.Printf("Signature 2 (%d bytes): %x\n", len(sig2), sig2)
	fmt.Printf("Signature 3 (%d bytes): %x\n", len(sig3), sig3)
	fmt.Println()

	// ── 3. Individual Verification 
	fmt.Println("--- Verification ---")
	fmt.Println("Sig1 valid:", Verify(pubKey1, msg1, sig1))
	fmt.Println("Sig2 valid:", Verify(pubKey2, msg2, sig2))
	fmt.Println("Sig3 valid:", Verify(pubKey3, msg3, sig3))
	fmt.Println("Sig1 with wrong message (should be false):", Verify(pubKey1, []byte("wrong message"), sig1))
	fmt.Println()

	// ── 4. Aggregation 
	fmt.Println("--- Aggregation ---")
	aggSig, err := Aggregate([]bls.Signature{sig1, sig2, sig3})
	if err != nil {
		fmt.Println("Error aggregating:", err)
		return
	}
	fmt.Printf("Aggregated signature (%d bytes): %x\n", len(aggSig), aggSig)
	fmt.Println()

	// ── 5. Aggregate Verification 
	fmt.Println("--- Aggregate Verification ---")
	pubs := []*bls.PublicKey[bls.G1]{pubKey1, pubKey2, pubKey3}
	msgs := [][]byte{msg1, msg2, msg3}

	fmt.Println("Aggregate valid:", VerifyAggregate(pubs, msgs, aggSig))
	fmt.Println("Aggregate with wrong message (should be false):", VerifyAggregate(pubs, [][]byte{msg1, []byte("wrong"), msg3}, aggSig))
}