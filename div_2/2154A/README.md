# 2154A. Notelook

<p style="text-align: left">
  <a href="#русский">Русский</a> ・ <a href="#english">English</a>
</p>

---

## Русский

Дана бинарная строка `s` длины `n` и число `k`. Тето проходит по строке слева направо и умеет обнулять `1`, если та не защищена и если в предыдущих `k-1` позициях нет ни одной `1`. Нужно найти минимальное число позиций, которые нужно защитить, чтобы строка не изменилась вообще — ни одна `1` не должна исчезнуть.

Поскольку от нас требуется, чтобы строка осталась неизменной с самого начала и до конца, проверять можно сразу по исходной строке — если ни одна операция не проходит, значения `1` никуда не двигаются, а значит условие «предыдущие `k-1` позиций пусты» либо выполняется для исходного расположения единиц, либо нет; пересчитывать что-либо по ходу дела не нужно.

### Идея решения

Единица на позиции `i` в опасности только тогда, когда рядом слева, в пределах `k-1` клеток, нет другой единицы. Если такая соседняя единица есть — она сама перекрывает условие обнуления для `i`, и защищать `i` не требуется, достаточно того, что соседняя единица останется на месте. А вот если ближайшая предыдущая единица находится дальше, чем `k-1` позиций (или её вообще нет), эта `1` ничем не прикрыта — её обязательно нужно защищать.

Отсюда и получается проход одним указателем: держим индекс последней встреченной `1` — неважно, защищённой или нет, — и сравниваем с текущей позицией. Если расстояние до неё `k` или больше, текущая единица остаётся без прикрытия, и её нужно защитить. Указатель на «последнюю единицу» при этом обновляется при любой встреченной `1`, а не только при защищённой — потому что даже незащищённая соседняя единица тоже физически стоит на месте (раз мы условились, что строка не меняется), и она так же перекрывает окно проверки для соседей.

### Использование

```go
s := "11010"
k := 3
ans := solution(s, k)
// i=0 ('1'): last=-3, 0-(-3)=3>=3 → защищаем, last=0
// i=1 ('1'): 1-0=1<3 → сосед прикрывает, last=1
// i=3 ('1'): 3-1=2<3 → сосед прикрывает, last=3
// ans = 1
```

### Операции

| Шаг | Сложность | Описание |
|---|---|---|
| проход по строке | O(n) | один линейный проход слева направо |
| проверка расстояния | O(1) | сравнение `i - last` с `k` на каждом шаге |
| обновление указателя | O(1) | `last = i` при встрече любой `1` |

Весь проход — один линейный обход строки, дополнительная память не требуется. Итоговая сложность `O(n)` на тест-кейс, что с запасом укладывается в лимит при `n ≤ 1000`.

---

## English

Given a binary string `s` of length `n` and a number `k`. Teto walks through the string left to right and can zero out a `1` if it's unprotected and the previous `k-1` positions contain no `1` at all. The task is to find the minimum number of positions to protect so the string never changes — not a single `1` disappears.

Since we need the string to stay exactly as it started, checking can be done directly on the original string — if no operation ever fires, the ones never move, so the condition "previous `k-1` positions are empty" either holds for the original layout of ones or it doesn't; there's nothing to recompute along the way.

### Solution idea

A `1` at position `i` is at risk only when there's no other `1` within the previous `k-1` cells. If there is a nearby `1`, it alone blocks the zeroing condition for `i`, so `i` itself doesn't need protection — it's enough that the neighbor stays in place. But if the nearest previous `1` is farther than `k-1` positions away (or there isn't one), this `1` has no cover at all and must be protected.

That gives a single-pointer scan: keep the index of the last seen `1` — protected or not — and compare it to the current position. If the distance to it is `k` or more, the current one is uncovered and needs protecting. The "last one" pointer is updated on every `1` encountered, not only on protected ones, because even an unprotected neighboring `1` physically stays put (given that the string doesn't change), and it covers its neighbors' check window just the same.

### Usage

```go
s := "11010"
k := 3
ans := solution(s, k)
// i=0 ('1'): last=-3, 0-(-3)=3>=3 → protect, last=0
// i=1 ('1'): 1-0=1<3 → covered by neighbor, last=1
// i=3 ('1'): 3-1=2<3 → covered by neighbor, last=3
// ans = 1
```

### Operations

| Step | Complexity | Description |
|---|---|---|
| string scan | O(n) | single left-to-right pass |
| distance check | O(1) | comparing `i - last` against `k` on each step |
| pointer update | O(1) | `last = i` on any `1` encountered |

The whole pass is a single linear scan, no extra memory needed. Total complexity is `O(n)` per test case, comfortably within the limit for `n ≤ 1000`.

---

<br>

> Единицу защищает не забота, а сосед — если рядом уже стоит другая, второй страж не нужен.
>
> *A `1` isn't kept safe by care, but by a neighbor — if another one already stands close by, a second guard is unnecessary.*