package closer

import (
	"io"
	"log"
)

const tmpl = "failed to close resource with index %d and error %v\n"

func MustClose(resources ...io.Closer) {
	for index, resource := range resources {
		if err := resource.Close(); err != nil {
			log.Panicf(tmpl, index, err)
		}
	}
}

func TryClose(resources ...io.Closer) {
	var errClose error
	for index, resource := range resources {
		if err := resource.Close(); err != nil {
			errClose = err
			log.Printf(tmpl, index, err)
		}
	}
	if errClose != nil {
		panic("failed to close")
	}
}

func MayClose(resources ...io.Closer) {
	for index, resource := range resources {
		if err := resource.Close(); err != nil {
			log.Printf(tmpl, index, err)
		}
	}
}
