package rocket

import "fmt"

func (r *Rocket) checkDotEnv(path string) error {
	err := r.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))

	if err != nil {
		return err
	}

	return nil
}
