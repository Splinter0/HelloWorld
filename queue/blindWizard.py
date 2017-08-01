from array import array

class Queue(object):
    def __init__(self, order):
        self.deck = array("B")
        for x in order:
            self.deck.append(x)
    def enqueue(self, value):
        self.deck.append(value)
    def dequeue(self):
        self.deck.remove(self.deck[0])
        return self.deck[0]
    def emptyQueue(self):
        if len(self.deck) == 0 : return True
        else : return False

def wizard():
    userOrder = raw_input("Input order of the cards (split numbers by whitespaces) : ")
    userOrder = userOrder.split(" ")
    order = []
    for x in userOrder:
        order.append(int(x))

    print "Order:", order
    w = Queue(order)
    empty = w.emptyQueue()

    while not empty:
        """
        print("Deck not empty")
        w.enqueue(order[0])
        print("Deck :", w.deck,"Table :", w.table)
        w.dequeue()
        print("Deck :", w.deck,"Table :", w.table)
        """
        try :
            value = w.dequeue()
            w.enqueue(value)
            value = w.dequeue()
        except :
            pass
        print(value)
        empty = w.emptyQueue()
    #print("Empty deck, table cards:", w.table)

wizard()
