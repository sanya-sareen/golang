package main

import (
	"fmt"
	"bufio"
	"strings"
	"log"
	"net"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMx, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")

	for scanner.Scan(){
		validateDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil{
		log.Fatal("Error: could not read from input: %v\n", err)
	}


}

func validateDomain(domainName string){
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domainName)

	if err != nil{
		log.Printf("error: %v\n", err)

	}

	if len(mxRecords) > 0{
		hasMX = true
	}

	txtRecords, err:= net.LookupTXT(domainName)
	if err != nil{
		log.Printf("Error:%v\n", err)
	}
	for _, record := range txtRecords{
		if strings.HasPrefix(record, "v=spf1"){
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc."+ domainName)
	if err != nil {
		log.Printf("Error%v\n", err)
	}

	for _, record:= range dmarcRecords{
		if strings.HasPrefix(record, "v=DMARC1"){
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v", domainName, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

}