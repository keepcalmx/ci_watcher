package orm

type OwnerInfo struct {
	ID   string `yaml:"id" json:"id" bson:"id"`
	Name string `yaml:"name" json:"name" bson:"name"`
}

type CaseInfo struct {
	ID       string `yaml:"id" json:"id" bson:"id"`
	Name     string `yaml:"name" json:"name" bson:"name"`
	OwnerID  string `yaml:"owner_id" json:"owner_id" bson:"owner_id"`
	EnvType  string `yaml:"env_type" json:"env_type" bson:"env_type"`
	WorkDir  string `yaml:"work_dir" json:"work_dir" bson:"work_dir"`
	Executor string `yaml:"executor" json:"executor" bson:"executor"`
	Status   string `json:"status" bson:"status"`
	Result   string `json:"result" bson:"result"`
}

type Schd struct {
	Owners []OwnerInfo `yaml:"owners" json:"owners" bson:"owners"`
	Cases  []CaseInfo  `yaml:"cases" json:"cases" bson:"cases"`
}
