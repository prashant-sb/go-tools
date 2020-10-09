package main

// CLI to calculate checksum of all files in given directory

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
    hasher "github.com/prashant-sb/go-utils/file_signatures/hash"
)

//
//  Options for CLI
//      dest: Destination dir / tmp will be default
//      sign: Checksum algorithm / md5 will be default
//
var (
	dest = flag.String("dest", "/tmp", "root direcory for calculate file hashes")
	sign = flag.String("sign", "md5", "Hashing algorithm")
)

// Worker thread for calculating checksum of file
// depending on algorithm provided by user

func checksumWorker(filePath string) error {
	var filehash func(filePath string) (string, error)

	switch *sign {

	case "crc":
		filehash = hasher.FileCrc32

	case "md5":
		filehash = hasher.FileMd5Sum

	case "sha256":
		filehash = hasher.FileSha256

	default:
		err := errors.New("Algorithm not supported.")
		return err
	}

	cs, err := filehash(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return err
	}

	fmt.Printf("%s :: %s\n", filePath, cs)
	return nil
}

// Callback for walking destination directory

func walkWith(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	go checksumWorker(path)

	return nil
}

func main() {
	flag.Parse()

	err := filepath.Walk(*dest, walkWith)
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
		return
	}
}
