# 1972A. Предложение контеста

<p style="text-align: left">
  <a href="#русский">Русский</a> ・ <a href="#english">English</a>
</p>

---

## Русский

### Постановка задачи

Дано два отсортированных массива: `a` — сложности уже предложенных задач контеста, `b` — допустимые верхние границы сложности для каждой позиции. Нужно добавить минимальное число новых задач, чтобы после каждой операции (вставка задачи, сортировка, удаление самой сложной) выполнялось `a[i] <= b[i]` для всех `i`.

### Идея решения

Каждая вставка новой задачи убирает из массива `a` текущий максимум. Значит, вставка задачи со сложностью `w` фактически позволяет заменить любую позицию в `a`, лишь бы после сортировки порядок сохранился. Вопрос не в том, какое `w` выбрать, а в том, сколько позиций массива `a` вообще придётся заменить.

Если при жадном сопоставлении слева направо оказывается `a[i] > b[i]`, эта позиция `b` не может быть закрыта существующей задачей — она сгорает, и одна вставка должна компенсировать именно её. Задача сводится к подсчёту таких сгоревших позиций.

### Как это реализовано в коде

```go
for i := 0; i < len(a); i++ {
	if i >= len(b) {
		break
	}

	if a[i] > b[i] {
		count++
		b = append(b[:i], b[i+1:]...)
		i--
	} else {
		continue
	}
}
```

Индекс `i` идёт по массиву `a`. На каждом шаге сравниваются `a[i]` и `b[i]`.

Если `a[i] <= b[i]`, задача на этой позиции устраивает — переходим дальше.

Если `a[i] > b[i]`, то `b[i]` слишком мала для этой задачи и её нельзя закрыть без вставки. Счётчик увеличивается, `b[i]` удаляется из массива (всё, что правее, сдвигается на одну позицию влево), и то же самое `a[i]` снова сравнивается — уже со значением, которое раньше стояло на позиции `i+1`.

Строка `i--` нужна именно для повторной проверки: цикл `for` после каждой итерации сам увеличивает `i`, и `i--` компенсирует это увеличение, оставляя индекс на месте.

Цикл заканчивается, когда `a` пройден до конца или `b` закончился раньше — на этом этапе оставшиеся элементы `a` уже не влияют на ответ.

### Пример

```text
a = [3, 4, 5]
b = [1, 2, 3]
```

`i=0`: `a[0]=3 > b[0]=1` → count=1, b становится `[2, 3]`, i остаётся 0.

`i=0`: `a[0]=3 > b[0]=2` → count=2, b становится `[3]`, i остаётся 0.

`i=0`: `a[0]=3 > b[0]=3`? нет (равно) — переходим дальше.

`i=1`: у `b` уже нет элемента на этом индексе, цикл прерывается.

Ответ: `2`.

Проверка вручную: чтобы `a[i] <= b[i]` выполнялось везде, из трёх исходных задач можно оставить только одну (третью, со сложностью `5`), а две первые придётся заменить — это и есть `2` вставки.

### Сложность

Удаление элемента из середины среза стоит `O(n)`, и в худшем случае это происходит для каждой позиции `a`. По одному набору входных данных:

```text
время: O(n^2)
память: O(n)
```

При `n <= 100` и `t <= 100` это укладывается в лимит по времени с большим запасом.

### Сборка и тестирование

```bash
go test -v ./...
```

---

## English

### Problem

Two sorted arrays are given: `a`, the difficulty of the problems already proposed for the contest, and `b`, the allowed upper bound of difficulty for each position. Find the minimum number of new problems needed so that after each operation (insert a problem, sort, drop the hardest one) `a[i] <= b[i]` holds for every `i`.

### Idea

Every insertion removes the current maximum of `a`. So proposing a problem with difficulty `w` effectively lets us replace any position of `a`, as long as sorted order is preserved afterward. The real question isn't which `w` to pick — it's how many positions of `a` need replacing at all.

If, during a greedy left-to-right comparison, `a[i] > b[i]` at some position, that slot of `b` can't be covered by an existing problem — it's wasted, and exactly one insertion is needed to cover it. The task reduces to counting these wasted slots.

### How the code does it

```go
for i := 0; i < len(a); i++ {
	if i >= len(b) {
		break
	}

	if a[i] > b[i] {
		count++
		b = append(b[:i], b[i+1:]...)
		i--
	} else {
		continue
	}
}
```

`i` walks through `a`. At each step, `a[i]` is compared with `b[i]`.

If `a[i] <= b[i]`, the current problem is fine as-is, and we move on.

If `a[i] > b[i]`, this `b[i]` is too tight for the current problem and can't be satisfied without an insertion. The counter is incremented, `b[i]` is removed from the slice (everything to the right shifts one position left), and the same `a[i]` is compared again — this time against the value that used to sit at `i+1`.

The `i--` line exists purely to re-check the same index: the `for` loop increments `i` after every iteration, and `i--` cancels that out.

The loop stops once `a` is fully scanned or `b` runs out first — leftover elements of `a` at that point no longer affect the answer.

### Example

```text
a = [3, 4, 5]
b = [1, 2, 3]
```

`i=0`: `a[0]=3 > b[0]=1` → count=1, b becomes `[2, 3]`, i stays 0.

`i=0`: `a[0]=3 > b[0]=2` → count=2, b becomes `[3]`, i stays 0.

`i=0`: `a[0]=3 > b[0]=3`? no (equal) — move on.

`i=1`: `b` has no element at this index, loop stops.

Answer: `2`.

Manual check: to satisfy `a[i] <= b[i]` everywhere, only one of the three original problems can be kept (the third one, `5`), the other two must be replaced — that's exactly `2` insertions.

### Complexity

Removing an element from the middle of a slice costs `O(n)`, and in the worst case this happens once per position of `a`. Per test case:

```text
time: O(n^2)
memory: O(n)
```

With `n <= 100` and `t <= 100`, this comfortably fits the time limit.

### Build and test

```bash
go test -v ./...
```

---

Жадное сопоставление здесь работает именно потому, что оба массива не убывают, а любая вставленная задача всё равно упорядочится сама — значит достаточно один раз пройти массивы слева направо и посчитать, сколько раз текущая задача оказалась строго сложнее допустимого предела.

The greedy matching works because both arrays are non-decreasing and any inserted value sorts itself into place — a single left-to-right pass counting how many times the current problem exceeds its allowed bound is enough.