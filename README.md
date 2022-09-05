# Find Jobs Nearby
This App web-app uses your current longitude and latitude location to find available jobs.
You can also view and search for all available job if no job found in your location

This task has two parts
- backend (golang)
- frontend (react)

## Technologies Used for backend
- Golang
- Postgis
- Gin-Gonic Router
- Docker / docker-compose 

## Scope of the Task
1.  Implement "find jobs nearby" feature based on user location (result : please show the job title from attached job titles document)
2.  Implement a map (choice of yours eg. open map, google map, etc..) with place marker for given coordinates (latitude and longitude)
3.  Search feature with job title
4.  Show up to 5 jobs within 5km radius of selected job title
5. Implement a very simple Front End interface to show the above (using any language of your choice)

## Running the application
### Backend
- user `docker-compose up --build` to start the app in a docker container
- If you have the right environment set up, you can also run the backend with the commands `go mod tidy` to sync all dependencies and `make run` to start the app


### Frontend
- Open a new terminal and navigate into the frontend directory with command `cd frontend` and command `cd locate-job` to enter the React directory
- install all dependencies with command `npm i`
- check your axios file to configure your port. Default port at axios.js is 8080
- run the frontend with the command `npm run start`


## Testing Backend
- Mock the Database with the command `make mock`
- Tests with mocked DB can be found in /internal/api/tests/search_test.go
- Use the command `make tests` to run all test files
