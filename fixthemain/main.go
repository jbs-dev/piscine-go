package main

import "github.com/01-edu/z01"

// Define a Door structure with a boolean field "state" to represent
// whether the door is open or closed.
type Door struct {
	state bool
}

// PrintStr prints the string to the console.
func PrintStr(str string) {
	// Iterate over each character in the string and print it using z01.
	for _, s := range str {
		z01.PrintRune(s)
	}
	// Print a newline character at the end of the string.
	z01.PrintRune('\n')
}

// OpenDoor sets the "state" field of the Door object to true to
// represent that the door is now open.
func OpenDoor(ptrDoor *Door) bool {
	// Print a message to indicate that the door is opening.
	PrintStr("Door Opening...")
	// Set the "state" field of the Door object to true.
	ptrDoor.state = true
	// Return true to indicate that the operation was successful.
	return true
}

// CloseDoor sets the "state" field of the Door object to false to
// represent that the door is now closed.
func CloseDoor(ptrDoor *Door) bool {
	// Print a message to indicate that the door is closing.
	PrintStr("Door Closing...")
	// Set the "state" field of the Door object to false.
	ptrDoor.state = false
	// Return true to indicate that the operation was successful.
	return true
}

// IsDoorOpen returns true if the "state" field of the Door object is false,
// which represents that the door is open.
func IsDoorOpen(ptrDoor *Door) bool {
	// Print a message to ask if the door is open.
	PrintStr("is the Door opened ?")
	// Return the opposite of the "state" field of the Door object.
	return !ptrDoor.state
}

// IsDoorClose returns true if the "state" field of the Door object is true,
// which represents that the door is closed.
func IsDoorClose(ptrDoor *Door) bool {
	// Print a message to ask if the door is closed.
	PrintStr("is the Door closed ?")
	// Return the "state" field of the Door object.
	return !ptrDoor.state
}

func main() {
	// Create a new Door object.
	door := &Door{}

	// Open the door.
	OpenDoor(door)

	// If the door is closed, open it.
	if IsDoorClose(door) {
		OpenDoor(door)
	}

	// If the door is open, close it.
	if IsDoorOpen(door) {
		CloseDoor(door)
	}

	// If the door is still open for some reason, close it.
	if door.state {
		CloseDoor(door)
	}
}
