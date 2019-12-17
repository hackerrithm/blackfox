package resolver

//go:generate go run github.com/99designs/gqlgen
import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/hackerrithm/blackfox/services/backend/api/pkg/authentication"

	"github.com/99designs/gqlgen/graphql"
	"github.com/go-redis/redis"
	"github.com/tinrab/retry"

	"github.com/hackerrithm/blackfox/services/backend/api/pkg"
	"github.com/hackerrithm/blackfox/services/backend/api/pkg/generated"

	auth "github.com/hackerrithm/blackfox/services/backend/auth/cmd/auth/client"
	geography "github.com/hackerrithm/blackfox/services/backend/geography/cmd/geography/client"
	goal "github.com/hackerrithm/blackfox/services/backend/goal/cmd/goal/client"
	group "github.com/hackerrithm/blackfox/services/backend/group/cmd/group/client"
	match "github.com/hackerrithm/blackfox/services/backend/match/cmd/match/client"
	post "github.com/hackerrithm/blackfox/services/backend/post/cmd/post/client"
	profile "github.com/hackerrithm/blackfox/services/backend/profile/cmd/profile/client"
	space "github.com/hackerrithm/blackfox/services/backend/space/cmd/space/client"
	task "github.com/hackerrithm/blackfox/services/backend/task/cmd/task/client"
	user "github.com/hackerrithm/blackfox/services/backend/user/cmd/user/client"
)

// GQLServer ...
type GQLServer struct {
	userClient      *user.Client
	authClient      *auth.Client
	postClient      *post.Client
	spaceClient     *space.Client
	taskClient      *task.Client
	profileClient   *profile.Client
	geographyClient *geography.Client
	matchClient     *match.Client
	groupClient     *group.Client
	goalClient      *goal.Client
	redisClient     *redis.Client
	Chats           map[string]*pkg.Chat
	mu              sync.Mutex
}

// NewGraphQLServer ...
func NewGraphQLServer(
	userURL,
	authURL,
	postURL,
	spaceURL,
	taskURL,
	profileURL,
	geographyURL,
	goalURL,
	matchURL,
	redisURL string) (*GQLServer, error) {
	userClient, err := user.NewClient(userURL)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(&redis.Options{
		Addr: redisURL,
	})

	retry.ForeverSleep(2*time.Second, func(_ int) error {
		_, err := client.Ping().Result()
		return err
	})

	postClient, err := post.NewClient(postURL)
	if err != nil {
		userClient.Close()
		log.Println("error user NewGraphQLServer")
		return nil, err
	}

	spaceClient, err := space.NewClient(spaceURL)
	if err != nil {
		userClient.Close()
		postClient.Close()
		log.Println("error user NewGraphQLServer")
		return nil, err
	}

	taskClient, err := task.NewClient(taskURL)
	if err != nil {
		userClient.Close()
		postClient.Close()
		spaceClient.Close()
		log.Println("error user NewGraphQLServer")
		return nil, err
	}

	profileClient, err := profile.NewClient(profileURL)
	if err != nil {
		userClient.Close()
		postClient.Close()
		spaceClient.Close()
		taskClient.Close()
		log.Println("error user NewGraphQLServer")
		return nil, err
	}

	geographyClient, err := geography.NewClient(geographyURL)
	if err != nil {
		userClient.Close()
		postClient.Close()
		spaceClient.Close()
		taskClient.Close()
		profileClient.Close()
		log.Println("error user NewGraphQLServer")
		return nil, err
	}

	goalClient, err := goal.NewClient(goalURL)
	if err != nil {
		userClient.Close()
		postClient.Close()
		spaceClient.Close()
		taskClient.Close()
		profileClient.Close()
		geographyClient.Close()
		log.Println("error user NewGraphQLServer")
		return nil, err
	}
	matchClient, err := match.NewClient(matchURL)
	if err != nil {
		userClient.Close()
		postClient.Close()
		spaceClient.Close()
		taskClient.Close()
		profileClient.Close()
		geographyClient.Close()
		goalClient.Close()
		// worldEventClient.Close()
		log.Println("error user NewGraphQLServer")
		return nil, err
	}

	groupClient, err := group.NewClient(matchURL)
	if err != nil {
		userClient.Close()
		postClient.Close()
		spaceClient.Close()
		taskClient.Close()
		profileClient.Close()
		geographyClient.Close()
		goalClient.Close()
		matchClient.Close()
		log.Println("error user NewGraphQLServer")
		return nil, err
	}

	authClient, err := auth.NewClient(authURL)
	if err != nil {
		userClient.Close()
		authClient.Close()
		postClient.Close()
		spaceClient.Close()
		taskClient.Close()
		profileClient.Close()
		geographyClient.Close()
		goalClient.Close()
		matchClient.Close()
		log.Println("error user NewGraphQLServer")
		return nil, err
	}

	return &GQLServer{
		userClient:      userClient,
		authClient:      authClient,
		postClient:      postClient,
		spaceClient:     spaceClient,
		taskClient:      taskClient,
		profileClient:   profileClient,
		geographyClient: geographyClient,
		goalClient:      goalClient,
		matchClient:     matchClient,
		groupClient:     groupClient,
		redisClient:     client,
		Chats:           map[string]*pkg.Chat{},
		mu:              sync.Mutex{},
	}, nil
}

