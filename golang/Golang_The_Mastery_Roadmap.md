# Golang Systems Architect: The Mastery Roadmap (Jan – May 2026)

**Objective:** Achieve idiomatic fluency in Go, specifically targeted at Systems Programming and Cloud-Native Architecture (Kubernetes foundations).

**Timeline:** January 14, 2026 – May 1, 2026
**Philosophy:** The Traffic Light System (🔴 Theory -> 🟡 Ability -> 🟢 Mastery).

---

## 🧱 Phase 1: The Idiomatic Foundation (Jan 10 – Feb 7)
**Theme:** *Thinking in Go ("The Go Way")*
*Goal: Break OOP habits. Understand memory layout and strict typing.*

* **Week 1 (Jan 10-18): Syntax & Behavior**
    * Variables (`var` vs `:=`), Functions, Multiple Returns.
    * Control Flow (`if`, `switch`, `for` loops).
    * **The Trap:** The "Main Exit" problem (Goroutine lifecycle).

* **Week 2 (Jan 19-25): Structs & Data Modeling**
    * Struct definitions and initialization.
    * **Crucial Concept:** Value Receivers (Copy) vs. Pointer Receivers (Reference).
    * The "Map Update Trap" (Why you can't address map elements directly).

* **Week 3 (Jan 26-Feb 1): Interfaces & Composition**
    * Implicit Interfaces ("If it walks like a duck...").
    * Composition over Inheritance (Embedding structs).
    * The `interface{}` (Empty interface) and Type Assertions.

* **Week 4 (Feb 2-8): Error Handling & Resilience**
    * The philosophy of `if err != nil`.
    * `panic` and `recover` (The "Do Not Touch" zone).
    * Creating custom error types.
    * **🏆 Capstone:** **Task Manager CLI**. A tool that adds/lists tasks in a JSON file using Structs and File I/O.

---

## ⚡ Phase 2: Concurrency & The Runtime (Feb 9 – Mar 8)
**Theme:** *Managing Chaos (The Kubernetes Engine)*
*Goal: Master the features that make Go famous. Learn to do many things safely.*

* **Week 5 (Feb 9-15): Goroutines & The Scheduler**
    * The `go` keyword.
    * The M:N Scheduler (How Go maps thousands of routines to OS threads).
    * Race Conditions (Detecting them with `go run -race`).

* **Week 6 (Feb 16-22): Channels (The Pipelines)**
    * Unbuffered Channels (Synchronization).
    * Buffered Channels (Queueing).
    * The "Producer-Consumer" Pattern.
    * Closing channels and the `range` loop over channels.

* **Week 7 (Feb 23-Mar 1): The Sync Package**
    * `sync.WaitGroup` (Waiting for a group of tasks).
    * `sync.Mutex` vs `sync.RWMutex` (Protecting shared memory).
    * Atomic operations (`sync/atomic`).

* **Week 8 (Mar 2-8): Orchestration**
    * `select` statement (Handling multiple channels).
    * Timeouts and Deadlines.
    * **🏆 Capstone:** **Concurrent Port Scanner**. A tool that scans 100 ports simultaneously using workers and reports results via a channel.

---

## 🛠️ Phase 3: Systems Tooling & IO (Mar 9 – Apr 5)
**Theme:** *Building Your Own "kubectl"*
*Goal: Interact with the OS, files, streams, and other processes.*

* **Week 9 (Mar 9-15): The Universal Interfaces**
    * `io.Reader` and `io.Writer`.
    * Chaining readers (`io.TeeReader`, `io.MultiWriter`).
    * Working with Files (`os.Open`, `ioutil` vs `os` package).

* **Week 10 (Mar 16-22): CLI Engineering**
    * Parsing flags (`flag` package).
    * Handling Arguments (`os.Args`).
    * Exit codes (`os.Exit`).

* **Week 11 (Mar 23-29): OS Interaction**
    * Executing external commands (`os/exec`).
    * Handling OS Signals (Graceful Shutdown on `Ctrl+C`).
    * Environment Variables.

* **Week 12 (Mar 30-Apr 5): Quality Assurance**
    * `go test` basics.
    * Table-Driven Tests (The idiomatic way to test).
    * Benchmarking (`go test -bench`).
    * **🏆 Capstone:** **Log Rotator**. A daemon that watches a log file; when it hits a size limit, it zips it and creates a new one.

---

## 🌐 Phase 4: Networking & Microservices (Apr 6 – May 1)
**Theme:** *The Cloud Native Layer*
*Goal: Networking, APIs, and the Context package.*

* **Week 13 (Apr 6-12): HTTP Client**
    * Making requests (`http.Get`, `http.Post`).
    * Setting Headers and Timeouts.
    * Parsing JSON responses (`encoding/json`).

* **Week 14 (Apr 13-19): HTTP Server**
    * Building a REST API with the Standard Library.
    * Routing and Middleware basics.
    * JSON Marshaling/Unmarshaling.

* **Week 15 (Apr 20-26): The Context Package**
    * `context.Background`, `context.TODO`.
    * Propagating cancellation signals.
    * Passign request-scoped values.

* **Week 16 (Apr 27-May 1): Final Integration**
    * **🏆 Final Project:** **System Monitor API**. An HTTP server that returns the host machine's real-time CPU/RAM usage as JSON.
