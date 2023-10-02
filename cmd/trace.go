package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(traceCmd)
}

var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP",
	Long:  `Trace the IP.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				showData(ip)
			}
		} else {
			fmt.Println("Provide IP!!")
		}
	},
}

type IP struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

func showData(ip string) {
	url := "https://ipinfo.io/" + ip + "/geo"
	respbody := getData(url)

	data := IP{}

	err := json.Unmarshal(respbody, &data)
	if err != nil {
		log.Fatal("Error unmarshalling the respbody")
	}

	c := color.New(color.FgRed).Add(color.Underline).Add(color.FgRed)

	c.Println("Ip provided : ")

	fmt.Printf(" IP : %s,\n City : %s,\n Region : %s,\n Country : %s,\n Loc : %s,\n Postal : %s,\n Timezone : %s,\n\n", data.IP, data.City, data.Region, data.Country, data.Loc, data.Postal, data.Timezone)
}

func getData(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Unable to get IP info")
	}

	responseByte, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Unable to read IP info")
	}

	return responseByte
}
