\# Group Anagrams (LeetCode 49)



Дата решения: 4 октября 2026

Тема: хеш-таблицы, группировка.



\## Условие

Дан массив строк strs. Сгруппировать анаграммы вместе.

Порядок групп и слов внутри группы не важен.



\## Решение (сортировка как ключ)



```python

class Solution(object):

&#x20;   def groupAnagrams(self, strs):

&#x20;       seen={}

&#x20;       for wrd in strs:

&#x20;           srt=''.join(sorted(wrd))

&#x20;           if srt in seen:

&#x20;               seen\[srt].append(wrd)

&#x20;           else:

&#x20;               seen\[srt]=\[wrd]

&#x20;       return \[x for x in seen.values()]

&#x20;       

