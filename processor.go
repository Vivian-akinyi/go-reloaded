package main

func ProcessText(text string) string {
	text = ProcessHex(text)
	text = ProcessBinary(text)
	text = ProcessUpper(text)
	text = ProcessLower(text)
	text = ProcessCap(text)
	text = ProcessUpperMultiple(text)
	return text
}
