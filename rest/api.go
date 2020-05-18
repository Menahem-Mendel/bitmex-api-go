package rest

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// exchange url
const (
	BitmexHost        = "www.bitmex.com"
	BitmexHostTestnet = "testnet.bitmex.com"
	BitmexAPIPath     = "api/v1/"
)

// timeframe
const (
	Minute         = "1m"
	ThreeMinutes   = "3m"
	FiveMinutes    = "5m"
	FifteenMinutes = "15m"
	ThirtyMinutes  = "30m"
	Hour           = "1h"
	TwoHours       = "2h"
	ThreeHours     = "3h"
	FourHours      = "4h"
	SixHours       = "6h"
	TwelveHours    = "12h"
	Day            = "1D"
	ThreeDays      = "3D"
	Week           = "1W"
	TwoWeeks       = "2W"
	Month          = "1M"
)

// pairs
const (
	XBTUSD    = "XBTUSD"
	XBTM20    = "XBTM20"
	XBTU20    = "XBTU20"
	ADAM20    = "ADAM20"
	BCHM20    = "BCHM20"
	EOSM20    = "EOSM20"
	ETHUSD    = "ETHUSD"
	ETHUSDM20 = "ETHUSDM20"
	ETHM20    = "ETHM20"
	LTCM20    = "LTCM20"
	TRXM20    = "TRXM20"
	XRPUSD    = "XRPUSD"
	XPRM20    = "XPRM20"
)

// indices
const (
	BXBT       = ".BXBT"
	BXBT30M    = ".BXBT30M"
	BXBTNEXT   = ".BXBT_NEXT"
	XBTBON     = ".XBTBON"
	XBTBON2H   = ".XBTBON2H"
	XBTBON8H   = ".XBTBON8H"
	XBTUSDPI   = ".XBTUSDPI"
	XBTUSDPI2H = ".XBTUSDPI2H"
	XBTUSDPI8H = ".XBTUSDPI8H"

	BADAXBT     = ".BADAXBT"
	BADAXBT30M  = ".BADAXBT30M"
	BADAXBTNEXT = ".BADAXBT_NEXT"

	BBCHXBT     = ".BBCHXBT"
	BBCHXBT30M  = ".BBCHXBT30M"
	BBCHXBTNEXT = ".BBCHXBT_NEXT"

	BVOL    = ".BVOL"
	BVOL24H = ".BVOL24H"
	BVOL7D  = ".BVOL7D"

	BEOSXBT     = ".BEOSXBT"
	BEOSXBT30M  = ".BEOSXBT30M"
	BEOSXBTNEXT = ".BEOSXBT_NEXT"

	BETH         = ".BETH"
	BETH30M      = ".BETH30M"
	BETHXBT      = ".BETHXBT"
	BETHXBT30M   = ".BETHXBT30M"
	BETHXBT_NEXT = ".BETHXBT_NEXT"
	BETHNEXT     = ".BETH_NEXT"
	ETHBON       = ".ETHBON"
	ETHBON2H     = ".ETHBON2H"
	ETHBON8H     = ".ETHBON8H"
	ETHUSDPI     = ".ETHUSDPI"
	ETHUSDPI2H   = ".ETHUSDPI2H"
	ETHUSDPI8H   = ".ETHUSDPI8H"

	EVOL7D = ".EVOL7D"

	BLTCXBT     = ".BLTCXBT"
	BLTCXBT30M  = ".BLTCXBT30M"
	BLTCXBTNEXT = ".BLTCXBT_NEXT"

	BTRXXBT     = ".BTRXXBT"
	BTRXXBT30M  = ".BTRXXBT30M"
	BTRXXBTNEXT = ".BTRXXBT_NEXT"

	BXRP        = ".BXRP"
	BXRPXBT     = ".BXRPXBT"
	BXRPXBT30M  = ".BXRPXBT30M"
	BXRPXBTNEXT = ".BXRPXBT_NEXT"
	BXRPNEXT    = ".BXRP_NEXT"
	XRPBON      = ".XRPBON"
	XRPBON2H    = ".XRPBON2H"
	XRPBON8H    = ".XRPBON8H"
	XRPUSDPI    = ".XRPUSDPI"
	XRPUSDPI2H  = ".XRPUSDPI2H"
	XRPUSDPI8H  = ".XRPUSDPI8H"

	USDBON   = ".USDBON"
	USDBON2H = ".USDBON2H"
	USDBON8H = ".USDBON8H"
)

// limits
const (
	MAXCount float32 = 1000
)

// endpoints
const (
	trade    = "trade"
	tradeBin = "trade/bucketed"
)

// type for context key value
type ContextKey string

// key for secret value context
const KeyCtxSecret ContextKey = "secret"

// headers
const (
	signatureAPI = "api-signature"
	keyAPI       = "api-key"
	expiresAPI   = "api-expires"

	remainingRequestNum = "x-ratelimit-remaining"
	limitRequestNum     = "x-ratelimit-limit"
	limitReset          = "x-ratelimit-reset"
)

var expires = time.Now().Add(time.Hour).Unix()

func (c *Client) get(ctx context.Context, path *url.URL) ([]byte, error) {
	var out []byte

	req, err := request(ctx, http.MethodGet, c.key, path, "")
	if err != nil {
		return out, fmt.Errorf("#Client.get req: %v", err)
	}

	resp, err := do(req)
	if err != nil {
		return out, fmt.Errorf("#Client.get resp: %v", err)
	}
	defer resp.Body.Close()

	out, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("#Client.get reading response: %v", err)
	}

	return out, nil
}

func request(ctx context.Context, method, key string, url *url.URL, dataPost string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("#request: %v", err)
	}

	req.Header.Set("content-type", "application/json; charset=utf-8")

	if v, ok := req.Context().Value(KeyCtxSecret).(string); ok {
		expires := strconv.FormatInt(expires, 10)

		sign := signature(v, method+url.RequestURI()+expires)

		if method == http.MethodPost {
			sign = signature(v, method+url.RequestURI()+expires+dataPost)
		}

		req.Header.Set(signatureAPI, sign)
		req.Header.Set(keyAPI, key)
		req.Header.Set(expiresAPI, expires)
	}

	return req, nil
}

func do(req *http.Request) (*http.Response, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("#do resp: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		bs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("#do err reading error message from response")
		}
		return nil, fmt.Errorf("#do got status: %v\n%s", resp.Status, bs)
	}

	// resp.Header.Get(limitRequestNum)
	if resp.Header.Get(remainingRequestNum) == "1" {
		num, err := strconv.ParseInt(resp.Header.Get(limitReset), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("#do parseint: %v", err)
		}

		fmt.Println("Wait Ratelimit Reset...")
		time.Sleep(time.Until(time.Unix(num, 0)))
	}

	return resp, nil
}

func signature(secret, data string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))

	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}
