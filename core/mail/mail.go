package mail

//TCP connection to ensure each message is received (by high chance) by the counterpart
var Connected []string

//Create connection to the specified IP address
func Connect_to(ip string) {

	Connected = append(Connected, ip)
}

//Send message to the specific IP address
func Send_to() {

}
