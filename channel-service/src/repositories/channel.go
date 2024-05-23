package repositories

type ChannelService struct {
	repo BaseRepo
}

func NewChannelService(repo BaseRepo) *ChannelService {
	return &ChannelService{repo: repo}
}