// Mutation ...
func (r *GQLServer) Mutation() generated.MutationResolver {
	return &mutationResolver{
		server: r,
	}
}

// Query ...
func (r *GQLServer) Query() generated.QueryResolver {
	return &queryResolver{
		server: r,
	}
}

// Subscription ...
func (r *GQLServer) Subscription() generated.SubscriptionResolver {
	return &subscriptionResolver{
		server: r,
	}
}

// ToExecutableSchema ...
func (r *GQLServer) ToExecutableSchema() graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: r,
	})
}

// func NewRootResolvers() generated.Config {
// 	c := generated.Config{}

// 	// Schema Directive
// 	c.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
// 		ctxUserID := ctx.Value(UserIDCtxKey)
// 		if ctxUserID != nil {
// 			return next(ctx)
// 		} else {
// 			return nil, errors.UnauthorisedError
// 		}
// 	}
// 	return c
// }

type mutationResolver struct{ server *GQLServer }

func (r *mutationResolver) RegisterUser(ctx context.Context, user pkg.RegisterInput) (string, error) {
	a, err := r.server.userClient.RegisterUser(ctx, user.Username, user.Password, user.Firstname, user.Lastname, user.EmailAddress, user.Gender)
	if err != nil {
		return "", err
	}

	return a, nil
}
func (r *mutationResolver) LoginUser(ctx context.Context, user pkg.LoginInput) (string, error) {
	a, err := r.server.userClient.LoginUser(ctx, user.Username, user.Password)
	if err != nil {
		return "", err
	}

	return a.Token, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	a, err := r.server.userClient.DeleteUser(ctx, id)
	if err != nil {
		return a, err
	}

	return a, nil
}

// Post

func (r *mutationResolver) AddPost(ctx context.Context, user pkg.PostPostInput) (string, error) {
	// log.Println("did the upload with [ ", user.File, " ]")
	a, err := r.server.postClient.Post(ctx, user.Author, user.Topic, user.Category, user.ContentText, user.ContentPhotoName)
	if err != nil {
		return "", err
	}

	return a, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, user pkg.UpdatePostInput) (string, error) {
	a, err := r.server.postClient.Put(ctx, user.Author, user.Topic, user.Category, user.ContentText, user.ContentPhotoName, user.ID)
	if err != nil {
		return "", err
	}

	return a, nil
}

// Space

func (r *mutationResolver) AddSpace(ctx context.Context, user pkg.PostSpaceInput) (string, error) {
	a, err := r.server.spaceClient.Post(ctx, user.Creator, user.Topic, user.Details, user.Description, user.Type, user.Managers, user.Followers, user.Tags)
	if err != nil {
		return "", err
	}

	return a, nil
}

func (r *mutationResolver) UpdateSpace(ctx context.Context, user pkg.UpdateSpaceInput) (string, error) {
	a, err := r.server.spaceClient.Put(ctx, user.ID, user.Creator, user.Topic, user.Details, user.Description, user.Type, user.Managers, user.Followers, user.Tags)
	if err != nil {
		return "", err
	}

	return a, nil
}

// Task

func (r *mutationResolver) AddTask(ctx context.Context, user string) (*pkg.Task, error) {
	a, err := r.server.taskClient.Post(ctx, user)
	if err != nil {
		return &pkg.Task{}, err
	}

	return &pkg.Task{
		Text: a,
	}, nil
}

