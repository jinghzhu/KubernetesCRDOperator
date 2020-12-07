package utils

import (
	"encoding/base64"
	"io/ioutil"
)

// DecodeBase64ToFile accepts the string of base64 code and a file path. It decodes
// the string to that file. Return error if any issue.
func DecodeBase64ToFile(code, dest string) error {
	buff, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dest, buff, 07440)

	return err
}

// DecodeBase64FromStrToStr decodes input base64 string to a string.
func DecodeBase64FromStrToStr(code string) (string, error) {
	byteArr, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return "", err
	}

	return string(byteArr), nil
}

// DecodeBase64FromStrToBytes decodes input base64 string to bytes.
func DecodeBase64FromStrToBytes(code string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(code)
}

// DecodeBase64FromBytesToStr decodes input bytes to a string.
func DecodeBase64FromBytesToStr(code []byte) (string, error) {
	return DecodeBase64FromStrToStr(string(code))
}

// DecodeBase64FromBytesToBytes decodes input bytes to bytes.
func DecodeBase64FromBytesToBytes(src []byte) (dst []byte, err error) {
	_, err = base64.StdEncoding.Decode(dst, src)

	return dst, err
}

// EncodeBase64FromStrToStr encodes input string into string.
func EncodeBase64FromStrToStr(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// EncodeBase64FromStrToBytes encodes input string to bytes.
func EncodeBase64FromStrToBytes(src string) (dst []byte) {
	base64.StdEncoding.Encode(dst, []byte(src))

	return dst
}

// EncodeBase64FromBytesToStr encodes input bytes to string.
func EncodeBase64FromBytesToStr(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// EncodeBase64FromBytesToBytes endoces input bytes to bytes.
func EncodeBase64FromBytesToBytes(src []byte) (dst []byte) {
	base64.StdEncoding.Encode(dst, src)

	return dst
}

// EncodeBase64 accepts a file path. Return its base64 code and any error.
func EncodeBase64(path string) (string, error) {
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buff), nil
}
