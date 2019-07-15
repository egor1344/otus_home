package go_envdir

import (
	"testing"
)

func TestGoEnvDir2(t *testing.T) {
	err := GoEnvDir("", "test_programm")
	if err == nil {
		t.Error("Wrong result ", err)
	}
	err = GoEnvDir("test_path", "")
	if err == nil {
		t.Error("Wrong result ")
	}
	err = GoEnvDir("/home/egor/GolandProjects/otus_home/go-envdir/test_env", "env")
	if err != nil {
		t.Error("Wrong result ", err)
	}
}
