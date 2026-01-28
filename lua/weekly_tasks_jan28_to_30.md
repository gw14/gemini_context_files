## 🏗️ Phase 1, Week 1: Variables, Scoping, and Truthiness

### 🌙 Jan 28 (Wednesday): The Global Leak Test

**Theme:** *Local by Default*

* **The Task:** Create a script that defines a global variable and a local variable of the same name inside a function.
* **The Goal:** Observe "shadowing." Print both to prove the global remains unchanged.
* **2-Minute Rule:** Type `local x = "secure"`.

### 🌙 Jan 29 (Thursday): The Truthiness Audit

**Theme:** *Breaking OOP Habits*

* **The Task:** Write a "Logic Checker" function that accepts an input and prints "It is True" or "It is False".
* **The Goal:** Test `0`, `""` (empty string), `{}`, `false`, and `nil`.
* **Architect's Note:** Prove to yourself that in Lua, only `false` and `nil` fail the test.

### 🌙 Jan 30 (Friday): Block Scoping & Loops

**Theme:** *Tightening the Engine*

* **The Task:** Use a `for` loop to increment a value, but define the counter as `local` within the loop block.
* **The Goal:** Attempt to print the counter *after* the loop ends to trigger an error, confirming the variable has been successfully garbage collected or is out of scope.

### 🌙 Jan 31 (Saturday): Integration Review (Neovim/Luau Focus)

**Theme:** *Contextual Application*

* **The Task:** Open your Neovim `init.lua` or a Roblox script and convert three global settings/variables to `local`.
* **The Goal:** Reduce "namespace pollution" in a real-world environment you actually use.

### ☀️ Feb 01 (Sunday): Weekly Audit & The "Minimum Viable Day"

**Theme:** *System Maintenance*

* **The Task:** Perform a **Friction Audit**. Was it hard to open your Lua interpreter?
* **The Goal:** Map out your **Minimum Viable Day** (15-minute version) for Week 2.
* **Identity Check:** Confirm you are becoming the developer who writes clean, scoped code.

---

### 🧠 Performance Protocol

* **Habit Stacking:** "After I finish my daily **CKA Speed Drill**, I will spend 10 minutes on my **Lua Task of the Day**".
* **Never Miss Twice:** If Saturday is a wash due to family or work, Sunday becomes a mandatory "Green" day for your Lua foundation.
