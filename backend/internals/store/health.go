package store

import "context"

func (s *postgresStore) Health(ctx context.Context) error {
	err := s.connect(ctx)
	if err != nil {
		return err
	}
	defer s.close()
	err = s.db.PingContext(ctx)
	return err
}
