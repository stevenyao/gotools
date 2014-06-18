package main

import (
	"fmt"
	"os"
	"flag"
	"crypto/md5"
	"crypto/sha1"
	"hash"
	)

var filename = flag.String("i", "", "target file name")
var hashtype = flag.String("t", "", "algorithm")
func main() {
	flag.Parse()

	if flag.NFlag() < 2 {
		flag.PrintDefaults()
		return;
	}

	f, err := os.OpenFile(*filename, os.O_RDONLY, 0666)
	if err != nil {
		panic("failed to open " + *filename)
	}
	defer f.Close()

	buf := make([]byte, 1024 * 64)
	h := func(t string) hash.Hash{
			if t == "md5" {
				return md5.New()
			} else if t == "sha1" {
				return sha1.New()
			} else {
				return nil
			}
		}(*hashtype)

	for{
		count, err := f.Read(buf)
		if err != nil {
			break
		}

		h.Write(buf[0:count])
	}

	fi, _:= f.Stat()
	fmt.Printf("Name: %s, Size: %d, hash: %x\n", fi.Name(), fi.Size(), h.Sum(nil))
}