package main

import "encoding/xml"

type NameSilo struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

type Request struct {
	XMLName   xml.Name `xml:"request"`
	Operation string   `xml:"operation"`
	Ip        string   `xml:"ip"`
}

type Reply struct {
	XMLName        xml.Name         `xml:"reply"`
	Code           string           `xml:"code"`
	Detail         string           `xml:"detail"`
	ResourceRecord []ResourceRecord `xml:"resource_record"`
}
type ResourceRecord struct {
	XMLName  xml.Name `xml:"resource_record"`
	RecordId string   `xml:"record_id"`
	Typee    string   `xml:"type"`
	Host     string   `xml:"host"`
	Value    string   `xml:"value"`
	Ttl      string   `xml:"ttl"`
	Distance string   `xml:"distance"`
}

type NamesiloUpdate struct {
	XMLNAME xml.Name      `xml:"namesilo"`
	Request RequestUpdate `xml:"request"`
	Reply   ReplyUpdate   `xml:"reply"`
}
type RequestUpdate struct {
	Operation string `xml:"operation"`
	Ip        string `xml:"ip"`
}
type ReplyUpdate struct {
	Code     string `xml:"code"`
	Detail   string `xml:"detail"`
	RecordId string `xml:"record_id"`
}
