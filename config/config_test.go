package config

import "testing"

func TestInit(t *testing.T) {
	t.Run("Should open an .env file ", func(t *testing.T) {
		err := Init("../.env.example")
		if err != nil && err != ErrDb {
			t.Errorf("Error on loading file: %s", err)
		}
	})
	t.Run("Should return an database error", func(t *testing.T) {
		err := Init("../.env.example")
		if err != ErrDb {
			t.Errorf("Error: error loading database expected, got : %s ", err)
		}
	})
	t.Run("Should occur an error", func(t *testing.T) {
		err := Init()
		if err == nil {
			t.Errorf("Should not open any file!")
		}
	})
}
