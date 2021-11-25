package models

type Client struct {
	Regno         int    
	Serial        string
	Status          string 
	Name          string
	DBname        string
	DBport        string
	LocalNetwork  string
	LocalPort     string
	PublicNetwork string 
	PublicPort    string 
}