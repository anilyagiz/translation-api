package main

import (
	"log"
	"net/http"
	"translation-api/translategooglefree"
)

func translateHandler(w http.ResponseWriter, r *http.Request) {
	sourceText := r.URL.Query().Get("text")
	sourceLang := r.URL.Query().Get("sourceLang")
	targetLang := r.URL.Query().Get("targetLang")

	if sourceText == "" || sourceLang == "" || targetLang == "" {
		http.Error(w, "Missing query parameters: text, sourceLang, targetLang", http.StatusBadRequest)
		return
	}

	translatedText, err := translategooglefree.Translate(sourceText, sourceLang, targetLang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(translatedText))
}

func main() {
	http.HandleFunc("/translate", translateHandler)
	log.Println("Translation API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
