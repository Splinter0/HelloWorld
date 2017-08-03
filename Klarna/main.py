
inp = raw_input("Input > ")
inputs = inp.split(" ")

with open("stats_EUROPARL-EN.txt", "rb") as f:
    words = []
    numbers = []
    coos = []
    for line in f:
        word = line.split("	")[0]
        words.append(word)

print(inputs)
for n in inputs:
    print(words[int(n)-1])
#for n in range(len(words)) :
#    if str(n-1) in inputs:
#        print(words[n-1])