func (r *mutationResolver) UpdateTask(ctx context.Context, user pkg.UpdateTaskInput) (string, error) {
	userID, err := strconv.ParseUint(user.ID, 10, 32)
	if err == nil {
		fmt.Printf("Type: %T \n", userID)
		fmt.Println(userID)
	}
	a, err := r.server.taskClient.Put(ctx, uint32(userID), user.Text)
	if err != nil {
		return "", err
	}

	return a, nil
}

// Chat

func (r *mutationResolver) PostUserMessage(ctx context.Context, text string, senderName string, receiverName string) (*pkg.UserMessage, error) {
	r.server.mu.Lock()
	receiver := r.server.Chats[receiverName]
	if receiver == nil {
		receiver = &pkg.Chat{
			ID:        receiverName,
			Observers: map[string]chan *pkg.UserMessage{},
		}
		r.server.Chats[receiverName] = receiver
	}
	r.server.mu.Unlock()

	message := pkg.UserMessage{
		ID:         randString(8),
		Sender:     senderName,
		Receiver:   receiverName,
		Type:       "text-only",
		IsSeen:     false,
		IsSent:     false,
		IsReceived: false,
		Timestamp:  time.Now(),
		Text:       text,
	}

	mj, _ := json.Marshal(message)
	if err := r.server.redisClient.LPush("messages", mj).Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	receiver.Messages = append(receiver.Messages, message)
	r.server.mu.Lock()
	for _, observer := range receiver.Observers {
		observer <- &message
	}
	r.server.mu.Unlock()
	return &message, nil
}

// Goal

func (r *mutationResolver) AddGoal(ctx context.Context, user pkg.PostGoalInput) (string, error) {
	a, err := r.server.goalClient.Post(ctx, user.Creator, user.Aim, user.Reason, user.Details, user.Journey.Type, user.Type, user.Tags)
	if err != nil {
		return "", err
	}

	return a, nil
}

func (r *mutationResolver) UpdateGoal(ctx context.Context, user pkg.UpdateGoalInput) (string, error) {
	a, err := r.server.goalClient.Put(ctx, user.ID, user.Creator, user.Aim, user.Reason, user.Details, user.Type, user.Inspiration,
		user.Tags, user.Likes, user.SimilarGoals, user.Watchers,
		user.IsAchieved, user.IsPrivate,
		user.Journey.Type, user.Journey.Details,
		user.Journey.IsStarted, user.Journey.IsInProgress, user.Journey.IsComplete,
		user.Journey.Steps)
	if err != nil {
		return "", err
	}

	return a, nil
}

// Group

func (r *mutationResolver) AddGroup(ctx context.Context, user pkg.PostGroupInput) (string, error) {
	a, err := r.server.groupClient.Post(ctx, user.Title, user.Details, user.Description, user.Type, user.People)
	if err != nil {
		return "", err
	}

	return a, nil
}

func (r *mutationResolver) UpdateGroup(ctx context.Context, user pkg.UpdateGroupInput) (string, error) {
	a, err := r.server.groupClient.Put(ctx, user.ID, user.Title, user.Details, user.Description, user.Type, user.People)
	if err != nil {
		return "", err
	}

	return a, nil
}

type queryResolver struct{ server *GQLServer }

func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*pkg.User, error) {
	// var gc *gin.Context
	// ctx := getContext(gc)
	// var k []string
	// var g []string
	var q pkg.User
	var z []*pkg.User

	if user := authentication.ForContext(ctx); user == nil {
		return nil, fmt.Errorf("Access denied")
	}

	a, err := r.server.userClient.GetAllUsers(ctx, "")
	if err != nil {
		return nil, err
	}

	// log.Println(a)
	// var user1 = User{
	// 	ID:       "12324",
	// 	Username: "kg",
	// 	Password: "123",
	// }
	// var user2 = User{
	// 	ID:       "12325",
	// 	Username: "kg",
	// 	Password: "lake",
	// }
	// var user User
	// user.Username = a.User.Username
	// user.Password = a.User.Password
	// user.Firstname = a.User.Firstname
	// user.Lastname = a.User.Lastname
	// user.EmailAddress = a.User.Emailaddress
	// user.Gender = a.User.Gender
	// user.ID = a.User.Id

	// return user, nil

	for index := 0; index < len(a); index++ {
		q.Username = a[index].Username
		q.Password = a[index].Password

		z = append(z, &q)

	}
	// z = append(z, user1)
	// z = append(z, user2)

	// for index := 0; index < len(z); index++ {
	// 	q.Username = z[index].Username
	// 	q.Password = z[index].Password

	// 	z = append(z, q)

	// }
	return z, nil
}

