# 1D Ластик

<p style="text-align: left">
  <a href="#русский">Русский</a> ・ <a href="#english">English</a>
</p>

---

## Русский

Условие: полоска из `n` ячеек, каждая чёрная (`B`) или белая (`W`). За одну операцию можно взять любые `k` идущих подряд ячеек и покрасить их в белый. Нужно минимальное число операций, чтобы не осталось ни одной чёрной ячейки.

Ключевое наблюдение: операция — это отрезок фиксированной длины `k`, и его можно поставить где угодно, лишь бы он был внутри полоски. Если идти по строке слева направо и встретить чёрную ячейку, не покрытую ни одной уже поставленной операцией, эту ячейку всё равно придётся чем-то закрыть. Среди всех отрезков длины `k`, которые её закрывают, самый выгодный — тот, что начинается ровно в этой ячейке: он не может начаться левее (иначе что-то потрачено на уже белые или уже закрытые ячейки слева) и не может начаться правее (иначе сама ячейка останется не покрытой). Начало ровно в позиции текущей чёрной ячейки — это одновременно допустимо и даёт максимальный охват вправо, а значит с наибольшей вероятностью погасит и будущие чёрные ячейки бесплатно.

### Идея решения

Идём по строке одним проходом и держим переменную `move` — индекс начала последней поставленной операции. Пока текущая позиция `i` попадает в диапазон `[move, move+k-1]`, она уже белая по построению, и её не нужно ничего делать, даже если в исходной строке там стоит `B` — early в реальности такого не бывает, потому что зона `[move, move+k-1]` в строке уже вся закрашена предыдущей операцией.

Как только встречается `B`, для которой `i - move >= k` (то есть она вне диапазона последней операции), ставится новая операция, начинающаяся в `i`, и `move` обновляется на `i`. Условие `i - move >= k` — это ровно проверка «текущая позиция не покрыта последним отрезком».

Начальное значение `move = -k` — техническая деталь, чтобы для самой первой чёрной ячейки `i - move = i + k` было заведомо `>= k` при любом `i >= 0`, то есть первая операция всегда сработает, без отдельной проверки на «ещё ничего не ставили».

### Асимптотика

Один проход по строке — `O(n)` на тест-кейс, никакой дополнительной памяти, кроме двух счётчиков. По условию сумма `n` по всем тестам ограничена `2·10^5`, так что суммарная работа по всем `t` тестам — `O(N)`, где `N` — суммарная длина всех строк.

### Использование

```go
s := "BBWBBBWB"
k := 3
ans := solution(s, k)
// i=0: B, 0-(-3)=3>=3 → операция [0,2], move=0
// i=1,2: покрыты
// i=3: B, 3-0=3>=3 → операция [3,5], move=3
// i=4,5: покрыты
// i=7: B, 7-3=4>=3 → операция [7,9], move=7
// ans = 3
```

### Операции

| Шаг | Сложность | Описание |
|---|---|---|
| проход по строке | O(n) | одна переменная `move` хранит начало последней операции |
| проверка покрытия | O(1) | `i - move >= k` — вышли ли за пределы последнего отрезка |
| постановка операции | O(1) | сдвиг `move` на текущий индекс `i` |

---

## English

Setup: a strip of `n` cells, each black (`B`) or white (`W`). One operation picks any `k` consecutive cells and turns them white. Find the minimum number of operations to remove all black cells.

The key observation: an operation is a fixed-length segment of `k` cells that can start anywhere within the strip. Scanning left to right, once a black cell is found that isn't covered by any operation placed so far, something has to cover it. Among all length-`k` segments that cover this cell, starting the segment exactly at this cell is the best choice — starting earlier would waste part of the segment on cells to the left that are already white or already covered, and starting later would leave this very cell uncovered. Starting right at the cell is both valid and gives the segment maximum reach to the right, so it has the best chance of covering future black cells for free.

### Core idea

Walk the string once, keeping a variable `move` — the start index of the last operation placed. As long as the current position `i` falls inside `[move, move+k-1]`, it's already white by construction, so no action is taken even if the raw string still shows `B` there — in practice that never happens, because the whole range `[move, move+k-1]` was painted by the previous operation.

Once a `B` is found with `i - move >= k` (meaning it falls outside the last operation's range), a new operation starts at `i`, and `move` is updated to `i`. The condition `i - move >= k` is exactly the check "is the current position outside the last placed segment."

The initial value `move = -k` is a technical detail: for the very first black cell, `i - move = i + k` is guaranteed `>= k` for any `i >= 0`, so the first operation always fires without a separate "nothing placed yet" check.

### Complexity

A single pass over the string is `O(n)` per test case, with no extra memory beyond two counters. The problem guarantees the sum of `n` across all test cases is bounded by `2·10^5`, so the total work across all `t` test cases is `O(N)`, where `N` is the combined length of all strings.

### Usage

```go
s := "BBWBBBWB"
k := 3
ans := solution(s, k)
// i=0: B, 0-(-3)=3>=3 → operation [0,2], move=0
// i=1,2: covered
// i=3: B, 3-0=3>=3 → operation [3,5], move=3
// i=4,5: covered
// i=7: B, 7-3=4>=3 → operation [7,9], move=7
// ans = 3
```

### Operations

| Step | Complexity | Description |
|---|---|---|
| pass over the string | O(n) | one variable `move` holds the start of the last operation |
| coverage check | O(1) | `i - move >= k` — did we step outside the last segment |
| placing an operation | O(1) | shift `move` to the current index `i` |

---

<br>

> Операция не спрашивает, где ей хочется встать — она встаёт там, где стоит первая непокрытая чёрная ячейка, и тянется от неё как можно дальше вправо.
>
> *An operation doesn't ask where it wants to go — it starts at the first uncovered black cell and reaches as far right from there as it can.*