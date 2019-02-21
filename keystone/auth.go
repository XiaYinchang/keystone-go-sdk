package keystone

func NewAuth(username, password, domainName string) Auth {
	// create a password-based auth
	return Auth{
		Identity: Identity{
			Methods: []string{"password"},
			Password: IdentityUser{
				User: User{
					Name:     username,
					Password: password,
					Domain: Domain{
						Name: domainName,
					},
				},
			},
		},
	}
}
