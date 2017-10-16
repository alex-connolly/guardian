package ast

type DMap struct {
	keyMap map[string]IntPair
	array  []StringPair
	index  int
}

type IntPair struct {
	Node  Node
	Index int
}

type StringPair struct {
	Node Node
	Key  string
}

func (d *DMap) Next() StringPair {
	pair := d.array[d.index]
	d.index++
	return pair
}

func (d *DMap) Get(key string) Node {
	return d.keyMap[key].Node
}

func (d *DMap) Add(key string, node Node) {
	_, ok := d.keyMap[key]
	if ok {
		// already in map
		d.Remove(key)
		d.Add(key, node)
	} else {
		d.keyMap[key] = IntPair{node, len(d.array)}
		d.array = append(d.array, StringPair{node, key})
	}
}

func (d *DMap) Remove(key string) {
	pair, ok := d.keyMap[key]
	if ok {
		d.array = append(d.array[:pair.Index], d.array[pair.Index+1:]...)
		delete(d.keyMap, key)
	}
}
