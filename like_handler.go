package IGapi

//
// import (
// 	"fmt"
// 	"net/http"
//
// 	"github.com/gorilla/mux"
// 	"github.com/nsqio/go-nsq"
// )
//
// var (
// 	producer *nsq.Producer
// )
//
// func init() {
// 	var err error
// 	if producer, err = nsq.NewProducer("localhost:4150", nsq.NewConfig()); err != nil {
// 		panic(err)
// 	}
// }
//
// func LikeHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	fmt.Fprint(w, "found", vars["media_id"])
// 	if err := producer.Publish("instagram_likes", []byte(vars["media_id"])); err != nil {
// 		panic(err)
// 	}
