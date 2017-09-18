To build tictactoe execute:
	go build -o bin/tictactoe.exe ./main

This will create an executable in the bin folder which can then be run.
Enjoy.


To play use keys:
- w - up
- s - down
- a - left
- d - right

-------------------
Known problems:
- there never is a winner
- game only ends when the all cells have been used
- "index out of bounds" due to no checks being made on position
- and many more...

Future improvements:
- improve choosing cell when moving in some direction.
- add proper 2-player playing mode
- add player against computer playing mode
- add levels of difficulty when playing with computer