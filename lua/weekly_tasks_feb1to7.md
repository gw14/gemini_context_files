Based on your current progress in **Phase 1: The Idiomatic Foundation** and your focus on Neovim/Luau, here is your new **Lua Weekly Schedule** for **February 2 – February 8, 2026**.

I have moved your unfinished tasks—the **Block Scoping/Garbage Collection** test, the **Neovim/Luau Global Audit**, and the **System Friction Audit**—to the start of the week. This ensures your "Local by Default" mindset is solidified before moving into **Week 2: The Table Archetype**.

### 🌙 Lua Weekly Schedule: Feb 2 – Feb 8, 2026

**Theme:** *The Table Archetype (Arrays vs. Dictionaries)*

* **Feb 02 (Mon): Final Review – The Block Scope Kill-Test**
* **Task:** Use a `for` loop to increment a value, defining the counter as `local` within the block. Attempt to print it *after* the loop to confirm it is out of scope.
* **Goal:** Prove that Lua successfully cleans up block-local variables.


* **Feb 03 (Tue): Contextual Audit – Neovim/Luau Refactor**
* **Task:** Open your Neovim `init.lua` or a Roblox script and convert three global settings/variables to `local`.
* **Goal:** Reduce namespace pollution in a production environment.


* **Feb 04 (Wed): System Maintenance – Friction Audit**
* **Task:** Map out your **Minimum Viable Day** (15-minute version) for the rest of the month.
* **Question:** Was it hard to open your interpreter today? If so, automate the startup.


* **Feb 05 (Thu): The Dictionary Discovery**
* **Task:** Create a table representing a "Player" or "Neovim Config" using Key-Value pairs (e.g., `hp = 100` or `theme = "dark"`).
* **Focus:** Practice adding and removing keys dynamically using `table["key"] = nil`.


* **Feb 06 (Fri): The Array Archetype**
* **Task:** Create a list-style table and practice `table.insert` and `table.remove`.
* **Professional Standard:** Observe how the length operator `#myTable` changes as you add/remove items.


* **Feb 07 (Sat): Iteration Mastery – `pairs` vs. `ipairs**`
* **Task:** Write a script that iterates over a mixed table (containing both numeric indices and string keys) using both `pairs` and `ipairs`.
* **Goal:** Understand why `ipairs` stops at the first "hole" or non-numeric key.


* **Feb 08 (Sun): Weekend Integration – The Config Manager**
* **Task:** Build a small "Setting Manager" table for a hypothetical Neovim plugin.
* **Requirements:** Store default settings in one table and user overrides in another; write a function to merge them.



---

### 🧠 Performance Protocols

* **The 2-Minute Rule:** If you feel resistance, just type `local config = {}`.
* **Habit Stacking:** "After I finish my **CKA Speed Drill**, I will spend 15 minutes on my **Lua Table of the Day**".
* **Never Miss Twice:** If Saturday is a wash, Sunday is a mandatory "Green" day for your Lua foundation.
