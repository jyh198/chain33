package common

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

const (
	HashLength = 32
)

type Hash [HashLength]byte

func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetBytes(b)
	return h
}

func StringToHash(s string) Hash { return BytesToHash([]byte(s)) }
func BigToHash(b *big.Int) Hash  { return BytesToHash(b.Bytes()) }
func HexToHash(s string) Hash    { return BytesToHash(FromHex(s)) }

// Get the string representation of the underlying hash
func (h Hash) Str() string   { return string(h[:]) }
func (h Hash) Bytes() []byte { return h[:] }
func (h Hash) Hex() string   { return hexEncode(h[:]) }

// TerminalString implements log.TerminalStringer, formatting a string for console
// output during logging.
func (h Hash) TerminalString() string {
	return fmt.Sprintf("%x…%x", h[:3], h[29:])
}

// String implements the stringer interface and is used also by the logger when
// doing full logging into a file.
func (h Hash) String() string {
	return h.Hex()
}

// Format implements fmt.Formatter, forcing the byte slice to be formatted as is,
// without going through the stringer interface used for logging.
func (h Hash) Format(s fmt.State, c rune) {
	fmt.Fprintf(s, "%"+string(c), h[:])
}

// Sets the hash to the value of b. If b is larger than len(h), 'b' will be cropped (from the left).
func (h *Hash) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-HashLength:]
	}

	copy(h[HashLength-len(b):], b)
}

// Set string `s` to h. If s is larger than len(h) s will be cropped (from left) to fit.
func (h *Hash) SetString(s string) { h.SetBytes([]byte(s)) }

// Sets h to other
func (h *Hash) Set(other Hash) {
	for i, v := range other {
		h[i] = v
	}
}

func EmptyHash(h Hash) bool {
	return h == Hash{}
}

func hexEncode(b []byte) string {
	enc := make([]byte, len(b)*2+2)
	copy(enc, "0x")
	hex.Encode(enc[2:], b)
	return string(enc)
}

func ToHex(b []byte) string {
	hex := Bytes2Hex(b)
	// Prefer output of "0x0" instead of "0x"
	if len(hex) == 0 {
		hex = "0"
	}
	return "0x" + hex
}

func FromHex(s string) []byte {
	if len(s) > 1 {
		if s[0:2] == "0x" || s[0:2] == "0X" {
			s = s[2:]
		}
		if len(s)%2 == 1 {
			s = "0" + s
		}
		return Hex2Bytes(s)
	}
	return nil
}

// Copy bytes
//
// Returns an exact copy of the provided bytes
func CopyBytes(b []byte) (copiedBytes []byte) {
	if b == nil {
		return nil
	}
	copiedBytes = make([]byte, len(b))
	copy(copiedBytes, b)

	return
}

func HasHexPrefix(str string) bool {
	l := len(str)
	return l >= 2 && str[0:2] == "0x"
}

func IsHex(str string) bool {
	l := len(str)
	return l >= 4 && l%2 == 0 && str[0:2] == "0x"
}

func Bytes2Hex(d []byte) string {
	return hex.EncodeToString(d)
}

func Sha256(b []byte) Hash {
	return sha256.Sum256(b)
}

func Hex2Bytes(str string) []byte {
	h, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return h
}
