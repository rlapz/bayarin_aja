package utils

import (
	"errors"
	"io/fs"
	"log"
	"os"
)

func FileTest(path []string) {
	for _, v := range path {
		f, err := os.Open(v)
		if errors.Is(err, fs.ErrNotExist) {
			f1, err := os.Create(v)
			if err != nil {
				log.Panicf("cannot create: \"%s\" file\n", path)
			}

			defer f1.Close()
		} else {
			f.Close()
		}
	}
}
