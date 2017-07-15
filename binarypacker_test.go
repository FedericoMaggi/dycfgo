//
// packager crypto package
// Author: Guido Ronchetti <guido.ronchetti@nexo.cloud>
// v1.0 14/07/2017
//

package binarypacker

import (
	"reflect"
	"testing"
)

func TestStringEncoding(t *testing.T) {
	testString := "This is a test complete string."

	encoded, err := Marshal(testString)
	if err != nil {
		t.Fatalf("Unexpected error encoding string: %s.\n", err.Error())
	}
	decoded, err := Unmarshal(encoded)
	if err != nil {
		t.Fatalf("Unexpected error decoding string: %s.\n", err.Error())
	}
	if decoded.(string) != testString {
		t.Fatalf("Decoded string is different from original: \"%s\" != \"%s\".\n", decoded.(string), testString)
	}
}

func TestFloatEncoding(t *testing.T) {
	testFloat := 14538273.374625374992387625102

	encoded, err := Marshal(testFloat)
	if err != nil {
		t.Fatalf("Unexpected error encoding float: %s.\n", err.Error())
	}
	decoded, err := Unmarshal(encoded)
	if err != nil {
		t.Fatalf("Unexpected error decoding float: %s.\n", err.Error())
	}
	if decoded.(float64) != testFloat {
		t.Fatalf("Decoded float is different from original: \"%f\" != \"%f\".\n", decoded.(float64), testFloat)
	}
}

func TestIntegerEncoding(t *testing.T) {
	testInt := int64(-1453827337462537497)

	encoded, err := Marshal(testInt)
	if err != nil {
		t.Fatalf("Unexpected error encoding int: %s.\n", err.Error())
	}
	decoded, err := Unmarshal(encoded)
	if err != nil {
		t.Fatalf("Unexpected error decoding int: %s.\n", err.Error())
	}
	if decoded.(int64) != testInt {
		t.Fatalf("Decoded int is different from original: \"%d\" != \"%d\".\n", decoded.(int64), testInt)
	}
}

func TestDataEncoding(t *testing.T) {
	testData := []byte{0x2C, 0x4B, 0x7, 0x1, 0xA3, 0x2, 0xF, 0x1F,
		0x13, 0xA, 0x5A, 0x11, 0x2B, 0x4, 0x7, 0x1,
		0x2C, 0x4B, 0x7, 0x1, 0xA3, 0x2, 0xF, 0x1F,
		0x13, 0xA, 0x5A, 0x11, 0x2B, 0x4, 0x7, 0x1,
		0x2C, 0x4B, 0x7, 0x1, 0xA3, 0x2, 0xF, 0x1F,
		0x13, 0xA, 0x5A, 0x11, 0x2B, 0x4, 0x7, 0x1,
		0x2C, 0x4B, 0x7, 0x1, 0xA3, 0x2, 0xF, 0x1F,
		0x13, 0xA, 0x5A, 0x11, 0x2B, 0x4, 0x7, 0x1,
	}

	encoded, err := Marshal(testData)
	if err != nil {
		t.Fatalf("Unexpected error encoding data: %s.\n", err.Error())
	}
	decoded, err := Unmarshal(encoded)
	if err != nil {
		t.Fatalf("Unexpected error decoding data: %s.\n", err.Error())
	}
	if reflect.DeepEqual(decoded.([]byte), testData) != true {
		t.Fatalf("Decoded data is different from original: \"%v\" != \"%v\".\n", decoded.([]byte), testData)
	}
}
