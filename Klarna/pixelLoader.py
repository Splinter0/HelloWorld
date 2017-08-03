from __future__ import print_function
from PIL import Image
from io import open

im = Image.open("hint2.png")
pix = im.load()
pixels = []

with open("out.txt", "w", encoding="utf-8") as f:
    for i in range(1, 92):
        pixels.append(pix[i,0])
        print(chr(pix[i,0][2]), end="")
