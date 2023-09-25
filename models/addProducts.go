package models

import (
	"log"
	"encoding/json"
	"msg_quing_system/utility"
	"github.com/jmoiron/sqlx"
)
type AddProduct struct {
	UserId              int64 			`db:"user_id"`
	Id       			int64  			`db:"id"`
	ProductName     	string  	 	`db:"product_name"`
	ProductDescription 	string 			`db:"product_description"`
	ProductImages 		[]string 		`db:"product_images"`
	ProductPrice   		float64 		`db:"product_price"`
	Product_id 			int64
	ProductCompressedImg []string       `db:"compressed_product_images"`
}

type AddProductModel struct {
	DB *sqlx.DB
}
func (data AddProductModel) CreatedNewUser() (int64, error) {
	query := `INSERT INTO users (name, mobile, latitude, longitude)
	VALUES (:name, :mobile, :latitude, :longitude)
	`
	values := map[string]interface{}{
	    "name":      "2nd user",
	    "mobile":    1234567890,
	    "latitude":  37.7749,
	    "longitude": -122.4194,
	}

	id, err := utility.Db.NamedExec(query, values)
	if err != nil {
	    log.Println(err)
	}else{
		UserId,_:=id.LastInsertId()
		return UserId, err
		}
		return 0, err
	}

func (data AddProductModel) AddProductDb(insertData AddProduct) (int64, error) {
	var ProductVar AddProduct
	
	insertData.UserId,_= data.CreatedNewUser() //creating  user before creating a product
	imagesJSON ,_:=MarshalImg(insertData.ProductImages) //marshal img to save into db  

	id,err := utility.Db.NamedExec("INSERT INTO `products` (product_name, product_description, product_images, product_price,user_id) VALUES (:ProductName, :Product_description, :Product_images, :Product_price,:User_id)",
		map[string]interface{}{
		"ProductName":insertData.ProductName,
		"Product_description":insertData.ProductDescription,
		"Product_images":string(imagesJSON),
		"Product_price":insertData.ProductPrice,
		"User_id":insertData.UserId})

	if err != nil {
		log.Println(err)
	}else{
		ProductVar.Product_id,_=id.LastInsertId()
		return ProductVar.Product_id, err
	}
	return 0, err
}

func (user AddProductModel)GetProductImagesByID(id int64) ([]byte, error) {
    var productImages []byte
    query := "SELECT product_images FROM products WHERE product_id=?"
    err :=utility.Db.QueryRow(query, id).Scan(&productImages)
    if err != nil {
		log.Println(err)
        return nil, err
    }
    return productImages, nil
}

func (imgM AddProductModel)AddCompressedImagePath(compressImg []string,id int64)error{
	jsonImg,_:= MarshalImg(compressImg) //marshal img to save into db
	_, err := utility.Db.NamedExec("UPDATE products SET compressed_product_images = :Compressed_product_images WHERE product_id = :Product_id",
    map[string]interface{}{
        "Compressed_product_images": jsonImg,
        "Product_id": id,
    })
	if err != nil {
		log.Println(err)
	}
	return  err
}

func MarshalImg(imgPath []string)([]byte,error){
	productImagesJSON, err := json.Marshal(imgPath)
	if err != nil {
		log.Println(err)
	}
	return productImagesJSON,err
}