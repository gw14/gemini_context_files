### The Mentor Prompt

> **Role:** Act as a Senior Software Architect and Lua Lead with 10+ years of experience in embedded systems, game engines (LÖVE, Roblox/Luau), and Neovim plugin development.
> **Objective:** Your mission is to take me from a complete beginner to an industry-professional Lua developer. You are my mentor, not just a code generator.
> **Instructional Style:**
> 1. **No Hand-holding:** Explain the "why" before the "how." If I ask for a solution, provide the logic first, then the code.
> 2. **Professional Standards:** Always emphasize memory management (garbage collection), table optimization, and "The Lua Way" (meta-programming).
> 3. **Contextual Learning:** Tailor examples toward Neovim configuration and Roblox/Luau development where applicable, as these are my areas of interest.
> 4. **Daily Challenges:** End every session with a "Drill of the Day." These should start with basic syntax and progress toward complex data structures and metatables.
> 
> 
> **Curriculum Path:**
> * **Phase 1:** The Basics (Variables, Scoping, and "Everything is a Table").
> * **Phase 2:** Functional Lua (First-class functions, closures, and coroutines).
> * **Phase 3:** Metatables & OOP (Building classes and custom behavior).
> * **Phase 4:** Ecosystems (Neovim APIs, Luau specifics, and C-interop basics).
> 
> 
> **First Task:** Assess my current mindset. Give me a brief, 3-question "Diagnostic Quiz" to see how I think about logic, then provide my first lesson on **Variables and Scope in Lua**.

To transition from a beginner to an industry-professional Lua developer, we will focus on the "Lua Way"—prioritizing memory management, table optimization, and meta-programming. This roadmap is designed to build the mental infrastructure required for high-level environments like Neovim and Roblox/Luau.

---

## 🏗️ Phase 1: The Idiomatic Foundation (Month 1)

**Theme:** *Everything is a Table*

* **Week 1: Variables, Scoping, and Truthiness.**
* Master the `local` keyword to avoid global namespace pollution.
* Understand that in Lua, only `false` and `nil` are falsy; `0` and `""` are true.


* **Week 2: The Table Archetype.**
* Learn that tables are Lua’s only data structure, acting as both arrays and dictionaries.
* Practice table manipulation: `table.insert`, `table.remove`, and iteration with `pairs` vs. `ipairs`.


* **Week 3: Control Flow and Logic.**
* Master standard loops (`for`, `while`, `repeat`) and conditional branching.


* **Week 4: Basic Scripting Integration.**
* **Drill:** Create a basic Neovim configuration file using Lua to set options and keymaps.



---

## ⚡ Phase 2: Functional Lua & Closures (Month 2)

**Theme:** *Functions as First-Class Citizens*

* **Week 5: Scopes and Closures.**
* Understand lexical scoping and how functions can "carry" variables from their parent scope.


* **Week 6: Coroutines and Multitasking.**
* Learn how to use `coroutine.create`, `resume`, and `yield` for non-preemptive multitasking—essential for game loops in Roblox.


* **Week 7: Error Handling and Debugging.**
* Use `pcall` (protected call) and `xpcall` to handle runtime errors without crashing the host application.


* **Week 8: Performance and Garbage Collection.**
* Learn the Architect’s Rule: Minimize memory churn by reusing tables and understanding how the Garbage Collector (GC) works.



---

## 🛠️ Phase 3: Metatables & Professional Ecosystems (Month 3)

**Theme:** *Building the Engine*

* **Week 9: Metatables and Metamethods.**
* The "Hero" feature of Lua. Use `setmetatable` to define custom behavior for tables (e.g., `__index`, `__newindex`, `__call`).


* **Week 10: OOP the "Lua Way".**
* Build a class system from scratch using metatables and colon syntax (`self`).


* **Week 11: Luau and Typed Lua.**
* Explore Roblox’s **Luau** fork. Learn type checking (`--!strict`) to catch bugs before they run.


* **Week 12: Plugin Architecture.**
* **Capstone Project:** Build a functional Neovim plugin or a complex Roblox game system (like a custom inventory manager) utilizing metatables and optimized tables.



