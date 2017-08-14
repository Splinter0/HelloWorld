from array import array

class hashNode(object):
    def __init__(self, key, value):
        self.key = key
        self.value = value

class hashTable(object):
    def __init__(self, length):
        self.store = array(length)
        self.lenght = lenght
    def put(self, key, value):
        index = self.hash(key)
        n = hashNode(key, value)
        while self.store[index] != None:
            index += 1
        self.store[index] = n
    def hash(self, key):
        return key % lenght
    def get(self, key):
        index = self.hash(key)
        while self.store[index] != None:
            n = self.store[index]
            if key == n.key:
                return n.value
            index += 1
        raise KeyError

h = hashTable(5)
h.put(3, "halo")
h.put(2, "sacco")
h.put(6, "panza")
h.put(7, "crauti")
print(get(2))
