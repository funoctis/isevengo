//Package isevengo provides an easy-to-use wrapper over the isEven API (isevenapi.xyz)
package isevengo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const BASE_URL string = "https://api.isevenapi.xyz/api/iseven/"

func CallIsEvenAPI(number int) (*Response, error) {
	numAsString := strconv.Itoa(number)
	u := fmt.Sprintf("%s/%s", BASE_URL, numAsString)

	rawResp, err := http.Get(u)
	if err != nil {
		return &Response{}, err
	}
	defer rawResp.Body.Close()

	respData, err := ioutil.ReadAll(rawResp.Body)

	var resp Response
	err = json.Unmarshal(respData, &resp)
	if err != nil {
		return &Response{}, err
	}

	return &Response{Iseven: resp.Iseven, Ad: resp.Ad, Error: resp.Error}, nil
}

func IsEven(number int) (bool, error) {
	resp, err := CallIsEvenAPI(number)
	if err != nil {
		return false, err
	}

	if len(resp.Error) > 0 {
		return false, fmt.Errorf(resp.Error)
	}

	return resp.Iseven, nil
}
