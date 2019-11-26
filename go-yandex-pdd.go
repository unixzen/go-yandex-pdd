package yandex_pdd_api_go

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const baseURL string = "https://pddimp.yandex.ru/api2/admin/"

type Client struct {
	PddToken string
}

func AuthClient(pddToken string) *Client {
	return &Client{
		PddToken: pddToken,
	}
}

// Base method for making request
func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Add("PddToken", s.PddToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}

// Get records list
func (s *Client) GetRecordsList(domain string) ([]uint8, error) {
	data := url.Values{}
	data.Set("domain", domain)
	url := fmt.Sprint(baseURL, "dns/list?", data.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// Add dns record
func (s *Client) AddRecord(domain, recordType, subdomain, ttl, content string) ([]uint8, error) {
	data := url.Values{}
	data.Set("domain", domain)
	data.Set("type", recordType)
	data.Set("subdomain", subdomain)
	data.Set("ttl", ttl)
	data.Set("content", content)

	req, err := http.NewRequest("POST", fmt.Sprint(baseURL, "dns/add"), bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil

}

// Delete dns record
func (s *Client) DelRecord(domain, recordID string) ([]uint8, error) {
	data := url.Values{}
	data.Set("domain", domain)
	data.Set("record_id", recordID)

	req, err := http.NewRequest("POST", fmt.Sprint(baseURL, "dns/del"), bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// Delete domain
func (s *Client) DelDomain(domain string) ([]uint8, error) {
	data := url.Values{}
	data.Set("domain", domain)

	req, err := http.NewRequest("POST", fmt.Sprint(baseURL, "domain/delete"), bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

/*
Add domain
During register domain creating as domain of new organization.
*/

func (s *Client) RegisterDomain(domain string) ([]uint8, error) {
	data := url.Values{}
	data.Set("domain", domain)

	req, err := http.NewRequest("POST", fmt.Sprint(baseURL, "domain/register"), bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// Get status domain
func (s *Client) RegistrationStatus(domain string) ([]uint8, error) {
	data := url.Values{}
	data.Set("domain", domain)
	url := fmt.Sprint(baseURL, "domain/registration_status?", data.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// Get details domain
func (s *Client) DetailsDomain(domain string) ([]uint8, error) {
	data := url.Values{}
	data.Set("domain", domain)
	url := fmt.Sprint(baseURL, "domain/details?", data.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// Add email
func (s *Client) AddEmail(domain, login, password string) ([]uint8, error) {
	data := url.Values{}
	data.Set("domain", domain)
	data.Set("login", login)
	data.Set("password", password)

	req, err := http.NewRequest("POST", fmt.Sprint(baseURL, "email/add"), bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// Get domains list
func (s *Client) GetDomainsList() ([]uint8, error) {
	data := url.Values{}
	url := fmt.Sprint(baseURL, "domain/domains?", data.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// Get emails list
func (s *Client) GetEmailsList(domain string) ([]uint8, error) {
	data := url.Values{}
	data.Set("domain", domain)
	url := fmt.Sprint(baseURL, "email/list?", data.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// Delete email
func (s *Client) DelEmail(domain, login, uid string) ([]uint8, error) {
	data := url.Values{}
	data.Set("domain", domain)
	data.Set("login", login)
	data.Set("uid", uid)

	req, err := http.NewRequest("POST", fmt.Sprint(baseURL, "email/del"), bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// Get new emails at mail box
func (s *Client) GetEmailCounters(domain, login, uid string) ([]uint8, error) {
	data := url.Values{}
	data.Set("domain", domain)
	data.Set("login", login)
	data.Set("uid", uid)
	url := fmt.Sprint(baseURL, "email/counters?", data.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

/*
Get mailing list.
Management of Mailing list. At this not implemented methods GetStatusSubscriberMailingList (description of method
https://yandex.ru/dev/pdd/doc/reference/email-ml-getcansend-docpage/) and SetStatusSubscriberMailingList (description of method
https://yandex.ru/dev/pdd/doc/reference/email-ml-setcansend-docpage/), because option "can_send_on_behalf" always set at "no".
See comment method of SubscribeMailingList.
*/

func (s *Client) GetMailingList(domain string) ([]uint8, error) {
	data := url.Values{}
	data.Set("domain", domain)
	url := fmt.Sprint(baseURL, "email/ml/list?", data.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// Add mailing list
func (s *Client) AddMailingList(domain, mailList string) ([]uint8, error) {
	data := url.Values{}
	data.Set("domain", domain)
	data.Set("maillist", mailList)

	req, err := http.NewRequest("POST", fmt.Sprint(baseURL, "email/ml/add"), bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// Delete mailing list
func (s *Client) DelMailingList(domain, mailList string) ([]uint8, error) {
	data := url.Values{}
	data.Set("domain", domain)
	data.Set("maillist", mailList)

	req, err := http.NewRequest("POST", fmt.Sprint(baseURL, "email/ml/del"), bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

/*
Subscribe mailing list.
Function for subscribe to a specific email mailing list. Option "can_send_on_behalf" not included at that function because always set at "no".
*/

func (s *Client) SubscribeMailingList(domain, mailList, subscriber string) ([]uint8, error) {
	data := url.Values{}
	data.Set("domain", domain)
	data.Set("maillist", mailList)
	data.Set("subscriber", subscriber)

	req, err := http.NewRequest("POST", fmt.Sprint(baseURL, "email/ml/subscribe"), bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// Get list of subscribers
func (s *Client) GetListSubscribers(domain, mailList string) ([]uint8, error) {
	data := url.Values{}
	data.Set("domain", domain)
	data.Set("maillist", mailList)
	url := fmt.Sprint(baseURL, "email/ml/subscribers?", data.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// Unsubscribe from mailing list
func (s *Client) UnsubscribeMailingList(domain, mailList, subscriber string) ([]uint8, error) {
	data := url.Values{}
	data.Set("domain", domain)
	data.Set("maillist", mailList)
	data.Set("subscriber", subscriber)

	req, err := http.NewRequest("POST", fmt.Sprint(baseURL, "email/ml/unsubscribe"), bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

/* Add deputy.
Restriction. Only a user with an independent Yandex account can be assigned as a deputy administrator (an account of the form <login>@yandex.ru).)
*/
func (s *Client) AddDeputy(domain, deputyLogin string) ([]uint8, error) {
	data := url.Values{}
	data.Set("domain", domain)
	data.Set("login", deputyLogin)

	req, err := http.NewRequest("POST", fmt.Sprint(baseURL, "deputy/add"), bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil

}

// Get deputy list
func (s *Client) GetDeputyList(domain string) ([]uint8, error) {
	data := url.Values{}
	data.Set("domain", domain)
	url := fmt.Sprint(baseURL, "deputy/list?", data.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// Delete deputy
func (s *Client) DelDeputy(domain, deputyLogin string) ([]uint8, error) {
	data := url.Values{}
	data.Set("domain", domain)
	data.Set("login", deputyLogin)

	req, err := http.NewRequest("POST", fmt.Sprint(baseURL, "deputy/delete"), bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// Get DKIM status
func (s *Client) GetDkimStatus(domain string) ([]uint8, error) {
	data := url.Values{}
	data.Set("domain", domain)
	url := fmt.Sprint(baseURL, "dkim/status?", data.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// Enable DKIM
func (s *Client) EnableDkim(domain string) ([]uint8, error) {
	data := url.Values{}
	data.Set("domain", domain)

	req, err := http.NewRequest("POST", fmt.Sprint(baseURL, "dkim/enable"), bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// Disable DKIM
func (s *Client) DisableDkim(domain string) ([]uint8, error) {
	data := url.Values{}
	data.Set("domain", domain)

	req, err := http.NewRequest("POST", fmt.Sprint(baseURL, "dkim/disable"), bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}
