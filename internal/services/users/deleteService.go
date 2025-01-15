package users

func (s *ServiceImpl) DeleteService(userId string) error {
	return s.userRepo.Delete(userId)
}