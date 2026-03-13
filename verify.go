package main

import(
	"github.com/cloudflare/circl/sign/bls"
)


// func Verify[K KeyGroup](pub *PublicKey[K], msg []byte, sig Signature) bool
func Verify(pubKey *bls.PublicKey[bls.G1],msg []byte,sig bls.Signature) bool{
	return bls.Verify[bls.G1](pubKey,msg,sig)
}