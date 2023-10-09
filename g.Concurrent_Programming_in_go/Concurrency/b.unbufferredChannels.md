## Unbuffered Channels

> This type of channel only allows to send one piece of data and blocks the current goroutine until another one performs a receive operation on the channel. The same thing will happen if a receive operation on a channel is performed before a send operation, the goroutine where the receive operation was made will be blocked until another goroutine sends a message through the same channel. lets see it with an example

``` go
package main

import (
	"fmt"
	"time"
)

var (
	defaultTags = []string{"SystemUser", "User", "NewUser", "System"}
)

type Tag struct {
	Name, Type string
}

type User struct {
	Id, Name, LastName, Status string
	Tags                       []*Tag
}

type Post struct {
	Title  string
	Status string
	UserId string
}

func main() {
	blocking()	
}
 
/*
Main goroutine will be blocked until second goroutine
sends a message letting the main goroutine know that has finished its work
and so the main go routine can continue
*/
func blocking() {
	user := &User{}
	done := make(chan bool) // unbuffered channel

	go func() {
		fmt.Println("[Second-GoRoutine] Start Building User")
		buildingUser(user)
		fmt.Println("[Second-GoRoutine] Finished Building User")
		done <- true

		fmt.Println("[Second-GoRoutine] Set default user tags")
		setDefaultTags(user)
	}()

	fmt.Println("[Main-Goroutine] Start importing Posts")
	posts := importingPosts()
	fmt.Println("[Main-Goroutine] Finish importing Posts")
	fmt.Println("[Main-Goroutine] -----waiting------")
	<-done

	mergeUserPosts(user, posts)
	fmt.Println("Done!!")
	fmt.Printf("User %v\n", user)
	for _, post := range posts {
		fmt.Printf("Post %v\n", post)
	}
}

func mergeUserPosts(user *User, posts []*Post) {
	fmt.Println("[Main-Goroutine] Start merging user posts")
	for _, post := range posts {
		post.UserId = user.Id
	}
	fmt.Println("[Main-Goroutine] Finished merging user posts")
}

func importingPosts() []*Post {
	time.Sleep(1 * time.Second)
	titles := []string{"Post 1", "Random Post", "Second Post"}
	posts := []*Post{}
	for _, title := range titles {
		posts = append(posts, &Post{Title: title, Status: "draft"})
	}

	return posts
}

func buildingUser(user *User) {
	time.Sleep(2 * time.Second)
	user.Name = "John"
	user.LastName = "Doe"
	user.Status = "active"
	user.Id = "1"
}

func setDefaultTags(user *User) {
	time.Sleep(1 * time.Second)
	for _, tagName := range defaultTags {
		user.Tags = append(user.Tags, &Tag{Name: tagName, Type: "System"})
	}
}
```
Output
> main goroutine start creating features
> created feature=> 0x1400011a020
> main go routine waits 
> create product go routine starts
> this is app &{ []}
> creating app in main go routine &{laptop this is laptop 190 false []} &{ []}
> {test app [0x1400012c000]}

Let's understand the output

> * Let’s understand the output. When the program starts, an empty user object is created and a channel of type boolean(unbuffered channel) in lines 36 and 37 respectively, then a goroutine is created in line 45 which means that the piece of code within that function will be running in a separate goroutine.
> * The execution of the main goroutine continues and in line 49 we have a print statement, then in the second goroutine, since it is running concurrently at this point, it reaches line 40 and also executes a print statement.
> * The main goroutine continues and calls the method importingPosts and makes also two more print statements, the last one being [Main-Goroutine] ——--waiting------ , this is where the blocking concept that we talked about earlier comes into play, in line 59 we see that the main goroutine is reading from the donechannel, this basically means that the main goroutine will not continue its execution until the second goroutine sends a message to this channel.
> * In the second goroutine, the buildUser function is called and it prints [Second-GoRoutine] Finished Building User , then in the next line, it sends a message to the channel. At this point, the main goroutine will detect this and it will continue its execution, as well as the second goroutine.
> * The methods mergeUserPosts and setDefaultTags are called in the main and second goroutine respectively and we get their corresponding logs.
> * When we get to lines 63 to 66, the user and its posts are printed out, but if you check the tags array in the user struct is empty. The reason is that after the second goroutine sent a message to the main goroutine, both goroutines continued executing concurrently and as I previously mention the main goroutine will not wait until other goroutines finished executing, that being said, the second goroutine did not complete its work appending the user tags into the struct before the main goroutine finished and that is why the array is empty. If we remove line 97, we’ll be able to see the tags array is now filled in.

