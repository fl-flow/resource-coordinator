package resourcenode


type ResourceNodeRegisterSerializer struct {
  ResourceName    string    `json:"resource_name" binding:"required"`
  Max             uint      `json:"max" binding:"required"`
  Min             uint      `json:"min"`
  Init            uint      `json:"init"`
  Node            string    `json:"node" binding:"required"`
}


type ResourceNodeUpSerializer struct {
  ResourceName    string    `json:"resource_name" binding:"required"`
  Value           uint      `json:"value" binding:"required"`
  Uid             string    `json:"uid" binding:"required"`
  Node            string    `json:"node" binding:"required"`
}


type ResourceNodeDownSerializer struct {
  ResourceName    string    `json:"resource_name" binding:"required"`
  Uid             string    `json:"uid" binding:"required"`
  Node            string    `json:"node" binding:"required"`
}
