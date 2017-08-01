class ArrayQ(object):
    def __init__(self, noCards):
        self.deck = []
        self.noCards = noCards
        self.table = []
        for i in range(1, self.noCards+1):
            self.deck.append(i)
        print self.deck
    def enqueue(self, pointer):
        for x in self.deck:
            if pointer == x :
                self.deck.remove(x)
                self.deck.append(x)
                break
        print self.deck
    def dequeue(self, pointer):
        pass

a = ArrayQ(13)
a.enqueue(1)
