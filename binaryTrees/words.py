from __future__ import print_function
from tree import BinTree
from io import open

def main():
    swedish = BinTree()
    with open("word3.txt", "r", encoding="utf-8") as f:
        for line in f :
            word = line.strip()
            if word in swedish :
                print(word, end=" ")
            else :
                swedish.put(word)
    print("")

main()
