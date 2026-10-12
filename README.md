<h1 align="center"> 🚀 Go Concurrent Uptime Monitor </h1>

<p align="center">
  <img src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg?logo=go&color=00ADD8" alt="Made with Go">
  <img src="https://img.shields.io/badge/Concurrency-Goroutines-brightgreen.svg?style=flat" alt="Goroutines">
  <img src="https://img.shields.io/badge/Status-Awesome-ff69b4.svg" alt="Status">
</p>

> **A blazing fast, concurrent website status checker built entirely in Go.**  
> This project serves as the perfect "First Go Project," demonstrating the true power of Go's concurrency model (Goroutines and Channels) while covering essential backend concepts.

---

## ✨ What It Does

Provide a list of website URLs, and this tool will check their HTTP status **all at the exact same time**. It catches timeouts, handles broken links gracefully, and neatly packages the results into a formatted JSON file. 

## 🧠 Go Concepts Mastered

This isn't just a simple script; it's a tour of Go's core features. Here is what's running under the hood:

| Concept | How it's used in this project |
| :--- | :--- |
| 🏎️ **Goroutines** | Spawning lightweight threads to check multiple URLs simultaneously. |
| 📡 **Channels** | Safely passing HTTP responses back to the main thread. |
| ⏱️ **Select & Timeouts** | Preventing infinite hangs using `select` and `time.After`. |
| 📦 **Packages/Modules** | Structuring code cleanly by separating logic (`main` vs `checker`). |
| 🎯 **Pointers** | Passing configuration without wasting memory by duplicating structs. |
| 📝 **JSON Encoding** | Converting Go maps into beautifully formatted JSON files. |
| 🛡️ **Defers** | Preventing memory leaks by auto-closing HTTP bodies and files. |

---

## 📂 Project Structure

```text
uptime-monitor/
├── 📄 go.mod               # Go module tracking
├── 🚀 main.go              # Main execution, channels, and JSON I/O
└── 📁 checker/
    └── 🛠️ checker.go       # Structs, interfaces, and HTTP logic
```

---

## 🚦 Getting Started

### 1. Clone & Initialize
Create a folder and initialize the Go module:
```bash
mkdir uptime-monitor
cd uptime-monitor
go mod init uptime-monitor
```

### 2. Add the Code
Copy the `main.go` and `checker/checker.go` files into your directory as shown in the project structure above.

### 3. Run the Monitor
Execute the program from your terminal:
```bash
go run main.go
```

---

## 📊 Sample Output

### Terminal (Real-time feedback)
```text
[DOWN] https://fake-url.invalid - Get "https://fake-url.invalid": no such host
[UP] https://golang.org
[UP] https://google.com

[!] Global timeout reached! Aborting remaining checks.

Results saved neatly to status.json
```

### 🗂️ `status.json` (Pretty-printed results)
```json
{
    "https://fake-url.invalid": {
        "url": "https://fake-url.invalid",
        "status": "Get \"https://fake-url.invalid\": dial tcp: lookup fake-url.invalid: no such host",
        "up": false
    },
    "https://golang.org": {
        "url": "https://golang.org",
        "status": "200 OK",
        "up": true
    }
}
```

---

<p align="center">
  <i>Built to learn, scale, and understand the magic of Go. 💙</i>
</p>
