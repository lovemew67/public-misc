package main

import (
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
)

const (
	ENCRYPT_MODE_DISABLED = "disabled"
	ENCRYPT_MODE_ENABLED  = "enabled"
	ENCRYPT_MODE_BOTH     = "both"
)

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds)
}

func main() {
	// init viper
	viper.AutomaticEnv()
	viper.SetConfigFile("./local.toml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("viper read file error: %+v", err)
	}

	// run ran server
	switch viper.GetString("http.encrypt_mode") {
	case ENCRYPT_MODE_DISABLED:
		cmd := exec.Command(
			viper.GetString("http.ran_binary"),
			"-r", viper.GetString("http.web_root"),
			"-p", viper.GetString("http.port"),
		)
		cmd.CombinedOutput()
	case ENCRYPT_MODE_ENABLED:
		cmd := exec.Command(
			viper.GetString("http.ran_binary"),
			"-r", viper.GetString("http.web_root"),
			"-cert", viper.GetString("http.cert"),
			"-key", viper.GetString("http.key"),
			"-tls-port", viper.GetString("http.tls_port"),
			"-tls-policy", "only",
		)
		cmd.CombinedOutput()
	case ENCRYPT_MODE_BOTH:
		cmd := exec.Command(
			viper.GetString("http.ran_binary"),
			"-r", viper.GetString("http.web_root"),
			"-p", viper.GetString("http.port"),
			"-cert", viper.GetString("http.cert"),
			"-key", viper.GetString("http.key"),
			"-tls-port", viper.GetString("http.tls_port"),
			"-tls-policy", "both",
		)
		cmd.CombinedOutput()
	default:
		log.Panicf("unknown http.encrypt: %s", viper.GetString("http.encrypt_mode"))
	}
}
