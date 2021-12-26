package main

import (
	"fmt"
	"math"
	"net/http"
	"log"
)


func expCalculator() float64 {
	x := 0.0001;
	for i := 0 ; i < 1000000000 ;  i++ {
		x += math.Sqrt(x);
	}

	return x;
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%f\n", expCalculator());
        fmt.Fprintf(w, "%s","Code Education Rocks");
    })

	log.Fatal(http.ListenAndServe(":80", nil))
}
