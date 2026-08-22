# B. Seating in a Bus

<p style="text-align: left">
  <a href="#русский">Русский</a> ・ <a href="#english">English</a>
</p>

---

## Русский

Пассажиры заходят в автобус по одному, места им уже назначены заранее — известен порядок, в котором они займут места. Первый пассажир садится куда угодно. Каждый следующий обязан сесть рядом с кем-то, кто уже сидит: слева или справа от его места должно быть занято хотя бы одно кресло. Если для очередного пассажира оба соседних места свободны — рассадка невозможна, ответ `NO`. Если удалось рассадить всех — `YES`.

### Идея решения

Заводим булев массив `bus` размера `n + 2` — запас по краям нужен, чтобы не проверять границы отдельно при обращении к `seat-1` и `seat+1`. Первого пассажира сажаем без проверок, для остальных смотрим оба соседних места: если оба `false`, сразу возвращаем `NO` и прекращаем обработку запроса. Иначе помечаем место занятым и идём дальше.

### Асимптотика

Один запрос из `n` пассажиров обрабатывается за `O(n)` — по одному проходу по местам без вложенных циклов. Память — тоже `O(n)` на массив `bus`. Суммарно по всем запросам: `O(Σn)`.

или сразу без сборки:

```bash
go run main.go < input.txt
```

### Операции

Формат ввода:

```
q
n
p_1 p_2 ... p_n
```

`q` — число независимых запросов, для каждого сначала идёт `n` — количество пассажиров, затем `n` чисел — номера мест в порядке посадки.

Формат вывода: на каждый запрос одна строка, `YES` или `NO`.

Ввод/вывод читается и пишется через буферизованные `bufio.Scanner` и `bufio.Writer` — для больших `n` и `q` это существенно быстрее, чем построчный `fmt.Scan`.

## English

Passengers board a bus one at a time, and the seat each one will take is known in advance, along with the boarding order. The first passenger can sit anywhere. Every passenger after that must sit next to someone already seated — the seat immediately to the left or right has to be occupied. If neither neighboring seat is taken when it's someone's turn, the seating is impossible and the answer is `NO`. If everyone gets seated under this rule, the answer is `YES`.

### Core idea

Allocate a boolean array `bus` of size `n + 2` — the extra slots on both ends avoid separate bounds checks when looking at `seat-1` and `seat+1`. The first passenger is seated unconditionally. For every passenger after that, check both neighboring seats: if both are empty, the query answer is `NO` right away. Otherwise mark the seat as taken and move on.

### Complexity

Each query with `n` passengers is processed in `O(n)` — a single pass over the boarding order, no nested loops. Memory is `O(n)` for the `bus` array. Across all queries: `O(Σn)`.

### Operations

Input format:

```
q
n
p_1 p_2 ... p_n
```

`q` is the number of independent queries; each one starts with `n`, the number of passengers, followed by `n` seat numbers in boarding order.

Output format: one line per query, either `YES` or `NO`.

I/O goes through buffered `bufio.Scanner` and `bufio.Writer` instead of plain `fmt.Scan` — with large `n` and `q` this matters for the time limit.

<br>

> Место у окна не обещали никому, обещали только соседа.
>
> No one promised a window seat — only a neighbor.