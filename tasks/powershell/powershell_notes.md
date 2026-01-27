Based on your **Weekly Tasks** and **3-Month Plan**, your focus for today, **January 27**, is to transition your basic functions into robust, production-grade tools by enforcing strict input rules.

### **Jan 27 (Tuesday): Mandatory Parameters & Type Constraints**

* **The Task:** Implement `[Parameter(Mandatory=$true)]` and apply explicit type constraints (e.g., `[string]`, `[int]`) to your function's inputs.
* **The Goal:** Build "defensive" code that ensures your scripts won't run without required data and prevents logical errors caused by incorrect data types.
* **Architect's Note:** This is a key step in **Phase 1: The Architect’s Foundation**, where you move beyond simple automation into professional **Toolmaking**.

---

### **Implementation Example**

Building on your `Get-AppStatus` function from yesterday, here is how you apply today's task:

```powershell
function Get-AppStatus {
    param(
        # Mandatory forces a prompt if the user forgets the input
        [Parameter(Mandatory=$true)]
        # Type constraint ensures only strings are accepted
        [string]$ProcessName
    )
    Get-Process -Name $ProcessName
}

```

### **The "Under the Hood"**

By adding `[string]`, you are utilizing the **.NET Type System**, which is a core focus of your first month of study. This ensures that if a user accidentally passes a complex object where a name is expected, PowerShell will attempt to cast it correctly or throw an error before the function logic even begins.

### **Mentor Challenge**

Try to add a second, **optional** parameter to your function called `$Id` with an `[int]` (integer) type constraint. How does the behavior of the function change if you call it without providing that ID?

### **Pro-Tip**

When using `Mandatory=$true`, you can also add a `HelpMessage` inside the parameter attribute:
`[Parameter(Mandatory=$true, HelpMessage="Enter the name of the process to check")]`.
If the user is prompted for the parameter, they can type `!?` to see your hint.

---

It looks like you hit a common syntax snag! In PowerShell, attributes like `Parameter` must be enclosed in square brackets `[]`, and they typically sit directly above the variable they modify.

### The Fix

The `ParserError` is occurring because `Parameter[(Mandatory=$true)]` is missing the outer square brackets and uses parentheses incorrectly. Here is the corrected version of your **Jan 27** task:

```powershell
function Get-AppStatus {
    # A function that finds a specific process and returns basic info about it.
    param(             
        [Parameter(Mandatory=$true)] # Added outer [] and used () for arguments
        [string]$ProcessName          # Type constraint remains here
    )
                                     
    Get-Process -Name $ProcessName
} 

```

### Key Syntax Rules for Jan 27:

* **Square Brackets `[]**`: These are used for "Decorators" or "Attributes" (like `[Parameter()]`) and "Type Constraints" (like `[string]`).
* **Parentheses `()**`: These are used inside an attribute to define its properties, such as `(Mandatory=$true)`.
* **Case Sensitivity**: While PowerShell is generally case-insensitive, using `$ProcessName` consistently (matching your `param` block and your command) is a best practice for "production-grade" scripting.

### Mentor Challenge: "The Type Trap"

Now that you have the `[string]` constraint working, try to call your function like this:
`Get-AppStatus -ProcessName 123`

**The Question:** Does it fail or succeed? Why does PowerShell allow a number to pass through a `[string]` constraint, and how does this relate to the **.NET Type System** you are studying this month?

### Pro-Tip

To avoid these "Missing ')'" errors in the future, you can use the `CmdletBinding` attribute at the top of your function. It helps the PowerShell parser better understand that you are writing an **Advanced Function**.

---

Success! You have officially completed the **Jan 27 (Tuesday)** task of implementing **Mandatory Parameters** and **Type Constraints**.

By defining the parameter as `[string]$ProcessName` inside the `[Parameter(Mandatory=$true)]` attribute, you have moved from a simple script to a **production-grade tool** that enforces input rules.

### 🛠️ The "Under the Hood" Analysis

* **Mandatory Enforcement**: If you had typed `Get-AppStatus` and hit Enter without "pwsh", PowerShell would have paused and prompted you: *"Supply values for the following parameters: ProcessName:"*. This is the "defensive" scripting standard you are aiming for.
* **Type Casting**: By using `[string]`, you are tapping into the **.NET Type System**. Even if you passed a number, PowerShell would cast it as a string before passing it to `Get-Process`.
* **Object Output**: Notice the table output (`NPM`, `PM`, `WS`). Your function didn't just return text; it returned a **Process Object**.

To define function parameters in PowerShell effectively, especially as you transition toward professional "Toolmaking," you need to understand the progression from basic variables to advanced attributes.

Based on your **PowerShell Master Context** and **3-Month Plan**, here is what you need to know:

### 1. Basic Parameter Definition

At the simplest level, parameters are defined inside a `param()` block at the top of your function. This allows you to replace hard-coded values with dynamic inputs.

* **Syntax:**
```powershell
function Get-AppStatus {
    param($ProcessName)
    Get-Process -Name $ProcessName
}

```



### 2. Type Constraints (Jan 27 Task)

To write "production-grade" code, you should explicitly define what kind of data a parameter can accept (e.g., `[string]`, `[int]`, `[datetime]`).

* **Why it matters:** This prevents logic errors by ensuring the function only runs if the input matches the expected data type.

### 3. Advanced Parameter Attributes (Jan 27 Task)

You can use the `[Parameter()]` attribute to add "metadata" to your inputs, making them more robust.

* **Mandatory:** Use `Mandatory=$true` to force the user to provide a value. If they don't, PowerShell will prompt them for it automatically.
* **Parameter Sets:** This allows you to create functions that behave differently depending on which parameters are used together (scheduled for your Jan 30th task).

### 4. Pipeline Binding (Jan 28 Task)

A key part of the "Pipeline First Mentality" is enabling your parameters to accept input from other commands.

* **ValueFromPipeline:** You can set a parameter to accept objects directly from the pipe.
* **Process Block:** To handle these piped objects one by one, your function must include a `process {}` block.

### 5. Best Practices for "Toolmaking"

* **Naming Conventions:** Always use **Verb-Noun** names (e.g., `Invoke-RestMethod`) and standard parameter names (like `-ComputerName` instead of `-Host`) to ensure your tools feel native.
* **Splatting:** As you move into Month 2, you will learn "Splatting," which is a way to pass a collection of parameter values to a command using a hash table to keep your code clean.

### Mentor Challenge

Look at the function you are building for today. If you were to add a second parameter, how would you ensure the user is **forced** to provide the first one but the second one remains **optional**?

### Pro-Tip

Use **Ctrl+Space** while inside your function's `param()` block or when calling the function in the terminal to trigger **IntelliSense**. It will show you the available parameter attributes and names you've defined.