func (r *queryResolver) GetUser(ctx context.Context, id string) (*pkg.User, error) {
	// var gc *gin.Context
	// ctx := getContext(gc)

	a, err := r.server.userClient.GetUser(ctx, id)
	if err != nil {
		return &pkg.User{}, err
	}
	var user pkg.User
	user.Username = a.Username
	user.Password = a.Password

	return &user, nil
}

func (r *queryResolver) GetUserByUserName(ctx context.Context, username string) (*pkg.User, error) {
	// var gc *gin.Context
	// ctx := getContext(gc)

	a, err := r.server.userClient.GetUserByUserName(ctx, username)
	if err != nil {
		return &pkg.User{}, err
	}
	var user pkg.User
	user.Username = a.Username
	user.Password = a.Password
	user.Firstname = a.Firstname
	user.Lastname = a.Lastname
	user.EmailAddress = a.EmailAddress
	user.Gender = a.Gender
	user.ID = a.ID.Hex()

	return &user, nil
}

// GetUserByEmailAddress ...
func (r *queryResolver) GetUserByEmailAddress(ctx context.Context, email string) (*pkg.User, error) {
	// var gc *gin.Context
	// ctx := getContext(gc)

	a, err := r.server.userClient.GetUserByEmailAddress(ctx, email)
	if err != nil {
		return &pkg.User{}, err
	}
	var user pkg.User
	user.Username = a.Username
	user.Password = a.Password
	user.Firstname = a.Firstname
	user.Lastname = a.Lastname
	user.EmailAddress = a.EmailAddress
	user.Gender = a.Gender
	user.ID = a.ID.Hex()

	return &user, nil
}

func (r *queryResolver) GetAllPosts(ctx context.Context, pagination int) ([]*pkg.UserPost, error) {

	a, err := r.server.postClient.GetMultiple(ctx, 0, 0)
	if err != nil {
		return nil, err
	}

	posts := []*pkg.UserPost{}
	for _, a := range a {
		var pst = pkg.UserPost{}
		pst.ContentPhoto.Name = a.ContentPhoto.Name
		posts = append(posts, &pkg.UserPost{
			ID:           a.ID.Hex(),
			Author:       a.Author.Hex(),
			Topic:        a.Topic,
			Category:     a.Category,
			ContentText:  a.ContentText,
			ContentPhoto: pst.ContentPhoto,
		})
	}

	return posts, nil
}

func (r *queryResolver) GetPost(ctx context.Context, id string) (*pkg.UserPost, error) {

	a, err := r.server.postClient.Get(ctx, id, 0)
	if err != nil {
		return &pkg.UserPost{}, err
	}

	var pst = pkg.UserPost{}
	pst.ContentPhoto.Name = a.ContentPhoto.Name
	return &pkg.UserPost{
		ID:           a.ID.Hex(),
		Author:       a.Author.Hex(),
		Topic:        a.Topic,
		Category:     a.Category,
		ContentText:  a.ContentText,
		ContentPhoto: pst.ContentPhoto,
	}, nil
}

func (r *queryResolver) GetAllSpaces(ctx context.Context) ([]*pkg.Space, error) {
	var mngs, fllwrs []string

	a, err := r.server.spaceClient.GetMultiple(ctx, 0, 0)
	if err != nil {
		return nil, err
	}

	spaces := []*pkg.Space{}
	for _, a := range a {
		for _, m := range a.Managers {
			mngs = append(mngs, m.Hex())
		}
		for _, f := range a.Followers {
			fllwrs = append(fllwrs, f.Hex())
		}
		spaces = append(spaces, &pkg.Space{
			ID:          a.ID.Hex(),
			Creator:     a.Creator.Hex(),
			Topic:       a.Topic,
			Details:     a.Details,
			Description: a.Description,
			Type:        a.Type,
			Followers:   fllwrs,
			Managers:    mngs,
			Tags:        a.Tags,
		})
	}

	return spaces, nil
}

