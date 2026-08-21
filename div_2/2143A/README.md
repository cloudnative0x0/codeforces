# All Lengths Subtraction

<p style="text-align: left">
  <a href="#русский">Русский</a> ・ <a href="#english">English</a>
</p>

---

## Русский

Дана перестановка `p` длины `n`. Нужно выполнить ровно `n` операций — для каждого `k` от `1` до `n` выбрать подотрезок длины ровно `k` и вычесть из каждого его элемента единицу. Требуется определить, можно ли подобрать эти подотрезки так, чтобы после всех операций весь массив стал нулевым.

Порядок выполнения операций на итоговую сумму не влияет — вычитание коммутативно, важно только, какие именно подотрезки выбраны для каждой длины, а не в какой момент они применены. Поэтому задача сводится к чисто комбинаторному вопросу: существует ли набор из `n` подотрезков длин `1, 2, ..., n` (по одному на каждую длину), сумма индикаторов которых совпадает с `p`.

### Идея решения

Ключевое наблюдение: на шаге `k` до конца процедуры остаётся `n - k + 1` операций, включая текущую (это операции с длинами `k, k+1, ..., n`). Если у элемента `arr[i]` в этот момент осталось вычесть больше, чем `n - k` (то есть `arr[i] >= n - k + 1`), — единиц у него ровно столько же или больше, сколько операций впереди, а значит этот элемент обязан участвовать буквально в каждой из оставшихся операций, включая текущую. Пропустить его сейчас нельзя: тогда оставшихся операций не хватит, чтобы довести его до нуля.

Такие элементы находятся сканированием массива на каждом шаге `k`: отмечается самый левый и самый правый индекс, где `arr[i] > n - k`. Поскольку подотрезок обязан быть непрерывным, окно текущей операции обязано охватить весь диапазон от самого левого до самого правого такого индекса целиком — даже если между ними есть элементы, которым сейчас вычитание не обязательно, отказаться от них нельзя, отрезок нельзя разорвать.

Если этот вынужденный диапазон уже длиннее, чем `k`, задача неразрешима — ответ `NO`, потому что подотрезок длины `k` физически не способен покрыть более широкий обязательный диапазон. Если вынужденных элементов нет вовсе, окно берётся произвольно, по умолчанию с начала массива. Если вынужденный диапазон короче `k`, окно расширяется до нужной длины — сначала вправо, а когда упирается в конец массива, влево. Расширение безопасно: лишние элементы, попавшие под вычитание раньше срока, не портят решение, потому что все последующие операции только длиннее текущей и им проще дотянуться до дальних позиций.

После того как окно найдено, из каждого его элемента вычитается единица, и цикл переходит к следующему `k`. После `n` шагов проверяется, что весь массив обнулился; если нет — ответ `NO`, иначе — `YES`.

Отдельной структуры данных здесь не требуется. Вся работа идёт напрямую с массивом `arr` — копией `p` — и двумя индексами `start`/`end`, задающими текущее окно. Это классический двухуказательный (two-pointer) приём: окно не движется вдоль массива, а расширяется от вынужденного ядра наружу до нужного размера.

### Использование

```go
p := []int{1, 2, 3}
ans := solution(3, p)
// k=1: вынужден индекс 2 (значение 3 > 2), окно [2,2], arr = [1,2,2]
// k=2: вынуждены индексы 1,2 (оба >1), окно [1,2], arr = [1,1,1]
// k=3: вынуждены индексы 0,1,2 (все >0), окно [0,2], arr = [0,0,0]
// ans = "YES"
```

### Операции

| Шаг | Сложность | Описание |
|---|---|---|
| поиск вынужденного диапазона | O(n) | скан массива на поиск `arr[i] > n-k` |
| расширение окна | O(n) | доведение окна до длины `k` |
| вычитание по окну | O(n) | `arr[i]--` для `i` в `[start, end]` |
| проверка результата | O(n) | все ли элементы равны нулю |

Каждый из `n` шагов цикла — `O(n)`, итого `O(n²)` на тест-кейс, `O(n)` дополнительной памяти под копию массива. При `n ≤ 100` и `t ≤ 100` это укладывается в лимит на один тест (1 секунда) с большим запасом.

---

## English

Given a permutation `p` of length `n`. Exactly `n` operations must be performed — for each `k` from `1` to `n`, pick a subarray of length exactly `k` and subtract one from every element in it. The task is to decide whether these subarrays can be chosen so that after all operations the whole array becomes zero.

The order in which the operations are applied doesn't affect the final sum — subtraction is commutative, only which subarray is picked for each length matters, not when it's applied. So the task reduces to a purely combinatorial question: does there exist a set of `n` subarrays of lengths `1, 2, ..., n` (one per length) whose combined indicator sum equals `p`.

### Solution idea

The key observation: at step `k` there are `n - k + 1` operations left, including the current one (these are the operations of lengths `k, k+1, ..., n`). If an element `arr[i]` still needs more than `n - k` more subtractions at this point (that is, `arr[i] >= n - k + 1`), it needs as many or more decrements as there are operations left — so it must be included in literally every remaining operation, starting with this one. Skipping it now is fatal: the remaining operations alone would no longer be enough to bring it down to zero.

Such elements are found by scanning the array at each step `k`: the leftmost and rightmost index where `arr[i] > n - k` are recorded. Since a subarray must be contiguous, the current operation's window is forced to span the entire range from the leftmost to the rightmost such index — even elements between them that aren't themselves forced yet can't be skipped, because the segment can't be split.

If this forced range is already longer than `k`, the task is unsolvable — the answer is `NO`, because a length-`k` subarray can't physically cover a wider mandatory range. If no element is forced at all, the window is picked arbitrarily, defaulting to the start of the array. If the forced range is shorter than `k`, the window is expanded to the required length — first to the right, and once it hits the end of the array, to the left. Expanding is safe: elements pulled in early that didn't strictly need it yet aren't a problem, because every later operation is longer than the current one and has an easier time reaching far-off positions.

Once the window is set, one is subtracted from every element in it, and the loop moves to the next `k`. After `n` steps, the array is checked for being all zero; if not, the answer is `NO`, otherwise `YES`.

No auxiliary data structure is needed here. All the work happens directly on the array `arr` — a copy of `p` — and two indices `start`/`end` marking the current window. This is a classic two-pointer technique: the window doesn't slide along the array, it expands outward from a forced core to the required size.

### Usage

```go
p := []int{1, 2, 3}
ans := solution(3, p)
// k=1: index 2 is forced (value 3 > 2), window [2,2], arr = [1,2,2]
// k=2: indices 1,2 are forced (both >1), window [1,2], arr = [1,1,1]
// k=3: indices 0,1,2 are forced (all >0), window [0,2], arr = [0,0,0]
// ans = "YES"
```

### Operations

| Step | Complexity | Description |
|---|---|---|
| find forced range | O(n) | scan the array for `arr[i] > n-k` |
| expand window | O(n) | grow the window to length `k` |
| subtract over window | O(n) | `arr[i]--` for `i` in `[start, end]` |
| check result | O(n) | verify every element is zero |

Each of the `n` loop steps is `O(n)`, giving `O(n²)` per test case and `O(n)` extra memory for the array copy. With `n ≤ 100` and `t ≤ 100` this fits comfortably within the 1-second per-test limit.

---

<br>

> Окно не выбирает, кого пощадить — оно вынуждено охватить того, кому больше некуда деться. Всё остальное — просто заполнение до нужной длины.
>
> *The window doesn't choose who to spare — it's forced to cover whoever has nowhere left to go. Everything else is just padding to the required length.*