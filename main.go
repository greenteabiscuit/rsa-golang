package main

import (
	"fmt"
	"math/big"
)

type publicKey struct {
	E int
	N int
}

type privateKey struct {
	D int
	N int
}

func main() {
	publicKey, privateKey := generateKeys(101, 3259)
	plainText := "This is a secret"
	encryptedText := encrypt(plainText, publicKey)
	decryptedText := decrypt(encryptedText, privateKey)
	fmt.Println("Public Key:", publicKey.E, publicKey.N)
	fmt.Println("Private Key:", privateKey.D, privateKey.N)
	fmt.Println("Message:", plainText)
	fmt.Println("Encrypted Text:", encryptedText)
	fmt.Println("Decrypted Text:", decryptedText)
}

func generateKeys(p int, q int) (public publicKey, secret privateKey) {
	N := p * q
	L := lcm(p-1, q-1)
	var D, E int
	for i := 2; i < L; i++ {
		if gcd(i, L) == 1 {
			E = i
			break
		}
	}
	for i := 2; i < L; i++ {
		if (E*i)%L == 1 {
			D = i
			break
		}
	}
	return publicKey{E: E, N: N}, privateKey{D: D, N: N}
}

func lcm(p int, q int) int {
	return (p * q) / gcd(p, q)
}

func gcd(p int, q int) int {
	// https://play.golang.org/p/SmzvkDjYlb
	for q != 0 {
		t := q
		q = p % q
		p = t
	}
	return p
}

func encrypt(plainText string, publicKey publicKey) string {
	E, N := publicKey.E, publicKey.N
	resultString := ""
	for _, char := range plainText {
		bigChar := big.NewInt(int64(char))
		bigE := big.NewInt(int64(E))
		bigN := big.NewInt(int64(N))

		res := *bigChar.Exp(bigChar, bigE, bigN)
		resultString += string(rune(res.Int64()))
	}
	return resultString
}

func decrypt(encryptedText string, privateKey privateKey) string {
	D, N := privateKey.D, privateKey.N
	resultString := ""
	for _, char := range encryptedText {
		bigChar := big.NewInt(int64(char))
		bigD := big.NewInt(int64(D))
		bigN := big.NewInt(int64(N))

		res := *bigChar.Exp(bigChar, bigD, bigN)
		resultString += string(rune(res.Int64()))
	}
	return resultString
}
