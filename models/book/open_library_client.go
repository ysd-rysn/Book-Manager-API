package book

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var openLibraryAPI = "http://openlibrary.org/api/books?bibkeys=ISBN:%s&jscmd=data&format=json"

// fetch book data from open library api by ISBN
func FetchBookData(isbn string) (map[string]interface{}, error) {
	if !(len(isbn) == 10 || len(isbn) == 13) {
		return nil, errors.New("isbn must be 10 or 13 digits")
	}

	url := fmt.Sprintf(openLibraryAPI, isbn)
	fmt.Println(url)

	// fetch book data from open library api
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, err
	}

	resBody, _ := ioutil.ReadAll(res.Body)

	var jsonData map[string]interface{}
	if err := json.Unmarshal(resBody, &jsonData); err != nil {
		return nil, err
	}

	firstMapKey := "ISBN:" + isbn
	bookData := jsonData[firstMapKey].(map[string]interface{})

	// print book data from open library api
	if jsonString, err := PrettyJSONString(resBody); err != nil {
		return nil, err
	} else {
		fmt.Println(jsonString)
	}

	return bookData, nil
}

func PrettyJSONString(data []byte) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
