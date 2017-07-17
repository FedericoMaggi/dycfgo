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

func TestArrayEncoding(t *testing.T) {
	testArray := []interface{}{
		"this", int64(563), "is", 5.3847654, "a",
		[]byte{0x2B, 0xF, 0x56, 0x2F, 0x4C},
		"test", "array", true, nil, int64(53),
		[]interface{}{"this", "is", "another",
			"array"},
	}

	encoded, err := Marshal(testArray)
	if err != nil {
		t.Fatalf("Unexpected error encoding array: %s.\n", err.Error())
	}
	decoded, err := Unmarshal(encoded)
	if err != nil {
		t.Fatalf("Unexpected error decoding array: %s.\n", err.Error())
	}
	if reflect.DeepEqual(decoded.([]interface{}), testArray) != true {
		t.Fatalf("Decoded array is different from original: \"%v\" != \"%v\".\n", decoded.([]interface{}), testArray)
	}
}

func TestStringArrayEncoding(t *testing.T) {
	testArray := []interface{}{
		"this", 563, "is", 5.3847654, "a",
		[]byte{0x2B, 0xF, 0x56, 0x2F, 0x4C},
		"test", "array", true, nil, int64(53),
		[]string{"this", "is", "another",
			"array"},
	}
	referencerray := []interface{}{
		"this", int64(563), "is", 5.3847654, "a",
		[]byte{0x2B, 0xF, 0x56, 0x2F, 0x4C},
		"test", "array", true, nil, int64(53),
		[]interface{}{"this", "is", "another",
			"array"},
	}

	encoded, err := Marshal(testArray)
	if err != nil {
		t.Fatalf("Unexpected error encoding array: %s.\n", err.Error())
	}
	decoded, err := Unmarshal(encoded)
	if err != nil {
		t.Fatalf("Unexpected error decoding array: %s.\n", err.Error())
	}
	if reflect.DeepEqual(decoded.([]interface{}), referencerray) != true {
		t.Fatalf("Decoded array is different from original: \"%v\" != \"%v\".\n", decoded.([]interface{}), referencerray)
	}
}

func TestIntegerArrayEncoding(t *testing.T) {
	testArray := []interface{}{
		"this", 563, "is", 5.3847654, "a",
		[]byte{0x2B, 0xF, 0x56, 0x2F, 0x4C},
		"test", "array", true, nil, int64(53),
		[]int64{5, 10, 67, 48576, 37362},
	}
	referencerray := []interface{}{
		"this", int64(563), "is", 5.3847654, "a",
		[]byte{0x2B, 0xF, 0x56, 0x2F, 0x4C},
		"test", "array", true, nil, int64(53),
		[]interface{}{int64(5), int64(10), int64(67),
			int64(48576), int64(37362)},
	}

	encoded, err := Marshal(testArray)
	if err != nil {
		t.Fatalf("Unexpected error encoding array: %s.\n", err.Error())
	}
	decoded, err := Unmarshal(encoded)
	if err != nil {
		t.Fatalf("Unexpected error decoding array: %s.\n", err.Error())
	}
	if reflect.DeepEqual(decoded.([]interface{}), referencerray) != true {
		t.Fatalf("Decoded array is different from original: \"%v\" != \"%v\".\n", decoded.([]interface{}), referencerray)
	}
}

func TestFloatArrayEncoding(t *testing.T) {
	testArray := []interface{}{
		"this", 563, "is", 5.3847654, "a",
		[]byte{0x2B, 0xF, 0x56, 0x2F, 0x4C},
		"test", "array", true, nil, int64(53),
		[]float64{5.182, 10.283, 67.8, 48576.2387746,
			37362.38374},
	}
	referencerray := []interface{}{
		"this", int64(563), "is", 5.3847654, "a",
		[]byte{0x2B, 0xF, 0x56, 0x2F, 0x4C},
		"test", "array", true, nil, int64(53),
		[]interface{}{5.182, 10.283, 67.8, 48576.2387746,
			37362.38374},
	}

	encoded, err := Marshal(testArray)
	if err != nil {
		t.Fatalf("Unexpected error encoding array: %s.\n", err.Error())
	}
	decoded, err := Unmarshal(encoded)
	if err != nil {
		t.Fatalf("Unexpected error decoding array: %s.\n", err.Error())
	}
	if reflect.DeepEqual(decoded.([]interface{}), referencerray) != true {
		t.Fatalf("Decoded array is different from original: \"%v\" != \"%v\".\n", decoded.([]interface{}), referencerray)
	}
}

func TestBoolArrayEncoding(t *testing.T) {
	testArray := []interface{}{
		"this", 563, "is", 5.3847654, "a",
		[]byte{0x2B, 0xF, 0x56, 0x2F, 0x4C},
		"test", "array", true, nil, int64(53),
		[]bool{true, false, true, true},
	}
	referencerray := []interface{}{
		"this", int64(563), "is", 5.3847654, "a",
		[]byte{0x2B, 0xF, 0x56, 0x2F, 0x4C},
		"test", "array", true, nil, int64(53),
		[]interface{}{true, false, true, true},
	}

	encoded, err := Marshal(testArray)
	if err != nil {
		t.Fatalf("Unexpected error encoding array: %s.\n", err.Error())
	}
	decoded, err := Unmarshal(encoded)
	if err != nil {
		t.Fatalf("Unexpected error decoding array: %s.\n", err.Error())
	}
	if reflect.DeepEqual(decoded.([]interface{}), referencerray) != true {
		t.Fatalf("Decoded array is different from original: \"%v\" != \"%v\".\n", decoded.([]interface{}), referencerray)
	}
}
