# 2025A. Два экрана

## Русский

Наша задача — найти минимальное количество секунд, за которое первый экран будет содержать строку `s1`, а второй — строку `s2`.

За одну секунду мы можем либо добавить одну букву на экран, либо полностью скопировать содержимое одного экрана на другой.

### Идея решения

Главное здесь — найти **общий префикс** двух строк, то есть сколько первых символов у них совпадает подряд.

Например:

```text
s1 = ABCDE
s2 = ABCXY
```

Общий префикс здесь равен `3`, потому что первые три символа одинаковые:

```text
ABC
```

Дальше есть два случая.

Если общего префикса нет (`k == 0`), то копирование вообще не помогает. Нам придётся независимо набрать обе строки:

```text
len(s1) + len(s2)
```

Если общий префикс есть, то сначала можно набрать этот префикс на одном из экранов, затем скопировать его на второй экран. Копирование занимает одну дополнительную секунду.

После этого оставшиеся символы каждой строки дописываются отдельно.

Поэтому из общего количества символов можно вычесть `k`, которые больше не нужно набирать дважды, и добавить одну секунду на копирование:

```text
len(s1) + len(s2) - k + 1
```

### Пример

Пусть:

```text
s1 = ABCDE
s2 = ABCXY
```

Общий префикс:

```text
ABC
```

`k = 3`.

Сначала набираем `ABC` на первом экране — `3` секунды.

Копируем `ABC` на второй экран — ещё `1` секунда.

Теперь:

```text
screen 1 = ABC
screen 2 = ABC
```

Остаётся дописать:

```text
DE
XY
```

Это ещё `2 + 2 = 4` секунды.

Итого:

```text
3 + 1 + 4 = 8
```

Формула даёт тот же результат:

```text
5 + 5 - 3 + 1 = 8
```

### Почему достаточно искать только префикс

Копировать можно только **всю текущую строку целиком**. Поэтому выгодно копировать только ту часть, которая одинаковая у обеих строк начиная с первого символа.

Если строки отличаются уже на первой позиции, никакого полезного копирования сделать нельзя — именно поэтому при `k == 0` ответ равен сумме длин строк.

В коде это реализовано следующим циклом:

```go
for i := 0; i < min(len(s1), len(s2)); i++ {
	if s1[i] == s2[i] {
		k++
	} else {
		break
	}
}
```

Мы идём от начала обеих строк и увеличиваем `k`, пока символы совпадают. Как только встретили различие, останавливаемся.

После этого:

```go
if k == 0 {
	return len(s1) + len(s2)
}
```

А если общий префикс есть:

```go
return len(s1) + len(s2) - k + 1
```

### Сложность

Для каждого теста мы один раз проходим по общему началу двух строк.

В худшем случае:

```text
O(min(|s1|, |s2|))
```

Дополнительная память:

```text
O(1)
```

При ограничении `|s1|, |s2| ≤ 100` решение работает практически мгновенно.

### Использование

```go
s1 := "ABCDE"
s2 := "ABCXY"

ans := solution(s1, s2)
// k = 3
// ans = 5 + 5 - 3 + 1 = 8
```

### Операции

| Шаг                   | Сложность    | Описание                           |
| --------------------- | ------------ | ---------------------------------- |
| поиск общего префикса | O(min(n, m)) | сравниваем символы с начала строк  |
| вычисление ответа     | O(1)         | используем длины строк и `k`       |
| проверка `k == 0`     | O(1)         | определяем, выгодно ли копирование |

Итоговая сложность для одного теста:

```text
O(min(n, m))
```

Дополнительная память:

```text
O(1)
```

### Сборка и тестирование

```bash
go test -v ./...
```

---

## English

The task is to find the minimum number of seconds needed to make the first screen display `s1` and the second screen display `s2`.

In one second we can either append one character to a screen or copy the whole string from one screen to the other.

### Solution idea

The main thing we need to find is the **common prefix** of the two strings — the number of equal characters starting from the first position.

For example:

```text
s1 = ABCDE
s2 = ABCXY
```

The common prefix is:

```text
ABC
```

so `k = 3`.

There are two cases.

If there is no common prefix (`k == 0`), copying does not help. We have to build both strings independently:

```text
len(s1) + len(s2)
```

If the common prefix exists, we can build it once, copy it to the other screen, and then add the remaining characters separately.

Therefore, we can save `k` character additions and only need one extra second for the copy operation:

```text
len(s1) + len(s2) - k + 1
```

### Example

Suppose:

```text
s1 = ABCDE
s2 = ABCXY
```

The common prefix is:

```text
ABC
```

so `k = 3`.

First, build `ABC` on one screen — `3` seconds.

Then copy it to the second screen — `1` second.

Now both screens contain:

```text
ABC
```

We only need to append:

```text
DE
XY
```

which takes `4` more seconds.

Total:

```text
3 + 1 + 4 = 8
```

The formula gives the same result:

```text
5 + 5 - 3 + 1 = 8
```

### Why do we only need the prefix?

The copy operation copies the **whole current string**. Because of that, the useful part that can be shared must start from the first character.

If the strings differ at the first position, there is nothing useful to copy, so the answer is simply the sum of their lengths.

This is exactly what the code does:

```go
for i := 0; i < min(len(s1), len(s2)); i++ {
	if s1[i] == s2[i] {
		k++
	} else {
		break
	}
}
```

We compare the strings from the beginning and increase `k` while the characters are equal. As soon as we find a different character, we stop.

Then:

```go
if k == 0 {
	return len(s1) + len(s2)
}
```

Otherwise:

```go
return len(s1) + len(s2) - k + 1
```

### Complexity

For each test case, we scan the common beginning of the two strings once.

Time complexity:

```text
O(min(|s1|, |s2|))
```

Additional memory:

```text
O(1)
```

Since both strings have length at most `100`, the solution easily fits within the limits.

### Usage

```go
s1 := "ABCDE"
s2 := "ABCXY"

ans := solution(s1, s2)
// k = 3
// ans = 5 + 5 - 3 + 1 = 8
```

### Operations

| Step               | Complexity   | Description                           |
| ------------------ | ------------ | ------------------------------------- |
| find common prefix | O(min(n, m)) | compare characters from the beginning |
| calculate answer   | O(1)         | use string lengths and `k`            |
| check `k == 0`     | O(1)         | determine whether copying is useful   |

Overall complexity per test case:

```text
O(min(n, m))
```

Additional memory:

```text
O(1)
```

### Build and test

```bash
go test -v ./...
```

> Смысл решения очень простой: мы один раз набираем ту часть строк, которая у них совпадает с самого начала, копируем её на второй экран, а всё остальное дописываем отдельно.
> 
> The logic behind the solution is very simple: we first extract the portion of the strings that matches from the very beginning, copy it to the second screen, and then add the rest separately.
