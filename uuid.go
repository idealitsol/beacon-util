/*
 * @author    Emmanuel Kofi Bessah
 * @email     bekinsoft@gmail.com
 */

package util

import (
	"crypto/rand"
	"io"
	"net"
	"sync/atomic"
	"time"
)

// UUID ...
type UUID [16]byte

var hardwareAddr []byte
var clockSeq uint32

// const (
// 	VariantNCSCompat = 0
// 	VariantIETF      = 2
// 	VariantMicrosoft = 6
// 	VariantFuture    = 7
// )

func init() {
	if interfaces, err := net.Interfaces(); err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagLoopback == 0 && len(i.HardwareAddr) > 0 {
				hardwareAddr = i.HardwareAddr
				break
			}
		}
	}
	if hardwareAddr == nil {
		// If we failed to obtain the MAC address of the current computer,
		// we will use a randomly generated 6 byte sequence instead and set
		// the multicast bit as recommended in RFC 4122.
		hardwareAddr = make([]byte, 6)
		_, err := io.ReadFull(rand.Reader, hardwareAddr)
		if err != nil {
			panic(err)
		}
		hardwareAddr[0] = hardwareAddr[0] | 0x01
	}

	// initialize the clock sequence with a random number
	var clockSeqRand [2]byte
	io.ReadFull(rand.Reader, clockSeqRand[:])
	clockSeq = uint32(clockSeqRand[1])<<8 | uint32(clockSeqRand[0])
}

var timeBase = time.Date(1582, time.October, 15, 0, 0, 0, 0, time.UTC).Unix()

// TimeUUID generates a new time based UUID (version 1) using the current
// time as the timestamp.
func TimeUUID() UUID {
	return UUIDFromTime(time.Now())
}

// UUIDFromTime generates a new time based UUID (version 1) as described in
// RFC 4122. This UUID contains the MAC address of the node that generated
// the UUID, the given timestamp and a sequence number.
func UUIDFromTime(aTime time.Time) UUID {
	utcTime := aTime.In(time.UTC)
	t := int64(utcTime.Unix()-timeBase)*10000000 + int64(utcTime.Nanosecond()/100)
	clock := atomic.AddUint32(&clockSeq, 1)

	return TimeUUIDWith(t, clock, hardwareAddr)
}

// TimeUUIDWith generates a new time based UUID (version 1) as described in
// RFC4122 with given parameters. t is the number of 100's of nanoseconds
// since 15 Oct 1582 (60bits). clock is the number of clock sequence (14bits).
// node is a slice to gurarantee the uniqueness of the UUID (up to 6bytes).
// Note: calling this function does not increment the static clock sequence.
func TimeUUIDWith(t int64, clock uint32, node []byte) UUID {
	var u UUID

	u[0], u[1], u[2], u[3] = byte(t>>24), byte(t>>16), byte(t>>8), byte(t)
	u[4], u[5] = byte(t>>40), byte(t>>32)
	u[6], u[7] = byte(t>>56)&0x0F, byte(t>>48)

	u[8] = byte(clock >> 8)
	u[9] = byte(clock)

	copy(u[10:], node)

	u[6] |= 0x10 // set version to 1 (time based uuid)
	u[8] &= 0x3F // clear variant
	u[8] |= 0x80 // set to IETF variant

	return u
}

// String returns the UUID in it's canonical form, a 32 digit hexadecimal
// number in the form of xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx.
func (u UUID) String() string {
	var offsets = [...]int{0, 2, 4, 6, 9, 11, 14, 16, 19, 21, 24, 26, 28, 30, 32, 34}
	const hexString = "0123456789abcdef"
	r := make([]byte, 36)
	for i, b := range u {
		r[offsets[i]] = hexString[b>>4]
		r[offsets[i]+1] = hexString[b&0xF]
	}
	r[8] = '-'
	r[13] = '-'
	r[18] = '-'
	r[23] = '-'
	return string(r)

}
