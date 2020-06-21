package main

func Sync(client Client) {

	var inList bool = false
	for i, myclient := range myclients {
		if myclient.ip == client.ip {
			myclients = RemoveIndex(myclients, i)
			myclients = append(myclients, client)
			inList = true
			break
		}
	}
	if inList == false {
		myclients = append(myclients, client)
	}
}

func RemoveIndex(s []Client, index int) []Client {

	return append(s[:index], s[index+1:]...)
}
