package main

import (
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
    if _, err := io.WriteString(w, runtime.GOOS); err != nil {
        log.Print(err)
    }
}

func main() {
    fs := http.FileServer(http.Dir("./public"))
    http.Handle("/", fs)
    http.HandleFunc("/clicked", getClicked)

    OpenBrowser("http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
