package forms
const(
        EmptyFormErr  = "this field cannot be empty"
        LengthFormErr = "this field must be more than"
        EmailFormErr = "invalid email address"
        StartDateFormErr = "please enter valid start date"
        EndDateFormErr = "please enter valid end date"
)


 

type errors map[string][]string

func(e errors) Add(field, message string){
    
    e[field] = append(e[field], message)
}
func (e errors)Get(field string) string{
   es := e[field]
   if len(es) == 0{
    return ""
   }
    return es[0]
} 