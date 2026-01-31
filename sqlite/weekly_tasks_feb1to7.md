Based on your current progress and the "Foundations & The Terminal" phase of your curriculum, here is your new **SQLite Weekly Schedule** for **February 2 – February 8, 2026**.

I have moved your unfinished tasks—the **Relational Modeling** challenge and the **Weekly Planning/Integration**—to the beginning of this week. This ensures your conceptual foundation of **ACID** is solid before you master the **sqlite3 CLI** and **Vim** integration.

### 🗄️ SQLite Weekly Schedule: Feb 2 – Feb 8, 2026

**Theme:** *The CLI & The Filesystem*

* **Feb 02 (Mon): Final Review – Relational Schema Design**
* **Task:** Design a 3-table schema for your CLI combat game (e.g., `Players`, `Inventory`, `Items`).
* **Focus:** Implement `PRIMARY KEY` and `NOT NULL` constraints to enforce **Consistency** at the database level.


* **Feb 03 (Tue): Weekly Audit & Integration**
* **Task:** Perform the carried-over **Financial Review**. Analyze how your bank uses **Isolation** to prevent two transactions from updating your account at the same time.
* **Shift Task:** Formally finalize your Week 1 concepts to move into Week 2: The CLI.


* **Feb 04 (Wed): CLI Mastery – The `.dot` Commands**
* **Task:** Open the `sqlite3` terminal and master formatting commands: `.mode box`, `.headers on`, and `.nullvalue 'NULL'`.
* **Focus:** Practice creating a temporary in-memory database using `sqlite3 :memory:` to test schemas quickly.


* **Feb 05 (Thu): Neovim SQL Integration**
* **Task:** Configure your Neovim environment to work with `.sql` files.
* **Vim Practice:** Use `!sqlite3 mygame.db < schema.sql` from within Vim to execute your scripts without leaving the editor.


* **Feb 06 (Fri): Data Ingestion – `.import` and `.dump**`
* **Task:** Create a `.csv` file with dummy player data and use the `.import` command to load it into your `Players` table.
* **Reverse:** Use `.dump` to export your entire database schema and data to a text file for backup.


* **Feb 07 (Sat): Weekend Review – Filesystem Interaction**
* **Task:** Locate your `.db` file on your disk and observe its size.
* **Lab:** Perform a massive `INSERT` (1,000+ rows) and then a `DELETE`. Observe the file size and research why the size does not automatically decrease (Hint: `VACUUM`).


* **Feb 08 (Sun): Integration Lab – The Game DB Prototype**
* **Task:** Build a fully functional `game.db` using only the CLI and Vim scripts.
* **Goal:** Successfully join your `Players` and `Inventory` tables to show a "Player Loadout" report using the `.output` command to save the results to a file.



---

### 🧠 Performance Protocols

* **The 2-Minute Rule:** If you feel resistance, your only goal is to type `sqlite3 game.db` in your terminal.
* **Never Miss Twice:** If Monday’s schema design is a wash, Tuesday is a mandatory "Green" day to maintain your trajectory toward industry-professional DBA status.

**Shift Task:** Can you explain how SQLite’s `.mode` commands differ from standard SQL syntax, and why they are essential for a terminal-heavy workflow?
