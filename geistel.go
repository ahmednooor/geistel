package main

/*GeistelEncipherBytes (inputBytes []byte, keyBytes []byte) []byte*/
func GeistelEncipherBytes(inputBytes, keyBytes []byte) []byte {
	if len(inputBytes)%2 != 0 {
		byteToAppend := byte(0)
		if len(keyBytes) > 0 {
			byteToAppend = keyBytes[0]
		}
		inputBytes = append(inputBytes, byteToAppend)
		outputBytes := applyFeistelNetwork(inputBytes, keyBytes, false)
		return append(outputBytes, byte(0))
	}

	return applyFeistelNetwork(inputBytes, keyBytes, false)
}

/*GeistelDecipherBytes (inputBytes []byte, keyBytes []byte) []byte*/
func GeistelDecipherBytes(inputBytes, keyBytes []byte) []byte {
	if len(inputBytes)%2 != 0 && inputBytes[len(inputBytes)-1] == byte(0) {
		inputBytes = inputBytes[0 : len(inputBytes)-1]
		outputBytes := applyFeistelNetwork(inputBytes, keyBytes, true)
		return outputBytes[0 : len(outputBytes)-1]
	}

	return applyFeistelNetwork(inputBytes, keyBytes, true)
}

func applyFeistelNetwork(inputBytes, keyBytes []byte, isDecipher bool) []byte {
	if isDecipher {
		keyBytes = reverse(keyBytes)
	}

	inputBytesLen := len(inputBytes)
	midPoint := int(inputBytesLen / 2)

	leftBytes := inputBytes[0:midPoint]
	rightBytes := inputBytes[midPoint:len(inputBytes)]

	for _, keyByte := range keyBytes {
		leftBytes, rightBytes = applyFeistelFlip(leftBytes, rightBytes, keyByte)
	}

	return append(rightBytes, leftBytes...)
}

func applyFeistelFlip(leftBytes []byte, rightBytes []byte, keyByte byte) ([]byte, []byte) {
	nextLeftBytes := rightBytes
	roundFuncOutBytes := roundFunction(rightBytes, keyByte)
	nextRightBytes := xorFunction(leftBytes, roundFuncOutBytes)

	return nextLeftBytes, nextRightBytes
}

func xorFunction(leftBytes, roundFuncOutBytes []byte) []byte {
	xoredBytes := []byte{}

	for i := range leftBytes {
		xoredBytes = append(xoredBytes, leftBytes[i]^roundFuncOutBytes[i])
	}

	return xoredBytes
}

func roundFunction(inputBytes []byte, keyByte byte) []byte {
	outputBytes := []byte{}

	for _, inputByte := range inputBytes {
		outputBytes = append(outputBytes, getMod(inputByte, keyByte))
	}

	return outputBytes
}

func getMod(inputByte, keyByte byte) byte {
	if inputByte == 0 {
		return keyByte
	}

	if keyByte == 0 {
		return inputByte
	}

	if inputByte >= keyByte {
		return inputByte % keyByte
	}

	return keyByte % inputByte
}

func reverse(inputSlice []byte) []byte {
	newSlice := append([]byte{}, inputSlice...)

	// courtesy of internet
	for i := len(newSlice)/2 - 1; i >= 0; i-- {
		opp := len(newSlice) - 1 - i
		newSlice[i], newSlice[opp] = newSlice[opp], newSlice[i]
	}

	return newSlice
}
