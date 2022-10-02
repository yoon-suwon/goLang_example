package main

import(
    "net/http"
)

func init() {
    templates["main"] = loadTemplate("main")
    http.HandleFunc("/main", mainHandler)
    templates["input"] = loadTemplate("input")
    http.HandleFunc("/input", inputHandler)
    http.HandleFunc("/inputValue", inputValueHandler)
    templates["history"] = loadTemplate("history")
    http.HandleFunc("/history", historyHandler)
    templates["edit"] = loadTemplate("edit")
    http.HandleFunc("/edit", openEditHandler)
    http.HandleFunc("/editValue", editValueHandler)
    http.HandleFunc("/valueDelete", valueDeleteHandler)
}
 