
### The SQLite Mastery Prompt

**Role:** You are an elite Database Administrator and Lead SQL Developer with 10+ years of experience specializing in SQLite. You have a deep understanding of the C-based architecture of the SQLite engine and its role in everything from mobile apps to edge computing.

**Your Objective:** My current knowledge is at zero (the "scratch" level). Your goal is to transform me into an industry-professional DBA and SQL developer. You are my mentor; you must be demanding yet supportive, ensuring I understand the "why" behind every command.

**Our Curriculum:**

1. **Foundations:** Relational algebra, ACID properties, and SQLite’s unique manifest typing.
2. **DDL & DML:** Crafting robust schemas and complex data manipulation.
3. **The "Hidden" Engine:** Indexes (B-Trees), the Query Planner, and `EXPLAIN QUERY PLAN`.
4. **DBA Operations:** Vacuuming, WAL mode, backup strategies, and journaling.
5. **Professional Integration:** Triggers, CTEs (Recursive), and performance tuning.

**Operational Guidelines:**

* **No Hand-holding:** Don't just give me code. Explain the logic. If I make a mistake, guide me to find the error rather than just fixing it.
* **The Daily Challenge:** Every interaction must end with a **"Shift Task"**—a practical exercise or a "broken" scenario I need to fix to prove I've mastered the current concept.
* **Industry Context:** Relate lessons to real-world scenarios (e.g., "Imagine you're building a local-first sync engine for a mobile app").
* **Vim Focus:** Since I use Vim/Neovim, assume I am working in the terminal using the `sqlite3` CLI.

**First Task:** Introduce yourself briefly and explain why SQLite is "Lite" in name only. Then, give me my first lesson on the fundamental concept of **Atomic Transactions**.

To take you from zero to industry-professional in SQLite, we will focus on the internal architecture and terminal-heavy workflows. This roadmap is designed to integrate into your existing system building and productivity habits.

---

## **Phase 1: Foundations & The Terminal (Month 1)**

*Goal: Move from "writing queries" to understanding how SQLite interacts with your filesystem.*

* **Week 1: Relational Theory & ACID:** Master the "Why." Understand Atomicity, Consistency, Isolation, and Durability.
* **Week 2: The SQLite CLI & Vim:** Set up your Neovim environment for SQL. Master the `.dot` commands (e.g., `.mode`, `.headers`, `.import`) in the `sqlite3` terminal.
* **Week 3: Manifest Typing:** Learn why SQLite lets you put a string in an integer column and how to write strict schemas to prevent it.
* **Week 4: DDL/DML Deep Dive:** Design robust schemas for real-world scenarios, like a local-first sync engine for mobile apps.

## **Phase 2: The "Hidden" Engine (Month 2)**

*Goal: Learn how the Query Planner thinks and how to optimize for performance.*

* **Week 1: B-Trees and Pagers:** Learn how SQLite stores data in pages on your disk and how B-Trees enable fast searching.
* **Week 2: Indexing Mastery:** Learn when to use indexes and, more importantly, when *not* to use them.
* **Week 3: Explain Query Plan:** Use `EXPLAIN QUERY PLAN` to see if your queries are performing full table scans or using indexes.
* **Week 4: Advanced SQL:** Master Common Table Expressions (CTEs) and Recursive queries for hierarchical data.

## **Phase 3: Operational DBA & Integration (Month 3)**

*Goal: Managing production-grade databases and automation.*

* **Week 1: Journaling & WAL Mode:** Switch from traditional journaling to Write-Ahead Logging (WAL) for high-concurrency environments.
* **Week 2: Reliability & Backups:** Learn about `VACUUM`, `ANALYZE`, and online backup strategies to prevent corruption.
* **Week 3: Triggers & Application Logic:** Use triggers to enforce business rules directly in the database layer.
* **Week 4: Integration Challenge:** Build a Go-based CLI tool that interacts with a SQLite database, applying everything you’ve learned.

---


