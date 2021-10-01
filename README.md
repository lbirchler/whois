# whois
Fast bulk WHOIS lookup

# Installation
```
go get github.com/lbirchler/whois
```

# Usage
**Pipe a list of domain names into the tool to generate a csv file with whois data for each domain**
```
$ cat domains.txt | whois
```
results:
```
$ wc -l whois.csv
10001
```
fields returned for each domain:
```
$ csvtool head 1 whois.csv | tr -s ',' '\n'
query
domainName
registryDomainId
registrarWhoisServer
updatedDate
creationDate
registryExpiryDate
registrar
registrarIANAId
domainStatus
registrantOrganization
registrantStateProvince
registrantCountry
registrantEmail
adminEmail
techEmail
nameServer
dnsSec
billingEmail
registrarAbuseContactEmail
registrarAbuseContactPhone
```
speed:
```
$ time cat domains.txt | whois && wc -l domains.txt && wc -l whois.csv

real    0m11.942s
user    0m5.609s
sys     0m0.800s
1000 domains.txt
1001 whois.csv
```
**Pass -d flag to print raw whois data to stdout**
```
$ cat domains.txt | whois -d
   Domain Name: GOOGLE.COM
   Registry Domain ID: 2138514_DOMAIN_COM-VRSN
   Registrar WHOIS Server: whois.markmonitor.com
   Registrar URL: http://www.markmonitor.com
   Updated Date: 2019-09-09T15:39:04Z
   Creation Date: 1997-09-15T04:00:00Z
   Registry Expiry Date: 2028-09-14T04:00:00Z
   Registrar: MarkMonitor Inc.
   Registrar IANA ID: 292
   Registrar Abuse Contact Email: abusecomplaints@markmonitor.com
   Registrar Abuse Contact Phone: +1.2083895740
   Domain Status: clientDeleteProhibited https://icann.org/epp#clientDeleteProhibited
   Domain Status: clientTransferProhibited https://icann.org/epp#clientTransferProhibited
   Domain Status: clientUpdateProhibited https://icann.org/epp#clientUpdateProhibited
   Domain Status: serverDeleteProhibited https://icann.org/epp#serverDeleteProhibited
   Domain Status: serverTransferProhibited https://icann.org/epp#serverTransferProhibited
   Domain Status: serverUpdateProhibited https://icann.org/epp#serverUpdateProhibited
   Name Server: NS1.GOOGLE.COM
   Name Server: NS2.GOOGLE.COM
   Name Server: NS3.GOOGLE.COM
   Name Server: NS4.GOOGLE.COM
   DNSSEC: unsigned
   URL of the ICANN Whois Inaccuracy Complaint Form: https://www.icann.org/wicf/
>>> Last update of whois database: 2021-09-28T09:58:51Z <<<

...
```
# Flags
```
$ ./whois -h

Usage of ./whois:
  -c string
        output results to csv file (default "whois.csv")
  -d    print raw whois response to stdout
  -w int
        max workers (default 8)
```



