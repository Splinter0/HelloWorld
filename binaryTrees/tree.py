class Node(object):
    def __init__(self, value):
        self.left = None
        self.right = None
        self.value = value

def printOut(root):
    if root != None :
        printOut(root.left)
        print(root.value)
        printOut(root.right)

def addToTree(root, value):
    if root == None:
        n = Node(value)
        return n
    if exists(root, value):
        return root
    if root.value < value :
        root.right = addToTree(root.right, value)
    else :
        root.left = addToTree(root.left, value)
    return root

def exists(root, value):
    if root != None :
        if root.value != value:
            if root.value > value :
                return exists(root.left, value)
            else :
                return exists(root.right, value)
        else :
            return True
    else :
        return False


class BinTree(object):
    def __init__(self):
        self.root = None
    def put(self, value):
        self.root = addToTree(self.root, value)
    def __contains__(self, value):
        return exists(self.root, value)
    def write(self):
        printOut(self.root)
        print("\n")
"""TEST
t = BinTree()
t.put("gurka")
t.put("alla")
t.put("potatis")
t.put("potatis")
print(exists(t.root, "alla"))
t.write()
"""
