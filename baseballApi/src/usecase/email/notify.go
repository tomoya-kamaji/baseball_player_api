package usecase

type Message interface {
	Send()
	Format() string
}
