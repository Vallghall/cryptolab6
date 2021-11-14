package rsa

import (
	"crypto/rand"
	"math/big"
)

var _1 = big.NewInt(1)

type PublicKey struct {
	e *big.Int
	n *big.Int
}

type PrivateKey struct {
	d *big.Int
	n *big.Int
}

func (pk *keyPair) D() *big.Int {
	return pk.d
}

func (pk *keyPair) E() *big.Int {
	return pk.e
}

type keyPair struct {
	*PublicKey
	*PrivateKey
}

func (pk *PrivateKey) Decrypt(enc []int64) []int64 {
	dec := make([]int64, len(enc), len(enc))
	for i, sym := range enc {
		eSym := big.NewInt(sym)
		eSym.Exp(eSym, pk.d, nil).Mod(eSym, pk.n)
		dec[i] = eSym.Int64()
	}
	return dec
}

func (pk *PublicKey) Encrypt(txt []int64) (enc []int64) {
	for _, sym := range txt {
		eSym := big.NewInt(sym)
		eSym.Exp(eSym, pk.e, nil).Mod(eSym, pk.n)
		enc = append(enc, eSym.Int64())
	}
	return
}

func (p *Primes) GenerateKeyPair() *keyPair {
	e := p.generatePublicKey()
	d := new(big.Int).ModInverse(e, p.Phi())
	return &keyPair{
		&PublicKey{e, p.N()},
		&PrivateKey{d, p.N()},
	}
}

func (p Primes) generatePublicKey() *big.Int {
	e, _ := rand.Int(rand.Reader, p.Phi())
	for i := new(big.Int); ; {
		if i.GCD(nil, nil, e, p.Phi()); i.Cmp(_1) == 0 && e.Cmp(_1) != 0 {
			break
		}
		e, _ = rand.Int(rand.Reader, p.Phi())
	}
	return e
}
