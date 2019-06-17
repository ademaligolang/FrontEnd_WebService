# FrontEnd_WebService

This front end webservice has the responsibility of communicating with the Song webservice.
There are two URL's of importance to use it.

http://localhost:8080/AddSongs - This method will automatically pick up song_list.json and one by one call AddSong in the Song webservice

http://localhost:8080/GetSongGroups - Grabs song group information from the Song webservice and returns a pretty text response in the browser.

HOW TO USE
1) Make sure you have downloaded all three repositories into your workspace (FrontEnd_WebService, Song_Definitions and Song_WebService
2) Make sure to use "go get" so that imported packages are installed on your system.
3) Make sure that BOTH FrontEnd_WebService and also Song_WebService are running, if you choose to run from separate machines you'll need to change main.go in FrontEnd_WebService to replace "localhost" with the address of your second machine.
4) Use the URL's as above, first AddSongs and then GetSongGroups
