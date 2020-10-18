package hash

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"hash/crc32"
	"io"
	"os"
)

// Polynomial seed for CRC calculation.
const polynomial = 0xedb88320

// Calculates md5sum of file.
// returns checksum or error
func FileMd5Sum(filePath string) (string, error) {
	var md5sum string

	file, err := os.Open(filePath)
	if err != nil {
		return md5sum, err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return md5sum, err
	}

	hashInBytes := hash.Sum(nil)[:16]
	md5sum = hex.EncodeToString(hashInBytes)

	return md5sum, nil
}

// Calculates sha256 of file.
// returns checksum or error
func FileSha256(filePath string) (string, error) {
	var shaCheckSum string

	file, err := os.Open(filePath)
	if err != nil {
		return shaCheckSum, err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return shaCheckSum, err
	}
	hashInBytes := hash.Sum(nil)[:32]
	shaCheckSum = hex.EncodeToString(hashInBytes)

	return shaCheckSum, nil
}

// Calculates the CRC of file, returns
// checksum or error
func FileCrc32(filePath string) (string, error) {
	var crcCheckSum string

	file, err := os.Open(filePath)
	if err != nil {
		return crcCheckSum, err
	}
	defer file.Close()

	tablePolynomial := crc32.MakeTable(polynomial)
	hash := crc32.New(tablePolynomial)
	if _, err := io.Copy(hash, file); err != nil {
		return crcCheckSum, err
	}

	hashInBytes := hash.Sum(nil)[:16]
	crcCheckSum = hex.EncodeToString(hashInBytes)

	return crcCheckSum, nil
}
