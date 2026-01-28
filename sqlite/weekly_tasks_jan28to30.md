Since we are already mid-week, we are going to compress the foundational theory of **Atomic Transactions** and **ACID** into a high-intensity sprint. This will align with your existing discipline of daily practice and system building.

Here is your **Phase 1: Week 1** task list for the remainder of the week:

---

### **Thursday, Jan 29: The ACID Deep Dive**

* **The Theory:** Research and write a brief summary in your Neovim environment regarding the four pillars of ACID (Atomicity, Consistency, Isolation, Durability).
* **The "Why":** Identify one scenario in your "Go Mastery" book or CLI combat game where a lack of **Atomicity** would corrupt your game state (e.g., a player loses gold but the item isn't added to inventory).
* **Vim Practice:** Use your daily dedicated Vim practice time to create a `.sql` file with these definitions, focusing on rapid navigation between the pillars.

### **Friday, Jan 30: Transactional Mechanics**

* **CLI Mastery:** Open `sqlite3` and practice the manual control of transactions using `BEGIN`, `COMMIT`, and `ROLLBACK`.
* **Journaling Research:** Investigate the "Rollback Journal." Find where this temporary file lives on your disk during an active transaction before it is committed.
* **Refinement:** Update your study notes to explain how SQLite ensures **Durability** even if your computer loses power mid-write.

### **Saturday, Jan 31: Relational Modeling**

* **Schema Design:** Design a basic 3-table relational schema for your CLI combat game (e.g., `Players`, `Inventory`, `Items`).
* **Constraints:** Implement `PRIMARY KEY` and `NOT NULL` constraints to enforce **Consistency** at the database level.
* **Review:** Prepare your summary of the week's progress for your Sunday planning session.

### **Sunday, Feb 1: Weekly Planning & Integration**

* **System Building:** During your scheduled Sunday weekly planning, allocate specific slots for Phase 1: Week 2 (The SQLite CLI & Vim).
* **Financial Review:** As part of your daily financial transaction review, consider how the bank uses the **Isolation** property of ACID to handle your account updates.
* **Next Step:** Finalize your "Shift Task" from our previous interaction to officially move into Week 2.

