Here is your consolidated **Go Concurrency Master Note**, designed for your Sunday Integration Lab. It combines the core concepts of Shared Memory, Goroutines, and Synchronization with code examples for each.

---

# **Go Concurrency: The Foundations**

## **1. The Philosophy: Shared Memory vs. Channels**

In Go, Goroutines run in the same **address space**, meaning they can see and touch the same variables.

* **The Risk:** Race conditions (two routines changing data at once).
* **The Rule:** "Do not communicate by sharing memory; instead, share memory by communicating."

### **The "Traditional" Way (sync.Mutex)**

Use this when you have a simple shared resource (like a counter) that multiple routines need to access.

```go
type SafeCounter struct {
    mu    sync.Mutex
    Value int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()         // Lock the "box"
    c.Value++           // Change the data
    c.mu.Unlock()       // Release the lock
}

```

---

## **2. Goroutines: Lightweight Execution**

A goroutine is a function that runs concurrently with other functions. It is much cheaper than a traditional OS thread.

### **Code Example: Concurrent Scanning**

```go
func main() {
    // Prefixing with 'go' starts it in the background
    go printMessage("Scan started...") 
    
    // Warning: If main exits here, the goroutine dies immediately!
}

```

---

## **3. sync.WaitGroup: The Orchestrator**

Since `main` doesn't wait for goroutines, we use a `WaitGroup` to act as a counter. The program only finishes when the counter reaches zero.

### **The Three-Step Pattern**

1. **`Add(n)`**: Tell the counter how many goroutines to expect.
2. **`Done()`**: Tell the counter one task is finished (use `defer`).
3. **`Wait()`**: Block `main` until the counter is zero.

### **Code Example: Waiting for Workers**

```go
func main() {
    var wg sync.WaitGroup
    files := []string{"log1.txt", "log2.txt"}

    for _, file := range files {
        wg.Add(1) // 1. Add before starting
        go func(f string) {
            defer wg.Done() // 2. Done when finished
            scanFile(f)
        }(file)
    }

    wg.Wait() // 3. Block here until all Dones are called
}

```

---

## **4. Pointer Receivers & Concurrency**

For your lab, you must pass the `WaitGroup` as a **pointer**. If you pass it by value, the goroutine gets a *copy* of the counter. It will decrement its local copy, while the one in `main` stays at 1 forever, causing a **deadlock**.

### **Code Example: Pointer Receiver with WaitGroup**

```go
type FileScanner struct {
    Query string
}

// Pass WaitGroup as a pointer (*sync.WaitGroup)
func (s *FileScanner) Find(path string, wg *sync.WaitGroup) {
    defer wg.Done()
    
    // Search logic here...
    fmt.Printf("Searching for %s in %s\n", s.Query, path)
}

```

---

## **5. Efficient Streaming: bufio.Scanner**

To search files without loading them entirely into RAM (Shared Memory efficiency), use a stream-based scanner.

### **Code Example: Memory-Safe Searching**

```go
func (s *FileScanner) Find(path string, wg *sync.WaitGroup) {
    defer wg.Done()

    file, _ := os.Open(path)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    lineNum := 1
    for scanner.Scan() {
        if strings.Contains(scanner.Text(), s.Query) {
            fmt.Printf("Found in %s: Line %d\n", path, lineNum)
        }
        lineNum++
    }
}

```

