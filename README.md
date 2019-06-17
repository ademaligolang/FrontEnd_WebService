# FrontEnd_WebService

This front end webservice has the responsibility of communicating with the Song webservice.
There are two URL's of importance to use it.

http://localhost:8080/AddSongs
http://localhost:8080/GetSongGroups

AddSongs - This method will automatically pick up song_list.json and one by one call AddSong in the Song webservice
GetSongGroups - Grabs song group information from the Song webservice and returns a pretty text response in the browser.
