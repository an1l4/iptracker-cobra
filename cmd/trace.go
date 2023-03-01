package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP",
	Long:  `Trace the IP`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				showData(ip)
			}

		} else {
			fmt.Println("Please provide IP to trace")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

type IP struct {
	Ip       string `json:"ip"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Region   string `json:"region"`
	Timezone string `json:"timezone"`
	Loc      string `json:"loc"`
	Postal   string `json:"postal"`
}

func showData(ip string) {
	url := "http://ipinfo.io/" + ip + "/geo"

	responsebyte := getData(url)

	data := IP{}

	err := json.Unmarshal(responsebyte, &data)

	if err != nil {
		fmt.Println("unable to unmarshal the response")
	}

	c := color.New(color.FgRed).Add(color.Underline).Add(color.Bold)

	c.Println("Data Found")

	fmt.Printf("IP : %s \n City :%s \n Country : %s \n Region : %s \n Timezone : %s \n Loc : %s \n Postal : %s \n", data.Ip, data.City, data.Country, data.Region, data.Timezone, data.Loc, data.Postal)

}

func getData(url string) []byte {

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("unable to get response")
	}

	responseByte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("unable to read response")
	}

	return responseByte

}
