Since you are wrapping up **Phase 1, Week 3** and entering **Week 4**, your new Go schedule focuses on transitioning from **Interfaces and Composition** into **Pointers, Advanced Error Handling, and Basic Concurrency**.

I have moved your unfinished items—the **Empty Interface Review** and the **Integration Lab (Pluggable Logger)**—to the beginning of this week to ensure you have a solid foundation before moving to memory management.

### **📅 Go Weekly Schedule: Feb 2 – Feb 8, 2026**

**Theme:** *Memory Mastery & The Concurrency Bridge*

* **Feb 02 (Mon): Final Interface Review & Type Assertions**
* **Task:** Finish the **Empty Interface** `interface{}` lab from last week.
* **Focus:** Practice safely converting an empty interface back to its concrete type using the `v, ok := i.(T)` pattern.


* **Feb 03 (Tue): Integration Lab – Pluggable Logger**
* **Task:** Complete the lab to build a logger that can write to multiple destinations (console, file, etc.) using interfaces.
* **Focus:** Decoupling your code so the logger doesn't care *where* it writes, only that the destination satisfies the `io.Writer` interface.


* **Feb 04 (Wed): Pointers & Memory Management**
* **Task:** Learn the difference between "Passing by Value" vs. "Passing by Reference".
* **Exercise:** Write a function that modifies a struct field using a pointer receiver vs. a value receiver to see the effect on the original data.


* **Feb 05 (Thu): Advanced Error Handling (The `error` Interface)**
* **Task:** Go beyond basic `fmt.Errorf`. Create a custom struct that implements the `Error()` interface.
* **Focus:** Use `errors.Is()` and `errors.As()` to handle specific custom error types in your logic.


* **Feb 06 (Fri): Goroutines – Your First Step into Concurrency**
* **Task:** Introduction to the `go` keyword. Run multiple functions simultaneously.
* **Exercise:** Use `sync.WaitGroup` to ensure your `main()` function waits for all goroutines to finish before exiting.


* **Feb 07 (Sat): Weekend Review – Memory & Concurrency Troubleshooting**
* **Task:** Refactor your Pluggable Logger (from Tuesday) to use a pointer receiver for its methods.
* **Performance Protocol:** Use the **2-Minute Rule**: if a task feels heavy, just type `func main() { go ... }`.


* **Feb 08 (Sun): Integration Lab – The Concurrent File Scanner**
* **Task:** Build a small tool that scans multiple files at once for a specific keyword using goroutines.
* **Goal:** Practice combining pointers (for shared configuration) and basic concurrency.



---

### **Strategic Performance Protocols**

* **The 2-Minute Rule:** If you feel resistance, your only goal is to type the package and import blocks.
* **Never Miss Twice:** If Monday is a wash, Tuesday becomes a mandatory "Green" day to maintain momentum in your Roadmap.
