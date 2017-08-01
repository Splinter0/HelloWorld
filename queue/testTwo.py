from array import array

class ArrayQ(object):
    def __init__(self):
        self.queue = array("B")
    def enqueue(self, value):
        self.queue.append(value)
    def dequeue(self):
        self.queue.remove(self.queue[0])


test = ArrayQ()
test.enqueue(1)
test.enqueue(2)
test.enqueue(3)
test.enqueue(4)
print test.queue
test.dequeue()
print test.queue
test.enqueue(1)
print test.queue
