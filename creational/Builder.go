package main

// components

type Chasis struct {
	chasisType string
}

type Bodywork struct {
	brand string
}

// Components builder

type ComponentCarBuilder interface {
	buildChasis() Chasis
	buildBodywork() Bodywork
}

type Mazda struct {
	chasis   Chasis
	bodywork Bodywork
}

func NewMazda() *Mazda {
	return &Mazda{}
}

type MazdaBuilder struct {
}

func (mazda MazdaBuilder) buildChasis() Chasis {
	return Chasis{"Chasis mazda"}
}

func (mazda MazdaBuilder) buildBodywork() Bodywork {
	return Bodywork{"Bodywork mazda"}
}

// Director

type Director struct {
}

func (director Director) buildMazda(builderCar MazdaBuilder) Mazda {
	mazda := NewMazda()
	mazda.chasis = builderCar.buildChasis()
	mazda.bodywork = builderCar.buildBodywork()
	return *mazda
}
