package controller
import
(
	"msg_quing_system/models"
	"msg_quing_system/utility"
)
type Env struct{
	product interface{
		AddProductDb(models.AddProduct)(int64,error)
		GetProductImagesByID(id int64)([]byte,error)
		AddCompressedImagePath(imgArr []string,id int64)error
	}
}

var Db *Env
func init(){
	Db=&Env{
		product: models.AddProductModel{DB: utility.Db},
	}
}
