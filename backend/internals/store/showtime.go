package store

// func (s *postgresStore) AddShowtime(ctx context.Context, newUser schemas.UserSignup) error {
// 	err := s.connect(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	defer s.close()

// 	if _, err := s.db.NamedExecContext(
// 		ctx,
// 		`INSERT INTO users
// 		(username, password_hash, role)
// 		VALUES
// 		(:username,:password_hash, 'regular')`,
// 		newUser); err != nil {
// 		if strings.Contains(err.Error(), "SQLSTATE 23505") {
// 			return fmt.Errorf("username already exists : %s", newUser.Username)
// 		}
// 		return err
// 	}
// 	return nil
// }

// func (s *postgresStore) UserLogin(ctx context.Context, newUser schemas.UserLogin) (string, error) {
// 	err := s.connect(ctx)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer s.close()

// 	// users := make([]schemas.UserToken, 5)
// 	users := []schemas.UserToken{{}}

// 	if err := s.db.SelectContext(
// 		ctx,
// 		&users,
// 		`SELECT id, username, role FROM users
// 			where username=$1 and password_hash=$2 LIMIT 1;`,
// 		newUser.Username, newUser.EncodedPass,
// 	); err != nil {
// 		return "", err
// 	}
// 	if len(users) == 0 {
// 		return "", errors.New("invalid credentials")
// 	}
// 	token := utils.GenerateToken(users[0])
// 	return token, nil
// }

// func (s *postgresStore) GetUserDetails(ctx context.Context, userID int) (*schemas.UserDetails, error) {
// 	err := s.connect(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer s.close()

// 	users := []schemas.UserDetails{}

// 	if err := s.db.SelectContext(
// 		ctx,
// 		&users,
// 		`SELECT id, username, role, created_at, updated_at FROM users
// 		WHERE users.id=$1 LIMIt 1;`,
// 		userID); err != nil {
// 		return nil, err
// 	}
// 	return &users[0], nil
// }
