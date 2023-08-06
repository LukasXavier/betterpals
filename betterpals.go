package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func OpenBrowser(url string) {
    var cmd *exec.Cmd
    switch runtime.GOOS {
    case "darwin":
        cmd = exec.Command("open", url)
    case "windows":
        cmd = exec.Command("cmd", "/c", "start", url)
    default:
        if os.Getenv("WSLENV") == "" {
            cmd = exec.Command("xdg-open", url)
        } else {
            cmd = exec.Command("/mnt/c/Windows/system32/cmd.exe", "/c", "start", url)
        }
    }

    if err := cmd.Run(); err != nil {
        log.Fatal(err)
    }
}

func getClicked(w http.ResponseWriter, r *http.Request) {
    res, err := getIndividualTeam("64775970bba9d14862bcf9ce");
    if err != nil {
        log.Print(err)
    }
    tmpl, _ := template.New("mcd-button").Parse(`<div class="stats">{{.}}</div>`)
    if err := tmpl.Execute(w, res); err != nil {
        log.Print(err)
    }
}

func changeColor(w http.ResponseWriter, r *http.Request) {
    log.Print("clicked")
    if _, err := io.WriteString(w, "mcd-button--raised mcd-button-primary"); err != nil {
        log.Print(err)
    }
}

func main() {
    fs := http.FileServer(http.Dir("./public"))
    http.Handle("/", fs)
    http.HandleFunc("/clicked", getClicked)
    http.HandleFunc("/change-color", changeColor)

    // OpenBrowser("http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
