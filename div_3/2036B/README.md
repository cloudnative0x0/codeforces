# Vending Machine

<p style="text-align: left">
  <a href="#русский">Русский</a> ・ <a href="#english">English</a>
</p>

---

## Русский

Условие: есть `n` полок и `k` бутылок, у каждой бутылки — бренд `b_i` и цена `c_i`. На одну полку можно поставить сколько угодно бутылок, но только одного бренда. Все бутылки, поставленные на полки, будут проданы. Нужно максимизировать выручку.

Ключевое наблюдение: полка — это не ограничение по количеству бутылок, а разрешение продать один бренд целиком. Раз бутылок одного бренда на полке можно поставить сколько угодно, выгоднее всего занять полку под бренд и выставить туда **все** бутылки этого бренда сразу — дробить бренд между несколькими полками или оставлять часть бутылок бренда за пределами полки бессмысленно: если бренд уже занял полку, недостающие бутылки того же бренда либо тоже помещаются туда бесплатно (полка не ограничена по числу бутылок), либо, если они не помещались бы, задача была бы про вместимость полки, а её здесь нет.

Из этого наблюдения выходит: реальный выбор — не «какие бутылки взять», а «какие бренды взять», и ценность бренда — это сумма цен всех его бутылок. Полок `n`, брендов может быть больше — значит нужно выбрать `min(n, число_брендов)` брендов с наибольшей суммарной ценой.

### Идея решения

Сначала все `k` бутылок группируются по бренду в хеш-таблицу `map[бренд]сумма_цен` — один проход, каждая бутылка прибавляет свою цену к сумме своего бренда. Это даёт список сумм по каждому уникальному бренду, независимо от того, сколько бутылок в нём было.

Дальше эти суммы сортируются по убыванию, и берутся первые `min(n, m)` из них, где `m` — число уникальных брендов. Если брендов меньше, чем полок, все бренды помещаются и лишние полки просто не используются — это не ошибка, а часть ответа `min(n, m)`.

Хеш-таблица здесь нужна ровно для одной вещи — свести повторяющиеся бренды в одну сумму за один проход, без чего пришлось бы сначала сортировать все `k` бутылок по бренду и затем схлопывать соседние группы, что асимптотически не хуже, но требует лишней сортировки исходного массива вместо сортировки уже свёрнутых сумм.

### Асимптотика

Группировка через `map[int]int` — `O(k)` в среднем на вставку, итого `O(k)` на все `k` бутылок. Сортировка сумм по убыванию — `O(m log m)`, где `m ≤ k` — число уникальных брендов. Суммирование топ-`min(n, m)` значений — `O(min(n, m))`, что не превышает `O(m)`. Итоговая сложность на один тест-кейс — `O(k log k)` в худшем случае (когда все бренды уникальны, `m = k`), доминирует сортировка, а не группировка.

По условию сумма `k` по всем тестам ограничена `2·10^5`, так что суммарная сложность по всем `t` тестам — `O(K log K)`, где `K` — суммарное число бутылок, а не `O(t · k log k)` с каждым `k` максимальным — иначе `t = 10^4` тестов по `k = 2·10^5` дали бы недопустимую сумму.

### Использование

```go
b := []int{1, 1, 2, 3, 3, 3}
c := []int{10, 20, 5, 1, 1, 1}
ans := solution(2, len(b), b, c)
// бренд 1: 10+20=30, бренд 2: 5, бренд 3: 1+1+1=3
// топ-2: 30 + 5 = 35
```

### Операции

| Шаг | Сложность | Описание |
|---|---|---|
| группировка по бренду | O(k) | сумма цен на бренд через `map[int]int` |
| сортировка сумм | O(m log m) | суммы брендов по убыванию |
| выбор топ-N | O(min(n, m)) | суммирование первых `min(n, m)` сумм |

---

## English

Setup: `n` shelves and `k` bottles, each bottle has a brand `b_i` and a price `c_i`. Any number of bottles can go on one shelf, but only bottles of the same brand. Every bottle placed on a shelf gets sold. The task is to maximize revenue.

The key observation: a shelf isn't a capacity constraint, it's a license to sell an entire brand. Since a shelf holds unlimited bottles of one brand, the best move is always to dedicate a shelf to a brand and put **all** of that brand's bottles on it — splitting a brand across shelves or leaving some of its bottles unplaced never helps, because a shelf that already holds the brand can take the rest of that brand for free.

From this, the actual decision isn't "which bottles to take" but "which brands to take," and a brand's value is the sum of the prices of all its bottles. There are `n` shelves and possibly more brands than that, so the answer is the `min(n, number_of_brands)` brands with the highest total price.

### Core idea

First, all `k` bottles are grouped by brand into a hash map `map[brand]price_sum` in a single pass — each bottle adds its price to its brand's running sum. This produces one summed value per unique brand, regardless of how many bottles that brand had.

These sums are then sorted in descending order, and the top `min(n, m)` of them are taken, where `m` is the number of unique brands. If there are fewer brands than shelves, every brand fits and the leftover shelves simply go unused — that's not an edge case to special-case, it's exactly what `min(n, m)` already captures.

The hash map does exactly one job here: collapse repeated brands into a single sum in one pass, instead of first sorting all `k` bottles by brand and then merging adjacent groups — which is no better asymptotically but sorts the raw bottle array instead of the already-collapsed sums.

### Complexity

Grouping via `map[int]int` is `O(k)` average per insert, `O(k)` total for all `k` bottles. Sorting the sums in descending order is `O(m log m)`, where `m ≤ k` is the number of unique brands. Summing the top `min(n, m)` values is `O(min(n, m))`, which never exceeds `O(m)`. Per test case, the total is `O(k log k)` in the worst case (all brands unique, `m = k`) — sorting dominates, not the grouping.

The problem guarantees the sum of `k` across all test cases is bounded by `2·10^5`, so the total work across all `t` test cases is `O(K log K)`, where `K` is the combined bottle count, not `O(t · k log k)` with each `k` at its maximum — that bound would let `t = 10^4` cases at `k = 2·10^5` each blow past the stated limit.

### Usage

```go
b := []int{1, 1, 2, 3, 3, 3}
c := []int{10, 20, 5, 1, 1, 1}
ans := solution(2, len(b), b, c)
// brand 1: 10+20=30, brand 2: 5, brand 3: 1+1+1=3
// top-2: 30 + 5 = 35
```

### Operations

| Step | Complexity | Description |
|---|---|---|
| group by brand | O(k) | sum prices per brand via `map[int]int` |
| sort sums | O(m log m) | brand sums in descending order |
| pick top-N | O(min(n, m)) | sum the first `min(n, m)` sums |


---

<br>

> Полка не спрашивает, сколько бутылок ты хочешь на неё поставить — она спрашивает, чей это бренд. Ответ определяет всё остальное.
>
> *A shelf doesn't ask how many bottles you want on it — it asks whose brand it is. That answer decides everything else.*