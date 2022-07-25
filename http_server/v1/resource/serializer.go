package resource


type ResourceRegisterSerializer struct {
  ResourceName    string    `json:"resource_name" binding:"required"`
}


type ResourceAllocSerializer struct {
  ResourceName    string    `json:"resource_name" binding:"required"`
  Value           uint      `json:"value"`
}
