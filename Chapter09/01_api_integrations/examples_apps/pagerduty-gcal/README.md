Note: This code is a copy included with the book for convenience.
For the latest version of this tool and project support please refer to [this project](https://github.com/corsc/pagerduty-gcal).

This app is an example of a complicated task that relies on the coordination and/or data from multiple systems.


# PagerDuty vs Google Calendar

Use Google Calendar "Out of office" events to detect PagerDuty schedule issues

## Setup

* Go to https://console.developers.google.com/projectselector/apis/credentials (must be logged into your Company Google Account)
* In the top left corner, click "Select a project" and choose "PagerDuty vs Google Calendar"
* Click on "Create credentials" - OAuth client ID
* Save the credentials in a file called `credentials.json` next to the binary (or the base of this repo if you are using `go run main.go`)

* Login to PagerDuty
* Click your avatar on the far right corner of the menu and choose "My Profile"
* Select the "User Settings" tag
* Click "Create API User token"
* Set the API key as an environment variable called PD_API_KEY

## Running this app

* Run the app using the format below (or use `go run main.go` in the base of this repo)
* During the first run you will be asked to follow a link in your browser. Do this and paste the response code into the terminal
* This will create a file called `token.json` in the same directory as the binary (or the `main.go` file).  Do not delete this file or the `credentials.json`

The full command for this app is:

`pdgcal -schedule=[scheduleID] -start=[date in format YYYY-MM-DD]`

`scheduleID` is the last part of the URL when viewing the schedule in PagerDuty

## Achieving a "follow the sun" schedule

In order to achieve this you will need:

1) Create 1 layer in the PD schedule for each "slot".  For example for 2 slots, Slot 1 from 00:00 to 12:00 and Slot 2 12:00 to 00:00
1) Assign one or more users to each slot.
	* Users should not be assigned to more than one slot or they will be scheduled more often than those that are not
	* When calculating swaps (potential overrides) this tool will only use schedule entries where the hour and minute match exactly
1) Require all users to add "Out of Office" events to their company calendar (under the same email address as configured in PagerDuty)
1) Run this tool to find schedule issues and propose swaps (overrides)

## Testing this code

Note: this was a quick hack, so I was lazy and the tests make calls to the real APIs.
They currently do not modify anything but this means you will need to configure somethings and have a working internet connection.

* Define an environment variable called `TEST_PD_API_KEY` which is your PagerDuty API key (see above)
* Define an environment variable called `TEST_PD_USER_ID` which is an PagerDuty User ID (the last few characters of the URL when viewing a user)
* Define an environment variable called `TEST_PD_SCHEDULE_ID` which is an PagerDuty Schedule ID (the last few characters of the URL when viewing a schedule)
* Define an environment variable called `TEST_GC_USER_EMAIL` which is the google calendar email that matches the `TEST_PD_USER_ID` user and the Google Calendar


## Other Notes:

* This app assumes that the email settings for users in PagerDuty match the emails in Google Calendar
* This app assumes that users add an "Out of Office" event to their Google Calendar (calendar event must be public and contain the word `out`; these are defaults when using the "Out of Office" feature via Google Calendar UI) 
* This app also supports exclusions from scheduling.  Users must add a public calendar event with the title "xoncall" to their Google Calendar 
* The period this app works on is determined by the `-start` flag plus 30 days


## Problems?

* If your OAuth expires, then delete the token.json file and re-run the app (follow the prompts about allowing auth and copying the token)
