import os

class Pokedex(object):
    def __init__(self, fileName):
        self.pokedex = []
        with open(fileName, "rb") as f:
            lineNum = 0
            for line in f.readlines():
                lineNum += 1
                if lineNum == 1 :
                    pass
                else :
                    p = Pokemon(line)
                    self.pokedex.append(p)

    def searchPokemon(self, pokeInput):
        specs = None
        for pokemon in self.pokedex:
            if pokeInput == pokemon.name.lower():
                specs = pokemon.specs
                break
        if specs is None:
            return False
        else :
            parseSpecs = {"Per":str(specs[0]),"HP":str(specs[1]),"Atk":str(specs[2]),"Def":str(specs[3]),"Type":str(specs[4])}
            return parseSpecs

class Pokemon(object):
    def __init__(self, line):
        self.name = line.split(",")[2]
        self.per = line.split(",")[0]
        self.hp = line.split(",")[3]
        self.atk = line.split(",")[4]
        self.defense = line.split(",")[5]
        self.type = line.split(",")[10]
        self.specs = [self.per,self.hp,self.atk,self.defense,self.type]

    def __str__(self) :
        return str(self.name)

def main():
    dex = Pokedex("pokedex.csv")
    print("""\033[33m
    ,                           .::.
                             .;:**'            AMC
                              `                  0
  .:XHHHHk.              db.   .;;.     dH  MX   0
oMMMMMMMMMMM       ~MM  dMMP :MMMMMR   MMM  MR      ~MRMN
QMMMMMb  "MMX       MMMMMMP !MX' :M~   MMM MMM  .oo. XMMM 'MMM
  `MMMM.  )M> :X!Hk. MMMM   XMM.o"  .  MMMMMMM X?XMMM MMM>!MMP
   'MMMb.dM! XM M'?M MMMMMX.`MMMMMMMM~ MM MMM XM `" MX MMXXMM
    ~MMMMM~ XMM. .XM XM`"MMMb.~*?**~ .MMX M t MMbooMM XMMMMMP
     ?MMM>  YMMMMMM! MM   `?MMRb.    `""    !L"MMMMM XM IMMM
      MMMX   "MMMM"  MM       ~%:           !Mh.""  dMI IMMP
      'MMM.                                             IMX
       ~M!M                                             IMP\033[0m""")
    print("\n\tWelcome to the Pokedex! Type a pokemon's name please...\n\tType \"quit\" to exit the program\n")
    while True :
        pokeInput = str(raw_input("\t\033[41m-->\033[0m ")).lower()
        if pokeInput == "quit":
            break
        elif pokeInput.startswith("list") and pokeInput != "list":
            pokeRange = pokeInput[5:]
            firstNum = int(pokeRange.split("-")[0])
            secNum = int(pokeRange.split("-")[1])
            print("")
            for i in range(firstNum, secNum):
                print("\t"+dex.pokedex[i].name)
            print("")
        elif pokeInput == "list":
            for p in dex.pokedex:
                print("\t"+p.name)
            print("")
        elif pokeInput.startswith("image"):
            pokeName = pokeInput[6:]
            result = dex.searchPokemon(pokeName)
            if result:
                pokeId = result["Per"]
                if len(pokeId) == 1:
                    pokeId = "00"+pokeId
                elif len(pokeId) == 2:
                    pokeId = "0"+pokeId
                os.system("start chrome http://assets.pokemon.com/assets/cms2/img/pokedex/full/"+pokeId+".png")
            else :
                print("\n\tPokemon doesn't exist!\n")
        elif pokeInput == "help":
            print("")
            print("\tThese are the avaiable commands:")
            print("\t\tlist : list pokemon in a range e.g. list 8-10")
            print("\t\timage : show and image of the pokemon e.g. image charmander")
            print("\t\tquit : close the program")
            print("")
        else:
            specs = dex.searchPokemon(pokeInput)
            if not specs :
                print("\n\tPokemon doesn't exist!\n")
            else:
                print("")
                for k in specs :
                    print("\t"+str(k)+" : "+str(specs[k]))
                print("")


if __name__=="__main__":
    try :
        main()
    except KeyboardInterrupt :
        print("\nExiting...")
