package main

import (
	"fmt"
	"net/http"
	"vocabula/common"
	"vocabula/handlers"
)

func ServerMain() error {
	http.HandleFunc("/vocabulary", handlers.VocabularyHandler)

	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", common.GlobalServerConfig.ServicePort), nil)
	if err != nil {
		return err
	}

	return nil
}
