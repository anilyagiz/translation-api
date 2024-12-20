package translategooglefree

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

func encodeURI(s string) string {
	return url.QueryEscape(s)
}

func Translate(source, sourceLang, targetLang string) (string, error) {
	var response []interface{}
	encodedSource := encodeURI(source)
	apiURL := "https://translate.googleapis.com/translate_a/single?client=gtx&sl=" + sourceLang + "&tl=" + targetLang + "&dt=t&q=" + encodedSource

	resp, err := http.Get(apiURL)
	if err != nil {
		return "", errors.New("error fetching data from translation API")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("error reading response body")
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", errors.New("error parsing JSON response")
	}

	var translatedText string
	for _, segment := range response[0].([]interface{}) {
		translatedText += segment.([]interface{})[0].(string)
	}

	return translatedText, nil
}
