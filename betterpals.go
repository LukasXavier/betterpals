package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/LukasXavier/betterpals/api/store"
	"github.com/LukasXavier/betterpals/api/team"
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
    query := r.URL.Query()
    log.Print(query.Get("id"))
    res, err := team.New(query.Get("id"))
    if err != nil {
        log.Print(err)
    }
    tmpl, err := template.ParseFiles("./templates/team.html")
    if err != nil {
        log.Print(err)
    }
    if err := tmpl.Execute(w, res); err != nil {
        log.Print(err)
    }
}

func getSchedule(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    log.Print(query.Get("id"))
    STORE.Sync(query.Get("id"))
    res := STORE.Teams
    /* res, err := schedule.New(query.Get("id"))
    if err != nil {
        log.Print(err)
    } */
    tmpl, err := template.ParseFiles("./templates/schedule.html")
    if err != nil {
        log.Print(err)
    }

    val := make([]struct{ Id string; Name string }, 0, len(res))
    for k, v := range res {
        val = append(val, struct{ Id string; Name string }{Id: k, Name: v.Name})
    }

    if err := tmpl.Execute(w, val); err != nil {
        log.Print(err)
    }
}

var STORE = store.New()
func main() {
    fs := http.FileServer(http.Dir("./public"))
    http.Handle("/", fs)
    http.HandleFunc("/clicked", getClicked)
    http.HandleFunc("/schedule", getSchedule)

    // OpenBrowser("http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
