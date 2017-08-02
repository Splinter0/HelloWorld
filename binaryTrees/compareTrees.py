from __future__ import print_function
from tree import BinTree
from io import open

def main():
    swedish = BinTree()
    with open("word3.txt", "r", encoding="utf-8") as f:
        for line in f :
            word = line.strip()
            swedish.put(word)
    english = BinTree()
    with open("engelska.txt", "r", encoding="utf-8") as e:
        for line in e:
            words = (line.strip()).split(" ")
            for word in words:
                if word not in english:
                    english.put(word)
                    if word in swedish :
                        print(word, end=" ")
                
    print("")

main()
