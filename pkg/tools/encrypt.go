package tools

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
)

// HashAndSalt 加密密码
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {

	}
	return string(hash)
}

// ValidatePasswords 验证密码 false密码不一致 true密码一致
func ValidatePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}

// AesDecrypt AES加密
func AesDecrypt(cryted, key string) ([]byte, error) {
	//使用RawURLEncoding 不要使用StdEncoding
	//不要使用StdEncoding  放在url参数中回导致错误
	crytedByte, err := base64.StdEncoding.DecodeString(cryted)
	if err != nil {
		return nil, err
	}
	k := []byte(key)

	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	return orig, nil
}
