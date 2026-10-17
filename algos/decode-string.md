\# Decode String (LeetCode 394)



Дата решения: 14 октября 2026

Тема: стек, вложенные структуры.



\## Условие

Дана строка вида "3\[a2\[c]]". Раскодировать в "accaccacc".

Формат: k\[substring], где k — число, substring может содержать

вложенные k\[...].



\## Решение



```python

class Solution(object):

&#x20;   def decodeString(self, s):

&#x20;       count\_stack = \[]

&#x20;       string\_stack = \[]

&#x20;       current\_string = ""

&#x20;       current\_num = 0



&#x20;       for ch in s:

&#x20;           if ch.isdigit():

&#x20;               current\_num = current\_num \* 10 + int(ch)

&#x20;           elif ch == '\[':

&#x20;               count\_stack.append(current\_num)

&#x20;               string\_stack.append(current\_string)

&#x20;               current\_num = 0

&#x20;               current\_string = ""

&#x20;           elif ch == ']':

&#x20;               count = count\_stack.pop()

&#x20;               prev = string\_stack.pop()

&#x20;               current\_string = prev + current\_string \* count

&#x20;           else:

&#x20;               current\_string += ch



&#x20;       return current\_string

