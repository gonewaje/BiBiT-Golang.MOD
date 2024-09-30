// package main
package handler

import (
	"fmt"
	"net/http"
)

// func hello() {
// 	http.HandleFunc("/", HelloServer)
// 	// http.ListenAndServe(":8080", nil)
// }

func HelloServer(w http.ResponseWriter, r *http.Request) {
	// Serve the "Hello World" message in bold and centered using HTML
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Hello World</title>
			<style>
				body {
					display: flex;
					justify-content: center;
					align-items: center;
					height: 100vh;
					margin: 0;
					font-family: Arial, sans-serif;
				}
				h1 {
					font-weight: bold;
				}
			</style>
		</head>
		<body>
			<h1>Hello, World!</h1>
		</body>
		</html>
	`)
}
