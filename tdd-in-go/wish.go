s := reflect.MakeStruct()
reflect.AddFunc(s, "DoSomething", func(something reflect.Magic) {
	// do stuff
})