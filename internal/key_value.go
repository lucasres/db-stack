package internal

type KeyValue struct {
	Key   string
	Stack *Stack
}

func NewKeyValue(key string) *KeyValue {
	return &KeyValue{
		Key:   key,
		Stack: NewStack(),
	}
}

func (k *KeyValue) Get() string {
	return k.Stack.Get()
}

func (k *KeyValue) Pop() {
	k.Stack.Pop()
}

func (k *KeyValue) Push(v string) {
	k.Stack.Push(v)
}
