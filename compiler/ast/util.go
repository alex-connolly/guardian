package ast

type DMap struct {
	Map   map[string]IntPair
	Array []StringPair
	index int
}

// Goals:
// should be able to iterate over map
// should also be able to get ordered out of a map

type IntPair struct {
	Node  Node
	Index int
}

type StringPair struct {
	Node Node
	Key  string
}

func (d *DMap) Next() StringPair {
	pair := d.Array[d.index]
	d.index++
	return pair
}

func (d *DMap) Get(key string) Node {
	return d.Map[key].Node
}

func (d *DMap) Add(key string, node Node) {
	_, ok := d.Map[key]
	if ok {
		// already in map
		d.Remove(key)
		d.Add(key, node)
	} else {
		d.Map[key] = IntPair{node, len(d.Array)}
		d.Array = append(d.Array, StringPair{node, key})
	}
}

func (d *DMap) Remove(key string) {
	pair, ok := d.Map[key]
	if ok {
		d.Array = append(d.Array[:pair.Index], d.Array[pair.Index+1:]...)
		delete(d.Map, key)
	}
}
