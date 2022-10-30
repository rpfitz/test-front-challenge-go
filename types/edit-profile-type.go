package types

type UpdateProfileSuccessType struct {
	Success            bool   `json:"success"`
	Message            string `json:"message"`
	Email_From_Request string `json:"email_from_request"`
	Full_Name          string `json:"full_name"`
	Telephone          string `json:"telephone"`
	Email              string `json:"email"`
	Full_Name_Updated  bool   `json:"full_name_updated"`
	Telephone_Updated  bool   `json:"telephone_updated"`
	Email_Updated      bool   `json:"email_updated"`
	Token              string `json:"token"`
	Google             string `json:"google"`
	Error              string `json:"error"`
	Done               string `json:"done"`
}
