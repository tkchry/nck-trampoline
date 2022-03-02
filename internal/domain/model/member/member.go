package member

type Member struct {
	Id         *Id
	LineUserId *LineUserId
	Nickname   *Nickname
}

func NewMember(id *Id, lineUserId *LineUserId, nickname *Nickname) *Member {
	return &Member{
		Id:         id,
		LineUserId: lineUserId,
		Nickname:   nickname,
	}
}
