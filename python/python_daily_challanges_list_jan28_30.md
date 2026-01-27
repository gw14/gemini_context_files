To keep our momentum high, we are going to spend the rest of this week cementing your mental model of **Object Identity** and **Mutability**.

If you don't master these now, you will spend your career chasing "ghost bugs" where data changes in one part of your app for no apparent reason.

---

## 🗓️ Wednesday, Jan 28: The "Mutable Trap"

**The Goal:** Understand why some objects can change in place and others cannot.

* **Task 1 (Reading):** Research the difference between **Mutable** (Lists, Dicts, Sets) and **Immutable** (Strings, Integers, Tuples) objects in Python.
* **The Challenge:**
1. Define a string `s = "Python"`. Try to change the first letter to "J" using `s[0] = "J"`. Observe the error.
2. Define a list `l = ["P", "y", "t", "h", "o", "n"]`. Change the first letter using `l[0] = "J"`.
3. **The Question:** If strings are immutable, how does `s = s + "ic"` work? Does it change the original string, or create a new one? Use `id()` to prove your answer.



---

## 🗓️ Thursday, Jan 29: The Shadow Clone (Shallow vs. Deep)

**The Goal:** Learn how to copy data safely without leaving "tethers" to the original.

* **Task 1 (Reading):** Look up the `copy` module in the Python documentation, specifically `copy.copy()` vs `copy.deepcopy()`.
* **The Challenge:**
1. Create a nested list: `original = [[1, 2], [3, 4]]`.
2. Create a **shallow copy** named `shallow_copy`.
3. Create a **deep copy** named `deep_copy`.
4. Modify the *inner* list of the `original` (e.g., `original[0][0] = "X"`).
5. **The Analysis:** Print all three lists. Which ones changed? Explain why the shallow copy failed to protect the inner data.



---

## 🗓️ Friday, Jan 30: The "Interning" Mystery

**The Goal:** Discover how Python optimizes memory for small integers and strings.

* **Task 1 (Experiment):** Run the following code and look closely at the results:
```python
a = 256
b = 256
print(a is b)

x = 257
y = 257
print(x is y)

```


* **The Challenge:** 1. Why does Python say the first pair is the "same" object but the second pair is "different," even though the numbers are just one digit apart?
2. Research **"Integer Interning"** in Python.
3. **The Professional Takeaway:** Why is it dangerous to use `is` to compare values like numbers or user input? (Hint: Use `==` for values, `is` for identity).

---

### 💡 Pro-Tip for your Records

Create a folder named `python_mastery`. Inside, create a file for each day (e.g., `jan28_mutability.py`). Comment your code heavily with what you *think* will happen before you run it.
