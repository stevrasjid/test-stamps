package main

type ResultResponse struct {
	Cod int
	Message string
	Cnt int
	List []Weather
}	

type Weather struct {
	Main MainType
	Dt int64
	Dt_txt string
}

type MainType struct {
	Temp float32
}