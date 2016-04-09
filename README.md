# playCicleGo

Simple app that creates playlists with "zones", that allow you to select songs within each zone. 

For example, imagine a playlist for working out that has a warm-up, peak and warm-down zones, and in each zone you can select from a specific group of songs.
The ideia is to delimit playlists into "zones". You can define a zone by a time limit or a song limit - then you play songs until that limit is reached.
For example, I can have a playlist that has:
- Warm-up section (10 total songs, 3 song limit) (3 of these 10 songs will be played for warm-up);
- Running peak (25 total songs, 20:00 time limit) (Shuffled songs among these 25 will be played until 20:00 is reached);
- Warm-down section (5 total songs, 10:00 time limit) (Songs among these 5 will be played until the 10:00 limit is reached);
  
Currently playlists are "hardcoded". Either you play all songs in order or shuffle them all. Aren't you tired from this lack of control? Let's change  that.

Technology:
- Backend:
  - Golang for API server;
- Frontend:
  - Angular.js for interface;
  - Node.js for serving HTML/static files;
