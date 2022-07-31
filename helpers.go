package rocket

import "os"

func (r *Rocket) CreateDirIfNotExists(path string) error {
	const mode = 0755

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, mode)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Rocket) CreateFileIfNotExists(path string) error {
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
