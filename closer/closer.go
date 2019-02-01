package closer

import "io"

func MustClose(list ...io.Closer) {
	for _, r := range list {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}
}

func TryClose(list ...io.Closer) {
	var err error
	for _, r := range list {
		err = r.Close()
	}
	panic(err)
}

func MayClose(list ...io.Closer) {
	for _, r := range list {
		_ = r.Close()
	}
}