package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

func main() {
	inputBytes := []byte("abc")
	keyBytes := []byte("a")

	hasher := sha256.New()
	hasher.Write(keyBytes)
	hashedKeyBytes := hasher.Sum(nil)

	fmt.Printf("INPUT STRING:\t\t%s\n", string(inputBytes))
	fmt.Printf("KEY STRING:\t\t%s\n", string(keyBytes))
	fmt.Println()
	fmt.Printf("INPUT BYTES:\t\t%-4d\n", inputBytes)
	fmt.Printf("KEY BYTES:\t\t%-4d\n", keyBytes)
	fmt.Printf("HASHED KEY BYTES:\t%-4d\n", hashedKeyBytes)

	encipheredBytes := GeistelEncipherBytes(inputBytes, keyBytes)
	fmt.Println()
	fmt.Printf("ENCIPHERED BYTES:\t%-4d\n", encipheredBytes)
	fmt.Println("^ With raw key (" + strconv.Itoa(len(keyBytes)) + " Rounds).")

	decipheredBytes := GeistelDecipherBytes(encipheredBytes, keyBytes)
	fmt.Println()
	fmt.Printf("DECIPHERED BYTES:\t%-4d\n", decipheredBytes)
	fmt.Println("^ With raw key (" + strconv.Itoa(len(keyBytes)) + " Rounds).")

	encipheredBytesHashedKey := GeistelEncipherBytes(inputBytes, hashedKeyBytes)
	fmt.Println()
	fmt.Printf("ENCIPHERED BYTES:\t%-4d\n", encipheredBytesHashedKey)
	fmt.Println("^ With 32 byte SHA256 hash of the key (32 Rounds).")

	decipheredBytesHashedKey := GeistelDecipherBytes(encipheredBytesHashedKey, hashedKeyBytes)
	fmt.Println()
	fmt.Printf("DECIPHERED BYTES:\t%-4d\n", decipheredBytesHashedKey)
	fmt.Println("^ With 32 byte SHA256 hash of the key (32 Rounds).")
}
