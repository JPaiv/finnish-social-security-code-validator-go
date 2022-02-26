package Validate

import (
	"strconv"
)

func generateRemainderCheck() map[int]string {
	checkRemainders := make(map[int]string)
	var alphabet string
	alphabet = "ABCDEFHJKLMNOPRSTUVWXYZ"
	for i := 0; i < 32; i++ {
		if i < 10 {
			checkRemainders[i] = strconv.Itoa(i)
		} else {
			checkRemainders[i] = string(alphabet[0])
			alphabet = alphabet[1:]
		}
	}
	return checkRemainders
}
