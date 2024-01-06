#Application Requirements

- Application flow:
  - At a specified daily time, call the allergy API endpoints
  - Parse the response from each call and format it into a nice human-readable message
  - Send that message via Slack webhook
- All functions should have tests
- Application should be run from a Docker container
- The Docker container should be restarted on any failures or crashes
- An automated pipeline should be built that tests and then publishes the Docker container to Docker Hub
- Various important values, such as the API URL, Slack web hook URL, cron job time, and cron job timezone should all e configurable
