package product

import (
	"gopkg.in/jeevatkm/go-model.v1"
	
	"gorm.io/gorm"
	"fmt"
	"product/external"
)

type Product struct {
	gorm.Model
	Id int `gorm:"primaryKey" json:"id" type:"int"`
	Name string `json:"name"`
	Stock int `json:"stock"`

}

func (self *Product) onPostPersist() (err error){
	productChanged := NewProductChanged()
	model.Copy(productChanged, self)

	Publish(productChanged)

	return nil
}
func (self *Product) onPrePersist() (err error){ return nil }
func (self *Product) onPreUpdate() (err error){ return nil }
func (self *Product) onPostUpdate() (err error){ return nil }
func (self *Product) onPreRemove() (err error){ return nil }
func (self *Product) onPostRemove() (err error){ return nil }


func ItIsNotCommon(productChanged *ProductChanged){
	/** Example 1:  new item
	product := &Product{}
	productrepository.save(product)

	*/

	/** Example 2:  finding and process
	id, _ := strconv.ParseInt(productChanged.id, 10, 64)
	product, err := productrepository().FindById(int(id))
	if err != nil {

	}
	*/
}
