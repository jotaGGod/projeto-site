package entities

type Gender string

const (
	Masculino Gender = "MASCULINO"
	Feminino  Gender = "FEMININO"
	Outro     Gender = "OUTRO"
)

type FormEntity struct {
	Nome        string `bson:"Nome" validate:"required" json:"nome"`
	Email       string `bson:"Email" json:"email"`
	Password    string `bson:"Password" json:"password"`
	Description string `bson:"Description" json:"description"`
	//Gender      Gender   `bson:"Gender"`
	//Skill       []string `bson:"Skill"`
}
