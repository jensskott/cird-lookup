package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {

	region := flag.String("r", "", "region for amazon")
	tag := flag.String("t", "*", "what pattern to search for in name")
	flag.Parse()
	if *region == "" {
		log.Fatal("Region needs to be specified")
	}

	var nextOctet string
	var cidrList []string
	var octet []int

	sess := session.New(&aws.Config{
		Region: aws.String(*region),

		MaxRetries: aws.Int(5),
	})

	svc := ec2.New(sess)
	input := &ec2.DescribeVpcsInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{
				Name: aws.String("tag:Name"),
				Values: []*string{
					aws.String(*tag),
				},
			},
		},
	}

	result, err := svc.DescribeVpcs(input)
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range result.Vpcs {
		cidrList = append(cidrList, *i.CidrBlock)
	}

	for _, o := range cidrList {
		octet = append(octet, convertStrToInt(o))
	}

	if octet == nil {
		nextOctet = "28"
	} else {
		sort.Ints(octet)
		t := strconv.Itoa(octet[len(octet)-1] + 1)
		nextOctet = t
	}

	fmt.Printf(fmt.Sprintf("10.%s.0.0/16", nextOctet))
}

func convertStrToInt(o string) int {
	x := strings.Split(o, ".")
	i, _ := strconv.Atoi(x[1])
	return i
}
