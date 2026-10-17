\# Min Stack (LeetCode 155)



Дата решения: 13 октября 2026

Тема: стек, дизайн структуры данных.



\## Условие

Реализовать стек с push/pop/top/getMin за O(1).

getMin возвращает минимум всего стека.



\## Решение



```python

class MinStack(object):

&#x20;   def \_\_init\_\_(self):

&#x20;       self.stack = \[]

&#x20;       self.min\_stack = \[]



&#x20;   def push(self, value):

&#x20;       self.stack.append(value)

&#x20;       if not self.min\_stack or value <= self.min\_stack\[-1]:

&#x20;           self.min\_stack.append(value)

&#x20;       else:

&#x20;           self.min\_stack.append(self.min\_stack\[-1])



&#x20;   def pop(self):

&#x20;       self.stack.pop()

&#x20;       self.min\_stack.pop()



&#x20;   def top(self):

&#x20;       return self.stack\[-1]



&#x20;   def getMin(self):

&#x20;       return self.min\_stack\[-1]

