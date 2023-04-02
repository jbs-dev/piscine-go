package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
	valc := *******c // create a temporary variable to store the value pointed to by the 7-level deep pointer c
	*******c = ***a  // set the value pointed to by the 7-level deep pointer c to the value pointed to by the 3-level deep pointer a
	vald := ****d    // create a temporary variable to store the value pointed to by the 4-level deep pointer d
	****d = valc     // set the value pointed to by the 4-level deep pointer d to the value stored in the temporary variable valc
	valb := *b       // create a temporary variable to store the value pointed to by the pointer b
	*b = vald        // set the value pointed to by the pointer b to the value stored in the temporary variable vald
	***a = valb      // set the value pointed to by the 3-level deep pointer a to the value stored in the temporary variable valb
}

func Decript(a ***int, b *int, c *******int, d ****int) {
	vala := ***a    // create a temporary variable to store the value pointed to by the 3-level deep pointer a
	***a = *******c // set the value pointed to by the 3-level deep pointer a to the value pointed to by the 7-level deep pointer c
	valb := *b      // create a temporary variable to store the value pointed to by the pointer b
	*b = vala       // set the value pointed to by the pointer b to the value stored in the temporary variable vala
	vald := ****d   // create a temporary variable to store the value pointed to by the 4-level deep pointer d
	****d = valb    // set the value pointed to by the 4-level deep pointer d to the value stored in the temporary variable valb
	*******c = vald // set the value pointed to by the 7-level deep pointer c to the value stored in the temporary variable vald
}
