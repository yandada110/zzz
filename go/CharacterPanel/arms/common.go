package arms

type MainArticle struct {
	BaseAttackValue float64          // 基础攻击力
	MainArticle     float64          // 主词条值
	Type            string           // 主词条属性类型
	OtherBenefits   []*OtherBenefits // 被动效果增益
}

type OtherBenefits struct {
	Value    float64 // 被动效果增益值
	Type     string  // 被动效果增益类型
	GainForm int     // 局内增益，还是固定增益
}
