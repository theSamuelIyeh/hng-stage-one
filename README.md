# HNG Internship Stage One Project

This is a simple api endpoint that returns a JSON object containing various fields.

### Fields Returned from JSON

1. slack_name: this field is bound to the query parameter of the same name.
2. current_day: this field displays the current day based on the user's computer
3. utc_time: this field displays the current UTC time with a +/- 2 mins accuracy
4. track: this field is bound to the query parameter of the same name
5. github_file_url: this field displays the absolute link to the executable file contained in the repo
6. github_repo_url: this field displays the link for the repo
7. status_code: this field displays the status code of the request

### An example Request Url
http://127.0.0.1:3000/api/v1?slack_name=Samuel%20iyeh&track=backend

### An example of the Response
```json
{
    "slack_name": "Samuel iyeh",
    "current_day": "Saturday",
    "utc_time": "2023-09-09T00:06:00Z",
    "track": "backend",
    "github_file_url": "https://github.com/theSamuelIyeh/hng-stage-one/blob/main/hng-stage-1",
    "github_repo_url": "https://github.com/theSamuelIyeh/hng-stage-one.git",
    "status_code": 200
}
```