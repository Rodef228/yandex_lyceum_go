package ex5

import (
    "errors"
    "testing"
)

func TestGetUTFLength(t *testing.T) {
    tests := []struct {
        input    []byte
        expected int
        err      error
    }{
        {[]byte{0xff, 0xfe, 0xfd}, 0, ErrInvalidUTF8},
        {[]byte(""), 0, nil},
        {[]byte{0xc2}, 0, ErrInvalidUTF8},
        {[]byte{0xe2, 0x82}, 0, ErrInvalidUTF8},
        {[]byte{0xf0, 0x9f, 0x92}, 0, ErrInvalidUTF8},
        {[]byte{0x80}, 0, ErrInvalidUTF8},
        {[]byte{0xf8, 0x88, 0x80, 0x80}, 0, ErrInvalidUTF8},
        {[]byte{0xc1, 0x81}, 0, ErrInvalidUTF8},
        {[]byte{0x80}, 0, ErrInvalidUTF8},
        {[]byte{0xc0, 0x80}, 0, ErrInvalidUTF8},
        {[]byte{0xed, 0xa0, 0x80}, 0, ErrInvalidUTF8},
        {[]byte{0xed, 0xa0, 0x81}, 0, ErrInvalidUTF8},
        {[]byte{0xed, 0xaf, 0xbf}, 0, ErrInvalidUTF8},
        {[]byte{0xed, 0xb0, 0x80}, 0, ErrInvalidUTF8},
        {[]byte{0xed, 0xb0, 0x81}, 0, ErrInvalidUTF8},
        {[]byte{0xed, 0xbf, 0xbf}, 0, ErrInvalidUTF8},
    }

    for _, tt := range tests {
        result, err := GetUTFLength(tt.input)
        if result != tt.expected || !errors.Is(err, tt.err) {
            t.Errorf("For input %v, expected length %d and error %v, got length %d and error %v",
                tt.input, tt.expected, tt.err, result, err)
        }
    }
}
