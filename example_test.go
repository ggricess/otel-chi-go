// Copyright Splunk Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package otelchigo_test

import (
	"net/http"

	otelchigo "github.com/ggricess/otel-chi-go"
	"github.com/go-chi/chi/v5"
)

func Example() {
	router := chi.NewRouter()
	router.Use(otelchigo.Middleware())
	router.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World!\n"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
