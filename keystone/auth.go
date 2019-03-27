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

func NewSystemAuth(username, password, domainName string) SystemScopedAuth {
	// create a password-based auth
	return SystemScopedAuth{
		Auth: SysAuth{
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
			Scope: Scope{
				System: System{
					All: true,
				},
			},
		},
	}
}
