# B. Префиксная подпоследовательность

<p style="text-align: left">
  <a href="#русский">Русский</a> ・ <a href="#english">English</a>
</p>

---

## Русский

### Постановка задачи

Даны две двоичные строки `a` и `b`. Нужно найти максимальное `k`, такое что префикс строки `a` длины `k` является подпоследовательностью строки `b`.

### Идея решения

Подпоследовательность строится удалением символов, порядок оставшихся при этом не меняется. Значит, чтобы проверить, является ли префикс `a` подпоследовательностью `b`, достаточно идти по `b` слева направо и жадно "забирать" символы `a` по порядку, как только они встретились.

Жадность здесь корректна: если текущий символ `a` совпал с текущим символом `b`, нет смысла пропускать это совпадение и искать его позже — более раннее совпадение никогда не хуже. Символ `b`, который не подошёл, просто пропускается — на ответ он не влияет, но и не мешает.

Как только указатель по `a` дошёл до конца — весь `a` является подпоследовательностью `b`, ответ равен `n`. Если раньше кончился `b` — сколько символов `a` успели совпасть, столько и есть ответ.

### Как это реализовано в коде

```go
func solution(a, b string) int {
    moveA, moveB := 0, 0

    for moveA < len(a) && moveB < len(b) {
       if a[moveA] == b[moveB] {
          moveA++
       }

       moveB++
    }

    return moveA
}
```

`moveA` — сколько символов `a` уже сопоставлено, `moveB` — текущая позиция в `b`.

На каждом шаге сравниваются `a[moveA]` и `b[moveB]`. Если символы совпали, `moveA` увеличивается — очередной символ префикса `a` найден. `moveB` увеличивается в любом случае: пройденный символ `b` больше не понадобится, совпал он или нет.

Цикл останавливается, когда закончился `a` (тогда `moveA == len(a)`, весь `a` — подпоследовательность `b`) либо когда закончился `b` раньше (тогда `moveA` — это и есть максимальный найденный префикс).

Возвращаемое значение `moveA` и есть искомое `k`.

### Пример

```text
a = "1100"
b = "10"
```

`moveA=0, moveB=0`: `a[0]='1'`, `b[0]='1'` — совпали → `moveA=1`, `moveB=1`.

`moveA=1, moveB=1`: `a[1]='1'`, `b[1]='0'` — не совпали → `moveB=2`.

`moveB=2 == len(b)` — цикл завершается.

Ответ: `1`.

Проверка вручную: префикс `"1"` — подпоследовательность `"10"` (первый символ). Префикс `"11"` подпоследовательностью `"10"` уже не является — в `b` только одна `'1'`. Значит `k=1` — верно.

### Сложность

Каждый шаг цикла увеличивает `moveB`, а он ограничен `len(b)`, поэтому цикл делает не больше `m` итераций. На один набор входных данных:

```text
время: O(n + m)
память: O(1) не считая хранения самих строк
```

Суммарные ограничения задачи (сумма `n` и сумма `m` по всем наборам не превышают `2·10^5`) укладываются в этой оценке с большим запасом — весь ввод обрабатывается за линейное время.


---

## English

### Problem

Two binary strings `a` and `b` are given. Find the maximum `k` such that the prefix of `a` of length `k` is a subsequence of `b`.

### Idea

A subsequence is built by deleting characters while keeping the relative order of what's left. So checking whether a prefix of `a` is a subsequence of `b` comes down to walking through `b` left to right and greedily "claiming" characters of `a` in order, as soon as they show up.

The greedy approach is correct here: if the current character of `a` matches the current character of `b`, there's no reason to skip that match and look for a later one — an earlier match is never worse. A character of `b` that doesn't match is simply skipped — it doesn't affect the answer, but it doesn't get in the way either.

Once the pointer over `a` reaches the end, the whole `a` is a subsequence of `b`, and the answer is `n`. If `b` runs out first, however many characters of `a` matched by that point is the answer.

### How the code does it

```go
func solution(a, b string) int {
    moveA, moveB := 0, 0

    for moveA < len(a) && moveB < len(b) {
       if a[moveA] == b[moveB] {
          moveA++
       }

       moveB++
    }

    return moveA
}
```

`moveA` counts how many characters of `a` have been matched so far, `moveB` tracks the current position in `b`.

At each step, `a[moveA]` is compared with `b[moveB]`. If they match, `moveA` advances — one more prefix character of `a` has been found. `moveB` advances regardless: the current character of `b` won't be needed again, matched or not.

The loop stops either when `a` runs out (`moveA == len(a)`, the whole `a` is a subsequence of `b`) or when `b` runs out first (`moveA` at that point is the longest prefix found).

The returned `moveA` is exactly the answer `k`.

### Example

```text
a = "1100"
b = "10"
```

`moveA=0, moveB=0`: `a[0]='1'`, `b[0]='1'` — match → `moveA=1`, `moveB=1`.

`moveA=1, moveB=1`: `a[1]='1'`, `b[1]='0'` — no match → `moveB=2`.

`moveB=2 == len(b)` — loop ends.

Answer: `1`.

Manual check: the prefix `"1"` is a subsequence of `"10"` (its first character). The prefix `"11"` is not, since `b` only has one `'1'`. So `k=1` — correct.

### Complexity

Every loop iteration advances `moveB`, which is bounded by `len(b)`, so the loop runs at most `m` times. Per test case:

```text
time: O(n + m)
memory: O(1) aside from storing the strings themselves
```

The problem's overall constraints (the sum of `n` and the sum of `m` across all test cases each capped at `2·10^5`) fit comfortably within this bound — the whole input is processed in linear time.

---

Два указателя здесь работает потому, что строки бинарные и порядок символов в `b` фиксирован: каждое совпадение стоит забирать сразу, как только оно встретилось, потому что откладывать его смысла нет — более позднее совпадение с тем же символом ничем не лучше.

The greedy approach works here because the strings are binary and the order of characters in `b` is fixed: every match should be claimed the moment it appears, since postponing it makes no sense — a later match on the same character is never better.