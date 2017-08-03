from PIL import Image

im = Image.open("beiph9Ood.png")
pix = im.load()
string = ""

point = 0
size = im.size

while point < size:
    bites = ""
    for i in range(point, point+8):
        (r, g, b) = pix[i,0]
        if b % 2:
            bites += str(1)
        else :
            bites += str(0)
    point += 8
    character = chr(int(bites, 2))
    string += character
    print(string)
