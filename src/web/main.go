package web

import (
	"net/http"

	"github.com/epicnotgames/cli-flashcards/src/packages/utils"
	"github.com/epicnotgames/cli-flashcards/src/web/handlers"
)

func WebInit() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/flashcards", handlers.FlashcardsHandler)

	utils.ErrHandle(http.ListenAndServe(":8080", nil))
}