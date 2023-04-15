package utils

import (
	"errors"
	"io/fs"
	"log"
	"os"
)

const BASE_SCHEME = `{
	"count": 0,
	"tables": []
}
`

func FileJSONTest(path []string) {
	for _, v := range path {
		f, err := os.Open(v)
		if errors.Is(err, fs.ErrNotExist) {
			f, err = os.Create(v)
			if err != nil {
				log.Panicf("cannot create: \"%s\" file\n", path)
			}

			_, err = f.WriteString(BASE_SCHEME)
			if err != nil {
				log.Panicln("cannot write to:", path)
			}
		}

		f.Close()
	}
}
