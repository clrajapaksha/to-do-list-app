package cache

type Cache[K comparable, V any] struct {
	items    map[K]V
	channel  chan V
	capacity int
}

func (c *Cache[K, V]) readRoutine(key K) {
	c.channel <- c.items[key]
}

func (c *Cache[K, V]) safeRead(key K) (V, bool) {
	go c.readRoutine(key)
	_, found := c.items[key]
	output := <-c.channel
	return output, found
}

func (c *Cache[K, V]) writeRoutine(key K, value V) {
	c.items[key] = value
	c.channel <- c.items[key]
}

func (c *Cache[K, V]) safeWrite(key K, value V) V {
	go c.writeRoutine(key, value)
	output := <-c.channel
	return output
}

func (c *Cache[K, V]) deleteRoutine(key K) {
	value := c.items[key]
	delete(c.items, key)
	c.channel <- value
}

func (c *Cache[K, V]) safeDelete(key K) V {
	go c.deleteRoutine(key)
	output := <-c.channel
	return output
}

func New[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		items:    make(map[K]V),
		channel:  make(chan V),
		capacity: 30,
	}
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.safeWrite(key, value)
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	value, found := c.safeRead(key)
	return value, found
}

func (c *Cache[K, V]) Remove(key K) {
	c.safeDelete(key)
}

func (c *Cache[K, V]) Pop(key K) (V, bool) {
	value, found := c.safeRead(key)
	if found {
		c.safeDelete(key)
	}
	return value, found
}
