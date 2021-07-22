package helper

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

/**
 * @description:
 * @param {type}
 * @return:
 */
func DesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

/**
 * @description:
 * @param {type}
 * @return:
 */
func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

/**
 * @description:
 * @param {type}
 * @return:
 */
func DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

/**
 * @description:
 * @param {[]byte}
 * @return: []byte
 */
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
