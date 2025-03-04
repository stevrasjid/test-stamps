package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"time"
)

const APIKEY = "28b1ca81eb414b4086add00ba750ff96"
const FORMAT = "metric"
const DATEFORMAT = "Mon, 02 Jan 2006"
const TIMEFORMAT = "15:04:05"

func main() {
	//nomor 1
	getArrayListString(100)

	//nomor 2
	getWeatherForecast("Jakarta");
}

func getArrayListString(arrLength int) {
	var result string
	for i:= arrLength; i >= 0; i-- {
		isPrime := checkPrimeNumber(i)
		if isPrime {
			continue
		}
		if i == 0 {
			result += combineResult(i, "%")
			break
		}

		if i % 3 == 0 && i % 5 == 0 {
			result += combineResult(i, "FooBar")
		} else if i % 3 == 0 {
			result += combineResult(i, "Foo")		
		} else if i % 5 == 0 {
			result += combineResult(i, "Bar")						
		} else {
			result += combineResult(i, "")						
		}
	}

	fmt.Print(result)
}

func checkPrimeNumber(n int) bool {
	if (n <= 1) {
		return false
	}

	if (n <= 3) {
		return true
	}

	if (n % 2 == 0 || n % 3 == 0){
		return false
	}

	nSqrt := int(math.Sqrt(float64(n)))

	for i:= 5; i <= nSqrt; i+=6 {
		if (n % i == 0 || n % (n + 2) == 0) {
			return false
		}
	}

	return true
}

func combineResult(n int, text string) string {
	var result string
	if text != "" {
		result = text
	} else  {
		result = strconv.Itoa(n)
	} 

	if n > 0 {
		result += ", "
	}

	return result
}

func getWeatherForecast(city string) {
	
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?q=%s&appid=%s&units=%s", city, APIKEY, FORMAT)
	response, err := http.Get(url)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	defer response.Body.Close()
	
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	
	var resultResponse ResultResponse
	json.Unmarshal(responseBody, &resultResponse)

	resultMap := make(map[string]float32)
	for _, data := range resultResponse.List {
		unixTimeUTC := time.Unix(data.Dt, 0)
		date := unixTimeUTC.Format(DATEFORMAT)
		if _, ok := resultMap[date]; !ok {
			resultMap[date] = data.Main.Temp
		}
	}

	fmt.Println("\n \nNo 2 :")
	fmt.Println("Weather Forecast:")
	for key, value := range resultMap {
		fmt.Printf("%s: %.2f°C\n", key, value)
	}
}
