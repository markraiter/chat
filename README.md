# Chat Project for NIX Academy Golang+Vue.js course Junior Level

## The task

#### Environment preparation

Set up an environment using virtualization tools (docker, docker compose).
The project should contain containers: Golang, Nginx, MySQL and others you need to work with.
___

#### Practical task

1. Based on the tasks described in the section “DESCRIPTION OF THE SYSTEM FUNCTIONALITY” (below),
make a list of tasks that should be arranged in the form of a table, where there is:

- task sequence number,
- the name of the task,
- detailed description,
- time estimate of labor costs for this task (for example: "Implementation of authorization - 8 hours").

2. Create a project board in [Trello](https://trello.com/) and move all tasks there. The board should consist of the following columns:

- BACKLOG - all tasks,
- IN PROGRESS - the task you are currently working on,
- DONE - ready tasks.

3. Design a database and write migrations.

4. Requirements:

- Use MySQL as a database.
- Implement the frontend part using [Vue.js](https://vuejs.org/).
- To implement the chat, use web sockets.
- We recommend using the [Echo framework](https://echo.labstack.com/) when implementing the server part in Go. This framework also has functionality for working with [web sockets](https://echo.labstack.com/cookbook/websocket) and [JWT](https://echo.labstack.com/middleware/jwt).

___

#### DESCRIPTION OF THE SYSTEM FUNCTIONALITY

###### YOU HAVE TO IMPLEMENT A WEB CHAT WITH SOME ADDITIONAL FEATURES:

1. Registration (nickname and password)

2. Authorization. Implement with [JWT](https://jwt.io/)

3. Search for users in the system

4. Implementation of social connections:

- Ability to add as a friend

- Ability to add a user to the black list

- Ability to unfriend

The user cannot write to you if he is on the black list.

5. Chat:

- Page with chats

- Chat with a specific user

6. Settings:

- Possibility to change the password

- Ability to change nickname

- Ability to change avatar

7. _Additional task_:
_Implementing notifications in the browser when a new message is received_
___