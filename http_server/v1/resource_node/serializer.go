package resourcenode


type ResourceNodeRegisterSerializer struct {
  ResourceName    string    `json:"resource_name" binding:"required"`
  Max             uint      `json:"max" binding:"required"`
  Min             uint      `json:"min"`
  Init            uint      `json:"init"`
  Node            string    `json:"node" binding:"required"`
}


type ResourceNodeChangeSerializer struct {
  ResourceName    string    `json:"resource_name" binding:"required"`
  Value           uint      `json:"value" binding:"required"`
  Type            bool      `json:"type"`
  Node            string    `json:"node" binding:"required"`
}


// type ResourceListQuery struct {
//   Node            string    `json:"node" binding:"required"`
// }
