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
- when the last cell is to be filled, if the player moves, the cell becomes empty. 
- when the last cell is to be filled, the game should end without any action from the next player. 

Future improvements:
- add menu for choosing play mode
- add proper 2-player playing mode
- add player against computer playing mode
- add levels of difficulty when playing with computer
- add code tests

History:
- 23.09.2017 Added winner checking
- 23.09.2017 Improved choosing next cell when moving
- 22.09.2017 Fixed "index out of bounds" problem when moving out of the panel