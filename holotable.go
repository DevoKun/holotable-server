package main

import (
  "fmt"
  "net"
  "os"
  "strings"
  "strconv"
)

const (
  CONN_HOST = "0.0.0.0"
  CONN_PORT = "5320"
  CONN_TYPE = "tcp"
  CONN_URL  = CONN_HOST + ":" + CONN_PORT
)


func userIdFromUserName(username string) string {
  runes  := []rune(username)
  userid := ""
  for i := 0; i < len(runes); i++ {
    userid = userid + strconv.Itoa(int(runes[i]))
  } // for
  return userid
} // func userIdFromUserName


func main() {
  // Listen for incoming connections
  l, err := net.Listen(CONN_TYPE, CONN_URL)

  if err != nil {
    fmt.Println("Error listening:", err.Error())
    os.Exit(1)
  }

  // Close the listener when this application closes
  defer l.Close()


  fmt.Println("Listening on " + CONN_URL)
  for {
    // Listen for connections
    conn, err := l.Accept()

    if err != nil {
      fmt.Println("Error accepting connection:", err.Error())
      os.Exit(1)
    }

    go handleRequest(conn)
  }
}

func handleRequest(conn net.Conn) {
  fmt.Printf("Serving %s\n", conn.RemoteAddr().String())

  msg    := ""
  msglen := 0
  last6  := ""

  status := 4
  rating := 0
  wins   := 0
  losses := 0
  games  := 0
  completion := 0
  userid := ""

  userlist := ""

  for {
    buf := make([]byte, 1)
    buflen, err := conn.Read(buf)
    if err == nil {
      s := string(buf[:buflen])
      msg    = msg + s
      msglen = len(msg)
      last6  = msg
      if (msglen > 5) {
        last6 = msg[(msglen-6):]
      }
      //fmt.Println("[", s, "]:",msg,"(",last6,")")
      if (last6 == "endmsg") {
        msg = strings.Replace(msg, "endmsg", "", -1)
        msg = strings.Trim(msg, " ")
        fmt.Println("DEWIT: ["+msg+"]")

        if (msg[0:7] == "chatmsg") {
          conn.Write([]byte(string("chatmsg holotable: Thank you for saying hello endmsg")))
          conn.Write([]byte(string("chatmsg Scruffy looking nerf herder endmsg")))
        } else if (msg[0:7] == "version") {
          running_version := strings.Replace(msg, "version ", "", -1)
          conn.Write([]byte(string("chatmsg Client version is "+running_version+" endmsg")))
          if (running_version != "0.9.10") {
            conn.Write([]byte(string("chatmsg The latest holotable version is 0.9.10. Download the upgrade from www.holotable.com endmsg")))
          }
        } else if (msg[0:9] == "username:") {
          username_location := strings.Split(msg, "\f")
          username := username_location[0]
          username = strings.Replace(username, "username: ", "", -1)
          location := username_location[1]
          userid = userIdFromUserName(username)
          fmt.Println("* username: "+username)
          fmt.Println("* location: "+location)
          fmt.Println("* user id.: "+userid)
          conn.Write([]byte(string("chatmsg Hello, "+username+" endmsg")))
          conn.Write([]byte(string("chatmsg It is nice to have somebody logged in from "+location+" endmsg")))
          conn.Write([]byte(string("chatmsg Please enter your password to login.endmsg")))
          conn.Write([]byte(string("serverpassword endmsg")))
        } else if (msg[0:8] == "password") {
          conn.Write([]byte(string("chatmsg you have logged in endmsg")))
          conn.Write([]byte(string("serverchatmsg \n*****\n***** WELCOME TO HOLOTABLE \n***** Please enjoy yourself.\n***** \n endmsg")))
          // user status location rating wins losses #games completion%
          //
          // 0: online
          // 1: away
          // 2: ready for game
          // 3: busy
          // 4: inactive
          // 5: waiting for someone
          // 6: tournament play
          // 7: rated game
          //
          status     = 0
          userid     = userIdFromUserName("vader")
          rating     = 3000
          wins       = 1
          losses     = 2
          games      = 3 // total games
          completion = 3 // number of complete games
          userlist   = "userlist vader\fMustafar\b"+userid+" "+strconv.Itoa(status)+" "+strconv.Itoa(rating)+" "+strconv.Itoa(wins)+" "+strconv.Itoa(losses)+" "+strconv.Itoa(games)+" "+strconv.Itoa(completion)+" \tendmsg"
          conn.Write([]byte(string("serverchatmsg utinni entered lobby endmsg")))
          fmt.Println(userlist)
          conn.Write([]byte(string(userlist)))
          conn.Write([]byte(string("serveringame 0 endmsg")))
        } else if (msg == "keepalive") {
          conn.Write([]byte(string("keepalive endmsg")))
        } else if (msg[0:7] == "chatmsg") {
          conn.Write([]byte(string("server"+msg)))
        }


        msg   = ""
        last6 = ""
      } // if (last6 == "endmsg")

    } // if err == nil
  } // for {
} // func handleRequest(conn net.Conn)










