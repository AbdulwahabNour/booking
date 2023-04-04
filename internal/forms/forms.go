package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)


type Form struct{
    url.Values
    Errors errors
}

func New(data url.Values) *Form{

    return &Form{
        data,
        make(errors),
    }
}
// func(f *Form)TrimValues()
func (f *Form)Require(fields ... string)   {
  
    for _, field := range fields{

        input :=  f.Values.Get(field) 
        if input == ""{
            f.Errors.Add(field, EmptyFormErr)
        
        }
    }

} 
func (f *Form) MinLength(field string, length int){
        val := f.Values.Get(field)
        if len(val) < length{
            f.Errors.Add(field, fmt.Sprintf("%s %d",LengthFormErr,length))
        }
    
}
func (f *Form) IsEmail(field string){
    if !govalidator.IsEmail(f.Get(field)) {
           f.Errors.Add(field, EmailFormErr) 
    }
}

 
func (f *Form)Valid()bool{
   return len(f.Errors) == 0
}
func(f *Form)TrimValues(){
 
    for key, val := range f.Values {
        // Trim leading and trailing whitespace from the value.
        val[0] = strings.TrimSpace(val[0])
        // Update the value in the map.
        f.Values[key] = val
    }
}
