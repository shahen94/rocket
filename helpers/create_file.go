package helpers

import "os"

func CreateFileIfNotExists(path string) error {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return err
		}
		defer func(file *os.File) {
			file.Close()
		}(file)
	}
	return nil
}
