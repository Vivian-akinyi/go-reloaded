package main

func ProcessText(text string) string {
	text = ProcessHex(text)
	text = ProcessBinary(text)
	return text
}
