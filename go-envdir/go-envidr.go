package go_envdir

import (
	"errors"
)

func GoEnvDir(path_to_env string, programm string) error {
	if path_to_env == "" {
		return errors.New("Не указан path_to_env")
	}
	if programm == "" {
		return errors.New("Не указан programm")
	}
	return nil
}
