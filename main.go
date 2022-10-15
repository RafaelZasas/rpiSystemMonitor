package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rafaelzasas/rpiSystemMonitor/cpu"
	"github.com/rafaelzasas/rpiSystemMonitor/mem"
	// "github.com/rafaelzasas/rpiSystemMonitor/storage"
	// "github.com/rafaelzasas/rpiSystemMonitor/network"
)

func getMemUsed(w http.ResponseWriter, r *http.Request) {
	res, err := mem.GetUsed()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func getMemFree(w http.ResponseWriter, r *http.Request) {
	res, err := mem.GetFree()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func getSwpMemFree(w http.ResponseWriter, r *http.Request) {
	res, err := mem.GetSwpFree()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func getSwpMemUsed(w http.ResponseWriter, r *http.Request) {
	res, err := mem.GetSwpUsed()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func getCpuFreq(w http.ResponseWriter, r *http.Request) {
	res, err := cpu.GetFreq()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func getCpuTemp(w http.ResponseWriter, r *http.Request) {
	res, err := cpu.GetTemp()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func getCpuSummary(w http.ResponseWriter, r *http.Request) {
	res, err := cpu.GetTemp() // TODO:  cpu.GetSummary()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

//commonMiddleWare provides middleware like content type and headers
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)

	// CPU Routes
	router.HandleFunc("/cpu", getCpuSummary).Methods(http.MethodGet)
	router.HandleFunc("/cpu/temp", getCpuTemp).Methods(http.MethodGet)
	router.HandleFunc("/cpu/freq", getCpuFreq).Methods(http.MethodGet)

	// Mem Routes
	router.HandleFunc("/mem/free", getMemFree).Methods(http.MethodGet)
	router.HandleFunc("/mem/used", getMemUsed).Methods(http.MethodGet)
	router.HandleFunc("/mem/swp/free", getSwpMemFree).Methods(http.MethodGet)
	router.HandleFunc("/mem/swp/used", getSwpMemUsed).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":3000", router))
}

func main() {
	handleRequests()
}
