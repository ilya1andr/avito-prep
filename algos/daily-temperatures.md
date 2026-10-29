# Daily Temperatures (LeetCode 739)



Дата решения: 14 октября 2026

Тема: монотонный стек.



## Условие

Массив temperatures. Для каждого дня вернуть, через сколько дней

будет теплее. Если не будет — 0.



## Решение



```python

class Solution(object):

   def dailyTemperatures(self, temperatures):
       answer = [0]len(temperatures)
       stack = []   пары (температура, индекс)
       for i in range(len(temperatures)):
           while stack and stack[-1][0] < temperatures[i]:
               idx = stack[-1][1]
               answer[idx] = i - idx
               stack.pop()
           stack.append([temperatures[i], i])
       return answer
```

