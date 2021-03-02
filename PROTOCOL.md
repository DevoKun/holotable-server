Holotable Protocol
================

## Message delimeter

* `endmsg` seperates all messages.
* Any command, _or chat message_, et. al., is _ended_ by a ` endmsg` phrase.
* `endmsg` must **always** have a blank space (` `) before it.
* For example:<br />_(no newline is required, or used, in the protocol. the newline in the examples below is only for readability)_
```
keepalive endmsg
version x.y.z endmsg
chatmsg xxxxxxxxx endmsg
```



## Commands

| Command     | Description |
|-------------|-------------|
| `username:` | Allows client user to login to identify themselves to the server. |
| `serverpassword` | Sent from server to client to indicate the user should send their password. |
|  `password` | Sent from the client to the server with a login password. |
| `version` | Sent from the client to the server to indicate their version |




### `username:` _(User login)_

* After connecting to the server, the holotable client will send a `username:` command
```username: UUUUUUUUU\fLLLLLLLLLLL```

* The username value is a username _(`UUUUUUUUU`)_, seperated by a formfeed _(`\f`)_, with a location _(`LLLLLLLLLLL`)_.

### `password`

* When the server sends the `serverpassword` command to the client, the client responds with the `password` command.
* The `password` command is followed by an **MD5** hash of the user password:
```
password MD5SUM(PPPPPPPPP) endmsg
```

### `version`

```
version X.Y.Z endmsg
```

### `chatmsg`

* Sent from server to client, or client to server, to chat.
* Chat messages are either sent direct to the client/server without specifying a username.
```
chatmsg you have logged in. endmsg
```

* Or chat messages are sent with an originating username _(think group chat)_
```
chatmsg vader: I find your lack of faith disturbing. endmsg
```

### `keepalive`

* The client will assume that without a `keepalive` message from the server that it should disconnect.
* The client will send a `keepalive` message to the server every X seconds.


### `userlist`

* The userlist command seperates the `username` and `location` using the formfeed (`\f`) escape code.
* The userlist command seperates the `USERNAME+LOCATION` combo from the rest of the user stats using a backspace (`\b`) escape code.
* Unlike other `endmsg` statements, the `endmsg` used with `userlist` is a tab (`\t`) escape code.

```
userlist USERNAME\fLOCATION\bUSERID status rating wins losses games completion \tendmsg
userlist Vader\fMustafar\b2 0 3000 0 66 0 66 \tendmsg
```




### Phases

* `readyphase X`
* `undoreadyphase X`

Where **X** is an integer representing the phases:

| phase number  | phase name |
|---------------|------------|
| 0             | Activate |
| 1             | Control |
| 2             | Deploy |
| 3             | Battle |
| 4             | Move |
| 5             | Draw |
| 6             | End of Turn |


### Statuses

* A user's status is indicated by a number.

| Status ID | Status Name |
|-----------|-------------|
| 0         | online |
| 1         | away |
| 2         | ready for game |
| 3         | busy |
| 4         | inactive |
| 5         | waiting for someone |
| 6         | tournament play |
| 7         | rated game |




### Decks

* When loading a deck,  the holotable client will send the **CDF** file to the server.
* The deck load is started with the `loaddeckstart X` command, where **X** is an id number for the deck from the client.

* The entire CDF file is then sent to the server:

```
back imp.gif
card "/starwars/Endor-Dark/t_endor" "�Endor (0)\nDark Location - System [U] \nSet: Endor\nIcons: Planet, Parsec: 8 \n\nText: DARK (2): If you control, for each of your starships here, your total power is +1 in battles at Endor site.\n\nLight (2): If you have no Ewoks on Endor, Force drain -1 here. To move or deploy your starship to here requires +1 Force." 1
```

* Followed by:

```
loaddecksucceeded
loaddeckend
```



## Other commands not yet documented

```
requestgame X endmsg
requestaccepted X endmsg
tablesize X Y endmsg
tablesize 900 600 endmsg
sendsessionpassword endmsg
setsessionpassword PPPPPPPP endmsg
nextactions X endmsg
piles XXX XXX XXX XXX XXX XXX XXX XXX XX XXX X XX XXX endmsg
loaddeckstart XX endmsg
```









## User Name and Password

* The original **Holotable Server** leveraged **phpbb** as a user management system.
* The username and password originally came from the **phpbb** database.

```sql
SELECT user_id, username, user_password, user_passchg, user_email, user_login_attempts
  FROM phpbb3_users;
```

* The username is required for uniquely identifying a user.
* For the purposes of allowing the `holotable.exe` clients to _just work_, this Holotable server will accept any username and password combo.
* Passwords are **md5** hashes. Any future username+password support will need to user an **md5** hash comparison.





## User ID

* The original **Holotable Server** leveraged **phpbb** as a user management system.
* The user id originally came from the **phpbb** database.
```sql
SELECT user_id, username, user_password, user_passchg, user_email, user_login_attempts
  FROM phpbb3_users;
```
* The user id is required for establishing a game between two users.
* For the purposes of providing a user-specific unique ID, the username is converted in to their ASCII numeric value.

```go
func userIdFromUserName(username string) string {
  runes  := []rune(username)
  userid := ""
  for i := 0; i < len(runes); i++ {
    userid = userid + strconv.Itoa(int(runes[i]))
  } // for
  return userid
} // func userIdFromUserName
```

* For example, the username `vader` would produce user id `11897100101114`. Where:

```
v=118
a=97
d=100
e=101
r=114
```









***