# Go Project Starter

This is my personal stack and folder setup when creating modern, server 

### Tech Stack:
* [HTMX](https://htmx.org/)
* [Go](https://go.dev/)
* [go-chi](https://github.com/go-chi/chi)
* [Postgres](https://www.mongodb.com/) 
* [Docker](https://www.docker.com/)

## Features

* Users can set Month, Day and Time for the email to be sent out, in which one email will be sent out every week at that specific date.
* Users can upload their schedule, and have that be sent out based on the time set, or sent manually with a "Send Now" button.
* Oauth using microsoft 365 to connect the app to our outlook accounts so the emails sent out on behalf of the coworkers work email.

### Requirements:

* Clone repo using `git clone https://github.com/darienmiller88/ScheduleFlow.git`
* Migrate the necessary information to your local `.env` as described in the `.env.sample` file
* Run `go build` to create a root level `ScheduleFlow.exe` file, and then run `.\ScheduleFlow-V2` to run the executable. If an executable is not needed, simply input `go run main.go` instead.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Feel free to leave suggestions as well, I'm always looking for ways to improve!

<p align="right">(<a href="#top">back to top</a>)</p>

## License
[MIT](https://choosealicense.com/licenses/mit/)