# golive

golive is a Go wrapper around the [Infinite Flight Live API](https://infiniteflight.com/guide/developer-reference/live-api/overview).

#### Requirements

`go >= 1.18`

#### Usage

```golang
// 1. Import golive
import "github.com/sqeezelemon/golive"

// 2. Initialize the client
client := golive.NewClient("totally_an_api_key", &http.client{})

// 3. Done
try, catch := client.GetSessions()
```

#### Contacts
[**@sqeezelemon** on IFC](https://community.infiniteflight.com/u/sqeezelemon)

***
Special thanks to @sadfun for his consulting.