### another example
```go
package dummy3

import (
	"fmt"
	"time"
)

type Features struct {
	specs []string
}

type Product struct {
	name    string
	summary string
	price   float64
	active  bool
	feature []*Features
}

type App struct {
	name    string
	product []*Product
}

/*
Main goroutine will be blocked until second goroutine
sends a message letting the main goroutine know that has finished its work
and so the main go routine can continue
*/
func Blocking() {
	done := make(chan bool)
	var product Product
	var features Features
	var app App

	// fmt.Println(roduct)

	go func() {
		fmt.Println("create product go routine starts")
		createProduct(&product)
		done <- true
		makeProductActive(&product)
		addProductFeature(&product, &features)
		fmt.Println(product.feature[0].specs)
		fmt.Println("creat product go routine end")
	}()
	fmt.Println("main goroutine start creating features")
	createFeature(&features)
	fmt.Println("main go routine waits ")
	<-done
	fmt.Println("this is app", &app)
	createApp(&product, &features, &app)
	fmt.Println(app)
}

func createProduct(product *Product) {
	time.Sleep(1 * time.Second)
	product.name = "laptop"
	product.summary = "this is laptop"
	product.price = 190
}

func createFeature(feature *Features) {
	feature.specs = []string{"8gb ram"}
	fmt.Println("created feature=>", &feature)
}

func createApp(product *Product, feature *Features, app *App) {
	fmt.Println("creating app in main go routine", product, app)
	app.name = "test app"
	app.product = []*Product{product}
}

func makeProductActive(product *Product) {
	time.Sleep(time.Second * 1)
	product.active = true
}

func addProductFeature(product *Product, feature *Features) {
	fmt.Println("adding product feature")
	product.feature = []*Features{feature}
	fmt.Println("in addProductFeature", product)
}

```

> * With this example, we learned how to create an unbuffered channel using the built-in make function.
>> ```go 
>>  done := make(chan int) 
>> ```
> Also how to send and receive data from a channel
>> ``` go
>>  done <- true // send
>>  <-done // receive ignorting value
>> resp := <-done // receive storing value in a variable
>> ```
> Also, we saw how goroutines block execution if no other goroutine has sent/receive a message through the channel.

### channels can also be used to as a way of communicating multiple goroutines by using the result of one goroutines as the parameters for the another one 
#### lets take a look at the another example using multiple goroutines

```go
package main

import (
	"fmt"
	"time"
)

type Tag struct {
	Name, Type string
}

type Setting struct {
	NotificationEnabled bool
}

type User struct {
	Name, Status string
	Tag          []*Tag
	*Setting
}

type NotificationService struct {
}

func main() {
	usersToUpdate := make(chan []*User)
	userToNotify := make(chan *User)

	existingUser := []*User{
		{Name: "John", Status: "Active", Setting: &Setting{NotificationEnabled: true}},
		{Name: "Mohn", Status: "Active", Setting: &Setting{NotificationEnabled: true}},
		{Name: "Walker", Status: "Active", Setting: &Setting{NotificationEnabled: true}},
		{Name: "Paul", Status: "Active", Setting: &Setting{NotificationEnabled: true}},
	}

	newUsers := []*User{
		{Name: "Sachin", Status: "Active", Setting: &Setting{NotificationEnabled: true}},
		{Name: "Rahul", Status: "Active", Setting: &Setting{NotificationEnabled: false}},
		{Name: "Sourav", Status: "Active", Setting: &Setting{NotificationEnabled: true}},
		{Name: "VVS", Status: "Active", Setting: &Setting{NotificationEnabled: true}},
	}

	go filterUser(usersToUpdate, newUsers)
	go updateUsers(usersToUpdate, userToNotify, existingUser)
	notifyUser(userToNotify, existingUser)
}

func filterUser(usersToUpdate chan<- []*User, newUsers []*User) {
	defer close(usersToUpdate)
	fmt.Println("------------filterUsers called------------")
	filteredUsers := []*User{}
	for _, user := range newUsers {
		if user.Status == "Active" && user.NotificationEnabled {
			filteredUsers = append(filteredUsers, user)
		}
	}
	usersToUpdate <- filteredUsers
}

func updateUsers(usersToUpdate <-chan []*User, userToNotify chan<- *User, existingUser []*User) {
	defer close(userToNotify)
	fmt.Println("------------updateUsers called------------")
	for _, user := range existingUser {
		user.Tag = append(user.Tag, &Tag{Name: "NotificationEnabled", Type: "Notification"})
	}

	usersTobeUpdated := <-usersToUpdate

	for _, user := range usersTobeUpdated {
		time.Sleep(time.Second * 1)
		user.Tag = append(user.Tag, &Tag{Name: "notification", Type: "Notificaiton"})
		userToNotify <- user
	}
}

func notifyUser(userToNotify <-chan *User, existingUser []*User) {
	fmt.Println("------------NotifyUsers called------------")
	service := &NotificationService{}
	message := "You have received a new message"
	for _, user := range existingUser {
		service.SendNotification(user, message)
	}
	fmt.Println("==============will be blocked here until it receives=============")
	for userV := range userToNotify {
		service.SendNotification(userV, message)
	}
}

func (s *NotificationService) SendNotification(user *User, message string) {
	fmt.Println("------------SendEmailNotification called-----------")
	fmt.Printf("%v Hello %s, %s with tag %v\n", user, user.Name, message, user.Tag)
}
```

