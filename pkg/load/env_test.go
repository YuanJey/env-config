package load

import (
	"os"
	"testing"
)

func Test_loadEnv(t *testing.T) {
	os.Setenv("HOST", "localhost")
	os.Setenv("PORT", "8080")
	os.Setenv("ADDRS", "localhost,127.0.0.1")
	type config struct {
		Host  string   `json:"host" env:"HOST" def:"127.0.0.1"`
		Port  int      `json:"port" env:"PORT" def:"8080"`
		Addrs []string `json:"addrs" env:"ADDRS" def:"localhost,127.0.0.1"`
	}
	var c config
	if err := LoadEnv(&c); err != nil {
		t.Fatalf("failed to load env: %v", err)
	}
	t.Logf("%+v", c)
}

func Test_loadEnv_Default(t *testing.T) {
	type config struct {
		Host string `json:"host" env:"HOST_NOT_EXIST" def:"localhost"`
		Port int    `json:"port" env:"PORT_NOT_EXIST" def:"8080"`
	}
	var c config
	if err := LoadEnv(&c); err != nil {
		t.Fatalf("failed to load env: %v", err)
	}
	if c.Host != "localhost" {
		t.Errorf("expected host to be localhost, got %s", c.Host)
	}
	if c.Port != 8080 {
		t.Errorf("expected port to be 8080, got %d", c.Port)
	}
}
