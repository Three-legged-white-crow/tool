package main

import (
	"crypto/md5"
	"errors"
	"flag"
	"io"
	"log"
	"os"
)

const (
	OK = iota
	ErrFlagParse
	ErrChecksumAlgorithm
	ErrChecksumGenerate
	ErrChecksumCompare
	ErrChecksumFile

	readBuf = 1024

	md5Algorithm = "md5"
	md5Suffix    = ".md5"
)

var ErrWriteChecksumContent = errors.New("failed to write full content")

func main() {
	srcFilePath := flag.String("src", "", "src file abs path")
	destFilePath := flag.String("dest", "", "dest file abs path")
	checksumAlgorithm := flag.String("checksum", "md5", "checksum algorithm")
	flag.Parse()

	if len(*srcFilePath) == 0 {
		log.Println("[Checksum]Not specify src file")
		os.Exit(ErrFlagParse)
	}

	if len(*destFilePath) == 0 {
		log.Println("[Checksum]Not specify dest file")
		os.Exit(ErrFlagParse)
	}

	if *checksumAlgorithm != md5Algorithm {
		log.Println("[Checksum]Only support md5 at now")
		os.Exit(ErrChecksumAlgorithm)
	}

	srcChecksum, err := md5Checksum(*srcFilePath)
	if err != nil {
		log.Println("[Checksum]Failed to generate checksum of src file, err:", err.Error())
		os.Exit(ErrChecksumGenerate)
	}

	destChecksum, err := md5Checksum(*destFilePath)
	if err != nil {
		log.Println("[Checksum]Failed to generate checksum of dest file, err:", err.Error())
		os.Exit(ErrChecksumGenerate)
	}

	isEqual := compare(srcChecksum, destChecksum)
	if !isEqual {
		log.Println("[Checksum]Result of checksum with src file and dest file is not equal!!!")
		os.Exit(ErrChecksumCompare)
	}

	log.Println("[Checksum]Result of checksum with src file and dest file is equal...")

	md5DestFilePath := *destFilePath + md5Suffix
	err = md5File(md5DestFilePath, destChecksum)
	if err != nil {
		log.Println("[Checksum]Failed to write checksum file, err:", err.Error())
		os.Exit(ErrChecksumFile)
	}

	log.Println("[Checksum]Succeed to write checksum file:", md5DestFilePath)
	os.Exit(OK)
}

func md5Checksum(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	h := md5.New()
	buf := make([]byte, readBuf)

	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}

		h.Write(buf[:n])
	}

	return h.Sum(nil), nil

}

func compare(src, dest []byte) bool {
	if len(src) != len(dest) {
		return false
	}

	for i, v := range src {
		if v != dest[i] {
			return false
		}
	}

	return true
}

func md5File(path string, res []byte) error {
	// todo: check file is exists? if exists create a new file?
	l := len(res)

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	n, err := f.Write(res)
	if err != nil {
		return err
	}

	if n != l {
		return ErrWriteChecksumContent
	}

	return nil
}
