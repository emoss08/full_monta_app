package controllers

import (
	"net/http"
	"testing"
)

func TestGetRouteControl(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetRouteControl(tt.args.w, tt.args.r)
		})
	}
}

func TestUpdateRouteControl(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateRouteControl(tt.args.w, tt.args.r)
		})
	}
}
