package service

// dependancy injection, service depends on interfaces 
//not independent http cliends

type UserService struct {
	client client.APIClient 
}

func NewUserService (c c.client.APIClient) *UserService {
	return &UserService{client: c}
}

func (s *UserService) GetUsersWithPosts (ctx context.Context)
([]models.UserWithPosts, error){
	users, err := s.client.GetUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("get users with posts: %w", err)
	}

	posts, err := s.client.GetPosts(ctx)
	if err != nil {
		return nil, fmt.Errorf("get posts: %w", err)
	}

	postByUser := make(map[int][]modeld.Post)
	for _, post := range posts{
		postByUser[post.UserID] = append(posrByUser[post.UserID], post)
	}

	var result []models.GetUsersWithPosts
	for _, user := range users {
		result = append(result, models.UserWithPosts{
			User: user,
			PostCount: len(postByUser[user.Id]),
			Posts: postByUser[user.Id],
		})
	}
	return result, nil


}