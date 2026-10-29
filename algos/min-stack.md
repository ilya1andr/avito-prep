# Min Stack (LeetCode 155)



Дата решения: 13 октября 2026

Тема: стек, дизайн структуры данных.



## Условие

Реализовать стек с push/pop/top/getMin за O(1).

getMin возвращает минимум всего стека.



## Решение



```python

class MinStack(object):

   def __init__(self):
       self.stack = []
       self.min_stack = []



   def push(self, value):
       self.stack.append(value)
       if not self.min_stack or value <= self.min_stack[-1]:
           self.min_stack.append(value)
       else:
           self.min_stack.append(self.min_stack[-1])



   def pop(self):
       self.stack.pop()
       self.min_stack.pop()



   def top(self):
       return self.stack[-1]



   def getMin(self):
       return self.min_stack[-1]
   ```

