package handlers

func CreateUser(c *gin.Context){
	var req models.CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"Bad request": err.Error()})
        return
    }

    if err := storage.AddUser(req.Username, req.Email, req.Pswrd); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"UserCreation": "Succeeded to create a new user"})
}