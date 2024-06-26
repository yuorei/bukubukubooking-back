package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.48

import (
	"context"
	"fmt"

	"github.com/yuorei/bukubukubooking-back/graph/generated"
	"github.com/yuorei/bukubukubooking-back/src/domain"
)

// ID is the resolver for the id field.
func (r *messageResolver) ID(ctx context.Context, obj *domain.Message) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// Sender is the resolver for the sender field.
func (r *messageResolver) Sender(ctx context.Context, obj *domain.Message) (*domain.User, error) {
	panic(fmt.Errorf("not implemented: Sender - sender"))
}

// CreateMessage is the resolver for the createMessage field.
func (r *mutationResolver) CreateMessage(ctx context.Context, input domain.CreateMessageInput) (*domain.CreateMessagePayload, error) {
	panic(fmt.Errorf("not implemented: CreateMessage - createMessage"))
}

// UpdateMessage is the resolver for the updateMessage field.
func (r *mutationResolver) UpdateMessage(ctx context.Context, input domain.UpdateMessageInput) (*domain.UpdateMessagePayload, error) {
	panic(fmt.Errorf("not implemented: UpdateMessage - updateMessage"))
}

// DeleteMessage is the resolver for the deleteMessage field.
func (r *mutationResolver) DeleteMessage(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteMessage - deleteMessage"))
}

// Message is the resolver for the message field.
func (r *queryResolver) Message(ctx context.Context, id string) (*domain.Message, error) {
	panic(fmt.Errorf("not implemented: Message - message"))
}

// Messages is the resolver for the messages field.
func (r *queryResolver) Messages(ctx context.Context) ([]*domain.Message, error) {
	panic(fmt.Errorf("not implemented: Messages - messages"))
}

// MessageAdded is the resolver for the messageAdded field.
func (r *subscriptionResolver) MessageAdded(ctx context.Context, requestID string) (<-chan *domain.Message, error) {
	panic(fmt.Errorf("not implemented: MessageAdded - messageAdded"))
}

// ID is the resolver for the id field.
func (r *createMessagePayloadResolver) ID(ctx context.Context, obj *domain.CreateMessagePayload) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// Sender is the resolver for the sender field.
func (r *createMessagePayloadResolver) Sender(ctx context.Context, obj *domain.CreateMessagePayload) (*domain.User, error) {
	panic(fmt.Errorf("not implemented: Sender - sender"))
}

// ID is the resolver for the id field.
func (r *updateMessagePayloadResolver) ID(ctx context.Context, obj *domain.UpdateMessagePayload) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// Sender is the resolver for the sender field.
func (r *updateMessagePayloadResolver) Sender(ctx context.Context, obj *domain.UpdateMessagePayload) (*domain.User, error) {
	panic(fmt.Errorf("not implemented: Sender - sender"))
}

// Message returns generated.MessageResolver implementation.
func (r *Resolver) Message() generated.MessageResolver { return &messageResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

// CreateMessagePayload returns generated.CreateMessagePayloadResolver implementation.
func (r *Resolver) CreateMessagePayload() generated.CreateMessagePayloadResolver {
	return &createMessagePayloadResolver{r}
}

// UpdateMessagePayload returns generated.UpdateMessagePayloadResolver implementation.
func (r *Resolver) UpdateMessagePayload() generated.UpdateMessagePayloadResolver {
	return &updateMessagePayloadResolver{r}
}

type messageResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type createMessagePayloadResolver struct{ *Resolver }
type updateMessagePayloadResolver struct{ *Resolver }
