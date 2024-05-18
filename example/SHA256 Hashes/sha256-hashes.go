package main

// SHA256 hashes are frequently used to compute short identities for binary or text blobs. For example, TLS/SSL certificates use SHA256 to compute a certificate’s signature.
// Here’s how to compute SHA256 hashes in Go.

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"

	// The pattern for generating a hash is sha256.New(), sha256.Write(bytes), then sha256.Sum([]byte{}). Here we start with a new hash.
	h := sha256.New()

	// Write expects bytes. If you have a string s, use []byte(s) to coerce it into bytes.
	h.Write([]byte(s))

	// This gets the finalized hash result as a byte slice. The argument to Sum can be used to append to an existing byte slice: it usually isn’t needed.
	bs := h.Sum(nil)

	// SHA256 values are often printed in hex, for example in git commits. Use the %x format verb to convert a hash results to a hex string.
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
