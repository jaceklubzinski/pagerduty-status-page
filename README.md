# pagerduty-status-page
Public status page for PagerDuty business services (use only in local networks)

# Deployment
Public docker image available on docker hub https://hub.docker.com/repository/docker/jlubzinski/pagerduty-status-page

docker compose deployments/docker-compose.yml requires environment variables to work
```
# required
# PD auth token
PDSTATUS_PAGER_DUTY_AUTH_TOKEN:
# Resolved alerts that match a regular expression for analytical purposes 
# e.g. PDSTATUS_RESOLVED_ALERT_REGEX="^\[FIRING:.\](\w+)\s"
PDSTATUS_RESOLVED_ALERT_REGEX:
```
# UI
### Overview
- Search
- Number of warnings (yellow)
- Number of critical (red)
- Incident title
- How long ago was incident created
- Assignee
- Team 
### Analytics - Services
 - Search
 - Number of incidents for the last 14 days per service
 - Number of warning/critical for the last 14 per incident name for service (base on alert regex env)
 - Average time from open to resolve per incident name for service (base on alert regex env)
### Analytics - Incidents
 - Search
 - Number of warning/critical for the last 14 per incident name (base on alert regex env)
 - Average time from open to resolve per incident name (base on alert regex env)
