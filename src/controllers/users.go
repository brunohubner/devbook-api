package controllers

import "net/http"

func Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create"))
}

func FindAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("findAll"))
}

func FindOne(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("findOne"))
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update"))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("remove"))
}
