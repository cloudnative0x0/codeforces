# 1851B — Parity Sort

<p style="text-align: left">
  <a href="#русский">Русский</a> ・ <a href="#english">English</a>
</p>

---

## Русский

Условие: дан массив `a` из `n` чисел. Разрешена операция — менять местами два соседних элемента, но только если их чётность одинакова (оба чётные или оба нечётные). Нужно определить, можно ли такими операциями отсортировать массив по возрастанию.

Ключевое наблюдение: операция меняет местами только элементы одной чётности, значит чётный и нечётный элемент никогда не поменяются местами друг с другом. Их взаимный порядок зафиксирован с самого начала и никакими разрешёнными перестановками не меняется — если чётное число стояло раньше нечётного, оно так и останется раньше него.

Из этого следует: массив сортируем тогда и только тогда, когда после полной сортировки на каждой позиции стоит число той же чётности, что стояло там в исходном массиве. Проверять сам процесс перестановок не нужно — достаточно сравнить чётностный узор исходного массива с чётностным узором отсортированного.

### Идея решения

Массив копируется в `copyOfA`, чтобы сохранить исходный порядок. Сам массив `a` сортируется по возрастанию. Дальше идёт проход по всем позициям `i`: если чётность `a[i]` не совпадает с чётностью `copyOfA[i]`, ответ — `"no"`. Если расхождений не нашлось ни на одной позиции — ответ `"yes"`.

### Асимптотика

Сортировка массива — `O(n log n)`, сравнение чётностей после сортировки — `O(n)`. Итоговая сложность на один тест-кейс — `O(n log n)`, доминирует сортировка.

### Операции

| Шаг | Сложность | Описание |
|---|---|---|
| копирование массива | O(n) | сохранение исходного порядка в `copyOfA` |
| сортировка | O(n log n) | `a` сортируется по возрастанию |
| сравнение чётностей | O(n) | `a[i] % 2` против `copyOfA[i] % 2` на каждой позиции |

### Код

```go
// Позиция за позицией сверяем чётность отсортированного элемента
// с чётностью того, что стояло здесь до сортировки.
for i := 0; i < len(a); i++ {
    if (a[i] % 2) != (copyOfA[i] % 2) {
        // Чётный и нечётный элементы никогда не меняются местами
        // друг с другом — значит такое расхождение недостижимо.
        return "no"
    }
}
```

Ввод читается через `bufio.Scanner` с разбиением по словам (`ScanWords`) — это быстрее построчного чтения и снимает необходимость парсить пробелы вручную. Вывод буферизуется через `bufio.Writer` и сбрасывается один раз в конце через `defer writer.Flush()` — без этого при большом числе тестов вывод был бы медленным из-за системных вызовов на каждую строку.

---

## English

Setup: given an array `a` of `n` numbers. The allowed operation is swapping two adjacent elements, but only if they have the same parity (both even or both odd). The task is to determine whether the array can be sorted in ascending order using such operations.

The key observation: the operation only swaps elements of the same parity, so an even and an odd element never swap with each other. Their relative order is fixed from the start and stays that way — if an even number came before an odd one, it stays before it.

It follows that the array is sortable exactly when, after fully sorting it, every position holds a number of the same parity that was there originally. There's no need to simulate the swaps themselves — comparing the parity pattern of the original array with the parity pattern of the sorted one is enough.

### Core idea

The array is copied into `copyOfA` to preserve the original order. The array `a` itself is sorted in ascending order. Then every position `i` is checked: if the parity of `a[i]` doesn't match the parity of `copyOfA[i]`, the answer is `"no"`. If no mismatch is found anywhere, the answer is `"yes"`.

### Complexity

Sorting the array is `O(n log n)`, comparing parities after sorting is `O(n)`. Per test case the total is `O(n log n)`, dominated by the sort.

### Operations

| Step | Complexity | Description |
|---|---|---|
| copy array | O(n) | preserve original order in `copyOfA` |
| sort | O(n log n) | `a` sorted ascending |
| compare parities | O(n) | `a[i] % 2` vs `copyOfA[i] % 2` at each position |

### Code

```go
// Compare, position by position, the parity after sorting against
// the parity that was there before sorting.
for i := 0; i < len(a); i++ {
	if (a[i] % 2) != (copyOfA[i] % 2) {
		// Even and odd elements never swap with each other,
		// so this kind of mismatch is unreachable.
		return "no"
	}
}
```

Input is read via `bufio.Scanner` split on words (`ScanWords`) — faster than reading line by line and avoids manual whitespace parsing. Output is buffered through `bufio.Writer` and flushed once at the end via `defer writer.Flush()` — without that, a large number of test cases would make output slow due to a system call per line.

---

<br>

> Соседи меняются местами, только если они одной породы — а значит, чужой порядок между чётным и нечётным решён заранее.
>
> *Neighbors only swap when they're of a kind — so the order between even and odd is settled in advance.*