package controller

import (
	"encoding/json"
	"fmt"
	"go_brackets_validator/utils"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type BracketsController struct {
}

func (controller *BracketsController) ValidateAction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	//
	select {
	case <-time.After(1 * time.Second):
		//w.Write([]byte("request processed"))
	case <-ctx.Done():
		fmt.Fprint(os.Stderr, "request cancelled: "+r.URL.Path)
		return
	}

	var reqItems []string
	jsonData, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(jsonData, &reqItems)

	if err != nil {
		// Handle error
	}
	fmt.Println(reqItems)

	response := utils.NewApiResponse(w)
	response.SuccessJsonResponse(map[string]bool{"234234": false, "345345345": true})

}

func (controller *BracketsController) FixAction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	//
	select {
	case <-time.After(2 * time.Second):
		//w.Write([]byte("request processed"))
	case <-ctx.Done():
		fmt.Fprint(os.Stderr, "request cancelled: "+r.URL.Path)
		return
	}

	w.Write([]byte("request processed"))
}
