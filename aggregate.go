package main

import (
	"github.com/cloudflare/circl/sign/bls"
)

// Aggregate combines multiple BLS signatures into a single signature.
// func Aggregate[K KeyGroup](k K, sigs []Signature) (Signature, error)

func Aggregate(sigs []bls.Signature)(bls.Signature,error){
	return bls.Aggregate(bls.G1{},sigs)
}


// VerifyAggregate checks an aggregated signature against multiple public keys and messages.
// func VerifyAggregate[K KeyGroup](pubs []*PublicKey[K], msgs [][]byte, aggSig Signature) bool

func VerifyAggregate(pubs []*bls.PublicKey[bls.G1],msgs [][]byte,aggSig bls.Signature) bool {
	return bls.VerifyAggregate[bls.G1](pubs,msgs,aggSig)
}