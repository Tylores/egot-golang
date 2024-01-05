package ecs

import (
	"slices"
)

type World struct {
	entities     []*Entity
	new_entities []*Entity
	entity_map   map[string][]*Entity
	id_count     Id
}

func NewWorld() *World {
	return &World{
		entities:     nil,
		new_entities: nil,
		entity_map:   nil,
		id_count:     0,
	}
}

func (world *World) AddEntity(tag string) *Entity {
	world.id_count++
	e := &Entity{
		id:  world.id_count,
		tag: tag}
	world.new_entities = append(world.entities, e)
	return e
}

func (world *World) update() {
	for _, e := range world.new_entities {
		world.entities = append(world.entities, e)
		world.entity_map[e.tag] = append(world.entity_map[e.tag], e)
	}
	world.new_entities = nil

	world.entities = slices.DeleteFunc(world.entities, func(e *Entity) bool {
		return !e.active
	})
}

func (world *World) getEntities() []*Entity {
	return world.entities
}

func (world *World) getTaggedEntities(tag string) ([]*Entity, bool) {
	entities, found := world.entity_map[tag]
	return entities, found
}
