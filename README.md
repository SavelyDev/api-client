# api-client

A simple API client for getting information about IP addresses using the ipinfo package.

## Example

Get:
`go get -u -v  https://github.com/SavelyDev/api-client`

Code:

```
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/SavelyDev/api-client/ipinfo"
)

const API_TOKEN = "YOUR_TOKEN"

func main() {
	ipinfoClient, err := ipinfo.NewClient(time.Second*5, API_TOKEN)
	if err != nil {
		log.Fatal(err)
	}

	result, err := ipinfoClient.GetIPInfo("8.8.8.8")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.Info())
}
```

Output:

```
[Thu Aug 15 18:14:02 2024], GET https://ipinfo.io/8.8.8.8?token=*

IP: 8.8.8.8
Hostname: dns.google
City: Mountain View
Region: California
Country: US
Location: 37.4056,-122.0775
Organization: AS15169 Google LLC
Postal Code: 94043
Timezone: America/Los_Angeles
```
