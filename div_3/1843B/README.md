# 1843B — Long Long

<p style="text-align: left">
  <a href="#русский">Русский</a> ・ <a href="#english">English</a>
</p>

---

## Русский

Условие: дан массив `a` из `n` целых чисел. Разрешена операция — выбрать любой непрерывный отрезок массива и заменить в нём знак каждого числа на противоположный. Операцию можно применять сколько угодно раз. Нужно вывести два числа: максимально возможную сумму массива после операций и минимальное число операций, за которое эта максимальная сумма достигается.

Ключевое наблюдение: максимальная сумма всегда равна сумме модулей всех чисел — знак любого элемента можно поменять отдельной операцией на отрезке из одного числа, значит любое подмножество отрицательных чисел превращается в положительные, и лучше превратить их все. Вопрос только в минимальном числе операций, за которое это делается.

Второе наблюдение: одна операция накрывает целый непрерывный отрезок и переворачивает знак у всех чисел внутри разом. Значит выгоднее не переворачивать отрицательные числа по одному, а объединять в одну операцию целый подряд идущий блок отрицательных чисел. Ноль внутри такого блока не мешает объединению — переворот нуля ничего не меняет, поэтому два блока отрицательных чисел, разделённые только нулями, всё равно можно накрыть одной операцией, включив нули в отрезок бесплатно.

Из этого следует: минимальное число операций равно числу максимальных блоков отрицательных чисел, где блоки, разделённые исключительно нулями, считаются одним блоком, а положительное число всегда разрывает блок.

### Идея решения

Проход по массиву слева направо. Для каждого элемента его модуль сразу добавляется в `maxSum` — сумма модулей копится независимо от структуры блоков. Как только встречается отрицательное число, это начало нового блока: счётчик операций `count` увеличивается на единицу, и внутренним циклом индекс `i` двигается дальше по всем идущим подряд неположительным числам (отрицательным и нулям) — модули отрицательных чисел из этого хвоста тоже добавляются в `maxSum`, но `count` внутри хвоста больше не растёт, потому что весь блок покрывается той же самой операцией. Как только встречается положительное число, внутренний цикл останавливается, блок закрыт, и внешний цикл идёт дальше в поисках следующего блока.

### Асимптотика

Внешний и внутренний циклы вместе двигают индекс `i` только вперёд и никогда не возвращаются назад, поэтому суммарное число шагов по всему массиву не превышает `n`, несмотря на вложенность циклов. Сложность на один тест-кейс — `O(n)`.

### Операции

| Шаг | Сложность | Описание |
|---|---|---|
| накопление модулей | O(n) | каждый элемент ровно один раз добавляет `abs(a[i])` к `maxSum` |
| подсчёт блоков | O(n) | `count` растёт на 1 при входе в новый блок отрицательных чисел |
| пропуск хвоста блока | O(n) суммарно | внутренний цикл продвигает `i` по неположительным числам без повторного счёта |

### Код

```go
func solution(a []int) (int64, int) {
    var maxSum int64
    count := 0

    for i := 0; i < len(a); i++ {
        // Модуль текущего элемента в любом случае идёт в сумму —
        // положительные числа переворачивать не нужно, отрицательные
        // всё равно будут перевёрнуты одной из операций.
        if a[i] < 0 {
            maxSum += int64(-a[i])
        } else {
            maxSum += int64(a[i])
        }

        if a[i] < 0 {
            // Начало нового блока отрицательных чисел — на него
            // выделяется одна операция.
            count++

            // Блок продолжается, пока встречаются отрицательные числа
            // или нули: ноль не разрывает блок, потому что переворот
            // нуля не требует отдельной операции.
            for i+1 < len(a) && a[i+1] <= 0 {
                i++

                if a[i] < 0 {
                    maxSum += int64(-a[i])
                }
            }
        }
    }

    return maxSum, count
}
```

Сумма считается через `int64`, потому что сумма модулей `n` чисел может выйти за пределы `int32`, а обычный `int` в Go на некоторых платформах — 32-битный. Ввод и вывод организованы так же, как в предыдущем решении: `bufio.Scanner` с разбиением по словам и `bufio.Writer` с единственным `Flush()` в конце.

---

## English

Setup: given an array `a` of `n` integers. The allowed operation is choosing any contiguous segment of the array and flipping the sign of every number in it. The operation can be applied any number of times. Two numbers must be produced: the maximum possible sum of the array after operations, and the minimum number of operations needed to reach that maximum sum.

The key observation: the maximum sum always equals the sum of absolute values of all numbers — any single negative number can be flipped by an operation covering just that one number, so every negative number can be turned positive, and doing so for all of them is always the best move. The only real question is the minimum number of operations needed.

The second observation: one operation covers an entire contiguous segment and flips every number inside it at once. So it's better to merge a whole run of consecutive negative numbers into a single operation rather than flip them one at a time. A zero inside such a run doesn't break the merge — flipping a zero changes nothing, so two negative runs separated only by zeros can still be covered by one operation that includes the zeros for free.

It follows that the minimum number of operations equals the number of maximal runs of negative numbers, where runs separated only by zeros count as one run, and a positive number always breaks a run.

### Core idea

A single left-to-right pass over the array. For every element, its absolute value is added to `maxSum` right away — the sum of absolute values accumulates regardless of run structure. As soon as a negative number is found, that's the start of a new run: the operation counter `count` increases by one, and an inner loop advances the index `i` through all the following non-positive numbers (negatives and zeros) — absolute values of negatives in that tail are also added to `maxSum`, but `count` doesn't grow further inside the tail, because the whole run is covered by that same operation. Once a positive number appears, the inner loop stops, the run is closed, and the outer loop moves on to look for the next run.

### Complexity

The outer and inner loops together only ever move the index `i` forward and never backtrack, so the total number of steps across the whole array never exceeds `n`, despite the nested loops. Per test case, the complexity is `O(n)`.

### Operations

| Step | Complexity | Description |
|---|---|---|
| accumulate absolute values | O(n) | every element adds `abs(a[i])` to `maxSum` exactly once |
| count runs | O(n) | `count` increases by 1 on entering a new run of negative numbers |
| skip run tail | O(n) total | inner loop advances `i` through non-positive numbers without recounting |

### Code

```go
func solution(a []int) (int64, int) {
    var maxSum int64
    count := 0

    for i := 0; i < len(a); i++ {
        // The current element's absolute value goes into the sum either
        // way — positive numbers don't need flipping, negative ones will
        // be flipped by one of the operations regardless.
        if a[i] < 0 {
            maxSum += int64(-a[i])
        } else {
            maxSum += int64(a[i])
        }

        if a[i] < 0 {
            // Start of a new run of negative numbers — one operation
            // gets allocated to it.
            count++

            // The run keeps going as long as negatives or zeros appear:
            // a zero doesn't break the run, since flipping a zero
            // doesn't need a separate operation.
            for i+1 < len(a) && a[i+1] <= 0 {
                i++

                if a[i] < 0 {
                    maxSum += int64(-a[i])
                }
            }
        }
    }

    return maxSum, count
}
```

The sum is accumulated as `int64` because the sum of absolute values of `n` numbers can exceed the range of `int32`, and Go's plain `int` is 32-bit on some platforms. Input and output are set up the same way as in the previous solution: a `bufio.Scanner` split on words and a `bufio.Writer` flushed once at the end.

---

<br>

> Ноль не спорит ни с плюсом, ни с минусом — поэтому он бесплатно склеивает соседние минусы в одну операцию.
>
> *Zero doesn't argue with plus or minus — so it glues neighboring minuses into a single operation for free.*