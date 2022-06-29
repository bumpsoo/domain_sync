package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)


func timeStamp() {
	dt := time.Now()
	temp := strings.Split(dt.String(), ".")[0]
	fmt.Println(temp)
}

func checkErr(err error) {
	if err != nil {
		timeStamp()
		panic(err)
	}
}

func aDomain() []ResourceRecord {
	url := "https://www.namesilo.com/api/dnsListRecords?version=1&type=xml&key=" + apiKey + "&domain=" + myDomain
	res, err := http.Get(url)
	checkErr(err)
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	checkErr(err)
	var name NameSilo
	err = xml.Unmarshal(data, &name)
	checkErr(err)
	var aDomains []ResourceRecord

	for _, resource := range name.Reply.ResourceRecord {
		if resource.Typee == "A" {
			aDomains = append(aDomains, resource)
		}
	}
	return aDomains
}

func ipAddr() string {
	resp, err := http.Get("https://ifconfig.me")
	checkErr(err)

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	return string(data)
}

func update(aDomains []ResourceRecord, ip string) {
	for _, resource := range aDomains {
		var urlString string
		if resource.Host == myDomain {
			urlString = "https://www.namesilo.com/api/dnsUpdateRecord?version=1&type=xml&key=" + apiKey + "&domain=" + myDomain + "&rrid=" + resource.RecordId + "&rrvalue=" + ip
		} else {
			urlString = "https://www.namesilo.com/api/dnsUpdateRecord?version=1&type=xml&key=" + apiKey + "&domain=" + myDomain + "&rrid=" + resource.RecordId + "&rrhost=" + strings.Split(resource.Host, "."+myDomain)[0] + "&rrvalue=" + ip
		}

		respo, err := http.Get(urlString)
		checkErr(err)
		defer respo.Body.Close()
		dataaa, err := ioutil.ReadAll(respo.Body)
		checkErr(err)
		var updatesilo NamesiloUpdate
		err = xml.Unmarshal(dataaa, &updatesilo)
		checkErr(err)
		timeStamp()
		fmt.Println(updatesilo)
	}
}
func main() {
	aDomains := aDomain()
	ip := ipAddr()

	update(aDomains, ip)
}
