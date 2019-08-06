package main

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	logger, err := getLogger()
	defer logger.Sync()
	if err != nil {
		logger.Fatal(err)
	}
	err = readConfig()
	if err != nil {
		logger.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		logger.Info(r.URL.Path)
	})
	addr := viper.GetString("web.site") + ":" + viper.GetString("web.port")
	logger.Info("Listen in ", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		logger.Fatal(err)
	}

}

func readConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}
	return nil
}

func getLogger() (*zap.SugaredLogger, error) {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	return sugar, nil
}
