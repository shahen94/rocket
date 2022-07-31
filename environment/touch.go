package environment

import (
	"fmt"

	"github.com/shahen94/rocket/helpers"
)

func Touch(path string) error {
	err := helpers.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))

	if err != nil {
		return err
	}

	return nil
}
