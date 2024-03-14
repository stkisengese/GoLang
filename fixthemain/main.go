package main

import "github.com/01-edu/z01"

type Door struct {
	state string
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(ptrDoor *Door) bool {
	if ptrDoor.state == "OPEN" {
		PrintStr("Door Closing...\n")
		ptrDoor.state = "CLOSE"
		return true
	}
	PrintStr("Door is already closed!\n")
	return false
}

func OpenDoor(ptrDoor *Door) bool {
	if ptrDoor.state == "CLOSE" {
		PrintStr("Door Opening...\n")
		ptrDoor.state = "OPEN"
		return true
	}
	PrintStr("Door is already open!\n")
	return false
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?\n")
	return ptrDoor.state == "OPEN"
}

func IsDoorClosed(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?\n")
	return ptrDoor.state == "CLOSE"
}

func main() {
	door := &Door{"CLOSE"}
	OpenDoor(door)
	IsDoorClosed(door)
	IsDoorOpen(door)
	if door.state == "OPEN" {
		CloseDoor(door)
	}
}
