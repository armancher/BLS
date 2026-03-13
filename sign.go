package main


import (
	"github.com/cloudflare/circl/sign/bls"
)


func Sign(privKey *bls.PrivateKey[bls.G1], msg []byte) bls.Signature {
    return bls.Sign[bls.G1](privKey, msg)
}