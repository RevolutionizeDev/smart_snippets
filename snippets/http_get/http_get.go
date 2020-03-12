import "net/http"

// get request
resp, err := http.Get(url) //$1
if err != nil {
	//$2
}
defer resp.Body.Close()

// read response body
body, err := ioutil.ReadAll(resp.Body)
if err != nil {
	//$0
}

// access/use/print headers
fmt.Printf("%#v\n", resp.Header)

// access/use/print content of response body
fmt.Printf("%s\n", body)
