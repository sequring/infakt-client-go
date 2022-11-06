package infakt

import "net/http"

type InFaktClient struct {
	InfaktEndpoint string
	HTTPClient     *http.Client
	Token          string
	AuthHeader     string
}

type MetaInfo struct {
	Count      int    `json:"count"`
	TotalCount int    `json:"total_count"`
	Next       string `json:"next"`
	Previous   string `json:"previous"`
}

type Client struct {
	ID                   int    `json:"id"`
	CompanyName          string `json:"company_name"`
	Street               string `json:"street,omitempty"`
	StreetNumber         string `json:"street_number,omitempty"`
	FlatNumber           string `json:"flat_number,omitempty"`
	City                 string `json:"city,omitempty"`
	Country              string `json:"country,omitempty"`
	PostalCode           string `json:"postal_code,omitempty"`
	NIP                  string `json:"nip,omitempty"`
	PhoneNumber          string `json:"phone_number,omitempty"`
	WebSite              string `json:"web_site,omitempty"`
	Email                string `json:"email,omitempty"`
	Note                 string `json:"note,omitempty"`
	Receiver             string `json:"receiver,omitempty"`
	MailingCompanyName   string `json:"mailing_company_name,omitempty"`
	MailingStreet        string `json:"mailing_street,omitempty"`
	MailingCity          string `json:"mailing_city,omitempty"`
	MailingPostalCode    string `json:"mailing_postal_code,omitempty"`
	DaysToPayment        string `json:"days_to_payment,omitempty"`
	PaymentMethod        string `json:"payment_method,omitempty"`
	InvoiceNote          string `json:"invoice_note,omitempty"`
	SameForwardAddress   bool   `json:"same_forward_address"`
	FirstName            string `json:"first_name,omitempty"`
	LastName             string `json:"last_name,omitempty"`
	BusinessActivityKind string `json:"business_activity_kind,omitempty"`
}

type ClientRes struct {
	MetaInfo MetaInfo `json:"metainfo"`
	Clients  []Client `json:"entities,omitempty"`
}
