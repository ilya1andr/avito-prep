\# Longest Substring Without Repeating Characters (LeetCode 3)



Дата решения: 18 октября 2026

Тема: Строки, скользящее окно (Sliding Window), хеш-таблица.



\## Условие

Дана строка `s`. Нужно найти длину самой длинной подстроки, 

в которой нет повторяющихся символов.



\## Решение (O(n))



```python

class Solution(object):

&#x20;   def lengthOfLongestSubstring(self, s):

&#x20;       indx = {}

&#x20;       l = 0

&#x20;       mx = 0

&#x20;       for r in range(len(s)):

&#x20;           if s\[r] in indx and indx\[s\[r]] >= l:

&#x20;               l = indx\[s\[r]] + 1

&#x20;           indx\[s\[r]] = r

&#x20;           mx = max(mx, r - l + 1)

&#x20;       return mx

