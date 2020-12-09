package api

import (
	"datacache/models"
	"datacache/restapi/ops/data_cache_operations"
	"net/http"
	"testing"
)

func Benchmark_storecache(b *testing.B) {
	svc := NewService()
	for n := 0; n < b.N; n++ {
		svc.StoreData(data_cache_operations.PostStoreParams{HTTPRequest: &http.Request{}, Params: &models.Cache{Key: &[]string{"hello"}[0], Value: "there", Expiration: &[]string{"1s"}[0]}})
	}
}

func Benchmark_fetchcacheviakey(b *testing.B) {
	svc := NewService()
	for n := 0; n < b.N; n++ {
		svc.FetchData(data_cache_operations.GetFetchParams{HTTPRequest: &http.Request{}, Key: &[]string{"hello"}[0]})
	}
}

func Benchmark_fetchcacheviaindex(b *testing.B) {
	svc := NewService()
	for n := 0; n < b.N; n++ {
		svc.FetchData(data_cache_operations.GetFetchParams{HTTPRequest: &http.Request{}, Index: &[]int64{1}[0]})
	}
}
