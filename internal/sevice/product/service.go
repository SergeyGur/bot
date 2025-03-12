package product

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProduct
}

func (s *Service) Get(idx int) (*Product, error) {
	/*if idx < 0 || idx > len(allProduct) {
		return nil, fmt.Errorf("invalid index : %v", idx)
	}*/
	return &allProduct[idx], nil
}
