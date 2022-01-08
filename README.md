## Ways to run this program
### With Docker
To build the container: `docker-compose up -d --build`

Run the application after it's built: `docker-compose up -d`

### Standalone
To run this project as a standalone app type `go run .` in a command prompt in the root folder of this project.

To build an exectuable and run on Windows systems: `go build` in the root folder of this project. Then run the executable named `passwordgenerator.exe`.

To build an exectuable and run on Linux/Mac systems: `go build` in the root folder of this project. Then run the command `./passwordgenerator`.

## Viewing in a browser
This app runs on port 3000, to run in a local browser type: `localhost:3000`in the url textbox in the browser.

## Running the test suite
### With Docker
Make sure the password generator container is up and running. Then in the root folder of this project run this command `docker-compose exec passgen go test -v`
### Standalone
In the root folder of this project run this command `go test -v`