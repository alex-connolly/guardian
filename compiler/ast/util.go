package ast

type DMap struct {
	Map   map[string]pair
	Array []Node
	Index int
}

type pair struct {
	node  Node
	index int
}

func (d *DMap) Next() Node {
	node := d.Array[d.Index]
	d.Index++
	return node
}

func (d *DMap) Get(key string) Node {
	return d.Map[key].node
}

func (d *DMap) Add(key string, node Node) {
	_, ok := d.Map[key]
	if ok {
		// already in map
		d.Remove(key)
		d.Add(key, node)
	} else {
		d.Map[key] = pair{node, len(d.Array)}
		d.Array = append(d.Array, node)
	}
}

func (d *DMap) Remove(key string) {
	pair, ok := d.Map[key]
	if ok {
		d.Array = append(d.Array[:pair.index], d.Array[pair.index+1:]...)
		delete(d.Map, key)
	}
}
