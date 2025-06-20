package router

import (
	"github.com/dangLuan01/api_go_movie28/apis/collectionapi"
	"github.com/dangLuan01/api_go_movie28/apis/genreapi"
	"github.com/dangLuan01/api_go_movie28/apis/movieapi"
	"github.com/dangLuan01/api_go_movie28/apis/searchapi"
	"github.com/dangLuan01/api_go_movie28/apis/themeapi"
	"github.com/gorilla/mux"
)
func SetupRouter() *mux.Router{
	router := mux.NewRouter()
	
	router.HandleFunc("/api/v1/movie-hot", movieapi.GetMovieHot).Methods("GET")
	router.HandleFunc("/api/v1/category", movieapi.GetCategory).Methods("GET")
	router.HandleFunc("/api/v1/movies", movieapi.GetAllMovie).Methods("GET")
	router.HandleFunc("/api/v1/genre", genreapi.GetAllGenre).Methods("GET")
	router.HandleFunc("/api/v1/movie/{slug}", movieapi.GetMovieBySlug).Methods("GET")
	router.HandleFunc("/api/v1/genre-home", genreapi.GetAllGenreHome).Methods("GET")
	router.HandleFunc("/api/v1/genre/{slug}", genreapi.GetGenreInfo).Methods("GET")
	router.HandleFunc("/api/v1/collection", collectionapi.GetColletion).Methods("GET")
	router.HandleFunc("/api/v1/collection/{slug}", collectionapi.GetColletionBySlug).Methods("GET")
	router.HandleFunc("/api/v1/themes", themeapi.GetThemes).Methods("GET")
	router.HandleFunc("/api/v1/_search", searchapi.SearchES).Methods("GET")
	//router.HandleFunc("/api/v1/syn", searchapi.Syn).Methods("GET")

	return router
}