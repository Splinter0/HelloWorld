from array import array

class Queue(object):
    def __init__(self, order):
        self.deck = array("B")
        for x in order:
            self.deck.append(x)
    def enqueue(self, value):
        self.deck.append(value)
    def dequeue(self):
        t = self.deck[0]
        self.deck.remove(self.deck[0])
        return t
    def emptyQueue(self):
        if len(self.deck) == 0 : return True
        else : return False

#solution : 7 1 12 2 8 3 11 4 9 5 13 6 10

def wizard():
    userOrder = raw_input("Input order of the cards (split numbers by whitespaces) : ")
    userOrder = userOrder.split(" ")
    order = []
    for x in userOrder:
        order.append(int(x))

    print "Order:", order
    w = Queue(order)
    empty = w.emptyQueue()
    table = []
    while not empty:
        value = w.dequeue()
        w.enqueue(value)
        value = w.dequeue()
        table.append(value)
        empty = w.emptyQueue()
    print("Empty deck, table cards:", table)

wizard()
