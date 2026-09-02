# Prepend and Append

<p style="text-align: left">
  <a href="#русский">Русский</a> ・ <a href="#english">English</a>
</p>

---

## Русский

Условие: Тимур изначально имел бинарную строку `s`. За одну операцию он приписывал `0` к одному концу строки и `1` к другому. При этом можно было выбрать один из двух вариантов:

* `0` слева, `1` справа;
* `1` слева, `0` справа.

Дана строка, которая получилась после некоторого количества таких операций. Нужно определить **минимальную возможную длину исходной строки**.

### Ключевое наблюдение

Каждая выполненная операция добавляет **один `0` и один `1` на противоположные концы строки**.

Поэтому, если смотреть на итоговую строку с двух сторон, символы на концах должны быть разными:

```text
0 ... 1
1 ... 0
```

Именно такая пара символов могла появиться в результате последней операции.

Значит, чтобы восстановить максимально короткую исходную строку, нужно удалить как можно больше таких пар с концов.
После удаления одной пары проверяем новую пару крайних символов и повторяем процесс.

Как только крайние символы становятся одинаковыми:

```text
0 ... 0
```

или

```text
1 ... 1
```

дальше удалять уже нельзя. Такая строка не могла быть получена добавлением `0` и `1` на разные концы в последней операции.

Таким образом, задача сводится к подсчёту количества последовательных пар с концов строки, в которых символы различаются.

### Идея решения

В цикле рассматриваем:

```go
s[i]
s[len(s)-1-i]
```

Если они различаются, значит эти два символа могли быть добавлены одной операцией, поэтому можно мысленно удалить их оба.

Увеличиваем `res` — количество удалённых пар.

Если символы совпали, дальнейшее удаление невозможно, поэтому сразу останавливаемся.

После этого из исходной длины `len(s)` были удалены `res` пар, то есть всего:

```text
res * 2
```

символов.

Поэтому минимальная возможная длина исходной строки:

```text
len(s) - res * 2
```

### Реализация

```go
func solution(s string) int {
    res := 0

    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s)-1-i] {
            res++
        } else {
            break
        }
    }

    return len(s) - (res * 2)
}
```

Здесь `res` хранит количество операций, которые мы можем обратить назад, удаляя по одному символу с каждого конца.


### Асимптотика

Используется один проход от обоих концов строки.

В худшем случае проверяется `n / 2` пар символов:

```text
O(n)
```

Дополнительная память не используется:

```text
O(1)
```

### Операции

| Шаг                       | Сложность | Описание                         |
| ------------------------- | --------- | -------------------------------- |
| проверка крайних символов | O(1)      | сравниваем `s[i]` и `s[n-1-i]`   |
| удаление пары             | O(1)      | увеличиваем `res`                |
| весь проход               | O(n)      | рассматриваем не более `n/2` пар |

---

## English

Setup: Timur initially had a binary string `s`. In one operation, he appended `0` to one end of the string and `1` to the other end. There were two possible choices:

* `0` on the left and `1` on the right;
* `1` on the left and `0` on the right.

You are given the resulting string and need to find the **minimum possible length of the original string**.

### Key observation

Each operation adds exactly **one `0` and one `1` to opposite ends** of the string.

Therefore, when looking at the resulting string from both ends, the two outer characters of the latest operation must be different:

```text
0 ... 1
1 ... 0
```

So, to reconstruct the shortest possible original string, we should remove as many such pairs from the ends as possible.
After removing one pair, we check the new outer characters and continue.

As soon as the two outer characters become equal:

```text
0 ... 0
```

or:

```text
1 ... 1
```

we have to stop. Such a pair could not have been created by the last operation, because every operation adds one `0` and one `1`.

Therefore, the task is simply to count how many consecutive pairs from the two ends contain different characters.

### Core idea

In the loop, compare:

```go
s[i]
s[len(s)-1-i]
```

If they are different, these two characters could have been added by the same operation, so both can be removed.

We increment `res`, which stores the number of removed pairs.

If the characters are equal, we stop immediately because no more operations can be reversed.

After removing `res` pairs, we have removed:

```text
res * 2
```

characters in total.

Therefore, the minimum possible original length is:

```text
len(s) - res * 2
```

### Implementation

```go
func solution(s string) int {
    res := 0

    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s)-1-i] {
            res++
        } else {
            break
        }
    }

    return len(s) - (res * 2)
}
```

Here, `res` represents the number of operations that can be reversed by removing one character from each end.

### Complexity

We scan the string from both ends.

At most `n / 2` pairs are checked:

```text
O(n)
```

No additional memory is required:

```text
O(1)
```

### Operations

| Step                     | Complexity | Description                     |
| ------------------------ | ---------- | ------------------------------- |
| compare outer characters | O(1)       | compare `s[i]` and `s[n-1-i]`   |
| remove a pair            | O(1)       | increment `res`                 |
| entire scan              | O(n)       | at most `n/2` pairs are checked |

---

<br>

> Каждая операция оставляет на краях `0` и `1`. Поэтому, пока края различаются, мы можем отменять операцию. Первая одинаковая пара — граница исходной строки.
>
> *Every operation leaves `0` and `1` on opposite ends. So while the ends are different, we can undo an operation. The first equal pair marks the boundary of the original string.*
