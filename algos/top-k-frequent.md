\# Top K Frequent Elements (LeetCode 347)



Дата решения: 4 октября 2026

Тема: хеш-таблица + сортировка / bucket sort.



\## Условие

Дан массив nums и число k. Вернуть k самых частых элементов.

Порядок не важен.



\## Решение (сортировка)



```python

class Solution(object):

&#x20;   def topKFrequent(self, nums, k):

&#x20;       seen = {}

&#x20;       for num in nums:

&#x20;           seen\[num] = seen.get(num, 0) + 1

&#x20;       

&#x20;       srt = sorted(seen.keys(), key=lambda x: seen\[x], reverse=True)

&#x20;       return srt\[:k]

