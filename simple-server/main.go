package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const (
	ENCRYPT_MODE_DISABLED = "disabled"
	ENCRYPT_MODE_ENABLED  = "enabled"
	ENCRYPT_MODE_BOTH     = "both"
)

var (
	quit = make(chan os.Signal, 5)
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

	// init gin router
	gin.SetMode(viper.GetString("http.mode"))
	router := gin.Default()
	router.Any("", func(c *gin.Context) {
		s, _ := ioutil.ReadAll(c.Request.Body)
		c.JSON(200, gin.H{
			"params":                    c.Params,
			"keys":                      c.Keys,
			"accepted":                  c.Accepted,
			"request.body":              fmt.Sprintf("%s", s),
			"request.method":            c.Request.Method,
			"request.url":               c.Request.URL,
			"request.proto":             c.Request.Proto,
			"request.proto.major":       c.Request.ProtoMajor,
			"request.proto.minor":       c.Request.ProtoMinor,
			"request.header":            c.Request.Header,
			"request.content.length":    c.Request.ContentLength,
			"request.transfer.encoding": c.Request.TransferEncoding,
			"request.close":             c.Request.Close,
			"request.host":              c.Request.Host,
			"request.form":              c.Request.Form,
			"request.post.form":         c.Request.PostForm,
			"request.multipart.form":    c.Request.MultipartForm,
			"request.trailer":           c.Request.Trailer,
			"request.remote.addr":       c.Request.RemoteAddr,
			"request.request.uri":       c.Request.RequestURI,
			"request.tls":               c.Request.TLS,
		})
	})

	switch viper.GetString("http.encrypt_mode") {
	case ENCRYPT_MODE_DISABLED:
		go func() {
			router.Run(":" + viper.GetString("http.port"))
		}()
	case ENCRYPT_MODE_ENABLED:
		go func() {
			router.RunTLS(":"+viper.GetString("http.tls_port"), viper.GetString("http.cert"), viper.GetString("http.key"))
		}()
	case ENCRYPT_MODE_BOTH:
		go func() {
			router.Run(":" + viper.GetString("http.port"))
		}()
		go func() {
			router.RunTLS(":"+viper.GetString("http.tls_port"), viper.GetString("http.cert"), viper.GetString("http.key"))
		}()
	default:
		log.Panicf("unknown http.encrypt: %s", viper.GetString("http.encrypt_mode"))
	}
	log.Println("http(s) server(s) are running")

	// graeful shutdown
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-quit
	log.Printf("receive exit signal: %+v", sig)
}