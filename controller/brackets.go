package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go_brackets_validator/utils"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	jsonItemsLimit = 20 // ограничение на длину входящего json'a
)

type BracketsController struct {
}

func (controller *BracketsController) ValidateAction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	completeChan := make(chan bool, 1)

	go func(ctx context.Context) {
		defer close(completeChan)
		response := utils.NewApiResponse(w)
		var reqItems []string
		jsonData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(jsonData, &reqItems)

		if err != nil {
			response.ErrorJsonResponse([]error{errors.New("bad JSON")}, http.StatusBadRequest)
		} else if len(reqItems) > jsonItemsLimit {
			response.ErrorJsonResponse([]error{errors.New("items count must be less or equal " + strconv.Itoa(jsonItemsLimit))}, http.StatusBadRequest)
		} else {
			responseItems := make(map[string]bool)
			for _, item := range reqItems {
				select {
				case <-ctx.Done():
					return
				default:
					if _, isExist := responseItems[item]; !isExist {
						responseItems[item] = utils.ValidateBrackets(item)
						time.Sleep(1 * time.Second)
					}
				}
			}
			response.SuccessJsonResponse(responseItems)
		}

		completeChan <- true
	}(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("request cancelled: " + r.URL.Path)
		return
	case <-completeChan:
		fmt.Println("completed: " + r.URL.Path)
		return
	}
}

func (controller *BracketsController) FixAction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	completeChan := make(chan bool, 1)

	go func(ctx context.Context) {
		defer close(completeChan)
		response := utils.NewApiResponse(w)
		var reqItems []string
		jsonData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(jsonData, &reqItems)

		if err != nil {
			response.ErrorJsonResponse([]error{errors.New("bad JSON")}, http.StatusBadRequest)
		} else if len(reqItems) > jsonItemsLimit {
			response.ErrorJsonResponse([]error{errors.New("items count must be less or equal " + strconv.Itoa(jsonItemsLimit))}, http.StatusBadRequest)
		} else {
			responseItems := make(map[string]string)
			for _, item := range reqItems {
				select {
				case <-ctx.Done():
					return
				default:
					if _, isExist := responseItems[item]; !isExist {
						responseItems[item] = utils.FixBrackets(item)
					}
				}
			}
			response.SuccessJsonResponse(responseItems)
		}

		completeChan <- true
	}(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("request cancelled: " + r.URL.Path)
		return
	case <-completeChan:
		fmt.Println("completed: " + r.URL.Path)
		return
	}
}
