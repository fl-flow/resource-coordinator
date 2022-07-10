package resource


type ResourceRegisterSerializer struct {
  ResourceName    string    `json:"resource_name" binding:"required"`
}
