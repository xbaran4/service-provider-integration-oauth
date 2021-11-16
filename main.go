/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"spi-oauth/controllers"
)

func main() {
	start()
}

func start() {

	router := mux.NewRouter()

	router.HandleFunc("/github/authenticate", controllers.GitHubAuthenticate).Methods("GET")
	router.HandleFunc("/github/callback", controllers.GitHubCallback).Methods("GET")

	router.HandleFunc("/quay/authenticate", controllers.QuayAuthenticate).Methods("GET")
	router.HandleFunc("/quay/callback", controllers.QuayCallback).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
