# Longest Substring Without Repeating Characters (LeetCode 3)



Дата решения: 18 октября 2026

Тема: Строки, скользящее окно (Sliding Window), хеш-таблица.



## Условие

Дана строка `s`. Нужно найти длину самой длинной подстроки, 

в которой нет повторяющихся символов.



## Решение (O(n))



```python

class Solution(object):

   def lengthOfLongestSubstring(self, s):

       indx = {}

       l = 0

       mx = 0

       for r in range(len(s)):

           if s[r] in indx and indx[s[r]] >= l:

               l = indx[s[r]] + 1

           indx[s[r]] = r

           mx = max(mx, r - l + 1)

       return mx
```
