package main

import (
	"reflect"
	"strings"
)

type Record struct {
	query                      string
	domainName                 string
	registryDomainId           string
	registrarWhoisServer       string
	updatedDate                string
	creationDate               string
	registryExpiryDate         string
	registrar                  string
	registrarIANAId            string
	domainStatus               string
	registrantOrganization     string
	registrantStateProvince    string
	registrantCountry          string
	registrantEmail            string
	adminEmail                 string
	techEmail                  string
	nameServer                 string
	dnsSec                     string
	billingEmail               string
	registrarAbuseContactEmail string
	registrarAbuseContactPhone string
}

func (r *Record) GetHeaders() []string {
	var headers []string
	t := reflect.TypeOf(*r)
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			headers = append(headers, t.Field(i).Name)

		}
	}
	return headers
}

func (r *Record) GetRow() []string {
	var recStr = []string{
		r.query,
		r.domainName,
		r.registryDomainId,
		r.registrarWhoisServer,
		r.updatedDate,
		r.creationDate,
		r.registryExpiryDate,
		r.registrar,
		r.registrarIANAId,
		r.domainStatus,
		r.registrantOrganization,
		r.registrantStateProvince,
		r.registrantCountry,
		r.registrantEmail,
		r.adminEmail,
		r.techEmail,
		r.nameServer,
		r.dnsSec,
		r.billingEmail,
		r.registrarAbuseContactEmail,
		r.registrarAbuseContactPhone,
	}
	return recStr
}

func ParseResponse(query string, res string) []string {
	var rec Record
	rec.query = query
	for _, line := range strings.Split(strings.TrimRight(res, "\n"), "\n") {
		switch {
		case strings.Contains(line, "Domain Name:"):
			rec.domainName = cleanLine(line, "Domain Name:")

		case strings.Contains(line, "Registry Domain ID:"):
			rec.registryDomainId = cleanLine(line, "Registry Domain ID:")

		case strings.Contains(line, "Registrar WHOIS Server:"):
			rec.registrarWhoisServer = cleanLine(line, "Registrar WHOIS Server:")

		case strings.Contains(line, "Updated Date:"):
			rec.updatedDate = cleanLine(line, "Updated Date:")

		case strings.Contains(line, "Creation Date:"):
			rec.creationDate = cleanLine(line, "Creation Date:")

		case strings.Contains(line, "Registry Expiry Date:"):
			rec.registryExpiryDate = cleanLine(line, "Registry Expiry Date:")

		case strings.Contains(line, "Registrar:"):
			rec.registrar = cleanLine(line, "Registrar:")

		case strings.Contains(line, "Registrar IANA ID:"):
			rec.registrarIANAId = cleanLine(line, "Registrar IANA ID:")

		case strings.Contains(line, "Domain Status:"):
			rec.domainStatus = cleanLine(line, "Domain Status:")

		case strings.Contains(line, "Registrant Organization:"):
			rec.registrantOrganization = cleanLine(line, "Registrant Organization:")

		case strings.Contains(line, "Registrant State/Province:"):
			rec.registrantStateProvince = cleanLine(line, "Registrant State/Province:")

		case strings.Contains(line, "Registrant Country:"):
			rec.registrantCountry = cleanLine(line, "Registrant Country:")

		case strings.Contains(line, "Registrant Email:"):
			rec.registrantEmail = cleanLine(line, "Registrant Email:")

		case strings.Contains(line, "Admin Email:"):
			rec.adminEmail = cleanLine(line, "Admin Email:")

		case strings.Contains(line, "Tech Email:"):
			rec.techEmail = cleanLine(line, "Tech Email:")

		case strings.Contains(line, "Name Server:"):
			rec.nameServer = cleanLine(line, "Name Server:")

		case strings.Contains(line, "DNSSEC:"):
			rec.dnsSec = cleanLine(line, "DNSSEC:")

		case strings.Contains(line, "Billing Email:"):
			rec.billingEmail = cleanLine(line, "Billing Email:")

		case strings.Contains(line, "Registrar Abuse Contact Email:"):
			rec.registrarAbuseContactEmail = cleanLine(line, "Registrar Abuse Contact Email:")

		case strings.Contains(line, "Registrar Abuse Contact Phone:"):
			rec.registrarAbuseContactPhone = cleanLine(line, "Registrar Abuse Contact Phone:")

		}
	}
	return rec.GetRow()
}

func cleanLine(l string, s string) string {
	stripPre := strings.ReplaceAll(l, s, "")
	stripSpace := strings.TrimSpace(stripPre)
	return stripSpace

}
