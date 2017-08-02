class Node(object):
    def __init__(self, value):
        self.value = value
        self.next = None

class LinkedQ(object):
    def __init__(self):
        self.first = None
        self.last = None
    def enqueue(self, value):
        n = Node(value)
        if self.first == None:
            self.first = n
            self.last = n
        else :
            self.last.next = n
            self.last = n
    def dequeue(self):
        n = self.first
        self.first = self.first.next
        return n.value
    def isEmpty(self):
        return self.first == None
    def insert(self, value, index):
        new = Node(value)
        n = self.first
        for i in range(index-1):
            n = n.next
        new.next = n.next
        n.next = new

    def __str__(self):
        string = ""
        n = self.first
        while True:
            string = string+str(n.value)+" "
            n = n.next
            if n == None :
                break
        return string


q = LinkedQ()
q.enqueue(2)
q.enqueue(4)
q.enqueue(6)
print(q)
q.insert(5, 1)
print(q)
