package parrot

import "net"

// Users - connected users
type Users struct {
	users map[string]*net.Conn
}

func (users *Users) add(u string, c *net.Conn) {
	users.users[u] = c
}

func (users *Users) remove(u string) {
	delete(users.users, u)
}

func (users *Users) getAll() []string {
	result := make([]string, len(users.users))
	for key := range users.users {
		result = append(result, key)
	}
	return result
}
