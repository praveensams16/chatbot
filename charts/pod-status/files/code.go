package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
)

var (
	files []string
	//lines []string
)

func run(w http.ResponseWriter, r *http.Request) {

	cmd := exec.Command("kubectl", "get", "pods")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s := string(stdout)

	files := strings.Split(string(s), "\n")

	for i := 0; i < len(files); i++ {
		if !strings.Contains(files[i], "NAME") && len(files[i]) > 1 && !strings.Contains(files[i], "Completed") {

			lines := regexp.MustCompile("[^\\s]+")

			status := lines.FindAllString(files[i], -1)

			if status[2] == "Running" {
				fmt.Fprintf(w, "pod=\""+status[0]+"\" -> status=\"Running\" ->" +  "  ready=" + status[1]  + "  ->  restart=" + status[3]+  "\n")
			} else {
				fmt.Fprintf(w, "pod=\""+status[0]+"\" -> status=" +status[2] +"  ready=" + status[1]  + "  ->  restart=" + status[3]+ " \n")
			}
		}
	}

}

func handleRequests() {
	http.HandleFunc("/metrics", run)
	log.Fatal(http.ListenAndServe(":9101", nil))
}

func main() {
	handleRequests()
}


