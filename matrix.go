package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func calculateMatrix(a [][]int, b [][]int) (total int) {

	// init
	total = 0

	// calculate the matrix total
	for _, v := range a {
		for _, w := range v {
			total += w
		}
	}

	// calculate the matrix total
	for _, v := range b {
		for _, w := range v {
			total += w
		}
	}

	return
}

func main() {
	http.HandleFunc("/multi", func(w http.ResponseWriter, r *http.Request) {

		// read response body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		// close response body
		r.Body.Close()

		// marshal into a pair of matrices
		var mats [][][]int
		err = json.Unmarshal(body, &mats)

		// print
		fmt.Println(fmt.Sprintf("The body of the request new is %v", mats))

		// calculate the matrix total of values
		total := calculateMatrix(mats[0], mats[1])

		// encode the response
		j, err := json.Marshal(total)
		if err != nil {
			fmt.Println(err)
		}

		// write the response
		fmt.Fprintf(w, "%s", j)
	})

	http.ListenAndServe(":8097", nil)
}
