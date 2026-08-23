# 1646B. Quality vs Quantity

<p style="text-align: left">
  <a href="#русский">Русский</a> ・ <a href="#english">English</a>
</p>

---

## Русский

### Постановка задачи

Дан массив `a` из n целых чисел. Нужно определить, можно ли выбрать два непересекающихся подмножества массива — «количество» и «качество», — так, чтобы в «количестве» было строго больше элементов, чем в «качестве», а сумма «качества» была строго больше суммы «количества». Использовать все элементы массива не обязательно, часть может не войти ни в одну группу.

### Идея решения

Зафиксируем размер группы «качество» — `m`. Чтобы получить максимальную сумму при этом размере, туда нужно взять `m` самых больших элементов массива. «Количеству», чтобы всё равно проиграть по сумме, выгодно взять как можно меньше элементов с как можно меньшей суммой — то есть минимально допустимый размер `m+1`, заполненный `m+1` самыми маленькими элементами. Добавлять туда лишние элементы бессмысленно: каждый новый только увеличивает сумму, а не помогает выиграть.

Группы не пересекаются, пока верхние `m` элементов и нижние `m+1` элементов не накладываются друг на друга, то есть пока `m + (m+1) <= n`. Значит достаточно перебрать `m = 1, 2, 3, ...`, сравнивая сумму верхних `m` элементов с суммой нижних `m+1`, пока обе группы помещаются в массив без пересечения.

### Как это реализовано в коде

```go
func solution(a []int) string {
    sort.Ints(a)
    left, right := 1, len(a)-1

    blueSum := a[0] + a[1]
    redSum := a[right]

    for left < right {
       if redSum > blueSum {
          return "yes"
       }

       left++
       right--

       blueSum += a[left]
       redSum += a[right]
    }

    return "no"
}
```

После сортировки `left` — индекс последнего элемента, вошедшего в «количество» (blue), `right` — индекс первого элемента, вошедшего в «качество» (red). Blue всегда занимает префикс `a[0..left]`, red — суффикс `a[right..n-1]`. Размер blue равен `left+1`, размер red равен `n-right`, и по построению blue всегда на один элемент больше.

Пока `left < right`, префикс и суффикс не пересекаются, и сравнение корректно. Если `redSum > blueSum`, нашлось рабочее разбиение — ответ `"yes"`. Если нет, обе группы расширяются на один элемент: `left++` добавляет в blue следующий по величине маленький элемент, `right--` добавляет в red следующий по величине большой.

Цикл заканчивается, когда `left` и `right` сходятся — дальше расширять группы некуда без пересечения, а значит и проверять больше нечего. Если условие ни разу не сработало, ответ `"no"`.

### Пример

`a = [1, 2, 3, 10, 10]`

Сортировка ничего не меняет. `left=1, right=4`: blueSum = 1+2 = 3, redSum = a[4] = 10.

Проверка: 10 > 3 — да. Ответ `"yes"`: «качество» = {10}, «количество» = {1, 2} — один элемент перевешивает два.

`a = [1, 1, 1, 1, 1]`

`left=1, right=4`: blueSum=2, redSum=1. 1 > 2? нет. `left=2, right=3`: blueSum=3, redSum=2. 2 > 3? нет. `left=3, right=2` — цикл завершается.

Ответ `"no"`: при одинаковых значениях группа с бо́льшим числом элементов всегда набирает бо́льшую сумму, меньшим числом элементов её не обогнать.

### Сложность

Сортировка занимает `O(n log n)`, дальше один проход двух указателей — `O(n)`. На один тест:

```text
время: O(n log n)
память: O(n)
```

Память — это сам входной массив (сортировка на месте), указатели и суммы занимают `O(1)`.

---

## English

### Problem

Given an array `a` of n integers, decide whether it's possible to pick two disjoint subsets — "quantity" and "quality" — such that quantity has strictly more elements than quality, while quality's sum is strictly greater than quantity's sum. Not all elements have to be used; some can stay out of both groups.

### Idea

Fix the size of the quality group at `m`. To get the largest possible sum for that size, it should take the `m` largest elements in the array. For quantity to still lose on sum, it's best off with as few elements and as small a sum as possible — the minimum allowed size, `m+1`, filled with the `m+1` smallest elements. Adding extra elements to it is never useful: each one only adds to its sum without helping it win.

The two groups stay disjoint as long as the top `m` elements and the bottom `m+1` elements don't overlap, i.e. `m + (m+1) <= n`. So it's enough to try `m = 1, 2, 3, ...`, comparing the sum of the top `m` against the sum of the bottom `m+1`, for as long as both groups fit without overlapping.

### How the code does it

```go
func solution(a []int) string {
    sort.Ints(a)
    left, right := 1, len(a)-1

    blueSum := a[0] + a[1]
    redSum := a[right]

    for left < right {
       if redSum > blueSum {
          return "yes"
       }

       left++
       right--

       blueSum += a[left]
       redSum += a[right]
    }

    return "no"
}
```

After sorting, `left` is the index of the last element included in "quantity" (blue), `right` is the index of the first element included in "quality" (red). Blue always occupies the prefix `a[0..left]`, red the suffix `a[right..n-1]`. Blue's size is `left+1`, red's size is `n-right`, and by construction blue always has one more element.

While `left < right`, the prefix and suffix don't overlap, so the comparison is valid. If `redSum > blueSum`, a working split was found — the answer is `"yes"`. Otherwise both groups grow by one element: `left++` adds the next-smallest element to blue, `right--` adds the next-largest to red.

The loop stops once `left` and `right` meet — there's nowhere left to grow without the groups overlapping, so there's nothing left to check. If the condition never fired, the answer is `"no"`.

### Example

`a = [1, 2, 3, 10, 10]`

Sorting doesn't change it. `left=1, right=4`: blueSum = 1+2 = 3, redSum = a[4] = 10.

Check: 10 > 3 — yes. Answer `"yes"`: quality = {10}, quantity = {1, 2} — one element outweighs two.

`a = [1, 1, 1, 1, 1]`

`left=1, right=4`: blueSum=2, redSum=1. 1 > 2? no. `left=2, right=3`: blueSum=3, redSum=2. 2 > 3? no. `left=3, right=2` — loop ends.

Answer `"no"`: with equal values, the group with more elements always has a bigger sum, and fewer elements can never catch up.

### Complexity

Sorting costs `O(n log n)`, then a single two-pointer pass costs `O(n)`. Per test case:

```text
time: O(n log n)
memory: O(n)
```

The memory is just the input array itself (sorted in place); the pointers and running sums add `O(1)`.

---

Растущие с двух сторон группы работают именно потому, что при фиксированном размере выгоднее всего брать самые большие элементы в меньшую по размеру группу, а самые маленькие — в большую: проверять другие разбиения тех же размеров не нужно, они заведомо не лучше.

Growing both groups outward works because, for a fixed pair of sizes, taking the largest elements into the smaller group and the smallest into the larger one is always at least as good as any other split of those same sizes — no other partition needs checking.