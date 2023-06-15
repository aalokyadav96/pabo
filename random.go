package main

/*
func getDash() dashItems {
	var res dashItems
	res = dashItems{Content: "Sophistication"}
	return res
}

type dashItems struct {
	Content string
}

*/

func getDash() error{
	var foo Foo
	return conn.HGetAll("key").Scan(&foo)
}


type Foo struct {
     Key1 string
     Key2 int
}
