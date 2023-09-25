package controller
import (
"github.com/streadway/amqp"
	"net/http"
	"testing"
		"log"
	 "io/ioutil"
    "net/http/httptest"
)


//testing file compression and storage
func TestCompressImg(t *testing.T) {
	ProductImages := []string{"test_images/img1.jpeg", "test_images/img3.jpeg"}
	expectation := []string{"compressed_images/img1_compressed.jpg", "compressed_images/img3_compressed.jpg"}
	
	res, _ := CompressAndStoreInLocal(ProductImages)

	// Check if the lengths of the slices are the same
	if len(res) != len(expectation) {
		t.Errorf("Length mismatch: Expected %d items, but got %d", len(expectation), len(res))
	}

	// Compare each element of the slices
	for i := 0; i < len(res); i++ {
		if res[i] != expectation[i] {
			t.Errorf("Mismatch at index %d: Expected '%s', but got '%s'", i, expectation[i], res[i])
		}
	}
}
//testing  add product function
func TestAddProduct(t *testing.T) {
expected := `{"Status":"success","Message":"Product Created Successfully","Payload":[]}` // Update Payload as needed
         
    req := httptest.NewRequest(http.MethodPost, "/product", nil)
    w := httptest.NewRecorder()
    AddProduct(w, req)
    res := w.Result()
    defer res.Body.Close()
    data, err := ioutil.ReadAll(res.Body)
    if err != nil {
        t.Errorf("Error: %v", err)
    }
    if string(data) != expected {
        t.Errorf("Expected Status success but got %v", string(data))
    }else{
		log.Println("PASS TEST")
	}
}

//testing msg queue


func TestQueue(t *testing.T){
    connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err!=nil{
		t.Errorf("Error: %v", err)
	}
	defer connection.Close()
    // Call the Queue function with the  connection

     Producer(connection, 2)
    Consumer(connection) 

   
}