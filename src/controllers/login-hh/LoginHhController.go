package LoginHhController

import (
	"net/http"
	"fmt"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "123")
}