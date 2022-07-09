package node


type NodeRegisterSerializer struct {
  Node            string    `json:"node" binding:"required"`
}
