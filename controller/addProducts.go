package  controller
import (
	"encoding/json"
	"log"
	"net/http"
	"msg_quing_system/models"
	"msg_quing_system/utility"
	
)

func JsonDecoder(r *http.Request) (models.AddProduct, error) {
	var usersData models.AddProduct

	err := json.NewDecoder(r.Body).Decode(&usersData)
	if err != nil {
		log.Println(err)
	}
	return usersData, err
}

func AddProduct(w http.ResponseWriter, r *http.Request) utility.AjaxResponse {
	response := utility.AjaxResponse{Status: "success", Message: "Product Created Successfully",Payload:[]interface{}{}} // DEFAULT RESPONSE
	var result models.AddProduct
	// send json response by decode json
	result, err := JsonDecoder(r)
	if err != nil {
		log.Println(err)
	} else {
		product_id, err := Db.product.AddProductDb(result)
		if err != nil {
			response.Status = "failure"
			response.Message = "Something Went Wrong"
		}else{

			response,err = InvokeMesgQue(product_id) //invoking msg queue to start rabbitmq and send id to producer
			if err!=nil{
				log.Println(err)
			}
		}
	}
	utility.RenderTemplate(w,r,"",response)
	return response
}


