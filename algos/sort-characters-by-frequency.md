\# Sort Characters By Frequency (LeetCode 451)



Дата решения: 4 октября 2026

Тема: хеш-таблицы, частоты, сортировка.



\## Условие

Дана строка s. Отсортировать символы по убыванию частоты.

Символы с одинаковой частотой — в любом порядке.



\## Решение



```python

class Solution(object):

\\\&#x20;   def frequencySort(self, s):

\\\&#x20;       freq = {}

\\\&#x20;       for char in s:

\\\&#x20;           freq\\\\\\\[char] = freq.get(char, 0) + 1

\\\&#x20;       

\\\&#x20;       sorted\\\\\\\_chars = sorted(freq.keys(), key=lambda x: freq\\\\\\\[x], reverse=True)

\\\&#x20;       

\\\&#x20;       return ''.join(char \\\\\\\* freq\\\\\\\[char] for char in sorted\\\\\\\_chars)