func (r *queryResolver) GetSpace(ctx context.Context, id string) (*pkg.Space, error) {
	var mngs, fllwrs []string
	a, err := r.server.spaceClient.Get(ctx, id, 0)
	if err != nil {
		return &pkg.Space{}, err
	}

	for _, m := range a.Managers {
		mngs = append(mngs, m.Hex())
	}
	for _, f := range a.Followers {
		fllwrs = append(fllwrs, f.Hex())
	}

	return &pkg.Space{
		ID:          a.ID.Hex(),
		Creator:     a.Creator.Hex(),
		Topic:       a.Topic,
		Details:     a.Details,
		Description: a.Description,
		Type:        a.Type,
		Followers:   fllwrs,
		Managers:    mngs,
		Tags:        a.Tags,
	}, nil
}

func (r *queryResolver) GetAllTasks(ctx context.Context) ([]*pkg.Task, error) {
	a, err := r.server.taskClient.GetMultiple(ctx, 0, 0)
	if err != nil {
		return nil, err
	}

	tasks := []*pkg.Task{}
	for _, a := range a {
		tasks = append(tasks, &pkg.Task{
			ID:   a.ID,
			Text: a.Text,
		})
	}

	return tasks, nil
}

func (r *queryResolver) GetTask(ctx context.Context, id string) (*pkg.Task, error) {
	userID, err := strconv.ParseUint(id, 10, 32)
	if err == nil {
		fmt.Printf("Type: %T \n", userID)
		fmt.Println(userID)
	}
	a, err := r.server.taskClient.Get(ctx, uint32(userID), 0)
	if err != nil {
		return &pkg.Task{}, err
	}

	return &pkg.Task{
		ID:   a.ID,
		Text: a.Text,
	}, nil
}

func (r *queryResolver) GetProfile(ctx context.Context, id string) (*pkg.Profile, error) {
	var fllwrs, fllwing []string
	var prof = pkg.Profile{}

	a, err := r.server.profileClient.Get(ctx, id, 0)
	if err != nil {
		return &pkg.Profile{}, err
	}

	prof.ProfileImage.Name = a.ProfileImage.Name
	prof.BackgroundImage.Name = a.BackgroundImage.Name

	for _, f := range a.Followers {
		fllwrs = append(fllwrs, f.Hex())
	}

	for _, f := range a.Following {
		fllwing = append(fllwing, f.Hex())
	}

	return &pkg.Profile{
		ID:              a.ID.Hex(),
		About:           a.About,
		BackgroundImage: prof.BackgroundImage,
		ProfileImage:    prof.ProfileImage,
		Followers:       fllwrs,
		Following:       fllwing,
	}, nil
}

func (r *queryResolver) GetLocationDistance(ctx context.Context, lon, lat float64) (float64, error) {

	a, err := r.server.geographyClient.GetLocationDistance(ctx, lon, lat)
	if err != nil {
		return 0.0, err
	}
	log.Println(a)
	return a, nil
}

func (r *queryResolver) UserChat(ctx context.Context, username string) (*pkg.Chat, error) {
	r.server.mu.Lock()

	cmd := r.server.redisClient.LRange("messages", 0, -1)
	if cmd.Err() != nil {
		log.Println(cmd.Err())
		return nil, cmd.Err()
	}
	res, err := cmd.Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	userChat := r.server.Chats[username]
	messages := []pkg.UserMessage{}
	log.Println("res ::: ", res)

	for _, mj := range res {
		var m pkg.UserMessage
		err = json.Unmarshal([]byte(mj), &m)
		messages = append(messages, m)
	}

	if userChat == nil {
		userChat = &pkg.Chat{
			ID:        username,
			Observers: map[string]chan *pkg.UserMessage{},
		}
		r.server.Chats[username] = userChat
		log.Println("error")
	}
	r.server.mu.Unlock()

	return userChat, nil
}

func (r *queryResolver) GetAllGoals(ctx context.Context) ([]*pkg.Goal, error) {
	var likes, participants, watchers []string

	a, err := r.server.goalClient.GetMultiple(ctx, 0, 0)
	if err != nil {
		return nil, err
	}

	goal := []*pkg.Goal{}

	for _, a := range a {
		journey := pkg.Journey{}
		for _, m := range a.Participants {
			participants = append(participants, m.Hex())
		}

		for _, m := range a.Likes {
			likes = append(likes, m.Hex())
		}

		for _, m := range a.Participants {
			watchers = append(watchers, m.Hex())
		}

		journey.Details = a.Journey.Details

		goal = append(goal, &pkg.Goal{
			ID:           a.ID.Hex(),
			Creator:      a.Creator.Hex(),
			Aim:          a.Aim,
			Type:         a.Type,
			IsAchieved:   a.IsAchieved,
			IsPrivate:    a.IsPrivate,
			Inspiration:  a.Inspiration,
			Journey:      journey,
			Likes:        likes,
			Reason:       a.Reason,
			SimilarGoals: a.SimilarGoals,
			Watchers:     watchers,
			Participants: participants,
			Tags:         a.Tags,
		})
	}

	return goal, nil
}

