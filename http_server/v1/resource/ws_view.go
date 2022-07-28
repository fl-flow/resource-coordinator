package resource

import (
  "log"
  "encoding/json"
  "github.com/gin-gonic/gin"
  "github.com/gorilla/websocket"

  "github.com/fl-flow/resource-coordinator/resource/pool"
)


func ResourceAllocWs(c *gin.Context){
  // TODO: error code and msg
  conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
  if err != nil {
    log.Println("upgrade:", err)
    return
  }
  log.Println("ws connect...")

  _, receive, err := conn.ReadMessage()
  if err != nil {
    return
  }
  var f ResourceAllocSerializer
  if err := json.Unmarshal(receive, &f); err != nil {
    conn.WriteMessage(websocket.TextMessage, []byte("false"))
    conn.Close()
    return
  }
  r, e := resourcepool.ResourceNodeMap.GetResource(f.ResourceName)
  if e != nil {
    conn.WriteMessage(
      websocket.TextMessage,
      []byte("false"),
    )
    conn.Close()
    return
  }
  node, uid, e := r.ResourceAlloc(f.Value)
  if e != nil {
    conn.WriteMessage(
      websocket.TextMessage,
      []byte("false"),
    )
    conn.Close()
    return
  }
  resource, _ := r.GetNodeResource(node)

  data, err := json.Marshal(map[string]string{
    "node": node,
    "uid": uid,
  })
  if err != nil {
    log.Fatal(err)
  }
  conn.WriteMessage(websocket.TextMessage, data)
  for {
    conn.WriteMessage(websocket.TextMessage, []byte("true"))
    _, receive, err := conn.ReadMessage()
    if err != nil {
      log.Println(err)
      resource.NodeRwMutex.Lock()
      defer resource.NodeRwMutex.Unlock()
      newAllocatedValue := resource.Allocated - f.Value
      if newAllocatedValue < resource.Min {
        log.Fatalf("error ws free resource")
      }
      delete(*((*resource).Stream), uid)
      (*resource).Allocated = newAllocatedValue
      return
    }
    log.Println("ws receive : ", string(receive))
  }
}
