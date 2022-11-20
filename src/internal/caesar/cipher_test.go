package caesar

import (
	"testing"
)

type testCase struct {
	input  string
	shifts int
	result string
}

func TestCipherLatinEncrypt(t *testing.T) {

	cipher := NewCipher([]rune("abcdefghijklmnopqrstuvwxyz"))
	for _, tc := range []testCase{
		{
			input:  "Maksym",
			shifts: 1,
			result: "Nbltzn",
		},
		{
			input:  "Maksym",
			shifts: 27,
			result: "Nbltzn",
		},
		{
			input:  "Maksym",
			shifts: 0,
			result: "Maksym",
		},
		{
			input:  "Maksym!",
			shifts: -1,
			result: "Lzjrxl!",
		},
	} {
		result := cipher.Encrypt(tc.input, tc.shifts)
		if result != tc.result {
			t.Errorf("result not much, expected:%s, actual: %s; input: %s; shifts: %d", tc.result, result, tc.input, tc.shifts)
		}
	}
}

func TestCipherLatinEncryptDecrypt(t *testing.T) {
	cipher := NewCipher([]rune("abcdefghijklmnopqrstuvwxyz"))
	for _, tc := range []testCase{
		{
			input:  "Maksym",
			shifts: 1,
		},
		{
			input:  "132a4TESTNJ(UJ)&Y*G",
			shifts: 5,
		},
		{
			input:  "Maksym",
			shifts: 27,
		},
		{
			input:  "Maksym",
			shifts: 0,
		},
		{
			input:  "Maksym!",
			shifts: -1,
		},
	} {
		decryptedResult := cipher.Decrypt(cipher.Encrypt(tc.input, tc.shifts), tc.shifts)
		if decryptedResult != tc.input {
			t.Errorf("result not much, input: %s; shifts: %d", tc.input, tc.shifts)
		}
	}
}

func TestCipherCyrillicEncryptDecrypt(t *testing.T) {
	cipher := NewCipher([]rune("абвгдеёжзийклмнопрстуфхцчшщъыьэюя"))
	for _, tc := range []testCase{
		{
			input:  "Максим",
			shifts: 3,
		},
	} {
		decryptedResult := cipher.Decrypt(cipher.Encrypt(tc.input, tc.shifts), tc.shifts)
		if decryptedResult != tc.input {
			t.Errorf("result not much, input: %s; shifts: %d", tc.input, tc.shifts)
		}
	}
}