func (r *queryResolver) GetGoal(ctx context.Context, id string) (*pkg.Goal, error) {
	var likes, participants, watchers []string
	a, err := r.server.goalClient.Get(ctx, id, 0)
	if err != nil {
		return &pkg.Goal{}, err
	}

	journey := pkg.Journey{}
	for _, m := range a.Participants {
		participants = append(participants, m.Hex())
	}

	for _, m := range a.Likes {
		likes = append(likes, m.Hex())
	}

	for _, m := range a.Participants {
		watchers = append(watchers, m.Hex())
	}

	journey.Details = a.Journey.Details

	return &pkg.Goal{
		ID:           a.ID.Hex(),
		Creator:      a.Creator.Hex(),
		Aim:          a.Aim,
		Type:         a.Type,
		IsAchieved:   a.IsAchieved,
		IsPrivate:    a.IsPrivate,
		Inspiration:  a.Inspiration,
		Journey:      journey,
		Likes:        likes,
		Reason:       a.Reason,
		SimilarGoals: a.SimilarGoals,
		Watchers:     watchers,
		Participants: participants,
		Tags:         a.Tags,
	}, nil
}

func (r *queryResolver) GetAllGroups(ctx context.Context) ([]*pkg.Group, error) {
	var people []string
	a, err := r.server.groupClient.GetMultiple(ctx, 0, 0)
	if err != nil {
		return nil, err
	}

	groups := []*pkg.Group{}

	for _, a := range a {
		for _, m := range a.People {
			people = append(people, m.Hex())
		}
		groups = append(groups, &pkg.Group{
			ID:          a.ID.Hex(),
			Title:       a.Title,
			People:      people,
			Description: a.Description,
			Details:     a.Details,
			Type:        a.Type,
		})
	}

	return groups, nil
}

func (r *queryResolver) GetGroup(ctx context.Context, id string) (*pkg.Group, error) {
	var people []string
	a, err := r.server.groupClient.Get(ctx, id, 0)
	if err != nil {
		return &pkg.Group{}, err
	}

	for _, m := range a.People {
		people = append(people, m.Hex())
	}

	return &pkg.Group{
		ID:          a.ID.Hex(),
		Title:       a.Title,
		People:      people,
		Description: a.Description,
		Details:     a.Details,
		Type:        a.Type,
	}, nil
}

func (r *queryResolver) GetAllMatches(ctx context.Context) ([]*pkg.MatchedUser, error) {
	// var people []string
	// a, err := r.server.groupClient.GetMultiple(ctx, 0, 0)
	// if err != nil {
	// 	return nil, err
	// }

	// groups := []*Group{}

	// for _, a := range a {
	// 	for _, m := range a.People {
	// 		people = append(people, m.Hex())
	// 	}
	// 	groups = append(groups, &Group{
	// 		ID:          a.ID.Hex(),
	// 		Title:       a.Title,
	// 		People:      people,
	// 		Description: a.Description,
	// 		Details:     a.Details,
	// 		Type:        a.Type,
	// 	})
	// }

	return nil, nil
}

type subscriptionResolver struct{ server *GQLServer }

func (r *subscriptionResolver) UserMessageAdded(ctx context.Context, chatName string) (<-chan *pkg.UserMessage, error) {
	r.server.mu.Lock()
	chat := r.server.Chats[chatName]
	if chat == nil {
		chat = &pkg.Chat{
			ID:        chatName,
			Observers: map[string]chan *pkg.UserMessage{},
		}
		r.server.Chats[chatName] = chat
	}
	r.server.mu.Unlock()

	id := randString(8)
	events := make(chan *pkg.UserMessage, 1)

	go func() {
		<-ctx.Done()
		r.server.mu.Lock()
		delete(chat.Observers, id)
		r.server.mu.Unlock()
	}()

	r.server.mu.Lock()
	chat.Observers[id] = events
	r.server.mu.Unlock()

	return events, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
