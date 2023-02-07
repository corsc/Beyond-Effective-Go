Note: This code is a copy included with the book for convenience.
For the latest version of this tool and project support please refer
to [this project](https://github.com/corsc/pagerduty-manager).

This app is an example of a tedious task that are likely to repeat over time.
It is also a task where a level of consistency is valuable.

# PagerDuty Manager

This app aims to provide a simple and efficient way to define and set up your PagerDuty users, teams, services, and
schedules.

It takes a simple JSON file as an input and uses the [PagerDuty API](https://developer.pagerduty.com/) to synchronize
the JSON with the PD
configuration.

This application requires an API token with sufficient permissions and for this token to be stored in an environment
variable
named `PD_TOKEN`.

Sample Input JSON:

```json
{
  "teams": [
	{
	  "name": "[string - required]",
	  "description": "[string - optional]",
	  "slack": "[string - required]",
	  "members": [
		{
		  "name": "[string - required]",
		  "email": "[string - required]",
		  "timezone": "[string - optional]",
		  "role": "[string - optional - default - member; other values: lead, observer, dept-head]"
		}
	  ],
	  "services": [
		{
		  "name": "[string - required]",
		  "dashboard": "[string - optional]"
		}
	  ]
	}
  ],
  "default_timezone": "[string - required]"
}
```

Roles:
`member` - User that participates in the on-call schedule.
`observer` - User that does not participate in the on-call schedule.
`lead` - User that participates in the on-call schedule and is also the first level escalation.
`dept-head` - User that does not participate in the on-call schedule but serves as the second level escalation.

Ideally every `members` list should include at least 1x lead, and 1 x dept-head.

## Usage

`$ pd-manager members.json`

### Other Options:

* `-v` - Verbose listing of actions and results (useful for debugging).
* `-d` - Perform a "dry run". The tool will parse the config file and interact with PagerDuty (to check existing state)
  but will only make no
  changes.