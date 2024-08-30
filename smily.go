package main

import (
	"strings"
)

func CheckContain(text string, sign1 string, sign2 string) bool {
	if strings.Contains(text, sign1) || strings.Contains(text, sign2) {
		return true
	}
	return false
}

func CountSmilyFace(text []string) int {
	// TODO : start your code here
	smiles := 0
	for _, v := range text {
		if len(v) == 3 {
			pos1 := CheckContain(string(v[0]), ":", ";")
			pos2 := CheckContain(string(v[1]), "-", "~")
			pos3 := CheckContain(string(v[2]), ")", "D")
			if pos1 && pos2 && pos3 {
				smiles++
			}
		} else if len(v) == 2 {
			pos1 := CheckContain(string(v[0]), ":", ";")
			pos3 := CheckContain(string(v[1]), ")", "D")
			if pos1 && pos3 {
				smiles++
			}
		}

	}
	return smiles
}
