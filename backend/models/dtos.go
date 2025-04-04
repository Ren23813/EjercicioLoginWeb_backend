package models

type LoginRequest struct {
Username string `json: "username"`
Password string `json "password"`
}

type RegisterRequest struct {
        Username string `json:"username"`
        Password string `json:"password"`
}

type ErrorDetail struct {
        Message string `json:"message"`
}

type APIResponse struct {
        Success bool         `json:"success"`
        Data    any          `json:"data,omitempty"`
        Error   *ErrorDetail `json:"error,omitempty"`
}

type UserResponse struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
}

type LoginSuccessData struct {
    UserID   string `json:"user_id"`
    Username string `json:"username"`
}

type RegisterSuccessData struct {
    UserID   int64  `json:"user_id"`
    Username string `json:"username"`
}

func NewSuccessResponse(data any) APIResponse {
        return APIResponse{
                Success: true,
                Data:    data,
                Error:   nil,
        }
}


func NewErrorResponse(errorMessage string) APIResponse {
    return APIResponse{
        Success: false,
        Data:    nil,
        Error: &ErrorDetail{
            Message: errorMessage,
        },
    }
}


