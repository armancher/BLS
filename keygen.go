package main

import (
	
	"crypto/rand"
	

	"github.com/cloudflare/circl/sign/bls"
	
)

//curveorder is the order of r
//this is represented in hexadecimal(based-16)
// var curveorder , _ =new(big.Int).SetString(
// 	"73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff00000001",
// 	16,
// )


// It uses the G1 key group, meaning public keys live in G1 and signatures in G2.
func KeyGen() (*bls.PrivateKey[bls.G1], *bls.PublicKey[bls.G1], error) {
	// IKM (Input Keying Material) must be at least 32 bytes 
	ikm := make([]byte,32)
	salt := make([]byte,32)

	if _, err := rand.Reader.Read(ikm[:]); err != nil {
		return nil, nil, err
	}
	if _, err := rand.Reader.Read(salt[:]); err != nil {
		return nil, nil, err
	}

	keyInfo := []byte("bls-ethereum-project")

	privKey, err := bls.KeyGen[bls.G1](ikm, salt, keyInfo)
    if err != nil {
        return nil, nil, err
    }

    pubKey := SkToPk(privKey)
    return privKey, pubKey, nil


	

	//L is the integer given by ceil((3 * ceil(log2(r))) / 16)
	// const L =48

	// = "bls-ethereum-project" + 0x00 + 0x30
	// keyInfowithL := append(keyInfo,byte(L>>8),byte(L))

	// 	IKM (random seed)
	// 	↓
	// HKDF-Extract  →  PRK  (Pseudorandom Key, intermediate)
	// 	↓
	// HKDF-Expand   →  OKM  (48 bytes, the actual output)
	// 	↓
	// mod r         →  SK   (your final secret key)

	// for{

	// 	// Step 2: HKDF-Extract(salt, IKM || I2OSP(0,1))
	// 	PRK:= hkdf.Extract(sha256.New,append(ikm[:], 0x00),salt[:])

	// 	// Step 3: HKDF-Expand(PRK, keyInfo || I2OSP(L,2), L)
	// 	//okm : Output Keying Material
	// 	okm := make([]byte, L)
	// 	reader := hkdf.Expand(sha256.New, PRK, keyInfowithL)
	// 	if _, err := reader.Read(okm); err != nil {
	// 		return nil, nil, errors.New("KeyGen: HKDF-Expand failed: " + err.Error())
	// 	}

		// Step 4: SK = OS2IP(OKM) mod r
		// sk:=new(big.Int).SetBytes(okm)
		// sk.Mod(sk,curveorder)

		// //step 5 : if SK == 0, rehash salt and retry
		// if sk.Cmp(big.NewInt(0))==0{
		// 	h:=sha256.Sum256(salt)
		// 	salt = h[:]
		// 	continue
		// }}

		// skbyte := make([]byte,32)
		// sk.FillBytes(skbyte
	
	
}
func SkToPk(privKey *bls.PrivateKey[bls.G1]) *bls.PublicKey[bls.G1] {
	return privKey.PublicKey()
}