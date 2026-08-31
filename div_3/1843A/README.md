# Color Max Cost

<p style="text-align: left">
  <a href="#русский">Русский</a> ・ <a href="#english">English</a>
</p>

---

Код, приложенный к условию, не решает задачу. В нём две проблемы: массив нигде не сортируется, и строка `maxVal = a[right] - a[left]` перезаписывает переменную вместо того, чтобы прибавлять к ней. Проверка: `a = [1, 2, 100, 3]`, правильный ответ 100, этот код вернёт 98.

The code attached to the statement doesn't solve the problem. Two things are broken: the array is never sorted, and `maxVal = a[right] - a[left]` overwrites the variable instead of adding to it. Check it on `a = [1, 2, 100, 3]`: correct answer is 100, this code returns 98.

---

## Русский

Массив нужно раскрасить в произвольное число цветов, каждый цвет непустой. Стоимость цвета — разница между максимумом и минимумом среди его элементов, стоимость раскраски — сумма по цветам. Нужен максимум этой суммы.

Смотреть удобно не на цвета, а на то, какую роль играет каждый элемент внутри своего цвета. У него их всего три: он максимум своего цвета, минимум своего цвета, или ни то ни другое — тогда он просто лежит где-то между и на сумму никак не влияет. Раз так, держать элемент в «нейтральной» роли невыгодно: лучше вообще не включать его в цвет с кем-то ещё, а если он всё же где-то нужен, то пусть будет либо максимумом, либо минимумом. Самый простой способ так расставить роли — разбить массив на пары и класть каждую пару в отдельный двухэлементный цвет: один элемент пары становится минимумом, другой максимумом.

Осталось понять, как выбирать пары. Если заранее решить, что будет ровно k пар, сумма получится максимальной, когда на роль максимумов идут k самых больших чисел массива, а на роль минимумов — k самых маленьких. После сортировки по возрастанию это как раз пары с двух концов: последний элемент с первым, предпоследний со вторым и так далее. Больше пар — больше сумма (лишняя пара никогда не уменьшает результат, в худшем случае добавляет 0), поэтому берём максимум пар, k = n/2 с округлением вниз. При нечётном n один элемент остаётся без пары, идёт в цвет сам по себе и вклада в сумму не даёт.

Отсюда алгоритм: отсортировать массив, поставить один указатель в начало, другой в конец, и пока они не сошлись, прибавлять к ответу разность между ними и сдвигать оба навстречу друг другу.

```go
func solution(a []int) int {
    sort.Ints(a)
    
    maxVal := 0
    left, right := 0, len(a)-1
    
    for left < right {
        maxVal = a[right] - a[left]
        
        left++
        right--
    }
    
    return maxVal
}
```

По сравнению с исходным кодом здесь два изменения: `sort.Ints(a)` перед циклом и `+=` вместо `=`. При нечётном n указатели встречаются на среднем элементе, цикл на этом заканчивается, и этот элемент нигде отдельно не учитывается — так и должно быть, у него роли нет.

С ограничениями задачи (n ≤ 50, элементы до 50) переменной `ans` хватает обычного `int`: сумма не превышает 25 × 50 = 1250. Сортировка при таком n стоит копейки, так что про асимптотику можно не думать — весь тест-кейс отрабатывает мгновенно. Чтение и вывод остались такими же, как в остальных решениях: сканер по словам на вход, буферизованный `Writer` с одним `Flush` в конце. В `main` только добавился импорт `sort`.

---

## English

The array has to be colored with any number of colors, each color non-empty. A color's cost is the difference between its maximum and minimum; the total cost is the sum over colors. We want that sum as large as possible.

It helps to think about the role of each element inside its color rather than about colors as groups. There are only three roles: an element is either its color's maximum, its minimum, or neither — sitting somewhere between them, contributing nothing. Since a neutral role adds nothing, there's no reason to spend an element on it: either leave it out of a shared color, or make sure it's a max or a min. The easiest way to arrange that is to split the array into pairs and put each pair in its own two-element color, one side as the minimum, the other as the maximum.

The question is which pairs to pick. Fix the number of pairs at k, and the sum is largest when the k biggest numbers take the maximum role and the k smallest take the minimum role. Sort the array ascending and those are exactly the pairs from opposite ends: last with first, second-to-last with second, and so on. More pairs never hurt (an extra pair adds at least 0), so take as many as possible, k = n/2 rounded down. If n is odd, one element is left without a partner, sits in a color by itself, and adds nothing.

That gives the algorithm: sort the array, put one pointer at the start and one at the end, and while they haven't met, add the difference between them to the answer and move both inward.

```go
func solution(a []int) int {
    sort.Ints(a)
    
    maxVal := 0
    left, right := 0, len(a)-1
    
    for left < right {
        maxVal = a[right] - a[left]
        
        left++
        right--
    }
    
    return maxVal
}
```

Two changes from the original code: `sort.Ints(a)` before the loop, and `+=` instead of `=`. When n is odd, the pointers meet on the middle element and the loop stops there; that element is never counted separately, which is correct since it has no role.

Given the constraints (n ≤ 50, values up to 50), a plain `int` is enough for `ans` — the sum can't exceed 25 × 50 = 1250. Sorting at this size costs nothing, so there's not much to say about complexity; each test case finishes instantly. Input and output stay the same as in the other solutions: word-split scanner in, buffered `Writer` with one `Flush` out. `main` just needs the `sort` import added.