package main

import (
	"crypto/tls"
	"fmt"
	"n00bctl/tui"
	"net/http"
	"os"

	"github.com/luthermonson/go-proxmox"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/") // config file path
	viper.AutomaticEnv()             // read value ENV variable

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	insecureHTTPClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	tokenID := viper.GetString("proxmox_api.id")
	secret := viper.GetString("proxmox_api.secret")

	client := proxmox.NewClient("https://10.0.0.11:8006/api2/json",
		proxmox.WithClient(&insecureHTTPClient),
		proxmox.WithAPIToken(tokenID, secret),
	)

	version, err := client.Version()
	if err != nil {
		panic(err)
	}
	fmt.Println(version.Version)
	cluster, _ := client.Cluster()
	// fmt.Printf("%+v", cluster)
	clusterRes, _ := cluster.Resources()

	for _, node := range clusterRes {
		fmt.Printf("\n%+v", node)
	}
	// cli.Execute()

	tui.Init(client)
}
