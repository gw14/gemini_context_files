Since we are building **mental infrastructure**, we aren’t just learning shortcuts; we are learning **process management**. This roadmap is designed to move you from "I use tmux for splits" to "I treat tmux as a programmable terminal multiplexer."

We will use the **man tmux** manual as our primary terrain. Here is your "Zero to Hero" architectural roadmap.

---

## 🗺️ The Tmux Architect Roadmap

### Phase 1: The Client-Server Foundations (The "Socket" Layer)

* **The Goal:** Understand that tmux is not just a program, but a persistent server.
* **Key Concepts:** The tmux server vs. client, sessions, and the `/tmp` socket.
* **The Manual Deep-Dive:** `GENERAL` and `COMMANDS` sections.
* **Skill Tutor Check:** Can you explain what happens to your processes when you detach a session?

### Phase 2: The Hierarchy (Sessions, Windows, Panes)

* **The Goal:** Master the 3-tier structure to organize your CKA study labs and Go projects.
* **Key Concepts:** Creating, naming, and switching between targets. Understanding `target-window` (`-t`) and `target-pane`.
* **The Manual Deep-Dive:** `WINDOWS AND PANES` and `SESSIONS` sections.
* **Skill Tutor Check:** Can you move a pane from one window in Session A to a specific window in Session B without using a mouse?

### Phase 3: Navigation & Buffer Mastery (The "Vim" Bridge)

* **The Goal:** Seamlessly move data between your terminal and your Neovim buffers.
* **Key Concepts:** Copy-mode, vi-keys, paste-buffers, and the `choose-buffer` command.
* **The Manual Deep-Dive:** `BUFFERS` and `KEY BINDINGS` (specifically `mode-keys`).
* **Skill Tutor Check:** Can you capture the last 100 lines of a failing Go build and pipe them directly into a new file without leaving the terminal?

### Phase 4: Scripting & Automation (The "Go" Integration)

* **The Goal:** Automate your entire CKA environment and Go workspace with a single command.
* **Key Concepts:** `source-file`, custom keybindings, and using `tmux` as a CLI tool (running tmux commands from *outside* tmux).
* **The Manual Deep-Dive:** `COMMANDS` and `FORMATS`.
* **Skill Tutor Check:** Can you write a shell script that checks if a session exists, and if not, creates it with 3 specific panes (Editor, Compiler, Logs)?

### Phase 5: Hooks, Formats, and Advanced Styling

* **The Goal:** Make the status bar an information dashboard for your system.
* **Key Concepts:** `set-hook`, status-line customization, and using `strftime` for the clock.
* **The Manual Deep-Dive:** `HOOKS`, `STATUS LINE`, and `STYLES`.
* **Skill Tutor Check:** Can you trigger a notification (or change the status bar color) automatically when a long-running Go test suite finishes?

---

## 🛠️ Let’s Start Phase 1: The Foundations

To kick this off, I’m putting on the **10-year Power User** hat.

**The Architect:** "You want to master the terminal? Then you must understand the ghost in the machine. Before you ever touch a 'prefix' key, you must understand the separation of the Server and the Client."

**Your First Mission:**
Open your terminal and type `man tmux`. I want you to find the **`DESCRIPTION`** section at the very beginning.

**Guidelines for the Read:**

* **Look out for:** The distinction between the **server**, **sessions**, **windows**, and **panes**. It's a nested hierarchy; if you don't visualize it now, your config will be a mess later.
* **The Nuance:** Pay attention to how tmux handles the **TERM** environment variable and the socket file.

**Socratic Question:** "Once you’ve looked at the `DESCRIPTION`, tell me: If you kill the terminal window (the GUI app) on your laptop, why does the process running inside tmux keep living? What part of the tmux architecture is responsible for that 'immortality'?"

**Would you like me to wait for your "Rubber Duck" explanation of the Server-Client relationship before we move to the first Speed Drill?**
