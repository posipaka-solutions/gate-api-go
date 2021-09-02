package gaterequest

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/posipaka-trade/gate-api-go/internal/pnames"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"net/http"
	"strconv"
	"time"
)

type SignStruct struct {
	Method   string
	Prefix   string
	EndPoint string
	Body     []byte
	Api      exchangeapi.ApiKey
}

func SetHeader(req *http.Request) {
	req.Header.Set(pnames.Accept, "application/json")
	req.Header.Set(pnames.ContentType, "application/json")

}

func MakeSign(signStruct SignStruct, req *http.Request) {
	h := sha512.New()
	h.Write(signStruct.Body)
	hashedPayload := hex.EncodeToString(h.Sum(nil))

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	timestampStr := fmt.Sprint(timestamp)

	whatToSign := fmt.Sprintf("%s\n%s\n%s\n%s\n%s", signStruct.Method,
		signStruct.Prefix+signStruct.EndPoint,
		"",
		hashedPayload,
		timestampStr)
	hash := hmac.New(sha512.New, []byte(signStruct.Api.Secret))
	hash.Write([]byte(whatToSign))
	sign := hex.EncodeToString(hash.Sum(nil))

	req.Header.Set(pnames.Key, signStruct.Api.Key)
	req.Header.Set(pnames.Timestamp, timestampStr)
	req.Header.Set(pnames.Signature, sign)
}
