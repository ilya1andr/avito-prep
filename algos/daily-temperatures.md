\# Daily Temperatures (LeetCode 739)



Дата решения: 14 октября 2026

Тема: монотонный стек.



\## Условие

Массив temperatures. Для каждого дня вернуть, через сколько дней

будет теплее. Если не будет — 0.



\## Решение



```python

class Solution(object):

&#x20;   def dailyTemperatures(self, temperatures):

&#x20;       answer = \[0] \* len(temperatures)

&#x20;       stack = \[]   # пары (температура, индекс)



&#x20;       for i in range(len(temperatures)):

&#x20;           while stack and stack\[-1]\[0] < temperatures\[i]:

&#x20;               idx = stack\[-1]\[1]

&#x20;               answer\[idx] = i - idx

&#x20;               stack.pop()

&#x20;           stack.append(\[temperatures\[i], i])



&#x20;       return answer

