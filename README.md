# GMS - Garage Management System

This is a personal project for me to improve with languages/frameworks I'm not as well versed in, namely Go for the backend and Vue/Vite frameworks for the frontend. I've already got (about eight years, that's flown by) years of experience JS/TS. As such, the front end is less of a priority for me, and building a nice, shiny, front end around a non-existant backend makes little sense in the first place.

This is *very much* a work in progress. Currently I'm focusing on the rear end, as you do.

GMS is based around handling customers, vehicles, jobs, inventory, etc. for garages. 

## Progress

### As of the time of writing, the following works:
- Docker, and containerisation.
- - I've setup a dev profile for this, which is what I'm basing all my current work in.
- - There's also a prod profile for, well, production. That's not seeing too much use right now.
- - The containers which build are:
- - - GMS
- - - - garage_redis
- - - - garage_db
- - - - garage_backend (_dev)
- - - - garage_frontend (_dev)
- Makefile
- - Just to shorthand a lot of things, especially for running container-specific commands like Go's tidy and init.

### The following definitely does not work right now:
- The entire front end, really. The Docker container builds it all and delivers a blank screen when you go to http://localhost:3000/. That's about it really, since I've just chucked in some skeleton code and focused on the backend.
- Backend contents. To be very granular, I've been focusing on routing and, by extention, the User package first and foremost.


### Plans/Ideas for the future:
- The most immediate plans are to get the core functions of the backend working.
- I've got some thought of maybe building a car viewer. I've no idea how feasible this is across different manufacturers since they like to play it all private, but some means of visualising a car and being able to display different things like their wireframe/colours/breakdowns/detailed information when clicking on a part/etc.