package logger

import (
	"errors"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	openErr := errors.New("file not found")

	t.Run("file open fail", func(t *testing.T) {
		osOpenFile = func(name string, flag int, perm os.FileMode) (*os.File, error) {
			return nil, openErr
		}

		_, err := Init("", false)
		if !errors.Is(err, openErr) {
			t.Errorf("expected %s, got %s", openErr, err)
		}
	})

	t.Run("production success", func(t *testing.T) {
		osOpenFile = func(name string, flag int, perm os.FileMode) (*os.File, error) {
			return os.NewFile(1, ""), nil
		}

		_, err := Init("", true)
		if err != nil {
			t.Errorf("expected %s, got %s", "nil", err)
		}
	})

}
