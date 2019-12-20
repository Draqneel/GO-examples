package models

type User struct {
	ID 			 string `json:"user_id"`
	Login 		 string `json:"login"`
	HashPassword string `json:"hash_password"`
	FullName 	 string `json:"full_name"`
	PhoneNumber  string `json:"phone_number"`
	Email  		 string `json:"email"`
}

type Branch struct {
	ID      string `json:"branch_id"`
	Address string `json:"address"`
}

type Employee struct {
	User
	UserId	   string `json:"user_id"`
	Branch     Branch
	BranchId   string `json:"branch_id"`
	Role 	   string `json:"role"`
	Experience string `json:"experience"`
}

type Client struct {
	User
	UserId			 string `json:"user_id"`
	Address 		 string `json:"address"`
	DigitalSignature string `json:"digital_signature"`
}

type Balance struct {
	ID 			 	string `json:"balance_id"`
	User 		 	User
	UserId			string `json:"user_id"`
	Type 		 	string `json:"type"`
	CreatingDate 	string `json:"creating_date"`
	Capitalization  string `json:"capitalization"`
	Balance 		string `json:"balance"`
}

type Transaction struct {
	ID 				string `json:"transaction_id"`
	From 			User
	FromId			string `json:"from_id"`
	Where 			User
	WhereId 		string `json:"where_id"`
	SecureSignature string `json:"secure_signature"`
	Value 			string `json:"value"`
	DateTime 		string `json:"date_time"`
}
