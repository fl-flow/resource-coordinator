package resource


type ResourceRegisterSerializer struct {
  ResourceName    string    `json:"resource_name" binding:"required"`
  Max             uint      `json:"max" binding:"required"`
  Min             uint      `json:"min"`
  Init            uint      `json:"init"`
  Node            string    `json:"node" binding:"required"`
}


type ResourceChangeSerializer struct {
  ResourceName    string    `json:"resource_name" binding:"required"`
  Value           uint      `json:"value" binding:"required"`
  Type            bool      `json:"type" binding:"required"`
  Node            string    `json:"node" binding:"required"`
}


type ResourceListQuery struct {
  Node            string    `json:"node" binding:"required"`
}
