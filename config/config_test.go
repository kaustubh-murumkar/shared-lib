package config

import "testing"

func TestInit(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		osReadFile = func(name string) ([]byte, error) {
			return []byte(""), nil
		}
		_, err := Init("", JSON, false)
		if err != nil {
			t.Errorf("expected nil, got error %s", err.Error())
		}
	})
}
