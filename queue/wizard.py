from nodeQueue import LinkedQ

def main():
    userOrder = raw_input("Input order of the cards (split numbers by whitespaces) : ")
    userOrder = userOrder.split(" ")
    q = LinkedQ()
    for x in userOrder:
        q.enqueue(int(x))

    print "Order:", q
    table = []
    while not q.isEmpty():
        n = q.dequeue()
        q.enqueue(n)
        n = q.dequeue()
        table.append(n)
    print("Deck is empty, table cards : ", table)

main()
