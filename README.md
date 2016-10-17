# Collab
#### A slackbot for distributed skill sharing

Within a distributed team it can be difficult to know what skills are in everyone's tool belt. Particularly when working on different projects, holding anything as formal as a weekly standup can be a real problem, so finding the opportunity to share must be down to the individual.

The aim of this project is to provide an easy mechanism for team members to list what skills they have or are interested in learning. Slack has been chosen as a medium for this on the basis that it is a tool the team is likely already using, to eliminate having a cumbersome skills profile service elsewhere.

#### Expected use

- Create a channel for skills

- Every time a person works on a new technology or develop an interest in learning about it you send a message in this channel. This may take the form of:
`tech: scala, circleci, newthing2.0` or
`learn: devops, cooking`

- It's public, so others can see what you're working on or interested in but also searchable via a slash command:
  `/collab tech: go`
 to find users who have previously said they're using Go.

 Current supported keywords are only `tech` and `learn` but this may expand in future.
