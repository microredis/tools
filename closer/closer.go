package closer

import (
	"io"
	"log"
)

func MustClose(resources ...io.Closer) {
	for index, resource := range resources {
		if err := resource.Close(); err != nil {
			log.Panicf("failed to close resource with index %d and error %v\n", index, err)
		}
	}
}

func TryClose(resources ...io.Closer) {
	var errClose error
	for index, resource := range resources {
		if err := resource.Close(); err != nil {
			errClose = err
			log.Printf("failed to close resource with index %d and error %v\n", index, err)
		}
	}
	if errClose != nil {
		log.Panicf("failed to close")
	}
}

func MayClose(resources ...io.Closer) {
	for index, resource := range resources {
		if err := resource.Close(); err != nil {
			log.Printf("failed to close resource with index %d and error %v\n", index, err)
		}
	}
}