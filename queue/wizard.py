from array import array

class Wizard(object):
    def __init__(self, order):
        self.deck = array("B")
        for x in order:
            self.deck.append(x)
        self.table = []
    def enqueue(self, value):
        self.deck.remove(value)
        self.deck.append(value)
    def dequeue(self):
        self.table.append(self.deck[0])
        self.deck.remove(self.deck[0])
    def emptyQueue(self):
        if len(self.deck) == 0 : return True
        else : return False

userOrder = raw_input("Input order of the cards (split numbers by whitespaces) : ")
userOrder = userOrder.split(" ")
order = []
for x in userOrder:
    order.append(int(x))

print "Order:", order
w = Wizard(order)
empty = w.emptyQueue()

while not empty:
    print("Deck not empty")
    w.enqueue(order[0])
    print("Deck :", w.deck,"Table :", w.table)
    w.dequeue()
    print("Deck :", w.deck,"Table :", w.table)
    empty = w.emptyQueue()
print("Empty deck, table cards:", w.table)
