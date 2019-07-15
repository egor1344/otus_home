package go_envdir

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type typeMapEnv map[string]string

func GoEnvDir(pathToEnv string, programs string) error {
	if pathToEnv == "" {
		return errors.New("не указан path_to_env")
	}
	if programs == "" {
		return errors.New("не указан programs")
	}
	var mapEnv = make(typeMapEnv)
	err := openReadEnvDir(pathToEnv, &mapEnv)
	if err != nil {
		return err
	}
	err = setEnv(&mapEnv)
	if err != nil {
		return err
	}
	err = runProgram(programs)
	if err != nil {
		return err
	}
	return nil
}

func openReadEnvDir(pathToEnv string, mapEnv *typeMapEnv) error {
	err := filepath.Walk(pathToEnv, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		r, err := ioutil.ReadFile(path)
		if err != nil {
			return nil
		}
		(*mapEnv)[info.Name()] = string(r)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func setEnv(mapEnv *typeMapEnv) error {
	for key, value := range *mapEnv {
		err := os.Setenv(key, value)
		if err != nil {
			return nil
		}
	}
	return nil
}

func runProgram(program string) error {
	cmd := exec.Command(program)
	cmd.Env = append(os.Environ())
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	log.Println("вывод программы ", string(out))
	return nil
}