> * In this example, we have two channels usersToUpdate and userToNotify, notice how the first channel accepts an array of users and the second one only one single user object. Then there are two arrays of users, one for existing users and one for new users.
> * In the first goroutine, we send the usersToUpdate channel and the slice of newUsers, so when the program gets to line 40 a new goroutine is created.
> * Notice the syntax in filterUser function for the usersToUpdate param.  usersToUpdate chan<- []*User
> * Channels by default are bi-directional meaning that you can send and receive information through them, but when passing a channel to a function you can change this behavior and tell the channel that in the context of the function it will only serve one purpose, either to receive information or to send information.
> * So in this case, we are telling the channel usersToUpdate that in the context of this function this channel will only accept sending information and not receiving it.
> * This function filterUser range over the newUsers and only selects the ones that are active and has the setting enabled for notifications. After that the filtered users are sent through the channel.
> * At this point, this channel will not be used anymore for sending data so it is important to close the channel. In this case, we are using the defer function to call the built-in close function and close the usersToUpdate channel.
> * In the second goroutine, we send the usersToUpdate channel, the userToNotify channel and the existingUsers slice. This is where the concept of using a channel’s results as the input for another goroutine comes into play.
> * In this function we are also defining for each channel if it will be used for receiving information or sending information, usersToUpdate will be used to only receive data and userToNotify to send data.
> * updateUsers function first updated the existing users by appending a new tag to each of them. it creates a new variable assigning it to the result of the usersToUpdate channel. This line will block the execution of this goroutine until the channel sends a message. In other words, if the filterUser takes a lot of time to send the filteredUsers, this goroutine will have to wait in this line before proceeding.
> * Once the data is received this goroutine ranges over the newUsers and also updates their tags, but also sends the user through the userToNotify channel.
> * The userToNotify will also need to be closed after this function completes its work, so we have a defer to close the channel.
> * there is a function notifyUsers that is called in the main goroutine, that will notify users, it takes the userToNotify channel and the existingUsers as parameters.
> * This function first initializes a service for sending notifications and then ranges over the existing users and sends an email notification to each of them.
> * Then it ranges over the userToNotify channel, and for each user that is sent through this channel, this function sends an email notification to that user. This syntax allows us to receive all the information sent through this channel and once the channel is closed, the for loop will break too. This will prevent us from reading from a closed channel, as I mentioned before this is one way of ensuring you don’t read from a closed channel. The other syntax is as follows: resp, ok := <-userToNofity
> * The ok variable will be false if we are reading from a closed channel and true otherwise, but it will not panic.
> * As you can see in this example the functions will run concurrently and they communicate to each other using channels to send information about the filtered users and the users to notify.
> * In this example, we learned how to close a channel using defer and close function. defer close(done)
> * Also how to make a channel uni-directional when it is passed to a function
> * userToNotify <-chan *User // read-only channel userToNotify chan<- *User // send-only channel
> * How to range over a channel, this is useful when we don’t know how many items will be sent through a channel but we want to read all of them. for user := range userToNotify {}
> * 

> Note that it is only necessary to close a channel if the receiver is looking for a close. Closing the channel is a control signal on the channel indicating that no more data follows.
