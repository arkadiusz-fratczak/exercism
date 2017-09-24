package secret

const testVersion = 2

const winkFlag = 1
const doubleBlinkFlag = 1 << 1
const closeYourEyesFlag = 1 << 2
const jumpFlag = 1 << 3
const reverseFlag = 1 << 4

func Handshake(code uint) []string {
	handshake := make([]string, 0, 4)
	if code&winkFlag == 1 {
		handshake = append(handshake, "wink")
	}
	if code&doubleBlinkFlag == 2 {
		handshake = append(handshake, "double blink")
	}
	if code&closeYourEyesFlag == 4 {
		handshake = append(handshake, "close your eyes")
	}
	if code&jumpFlag == 8 {
		handshake = append(handshake, "jump")
	}
	if code&reverseFlag == 16 {
		handshake = reverse(handshake)
	}
	return handshake
}

func reverse(words []string) []string {
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return words
}
