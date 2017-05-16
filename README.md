# ICA04


*Deltakere: Simen Fredriksen, Stian Blankenberg, Jone Manneråk, Kristian Hagberg, Tarjei Taxerås og Rasmus Sørby*


# Oppgave 1

# a)
 
I windows brukes 0D 0A for å bryte linjen, mens Mac/Unix brukes kun 0A for å bryte linjen. Dette er fordi at Windows-systemet må ha en Carriage Return før line-breaker-en, mens Mac OS ikke trenger det. Dette er for at det ikke var behov for en driver når man skulle printe. Dette fører til 1 mer byte per linjeskift. Problemet kan være at filen kan inneholde en del kode som kan mistolkes når det blir overført fra et operativsystem til et annet.  Eller andre veien, at det ikke inneholder nok kode for å kjøres ordentlig.


# b)

![Bilde1](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18518521_10158500805705411_1438605580_n.png?oh=5c6eb8280f2f5469352af141924ddb3b&oe=591E34A2)

bilde 1: Vi ser at en tekstfil laget i Windows har Carriage return.

![Bilde2](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18450210_10158500813615411_1181054389_n.png?oh=59b8431860d00592346163630d5382eb&oe=591D62CF)

bilde 2: Vi ser at en tekstfil laget i Mac ikke har Carriage return.


Vi skrev en kode som telte antall linjeskift, slik at vi skrev ut “Det er”, antall Carriage returns, “Carriage returns”. 
Vi visste fra oppgave 1a at linjeskift representeres ved 0A og 0D 0A. 
0A er for mac og unix, dette er lik LF som er \n. 
0D 0A er for windows, dette er lik CRLF som er \r\n.
I kodene over kan du se 


# Oppgave 2
# a)

Som vi ser her gir programmet navnet på filen, størrelsen i bytes, permission for unix og den viser også andre ting som som det er en symbolsk link, om det er en Unix block eller character fil, om den er append only, om filen er en directory og om det er en regular file.
 
For å komme fram til dette brukte vi Lstat og FileInfo fra “Os” i  Golang-biblioteket. På denne måten kunne vi vise informasjonen enkelt. Koden i seg selv er relativt enkel og baseres på at vi sjekker om noe stemmer ved å bruke if og else-setninger og printer ut svaret via Println. 


# b)

![Bilde3](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18579247_10158500805710411_1171721428_n.png?oh=cde07d453b48522dba4598f58439c929&oe=591D76EE)

** Bilde 3** (stdin). Vi endret filnavnet her til /dev/stdin i koden for å finne stdin i ubuntu. 


![Bilde4](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18554610_10158500805715411_2066292245_n.png?oh=b72c3ca9e68f3940fe92ab5aa6c410ad&oe=591D37AF)

**Bilde 4** (loop0). da det ikke var mulig å bruke ram0 så ble loop0 brukt i stedet.. 

Vi ser her at vi fikk veldig forskjellige resultater. Et av dem var at “loop0” er en block device file, mens “stdin” er en char device file. . 


# c)
Eventuelle forskjeller vi ser her fra det forrige er at størrelsen er større (10 mot 6). Dette er pga. måten windows bruker linjeskift. Bortsett fra størrelsen er alt likt som i ubuntu. 


# Oppgave 3

# a)

De forskjellige kategoriene innen metoder for å arbeide med filer er:
Grunnleggende Operasjoner
Her er helt grunnleggende metoder som create/get file info/open/delete
Lesing og Skriving
Her er kopiering av filer, flere nøyaktige metoder for å legge til tekst, samt flere nøyaktige metoder for å lese innholdet til en fil. Bruker ta typisk bytes for å komme fram til ønsket posisjon.
Arkivering
Zipper og unzipper filer. (Samler og “sprer” filer)
Komprimering
Komprimerer og dekomprimerer en fil.
Diverse
Diverse spesifikke metoder, som å hente en fil fra HTTP, eller opprette en midlertidig fil som blir slettet etter at dens funksjon er brukt opp.

# b)

![Bilde5](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18492924_10158500835360411_88144497_n.png?oh=23513638fa6de63268d540e56467ec69&oe=591CFDE2)

bilde 5: Her ser vi at programmet klarte å skrive ut antall runer og antall linjer i tekstfilen. 

Vi fikk ikke til å skrive ut de 5 mest brukte runene, men har skrevet ut antall runer. Vi brukte bufio.ScanLines og bufio.ScanRunes til å skrive antall runer og linjer i tekstfilen vi undersøkte.


# Oppgave 4

# a)

Beregning: Studenter i ett fakultet / Totalt antall studenter *100% = Sannsynlighet.

![Bilde6](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18515947_10158500835355411_1633395177_n.png?oh=fa2eef280bd3d5ad11337c032163a246&oe=591D1E4D)


# b)

Du får minst informasjon fra fakultet F (Økonomi og Samfunnsvitenskap) og D (Teknologi og realfag) ettersom deres fakultets-kode er henholdsvis 01 og 10 i motsetning til resten av fakultetene som har 3 bits kode.

# c)

![Bilde7](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18516066_10158500835365411_620334031_n.png?oh=c39b113b47659d4d7bbf00eb74d812b8&oe=591D2022)


![Bilde8](https://scontent-arn2-1.xx.fbcdn.net/v/t34.0-12/18515918_10158500835350411_2009112016_n.png?oh=bcb9eaccf049e1f92a358ed414f7c4b8&oe=591CE961)

Summert sammen blir dette 255,1 bits for en gjennomsnittlig lenke som inneholder fakultets-koder for 100 tilfeldige elever.


# e)
