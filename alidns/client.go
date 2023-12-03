package alidns

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/cert-manager/cert-manager/pkg/issuer/acme/dns/util"
)

type Client struct {
	dnsc *alidns.Client
}

func newClient(region string, cred auth.Credential) (*Client, error) {
	cfg := sdk.NewConfig()
	client, err := alidns.NewClientWithOptions(region, cfg, cred)
	if err != nil {
		return nil, err
	}

	return &Client{dnsc: client}, nil
}

func (c *Client) getHostedZone(zone string) (string, error) {
	request := alidns.CreateDescribeDomainsRequest()
	request.KeyWord = util.UnFqdn(zone)
	request.SearchMode = "EXACT"

	response, err := c.dnsc.DescribeDomains(request)
	if err != nil {
		return "", err
	}

	zones := response.Domains.Domain
	if len(zones) == 0 {
		return "", fmt.Errorf("zone %s does not exist", zone)
	}

	return zones[0].DomainName, nil
}

func (c *Client) addTxtRecord(zone, rr, value string) error {
	record := c.newTxtRecord(zone, rr, value)
	_, err := c.dnsc.AddDomainRecord(record)
	return err
}

const recordTypeTxt = "TXT"

func (c *Client) newTxtRecord(zone, rr, value string) *alidns.AddDomainRecordRequest {
	request := alidns.CreateAddDomainRecordRequest()
	request.Type = recordTypeTxt
	request.DomainName = zone
	request.RR = rr
	request.Value = value
	return request
}

func (c *Client) getTxtRecord(zone, rr string) (*alidns.Record, error) {
	request := alidns.CreateDescribeDomainRecordsRequest()
	request.Type = recordTypeTxt
	request.DomainName = zone
	request.RRKeyWord = rr

	response, err := c.dnsc.DescribeDomainRecords(request)
	if err != nil {
		return nil, err
	}

	records := response.DomainRecords.Record
	for _, r := range records {
		if r.RR == rr {
			return &r, nil
		}
	}

	return nil, fmt.Errorf("txt record does not exist: %v.%v", rr, zone)
}

func (c *Client) deleteDomainRecord(id string) error {
	request := alidns.CreateDeleteDomainRecordRequest()
	request.RecordId = id

	_, err := c.dnsc.DeleteDomainRecord(request)
	return err
}
